package repository

import (
	"errors"

	db "gwi/assignment/core/data/database"
	coreEnt "gwi/assignment/core/data/entity"
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/chart/data/entity"
	userEnt "gwi/assignment/feature/user/data/entity"
)

type ChartRepository struct{}

var (
	ErrChartNotFound         = errors.New("chart not found")
	ErrChartNotFavorited     = errors.New("chart is not favorited")
	ErrChartAlreadyFavorited = errors.New("chart is already favorited")
)

func (repo *ChartRepository) GetAllCharts(page int) (*[]ent.Chart, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	charts := &[]ent.Chart{}
	err = db.Preload("Points").
		Preload("Asset").
		Offset((page - 1) * 20).
		Limit(20).
		Find(charts).Error

	return charts, err
}

func (repo *ChartRepository) GetChartById(chartId string) (*ent.Chart, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	chart := &ent.Chart{}
	err = db.Preload("Points").
		Preload("Asset").
		Where("id = ?", chartId).
		First(chart).Error

	return chart, err
}

func (repo *ChartRepository) CreateChart(chart *ent.Chart) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	return db.Create(chart).Error
}

func (repo *ChartRepository) UpdateChart(chartId string, chart *ent.Chart) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	dbChart, err := repo.GetChartById(chartId)
	if err != nil {
		return err
	}

	chart.Id = dbChart.Id
	chart.AssetId = dbChart.AssetId
	chart.Asset.Id = dbChart.AssetId

	point := ent.ChartPoint{}
	if err := db.Where("chart_id = ?", chart.Id).Delete(point).Error; err != nil {
		return err
	}

	return db.Updates(chart).Error
}

func (repo *ChartRepository) DeleteChart(chartId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	chart, err := repo.GetChartById(chartId)
	if err != nil {
		return err
	}

	point := ent.ChartPoint{}
	if err := db.Where("chart_id = ?", chart.Id).Delete(point).Error; err != nil {
		return err
	}

	err = db.Where("id = ?", chart.Id).Delete(chart).Error
	if err != nil {
		return err
	}

	asset := &assetEnt.Asset{}

	return db.Where("id = ?", chart.AssetId).Delete(asset).Error
}

func (repo *ChartRepository) GetFavoriteCharts(userId string, page int) (*[]ent.Chart, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	charts := &[]ent.Chart{}
	err = db.Table("users_assets").
		Select("charts.*").
		Preload("Points").
		Preload("Asset").
		Joins("INNER JOIN assets ON assets.id = users_assets.asset_id").
		Joins("INNER JOIN charts ON charts.asset_id = assets.id").
		Where("users_assets.user_id = ?", userId).
		Offset((page - 1) * 20).
		Limit(20).
		Find(charts).Error

	return charts, err
}

func (repo *ChartRepository) FavoriteChart(userId, chartId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	assetId, err := repo.getAssetIdFromChartId(chartId)
	if err != nil {
		return ErrChartNotFound
	}

	user := &userEnt.User{
		BaseEntity: coreEnt.BaseEntity{
			Id: userId,
		},
	}

	asset := &assetEnt.Asset{
		BaseEntity: coreEnt.BaseEntity{
			Id: assetId,
		},
	}

	return db.Model(asset).Association("Users").Append(user)
}

func (repo *ChartRepository) UnfavoriteChart(userId, chartId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	assetId, err := repo.getAssetIdFromChartId(chartId)
	if err != nil {
		return ErrChartNotFound
	}

	user := &userEnt.User{
		BaseEntity: coreEnt.BaseEntity{
			Id: userId,
		},
	}

	asset := &assetEnt.Asset{
		BaseEntity: coreEnt.BaseEntity{
			Id: assetId,
		},
	}

	return db.Model(asset).Association("Users").Delete(user)
}

func (repo *ChartRepository) getAssetIdFromChartId(chartId string) (string, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return "", err
	}

	asset := &ent.Chart{}
	err = db.
		Select("asset_id").
		Where("id = ?", chartId).
		First(asset).Error

	return asset.AssetId, err
}

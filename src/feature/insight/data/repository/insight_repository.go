package repository

import (
	"errors"

	db "gwi/assignment/core/data/database"
	coreEnt "gwi/assignment/core/data/entity"
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/insight/data/entity"
	userEnt "gwi/assignment/feature/user/data/entity"
)

type InsightRepository struct{}

var (
	ErrInsightNotFound         = errors.New("insight not found")
	ErrInsightNotFavorited     = errors.New("insight is not favorited")
	ErrInsightAlreadyFavorited = errors.New("insight is already favorited")
)

func (repo *InsightRepository) GetAllInsights(page int) (*[]ent.Insight, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	insights := &[]ent.Insight{}
	err = db.Preload("Asset").
		Offset((page - 1) * 20).
		Limit(20).
		Find(insights).Error

	return insights, err
}

func (repo *InsightRepository) GetInsightById(insightId string) (*ent.Insight, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	insight := &ent.Insight{}
	err = db.Preload("Asset").
		Where("id = ?", insightId).
		First(insight).Error

	return insight, err
}

func (repo *InsightRepository) CreateInsight(insight *ent.Insight) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	return db.Create(insight).Error
}

func (repo *InsightRepository) UpdateInsight(insightId string, insight *ent.Insight) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	dbInsight, err := repo.GetInsightById(insightId)
	if err != nil {
		return err
	}

	insight.Id = dbInsight.Id
	insight.AssetId = dbInsight.AssetId
	insight.Asset.Id = dbInsight.AssetId

	return db.Updates(insight).Error
}

func (repo *InsightRepository) DeleteInsight(insightId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	insight, err := repo.GetInsightById(insightId)
	if err != nil {
		return err
	}

	if err = db.Where("id = ?", insight.Id).Delete(insight).Error; err != nil {
		return err
	}

	asset := &assetEnt.Asset{}

	return db.Where("id = ?", insight.AssetId).Delete(asset).Error
}

func (repo *InsightRepository) GetFavoriteInsights(userId string, page int) (*[]ent.Insight, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	insights := &[]ent.Insight{}
	err = db.Table("users_assets").
		Select("insights.*").
		Preload("Asset").
		Joins("INNER JOIN assets ON assets.id = users_assets.asset_id").
		Joins("INNER JOIN insights ON insights.asset_id = assets.id").
		Where("users_assets.user_id = ?", userId).
		Offset((page - 1) * 20).
		Limit(20).
		Find(insights).Error

	return insights, err
}

func (repo *InsightRepository) FavoriteInsight(userId, insightId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	assetId, err := repo.getAssetIdFromInsightId(insightId)
	if err != nil {
		return ErrInsightNotFound
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

func (repo *InsightRepository) UnfavoriteInsight(userId, insightId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	assetId, err := repo.getAssetIdFromInsightId(insightId)
	if err != nil {
		return ErrInsightNotFound
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

func (repo *InsightRepository) getAssetIdFromInsightId(insightId string) (string, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return "", err
	}

	insight := &ent.Insight{}
	err = db.
		Select("asset_id").
		Where("id = ?", insightId).
		First(insight).Error

	return insight.AssetId, err
}

package repository

import (
	"errors"

	db "gwi/assignment/core/data/database"
	coreEnt "gwi/assignment/core/data/entity"
	assetEnt "gwi/assignment/feature/asset/data/entity"
	ent "gwi/assignment/feature/audience/data/entity"
	userEnt "gwi/assignment/feature/user/data/entity"
)

type AudienceRepository struct{}

var (
	ErrAudienceNotFound         = errors.New("audience not found")
	ErrAudienceNotFavorited     = errors.New("audience is not favorited")
	ErrAudienceAlreadyFavorited = errors.New("audience is already favorited")
)

func (repo *AudienceRepository) GetAllAudiences(page int) (*[]ent.Audience, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	audiences := &[]ent.Audience{}
	err = db.Preload("StatType").
		Preload("Asset").
		Offset((page - 1) * 20).
		Limit(20).
		Find(audiences).Error

	return audiences, err
}

func (repo *AudienceRepository) GetAudienceById(audienceId string) (*ent.Audience, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	audience := &ent.Audience{}
	err = db.Preload("StatType").
		Preload("Asset").
		Where("id = ?", audienceId).
		First(audience).Error

	return audience, err
}

func (repo *AudienceRepository) CreateAudience(audience *ent.Audience) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	return db.Create(audience).Error
}

func (repo *AudienceRepository) UpdateAudience(audienceId string, audience *ent.Audience) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	dbAudience, err := repo.GetAudienceById(audienceId)
	if err != nil {
		return err
	}

	audience.Id = dbAudience.Id
	audience.AssetId = dbAudience.AssetId
	audience.Asset.Id = dbAudience.AssetId
	audience.StatTypeId = dbAudience.StatTypeId
	audience.StatType.Id = dbAudience.StatTypeId

	statType := &ent.AudienceStatType{
		Title:          audience.StatType.Title,
		TitleFormatted: audience.StatType.TitleFormatted,
	}
	statType.Id = audience.StatTypeId

	if err := db.Updates(statType).Error; err != nil {
		return err
	}

	return db.Updates(audience).Error
}

func (repo *AudienceRepository) DeleteAudience(audienceId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	audience, err := repo.GetAudienceById(audienceId)
	if err != nil {
		return err
	}

	if err = db.Where("id = ?", audience.Id).Delete(audience).Error; err != nil {
		return err
	}

	statType := &ent.AudienceStatType{}
	err = db.Where("id = ?", audience.StatTypeId).Delete(statType).Error
	if err != nil {
		return err
	}

	asset := &assetEnt.Asset{}

	return db.Where("id = ?", audience.AssetId).Delete(asset).Error
}

func (repo *AudienceRepository) GetFavoriteAudiences(userId string, page int) (*[]ent.Audience, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return nil, err
	}

	audiences := &[]ent.Audience{}
	err = db.Table("users_assets").
		Select("audiences.*").
		Preload("StatType").
		Preload("Asset").
		Joins("INNER JOIN assets ON assets.id = users_assets.asset_id").
		Joins("INNER JOIN audiences ON audiences.asset_id = assets.id").
		Where("users_assets.user_id = ?", userId).
		Offset((page - 1) * 20).
		Limit(20).
		Find(audiences).Error

	return audiences, err
}

func (repo *AudienceRepository) FavoriteAudience(userId, audienceId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	assetId, err := repo.getAssetIdFromAudienceId(audienceId)
	if err != nil {
		return ErrAudienceNotFound
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

func (repo *AudienceRepository) UnfavoriteAudience(userId, audienceId string) error {
	db, err := db.GetGormConnection()
	if err != nil {
		return err
	}

	assetId, err := repo.getAssetIdFromAudienceId(audienceId)
	if err != nil {
		return ErrAudienceNotFound
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

func (repo *AudienceRepository) getAssetIdFromAudienceId(audienceId string) (string, error) {
	db, err := db.GetGormConnection()
	if err != nil {
		return "", err
	}

	asset := &ent.Audience{}
	err = db.
		Select("asset_id").
		Where("id = ?", audienceId).
		First(asset).Error

	return asset.AssetId, err
}

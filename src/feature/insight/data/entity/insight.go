package entity

import (
	coreEntity "gwi/assignment/core/data/entity"
	assetEntity "gwi/assignment/feature/asset/data/entity"
)

type Insight struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	AssetId string            `sql:"type:uuid" gorm:"not null;unique_index"`
	Asset   assetEntity.Asset `gorm:"foreignKey:AssetId"`
}

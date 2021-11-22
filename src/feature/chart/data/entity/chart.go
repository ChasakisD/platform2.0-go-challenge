package entity

import (
	coreEntity "gwi/assignment/core/data/entity"
	assetEntity "gwi/assignment/feature/asset/data/entity"
)

type Chart struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	XAxes   string            `gorm:"not null;"`
	YAxes   string            `gorm:"not null;"`
	Points  []ChartPoint      `gorm:"foreignKey:ChartId"`
	AssetId string            `sql:"type:uuid" gorm:"not null;unique_index"`
	Asset   assetEntity.Asset `gorm:"foreignKey:AssetId"`
}

type ChartPoint struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	ChartId string  `gorm:"not null;"`
	XValue  float64 `gorm:"not null;"`
	YValue  float64 `gorm:"not null;"`
}

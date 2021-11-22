package entity

import (
	"fmt"
	"strings"

	coreEntity "gwi/assignment/core/data/entity"
	assetEntity "gwi/assignment/feature/asset/data/entity"
)

type Audience struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	Gender        string
	BirthCountry  string
	AgeGroupMin   int `gorm:"not null;"`
	AgeGroupMax   int `gorm:"not null;"`
	StatTypeValue float64
	StatTypeId    string
	StatType      AudienceStatType  `gorm:"foreignKey:StatTypeId"`
	AssetId       string            `sql:"type:uuid" gorm:"not null;unique_index"`
	Asset         assetEntity.Asset `gorm:"foreignKey:AssetId"`
}

type AudienceStatType struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	Title          string `gorm:"not null;"`
	TitleFormatted string `gorm:"not null;"`
}

func (audience *Audience) GenerateDescription() string {
	valueString := fmt.Sprintf("%.1f", audience.StatTypeValue)
	statFormatted := strings.Replace(audience.StatType.TitleFormatted, "%s", valueString, 1)
	return fmt.Sprintf("%ss from %d to %d on %s have %s",
		audience.Gender,
		audience.AgeGroupMin,
		audience.AgeGroupMax,
		audience.BirthCountry,
		statFormatted)
}

package entity

import (
	coreEntity "gwi/assignment/core/data/entity"
	userEnt "gwi/assignment/feature/user/data/entity"
)

type Asset struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	Description string         `gorm:"not null;unique_index;"`
	Users       []userEnt.User `gorm:"many2many:users_assets"`
}

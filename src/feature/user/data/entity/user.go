package entity

import (
	coreEntity "gwi/assignment/core/data/entity"
)

type User struct {
	coreEntity.BaseEntity
	coreEntity.AuditEntity
	Username string
	Email    string
	Password string
}

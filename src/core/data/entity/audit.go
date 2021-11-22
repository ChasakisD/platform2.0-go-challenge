package entity

import (
	"time"
)

type AuditEntity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

package dbmodels

import (
	"time"

	"gorm.io/gorm"
)

type DbModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

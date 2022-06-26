package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int            `gorm:"primaryKey;column:id" json:"id"`
	CreatedAt time.Time      `gorm:"primaryKey;column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"primaryKey;column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

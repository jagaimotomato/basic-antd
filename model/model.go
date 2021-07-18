package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type By struct {
	CreatedBy string `gorm:"size:64" json:"created_by"`
	UpdatedBy string `gorm:"size:64" json:"updated_by"`
}

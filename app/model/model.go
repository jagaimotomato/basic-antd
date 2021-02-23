package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type By struct {
	CreatedBy string `gorm:"size:64" json:"createdBy"`
	UpdatedBy string `gorm:"size:64" json:"updatedBy"`
}

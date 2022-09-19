package gormhelper

import (
	"time"

	"gorm.io/gorm"
)

type Pagination struct {
	Page     int
	PageSize int
}

type DateModel struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

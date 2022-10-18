package gormhelper

import (
	"time"

	"gorm.io/gorm"
)

type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	Skip     bool `json:"no_pagination"`
}

type DateModel struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

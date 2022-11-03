package sinkronisasi

import (
	"time"

	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"gorm.io/gorm"
)

type Sinkronisasi struct {
	Id                  uint64          `json:"id" gorm:"primaryKey"`
	TanggalSinkronisasi *time.Time      `json:"tanggalSinkronisasi"`
	File                *string         `json:"file"`
	JumlahTidakSinkron  *int            `json:"jumlahTidakSinkron"`
	Tbp_Id              *uint64         `json:"tbp_id"`
	Tbp                 *mt.Tbp         `json:"tbp,omitempty" gorm:"foreignKey:Tbp_Id"`
	User_Id             *uint64         `json:"user_id"`
	User                *mu.User        `json:"user,omitempty" gorm:"foreignKey:User_Id"`
	UpdatedAt           *time.Time      `json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

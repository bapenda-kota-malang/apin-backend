package sinkronisasi

import (
	"time"

	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sinkronisasi struct {
	Id                  uint64          `json:"id" gorm:"primaryKey"`
	TanggalSinkronisasi *time.Time      `json:"tanggalSinkronisasi"`
	JenisPajak          string          `json:"jenisPajak"`
	File                string          `json:"file"`
	JumlahTidakSinkron  *int            `json:"jumlahTidakSinkron"`
	Tbp_Id              *uint64         `json:"tbp_id"`
	Tbp                 *mt.Tbp         `json:"tbp,omitempty" gorm:"foreignKey:Tbp_Id"`
	User_Id             *uint64         `json:"user_id"`
	User                *mu.User        `json:"user,omitempty" gorm:"foreignKey:User_Id"`
	UpdatedAt           *time.Time      `json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type CreateDto struct {
	TanggalSinkronisasi *time.Time      `json:"tanggalSinkronisasi"`
	JenisPajak          string          `json:"jenisPajak"`
	File                string          `json:"file"`
	JumlahTidakSinkron  *int            `json:"jumlahTidakSinkron"`
	Tbp_Id              *uint64         `json:"tbp_id"`
	User_Id             *uint64         `json:"user_id"`
	UpdatedAt           *time.Time      `json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt `json:"deletedAt"`
}

type UpdateDto struct {
	Id                  *uint64         `json:"id"`
	TanggalSinkronisasi *time.Time      `json:"tanggalSinkronisasi"`
	JenisPajak          string          `json:"jenisPajak"`
	File                *string         `json:"file"`
	JumlahTidakSinkron  *int            `json:"jumlahTidakSinkron"`
	Tbp_Id              *uint64         `json:"tbp_id"`
	User_Id             *uint64         `json:"user_id"`
	UpdatedAt           *time.Time      `json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt `json:"deletedAt"`
}

type FilterDto struct {
	TanggalSinkronisasi *time.Time      `json:"tanggalSinkronisasi"`
	JenisPajak          *string         `json:"jenisPajak"`
	File                *string         `json:"file"`
	JumlahTidakSinkron  *int            `json:"jumlahTidakSinkron"`
	Tbp_Id              *uint64         `json:"tbp_id"`
	User_Id             *uint64         `json:"user_id"`
	UpdatedAt           *time.Time      `json:"updatedAt"`
	DeletedAt           *gorm.DeletedAt `json:"deletedAt"`
}

type SinkronisasiMerge struct {
	Spt_Id             *uuid.UUID `json:"spt_id"`
	Sspd_No            *uint64    `json:"sspd_no"`
	Sspd_Tanggal       *uint64    `json:"Sspd_Tanggal"`
	Spt_Nominal        *float32   `json:"spt_nominal"`
	Spt_NominalDenda   *float64   `json:"spt_nominalDenda"`
	KodeBayar          *string    `json:"kodeBayar"`
	Excel_Nominal      *float64   `json:"excel_nominal"`
	Excel_NominalDenda *float64   `json:"excel_nominalDenda"`
}

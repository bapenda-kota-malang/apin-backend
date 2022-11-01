package objekpajak

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type ObjekPajak struct {
	Id           uint64                  `json:"id" gorm:"primarykey"`
	Nama         *string                 `json:"nama" gorm:"size:200" validate:"required"`
	Nop          *string                 `json:"nop" gorm:"size:50"`
	Alamat       *string                 `json:"alamat" gorm:"size:200" validate:"required"`
	RtRw         *string                 `json:"rtRw" gorm:"size:10" validate:"required"`
	Kecamatan_Id *uint64                 `json:"kecamatan_id" validate:"required"`
	Kecamatan    *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id *uint64                 `json:"kelurahan_id" validate:"required"`
	Kelurahan    *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
	Longitude    *float64                `json:"longitude" gorm:"type:decimal(9,6)"`
	Latitude     *float64                `json:"latitude" gorm:"type:decimal(8,6)"`
	Telp         *string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL              `json:"status"`
	IsNpwpd      uint64                  `json:"isNpwpd"`
	gormhelper.DateModel
}

type ObjekPajakCreate struct {
	Nama         *string    `json:"nama" validate:"required"`
	Nop          *string    `json:"nop"`
	Alamat       *string    `json:"alamat"  validate:"required"`
	RtRw         *string    `json:"rtRw" validate:"required"`
	Kecamatan_Id *uint64    `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id *uint64    `json:"kelurahan_id" validate:"required"`
	Longitude    *float64   `json:"longitude" validate:"required"`
	Latitude     *float64   `json:"latitude" validate:"required"`
	Telp         *string    `json:"telp" validate:"nohp"`
	Status       t.StatusBL `json:"status"`
}

type ObjekPajakUpdate struct {
	Id           uint64     `json:"id" validate:"required"`
	Nama         *string    `json:"nama" validate:"required"`
	Nop          *string    `json:"nop"`
	Alamat       *string    `json:"alamat"  validate:"required"`
	RtRw         *string    `json:"rtRw" validate:"required"`
	Kecamatan_Id *uint64    `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id *uint64    `json:"kelurahan_id" validate:"required"`
	Longitude    *float64   `json:"longitude" validate:"required"`
	Latitude     *float64   `json:"latitude" validate:"required"`
	Telp         *string    `json:"telp" validate:"nohp"`
	Status       t.StatusBL `json:"status"`
}

type FilterDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

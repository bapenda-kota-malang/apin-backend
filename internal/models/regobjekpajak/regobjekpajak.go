package regobjekpajak

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegObjekPajak struct {
	Id           uint64                  `json:"id" gorm:"primarykey"`
	Nama         *string                 `json:"nama" gorm:"size:200"`
	Nop          *string                 `json:"nop" gorm:"size:50"`
	Alamat       *string                 `json:"alamat" gorm:"size:200"`
	RtRw         *string                 `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id *uint64                 `json:"kecamatan_id"`
	Kecamatan    *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id *uint64                 `json:"kelurahan_id"`
	Kelurahan    *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
	Telp         *string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL              `json:"status"`
	IsNpwpd      uint64                  `json:"isNpwpd"`
	gormhelper.DateModel
}

type RegObjekPajakCreateDto struct {
	Nama         *string    `json:"nama" validate:"required"`
	Nop          *string    `json:"nop"`
	Alamat       *string    `json:"alamat" validate:"required"`
	RtRw         *string    `json:"rtRw" validate:"required"`
	Kecamatan_Id *uint64    `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id *uint64    `json:"kelurahan_id" validate:"required"`
	Telp         *string    `json:"telp" validate:"nohp"`
	Status       t.StatusBL `json:"status"`
}

type RegObjekPajakUpdateDto struct {
	Id           uint64     `json:"id" gorm:"primaryKey"`
	Nama         *string    `json:"nama" validate:"required"`
	Nop          *string    `json:"nop"`
	Alamat       *string    `json:"alamat" validate:"required"`
	RtRw         *string    `json:"rtRw" validate:"required"`
	Kecamatan_Id *uint64    `json:"kecamatan_id" validate:"required"`
	Kelurahan_Id *uint64    `json:"kelurahan_id" validate:"required"`
	Telp         *string    `json:"telp" validate:"nohp"`
	Status       t.StatusBL `json:"status"`
}

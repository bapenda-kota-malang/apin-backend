package objekpajak

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type ObjekPajak struct {
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

type ObjekPajakCreate struct {
	Nama         *string    `json:"nama" gorm:"size:200"`
	Nop          *string    `json:"nop" gorm:"size:50"`
	Alamat       *string    `json:"alamat" gorm:"size:200"`
	RtRw         *string    `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id *uint64    `json:"kecamatan_id"`
	Kelurahan_Id *uint64    `json:"kelurahan_id"`
	Telp         *string    `json:"telp" gorm:"size:20"`
	Status       t.StatusBL `json:"status"`
}

type ObjekPajakUpdate struct {
	Id           uint64  `json:"id" gorm:"primaryKey"`
	Nama         *string `json:"nama" gorm:"size:200"`
	Nop          *string `json:"nop" gorm:"size:50"`
	Alamat       *string `json:"alamat" gorm:"size:200"`
	RtRw         *string `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Telp         *string `json:"telp" gorm:"size:20"`
}

type FilterDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

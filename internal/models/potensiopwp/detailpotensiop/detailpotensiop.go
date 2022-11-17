package detailpotensiop

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailPotensiOp struct {
	Id           uint                    `json:"id" gorm:"primaryKey"`
	Potensiop_Id uint                    `json:"potensiop_id"`
	Nama         string                  `json:"nama" gorm:"size:50"`
	Nop          *string                 `json:"nop" gorm:"size:50"`
	Alamat       string                  `json:"alamat" gorm:"size:50"`
	RtRw         string                  `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id uint                    `json:"kecamatan_id"`
	Kelurahan_Id uint                    `json:"kelurahan_id"`
	Longitude    *float64                `json:"longitude" gorm:"type:decimal(9,6)"`
	Latitude     *float64                `json:"latitude" gorm:"type:decimal(8,6)"`
	Telp         *string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL              `json:"status"`
	IsNpwpd      bool                    `json:"isNpwpd"`
	Kecamatan    *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan    *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
	gormhelper.DateModel
}

type CreateDto struct {
	Potensiop_Id uint       `json:"-"`
	Nama         string     `json:"nama" validate:"required"`
	Nop          *string    `json:"nop"`
	Alamat       string     `json:"alamat" validate:"required"`
	RtRw         string     `json:"rtRw" validate:"required"`
	Kecamatan_Id uint       `json:"kecamatan_id" validate:"required;min=1"`
	Kelurahan_Id uint       `json:"kelurahan_id" validate:"required;min=1"`
	Longitude    *float64   `json:"longitude"`
	Latitude     *float64   `json:"latitude"`
	Telp         string     `json:"telp" validate:"nohp"`
	Status       t.StatusBL `json:"status" validate:"required"`
	IsNpwpd      bool       `json:"-"`
}

type UpdateDto struct {
	Nama         *string     `json:"nama"`
	Nop          *string     `json:"nop"`
	Alamat       *string     `json:"alamat"`
	RtRw         *string     `json:"rtRw"`
	Kecamatan_Id *uint       `json:"kecamatan_id" validate:"min=1"`
	Kelurahan_Id *uint       `json:"kelurahan_id" validate:"min=1"`
	Longitude    *float64    `json:"longitude"`
	Latitude     *float64    `json:"latitude"`
	Telp         string      `json:"telp" validate:"nohp"`
	Status       *t.StatusBL `json:"status"`
}

package regnpwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

type RegNarahubung struct {
	Id           uint64                 `json:"id" gorm:"primaryKey"`
	RegNpwpd_Id  uint64                 `json:"regNpwpd_id"`
	RegNpwpd     *RegNpwpd              `json:"regNpwpd,omitempty" gorm:"foreignKey:RegNpwpd_Id;references:Id"`
	Nama         string                 `json:"nama" gorm:"size:50"`
	Alamat       string                 `json:"alamat" gorm:"size:50"`
	Daerah_Id    *uint64                `json:"daerah_id"`
	Daerah       *areadivision.Daerah   `json:"daerah" gorm:"foreignKey:Daerah_Id"`
	Kelurahan_Id uint64                 `json:"kelurahan_id"`
	Kelurahan    areadivision.Kelurahan `json:"kelurahan" gorm:"foreignKey:Kelurahan_Id"`
	Telp         string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL             `json:"status"`
	Nik          string                 `json:"nik" gorm:"size:20"`
	Email        string                 `json:"email" gorm:"type:varchar"`
}

type RegNarahubungCreateDto struct {
	Nama         string     `json:"nama" validate:"required"`
	Alamat       string     `json:"alamat" validate:"required"`
	Daerah_Id    *uint64    `json:"daerah_id" validate:"required"`
	Kelurahan_Id *uint64    `json:"kelurahan_id" validate:"required"`
	Telp         string     `json:"telp" validate:"required;notelp"`
	Nik          string     `json:"nik" validate:"required;nik"`
	Status       t.StatusBL `json:"status"`
	Email        string     `json:"email" validate:"required;emailvalid"`
}

type RegNarahubungUpdateDto struct {
	Id           uint64     `json:"id"`
	Nama         string     `json:"nama" validate:"required"`
	Alamat       string     `json:"alamat" validate:"required"`
	Daerah_Id    *uint64    `json:"daerah_id" validate:"required"`
	Kelurahan_Id *uint64    `json:"kelurahan_id" validate:"required"`
	Telp         string     `json:"telp" validate:"required;notelp"`
	Nik          string     `json:"nik" validate:"required;nik"`
	Status       t.StatusBL `json:"status"`
	Email        string     `json:"email" validate:"required;emailvalid"`
}

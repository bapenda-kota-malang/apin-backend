package npwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
)

type Narahubung struct {
	Id           uint64                 `json:"id" gorm:"primaryKey"`
	Npwpd_Id     uint64                 `json:"npwpd_id"`
	Npwpd        Npwpd                  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	Nama         string                 `json:"nama" gorm:"size:50"`
	Alamat       string                 `json:"alamat" gorm:"size:50"`
	Daerah_Id    *uint64                `json:"daerah_id"`
	Daerah       *areadivision.Daerah   `gorm:"foreignKey:Daerah_Id"`
	Kelurahan_Id uint64                 `json:"kelurahan_id"`
	Kelurahan    areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	Telp         string                 `json:"telp" gorm:"size:20"`
	Keterangan   *string                `json:"keterangan" gorm:"type:varchar(100)"`
	Nik          string                 `json:"nik" gorm:"size:20"`
	Email        string                 `json:"email" gorm:"type:varchar"`
}
type NarahubungCreate struct {
	Nama         string  `json:"nama" validate:"required"`
	Alamat       string  `json:"alamat" validate:"required"`
	Daerah_Id    uint64  `json:"daerah_id" validate:"required"`
	Kelurahan_Id uint64  `json:"kelurahan_id" validate:"required"`
	Telp         string  `json:"telp" validate:"required;nohp"`
	Keterangan   *string `json:"keterangan"`
	Nik          string  `json:"nik" validate:"required;nik"`
	Email        string  `json:"email" validate:"required;emailvalid"`
}

type NarahubungUpdate struct {
	Id           uint64  `json:"id" validate:"required"`
	Nama         string  `json:"nama" validate:"required"`
	Alamat       string  `json:"alamat" validate:"required"`
	Daerah_Id    uint64  `json:"daerah_id" validate:"required"`
	Kelurahan_Id uint64  `json:"kelurahan_id" validate:"required"`
	Telp         string  `json:"telp" validate:"required;nohp"`
	Keterangan   *string `json:"keterangan"`
	Nik          string  `json:"nik" validate:"required;nik"`
	Email        string  `json:"email" validate:"required;emailvalid"`
}

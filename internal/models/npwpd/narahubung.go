package npwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

type Narahubung struct {
	Id           uint64                 `json:"id" gorm:"primaryKey"`
	Npwpd_Id     uint64                 `json:"npwpd_id"`
	Npwpd        Npwpd                  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	Nama         string                 `json:"nama" gorm:"size:50"`
	Alamat       string                 `json:"alamat" gorm:"size:50"`
	RtRw         string                 `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id uint64                 `json:"kecamatan_id"`
	Kecamatan    areadivision.Kecamatan `gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id uint64                 `json:"kelurahan_id"`
	Kelurahan    areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	Telp         string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL             `json:"status"`
	Nik          string                 `json:"nik" gorm:"size:20"`
}
type NarahubungCreate struct {
	Nama         string `json:"nama" gorm:"size:50"`
	Alamat       string `json:"alamat" gorm:"size:50"`
	RtRw         string `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id uint64 `json:"kecamatan_id"`
	Kelurahan_Id uint64 `json:"kelurahan_id"`
	Telp         string `json:"telp" gorm:"size:20"`
	Nik          string `json:"nik" gorm:"size:20"`
}

type NarahubungUpdate struct {
	Id           uint64 `json:"id" gorm:"primaryKey"`
	Nama         string `json:"nama" gorm:"size:50"`
	Alamat       string `json:"alamat" gorm:"size:50"`
	RtRw         string `json:"rtRw" gorm:"size:10"`
	Kecamatan_Id uint64 `json:"kecamatan_id"`
	Kelurahan_Id uint64 `json:"kelurahan_id"`
	Telp         string `json:"telp" gorm:"size:20"`
	Nik          string `json:"nik" gorm:"size:20"`
}

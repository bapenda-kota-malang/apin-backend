package npwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

type Narahubung struct {
	Id           uint64                 `json:"id" gorm:"primaryKey"`
	Npwpd_Id     uint64                 `json:"npwpd_id"`
	Npwpd        Npwpd                  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	Nama         string                 `json:"nama" gorm:"size:50" validate:"required"`
	Alamat       string                 `json:"alamat" gorm:"size:50" validate:"required"`
	RtRw         string                 `json:"rtRw" gorm:"size:10" validate:"required"`
	Daerah_Id    *uint64                `json:"daerah_id" validate:"required"`
	Daerah       *areadivision.Daerah   `gorm:"foreignKey:Daerah_Id"`
	Kelurahan_Id uint64                 `json:"kelurahan_id" validate:"required"`
	Kelurahan    areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	Telp         string                 `json:"telp" gorm:"size:20" validate:"required"`
	Status       t.StatusBL             `json:"status"`
	Nik          string                 `json:"nik" gorm:"size:20" validate:"required"`
	Email        string                 `json:"email" gorm:"type:varchar"`
}
type NarahubungCreate struct {
	Nama         string `json:"nama" gorm:"size:50" validate:"required"`
	Alamat       string `json:"alamat" gorm:"size:50" validate:"required"`
	RtRw         string `json:"rtRw" gorm:"size:10" validate:"required"`
	Daerah_Id    uint64 `json:"daerah_id" validate:"required"`
	Kelurahan_Id uint64 `json:"kelurahan_id" validate:"required"`
	Telp         string `json:"telp" gorm:"size:20" validate:"required"`
	Nik          string `json:"nik" gorm:"size:20" validate:"required"`
	Email        string `json:"email" validate:"required;emailvalid"`
}

type NarahubungUpdate struct {
	Id           uint64 `json:"id" gorm:"primaryKey"`
	Nama         string `json:"nama" gorm:"size:50"`
	Alamat       string `json:"alamat" gorm:"size:50"`
	RtRw         string `json:"rtRw" gorm:"size:10"`
	Daerah_Id    uint64 `json:"daerah_id"`
	Kelurahan_Id uint64 `json:"kelurahan_id"`
	Telp         string `json:"telp" gorm:"size:20"`
	Nik          string `json:"nik" gorm:"size:20"`
	Email        string `json:"email" validate:"required;emailvalid"`
}

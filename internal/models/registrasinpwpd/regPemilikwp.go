package registrasinpwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

type RegPemilikWp struct {
	Id                 uint64                  `json:"id" gorm:"primaryKey"`
	RegistrasiNpwpd_Id uint64                  `json:"registrasiNpwpd_id"`
	RegistrasiNpwpd    *RegistrasiNpwpd        `json:"npwpd,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	Nama               *string                 `json:"nama" gorm:"size:50"`
	Alamat             *string                 `json:"alamat" gorm:"size:50"`
	RtRw               *string                 `json:"rtRw" gorm:"size:10"`
	Daerah_Id          *uint64                 `json:"daerah_id"`
	Daerah             *areadivision.Daerah    `json:"daerah" gorm:"foreignKey:Daerah_Id"`
	Kelurahan_Id       *uint64                 `json:"kelurahan_id"`
	Kelurahan          *areadivision.Kelurahan `json:"kelurahan" gorm:"foreignKey:Kelurahan_Id"`
	Telp               *string                 `json:"telp" gorm:"size:20"`
	Status             t.StatusBL              `json:"status"`
	NoIdPemilik        *string                 `json:"noIdPemilik" gorm:"size:20"`
	Nik                *string                 `json:"nik" gorm:"size:20"`
}

type RegPemilikWpCreate struct {
	Nama      *string `json:"nama" gorm:"size:50" validate:"required"`
	Alamat    *string `json:"alamat" gorm:"size:50" validate:"required"`
	RtRw      *string `json:"rtRw" gorm:"size:10" validate:"required"`
	Daerah_Id *uint64 `json:"daerah_id" validate:"required"`
	// Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id" validate:"required"`
	Telp         *string `json:"telp" gorm:"size:20" validate:"required"`
	Nik          *string `json:"nik" gorm:"size:20" validate:"required"`
}

type RegPemilikWpUpdate struct {
	Id        uint64  `json:"id" gorm:"primaryKey"`
	Nama      *string `json:"nama" gorm:"size:50"`
	Alamat    *string `json:"alamat" gorm:"size:50"`
	RtRw      *string `json:"rtRw" gorm:"size:10"`
	Daerah_Id *uint64 `json:"daerah_id"`
	// Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Telp         *string `json:"telp" gorm:"size:20"`
	Nik          *string `json:"nik" gorm:"size:20"`
}

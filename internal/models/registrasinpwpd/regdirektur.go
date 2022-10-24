package registrasinpwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
)

type RegDirektur struct {
	Id                 uint64                 `json:"id" gorm:"primaryKey"`
	RegistrasiNpwpd_Id uint64                 `json:"registrasiNpwpd_id"`
	RegistrasiNpwpd    RegistrasiNpwpd        `json:"npwpd,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	Nama               string                 `json:"nama" gorm:"size:50"`
	Alamat             string                 `json:"alamat" gorm:"size:50"`
	Daerah_Id          *uint64                `json:"daerah_id"`
	Daerah             *areadivision.Daerah   `gorm:"foreignKey:Daerah_Id"`
	Kelurahan_Id       uint64                 `json:"kelurahan_id"`
	Kelurahan          areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	Telp               string                 `json:"telp" gorm:"size:20"`
	Nik                string                 `json:"nik" gorm:"size:20"`
}

type RegDirekturCreate struct {
	Nama         string  `json:"nama" gorm:"size:50"`
	Alamat       string  `json:"alamat" gorm:"size:50"`
	Daerah_Id    *uint64 `json:"daerah_id"`
	Kelurahan_Id uint64  `json:"kelurahan_id"`
	Telp         string  `json:"telp" gorm:"size:20"`
	Nik          string  `json:"nik" gorm:"size:20"`
}

type RegDirekturUpdate struct {
	Id           uint64  `json:"id" gorm:"primaryKey"`
	Nama         string  `json:"nama" gorm:"size:50"`
	Alamat       string  `json:"alamat" gorm:"size:50"`
	Daerah_Id    *uint64 `json:"daerah_id"`
	Kelurahan_Id uint64  `json:"kelurahan_id"`
	Telp         string  `json:"telp" gorm:"size:20"`
	Nik          string  `json:"nik" gorm:"size:20"`
}

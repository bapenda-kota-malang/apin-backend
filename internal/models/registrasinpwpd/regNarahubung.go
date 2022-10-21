package registrasinpwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
)

type RegNarahubung struct {
	Id                 uint64               `json:"id" gorm:"primaryKey"`
	RegistrasiNpwpd_Id uint64               `json:"registrasiNpwpd_id"`
	RegistrasiNpwpd    RegistrasiNpwpd      `json:"npwpd,omitempty" gorm:"foreignKey:RegistrasiNpwpd_Id;references:Id"`
	Nama               string               `json:"nama" gorm:"size:50"`
	Alamat             string               `json:"alamat" gorm:"size:50"`
	RtRw               string               `json:"rtRw" gorm:"size:10"`
	Daerah_Id          *uint64              `json:"daerah_id"`
	Daerah             *areadivision.Daerah `gorm:"foreignKey:Daerah_Id"`
	// Kecamatan_Id       uint64                 `json:"kecamatan_id"`
	// Kecamatan          areadivision.Kecamatan `gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id uint64                 `json:"kelurahan_id"`
	Kelurahan    areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	Telp         string                 `json:"telp" gorm:"size:20"`
	Status       t.StatusBL             `json:"status"`
	Nik          string                 `json:"nik" gorm:"size:20"`
}

type RegNarahubungCreate struct {
	Nama      string  `json:"nama" gorm:"size:50"`
	Alamat    string  `json:"alamat" gorm:"size:50"`
	RtRw      string  `json:"rtRw" gorm:"size:10"`
	Daerah_Id *uint64 `json:"daerah_id"`
	// Kecamatan_Id uint64 `json:"kecamatan_id"`
	Kelurahan_Id uint64 `json:"kelurahan_id"`
	Telp         string `json:"telp" gorm:"size:20"`
	Nik          string `json:"nik" gorm:"size:20"`
}

type RegNarahubungUpdate struct {
	Id        uint64  `json:"id" gorm:"primaryKey"`
	Nama      string  `json:"nama" gorm:"size:50"`
	Alamat    string  `json:"alamat" gorm:"size:50"`
	RtRw      string  `json:"rtRw" gorm:"size:10"`
	Daerah_Id *uint64 `json:"daerah_id"`
	// Kecamatan_Id uint64 `json:"kecamatan_id"`
	Kelurahan_Id uint64 `json:"kelurahan_id"`
	Telp         string `json:"telp" gorm:"size:20"`
	Nik          string `json:"nik" gorm:"size:20"`
}

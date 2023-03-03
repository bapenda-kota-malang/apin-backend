package tempatpembayaranspptmasal

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/bankpersepsi"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/banktunggal"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaran"
)

type TempatPembayaranSPPTMasal struct {
	Id                uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode     *string `json:"provinsi_kode" gorm:"type:varchar(5)"`
	Dati2_Kode        *string `json:"dati2_kode" gorm:"type:varchar(10)"`
	Kecamatan_Kode    *string `json:"kecamatan_kode" gorm:"type:varchar(10)"`
	Kelurahan_Kode    *string `json:"kelurahan_kode" gorm:"type:varchar(10)"`
	Tahun             *string `json:"tahun" gorm:"type:varchar(10)"`
	Kanwil_Kode       *string `json:"kanwil_kode" gorm:"type:varchar(10)"`
	KPPBB_Kode        *string `json:"kppbb_kode" gorm:"type:varchar(10)"`
	BankTunggal_Kode  *string `json:"banktunggal_kode" gorm:"type:varchar(10)"`
	BankPersepsi_Kode *string `json:"bankpersepsi_kode" gorm:"type:varchar(10)"`
	TP_Kode           *string `json:"tp_kode" gorm:"type:varchar(10)"`
	gh.DateModel

	// Relations
	Kelurahan        *areadivision.Kelurahan            `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
	TempatPembayaran *tempatpembayaran.TempatPembayaran `json:"tempat_pembayaran,omitempty" gorm:"foreignKey:TP_Kode;references:Tp_Id"`
	BankTunggal      *banktunggal.BankTunggal           `json:"bank_tunggal,omitempty" gorm:"foreignKey:BankTunggal_Kode;references:Kode"`
	BankPersepsi     *bankpersepsi.BankPersepsi         `json:"bank_persepsi,omitempty" gorm:"foreignKey:BankPersepsi_Kode;references:Kode"`
}

// table name
func (TempatPembayaranSPPTMasal) TableName() string {
	return "TempatPembayaranSPPTMasal"
}

type CreateDto struct {
	Provinsi_Kode     *string `json:"provinsi_kode"`
	Dati2_Kode        *string `json:"dati2_kode"`
	Kecamatan_Kode    *string `json:"kecamatan_kode"`
	Kelurahan_Kode    *string `json:"kelurahan_kode"`
	Tahun             *string `json:"tahun"`
	Kanwil_Kode       *string `json:"kanwil_kode"`
	KPPBB_Kode        *string `json:"kppbb_kode"`
	BankTunggal_Kode  *string `json:"banktunggal_kode"`
	BankPersepsi_Kode *string `json:"bankpersepsi_kode"`
	TP_Kode           *string `json:"tp_kode"`
}

type UpdateDto struct {
	Id                *uint64 `json:"id"`
	Provinsi_Kode     *string `json:"provinsi_kode"`
	Dati2_Kode        *string `json:"dati2_kode"`
	Kecamatan_Kode    *string `json:"kecamatan_kode"`
	Kelurahan_Kode    *string `json:"kelurahan_kode"`
	Tahun             *string `json:"tahun"`
	Kanwil_Kode       *string `json:"kanwil_kode"`
	KPPBB_Kode        *string `json:"kppbb_kode"`
	BankTunggal_Kode  *string `json:"banktunggal_kode"`
	BankPersepsi_Kode *string `json:"bankpersepsi_kode"`
	TP_Kode           *string `json:"tp_kode"`
}

type FilterDto struct {
	Provinsi_Kode     *string `json:"provinsi_kode"`
	Dati2_Kode        *string `json:"dati2_kode"`
	Kecamatan_Kode    *string `json:"kecamatan_kode"`
	Kelurahan_Kode    *string `json:"kelurahan_kode"`
	Tahun             *string `json:"tahun"`
	Kanwil_Kode       *string `json:"kanwil_kode"`
	KPPBB_Kode        *string `json:"kppbb_kode"`
	BankTunggal_Kode  *string `json:"banktunggal_kode"`
	BankPersepsi_Kode *string `json:"bankpersepsi_kode"`
	TP_Kode           *string `json:"tp_kode"`
	Page              int     `json:"page"`
	PageSize          int64   `json:"page_size"`
}

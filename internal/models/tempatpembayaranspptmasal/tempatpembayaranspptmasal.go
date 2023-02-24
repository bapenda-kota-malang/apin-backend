package tempatpembayaranspptmasal

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
)

type TempatPembayaranSPPTMasal struct {
	Id                   uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode        *string `json:"provinsi_kode"`
	Dati_Kode            *string `json:"dati_kode"`
	Kecamatan_Kode       *string `json:"kecamatan_kode"`
	Kelurahan_Kode       *string `json:"kelurahan_kode"`
	Tahun                *string `json:"tahun"`
	Kanwil_Kode          *string `json:"kanwil_kode"`
	KPPBB_Kode           *string `json:"kppbb_kode"`
	BankTunggal_Kode     *string `json:"banktunggal_kode"`
	BankPersepsi_Kode    *string `json:"bankpersepsi_kode"`
	TP_Kode              *string `json:"tp_kode"`
	NamaTempatPembayaran *string `json:"nama_tempat_pembayaran"`
	gh.DateModel

	// Relations
	Kelurahan *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
}

// table name
func (TempatPembayaranSPPTMasal) TableName() string {
	return "TempatPembayaranSPPTMasal"
}

type ListDto struct {
	Id                   *uint64 `json:"id"`
	Kelurahan_Kode       *string `json:"kelurahan_kode"`
	NamaKelurahan        *string `json:"nama_kelurahan"`
	BankTunggal_Kode     *string `json:"banktunggal_kode"`
	BankPersepsi_Kode    *string `json:"bankpersepsi_kode"`
	TP_Kode              *string `json:"tp_kode"`
	NamaTempatPembayaran *string `json:"nama_tempat_pembayaran"`
}

type CreateDto struct {
	Provinsi_Kode        *string `json:"provinsi_kode"`
	Dati_Kode            *string `json:"dati_kode"`
	Kecamatan_Kode       *string `json:"kecamatan_kode"`
	Kelurahan_Kode       *string `json:"kelurahan_kode"`
	Tahun                *string `json:"tahun"`
	Kanwil_Kode          *string `json:"kanwil_kode"`
	KPPBB_Kode           *string `json:"kppbb_kode"`
	BankTunggal_Kode     *string `json:"banktunggal_kode"`
	BankPersepsi_Kode    *string `json:"bankpersepsi_kode"`
	TP_Kode              *string `json:"tp_kode"`
	NamaTempatPembayaran *string `json:"nama_tempat_pembayaran"`
}

type UpdateDto struct {
	Id                   *uint64 `json:"id"`
	Provinsi_Kode        *string `json:"provinsi_kode"`
	Dati_Kode            *string `json:"dati_kode"`
	Kecamatan_Kode       *string `json:"kecamatan_kode"`
	Kelurahan_Kode       *string `json:"kelurahan_kode"`
	Tahun                *string `json:"tahun"`
	Kanwil_Kode          *string `json:"kanwil_kode"`
	KPPBB_Kode           *string `json:"kppbb_kode"`
	BankTunggal_Kode     *string `json:"banktunggal_kode"`
	BankPersepsi_Kode    *string `json:"bankpersepsi_kode"`
	TP_Kode              *string `json:"tp_kode"`
	NamaTempatPembayaran *string `json:"nama_tempat_pembayaran"`
}

type FilterDto struct {
	Provinsi_Kode        *string `json:"provinsi_kode"`
	Dati_Kode            *string `json:"dati_kode"`
	Kecamatan_Kode       *string `json:"kecamatan_kode"`
	Kelurahan_Kode       *string `json:"kelurahan_kode"`
	Tahun                *string `json:"tahun"`
	Kanwil_Kode          *string `json:"kanwil_kode"`
	KPPBB_Kode           *string `json:"kppbb_kode"`
	BankTunggal_Kode     *string `json:"banktunggal_kode"`
	BankPersepsi_Kode    *string `json:"bankpersepsi_kode"`
	TP_Kode              *string `json:"tp_kode"`
	NamaTempatPembayaran *string `json:"nama_tempat_pembayaran"`
	Page                 int     `json:"page"`
	PageSize             int64   `json:"page_size"`
}

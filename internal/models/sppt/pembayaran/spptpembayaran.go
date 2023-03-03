package pembayaran

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SpptPembayaran struct {
	Id                  uint64           `json:"id" gorm:"primaryKey"`
	Provinsi_Kd         string           `json:"provinsi_kd" gorm:"type:varchar(2);not null"`
	Daerah_Kd           string           `json:"daerah_kd" gorm:"type:varchar(2);not null"`
	Kecamatan_Kd        string           `json:"kecamatan_kd" gorm:"type:varchar(3);not null"`
	Kelurahan_Kd        string           `json:"Kelurahan_kd" gorm:"type:varchar(3);not null"`
	Blok_Kd             string           `json:"blok_kd" gorm:"type:varchar(3);not null"`
	NoUrut_Kd           string           `json:"noUrut_kd" gorm:"type:varchar(4);not null"`
	JenisOp_Kd          string           `json:"jenisOp_kd" gorm:"type:varchar(1);not null"`
	TahunPajakSppt      string           `json:"tahunPajakSppt" gorm:"type:varchar(4);not null"`
	PembayaranSpptKe    uint8            `json:"pembayaranSpptKe" gorm:"not null"`
	KanwilBank_Kd       *string          `json:"kanwiBank_kd" gorm:"type:varchar(2);not null"`
	KppbbBank_Kd        *string          `json:"kppbbBank_kd" gorm:"type:varchar(2);not null"`
	BankTunggal_Kd      *string          `json:"bankTunggal_kd" gorm:"type:varchar(2);not null"`
	BankPersepsi_Kd     *string          `json:"bankPersepsi_kd" gorm:"type:varchar(2);not null"`
	TP_Kd               *string          `json:"tp_kd" gorm:"type:varchar(2);not null"`
	Denda_Sppt          int64            `json:"denda_sppt" gorm:"size:12"`
	JumlahSpptYgDibayar int64            `json:"JumlahSpptYgDibayar" gorm:"size:15"`
	TglPembayaranSppt   datatypes.Date   `json:"tglPembayaranSppt" gorm:"not null"`
	TglRekamBayarSppt   time.Time        `json:"tglRekamBayarSppt" gorm:"not null;index"`
	NipRekamBayarSppt   string           `json:"nipRekamBayarSppt" gorm:"size:30"`
	Pegawai             *pegawai.Pegawai `json:"pegawai,omitempty" gorm:"foreignKey:NipRekamBayarSppt;references:Nip"`
	gormhelper.DateModel
}

type CreateDto struct {
	Provinsi_Kd         string          `json:"provinsi_kd" validate:"required"`
	Daerah_Kd           string          `json:"daerah_kd" validate:"required"`
	Kecamatan_Kd        string          `json:"kecamatan_kd" validate:"required"`
	Kelurahan_Kd        string          `json:"Kelurahan_kd" validate:"required"`
	Blok_Kd             string          `json:"blok_kd" validate:"required"`
	NoUrut_Kd           string          `json:"noUrut_kd" validate:"required"`
	JenisOp_Kd          string          `json:"jenisOp_kd" validate:"required"`
	TahunPajakSppt      string          `json:"tahunPajakSppt" validate:"required"`
	Denda_Sppt          *int64          `json:"denda_sppt" validate:"required"`
	JumlahSpptYgDibayar *int64          `json:"JumlahSpptYgDibayar" validate:"required"`
	TglPembayaranSppt   *datatypes.Date `json:"tglPembayaranSppt" validate:"required"`
}

type FilterDto struct {
	Provinsi_Kd         *string         `json:"provinsi_kd"`
	Daerah_Kd           *string         `json:"daerah_kd"`
	Kecamatan_Kd        *string         `json:"kecamatan_kd"`
	Kelurahan_Kd        *string         `json:"Kelurahan_kd"`
	Blok_Kd             *string         `json:"blok_kd"`
	NoUrut_Kd           *string         `json:"noUrut_kd"`
	JenisOp_Kd          *string         `json:"jenisOp_kd"`
	TahunPajakSppt      *string         `json:"tahunPajakSppt"`
	PembayaranSpptKe    *uint8          `json:"pembayaranSpptKe"`
	KanwilBank_Kd       *string         `json:"kanwiBank_kd"`
	KppbbBank_Kd        *string         `json:"kppbbBank_kd"`
	BankTunggal_Kd      *string         `json:"bankTunggal_kd"`
	BankPersepsi_Kd     *string         `json:"bankPersepsi_kd"`
	TP_Kd               *string         `json:"tp_kd"`
	Denda_Sppt          *int64          `json:"denda_sppt"`
	JumlahSpptYgDibayar *int64          `json:"JumlahSpptYgDibayar"`
	TglPembayaranSppt   *datatypes.Date `json:"tglPembayaranSppt"`
	TglRekamBayarSppt   *time.Time      `json:"tglRekamBayarSppt"`
	NipRekamBayarSppt   *string         `json:"nipRekamBayarSppt"`
	Page                int             `json:"page"`
	PageSize            int             `json:"page_size"`
}

type UpdateDto struct {
	// not implemented according to business flow
}

package tandaterima

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SpptTandaTerima struct {
	Id                 uint64           `json:"id" gorm:"primaryKey"`
	Provinsi_Kd        string           `json:"provinsi_kd" gorm:"type:varchar(2);not null"`
	Daerah_Kd          string           `json:"daerah_kd" gorm:"type:varchar(2);not null"`
	Kecamatan_Kd       string           `json:"kecamatan_kd" gorm:"type:varchar(3);not null"`
	Kelurahan_Kd       string           `json:"Kelurahan_kd" gorm:"type:varchar(3);not null"`
	Blok_Kd            string           `json:"blok_kd" gorm:"type:varchar(3);not null"`
	NoUrut_Kd          string           `json:"noUrut_kd" gorm:"type:varchar(4);not null"`
	JenisOp_Kd         string           `json:"jenisOp_kd" gorm:"type:varchar(1);not null"`
	TahunPajakSppt     string           `json:"tahunPajakSppt" gorm:"type:varchar(4);not null"`
	NamaYgMenerimaSppt string           `json:"namaYgMenerimaSppt" gorm:"type:varchar(30);not null"`
	TglTerimaWpSppt    datatypes.Date   `json:"tglTerimaWpSppt" gorm:"not null"`
	TglRekamTtrSppt    time.Time        `json:"tglRekamTtrSppt" gorm:"not null;index"`
	NipRekamTtrSppt    string           `json:"nipRekamTtrSppt" gorm:"size:30"`
	Pegawai            *pegawai.Pegawai `json:"pegawai,omitempty" gorm:"foreignKey:NipRekamTtrSppt;references:Nip"`
	gormhelper.DateModel
}

type CreateDto struct {
	Provinsi_Kd        string          `json:"provinsi_kd" validate:"required"`
	Daerah_Kd          string          `json:"daerah_kd" validate:"required"`
	Kecamatan_Kd       string          `json:"kecamatan_kd" validate:"required"`
	Kelurahan_Kd       string          `json:"Kelurahan_kd" validate:"required"`
	Blok_Kd            string          `json:"blok_kd" validate:"required"`
	NoUrut_Kd          string          `json:"noUrut_kd" validate:"required"`
	JenisOp_Kd         string          `json:"jenisOp_kd" validate:"required"`
	TahunPajakSppt     string          `json:"tahunPajakSppt" validate:"required"`
	NamaYgMenerimaSppt string          `json:"namaYgMenerimaSppt" validate:"required"`
	TglTerimaWpSppt    *datatypes.Date `json:"tglTerimaWpSppt" validate:"required"`
}

type FilterDto struct {
	Provinsi_Kd        *string         `json:"provinsi_kd"`
	Daerah_Kd          *string         `json:"daerah_kd"`
	Kecamatan_Kd       *string         `json:"kecamatan_kd"`
	Kelurahan_Kd       *string         `json:"Kelurahan_kd"`
	Blok_Kd            *string         `json:"blok_kd"`
	NoUrut_Kd          *string         `json:"noUrut_kd"`
	JenisOp_Kd         *string         `json:"jenisOp_kd"`
	TahunPajakSppt     *string         `json:"tahunPajakSppt"`
	NamaYgMenerimaSppt *string         `json:"namaYgMenerimaSppt"`
	TglTerimaWpSppt    *datatypes.Date `json:"tglTerimaWpSppt"`
	TglRekamTtrSppt    *time.Time      `json:"tglRekamTtrSppt"`
	NipRekamTtrSppt    *string         `json:"nipRekamTtrSppt"`
	Page               int             `json:"page"`
	PageSize           int             `json:"page_size"`
}

type UpdateDto struct {
	NamaYgMenerimaSppt *string         `json:"namaYgMenerimaSppt"`
	TglTerimaWpSppt    *datatypes.Date `json:"tglTerimaWpSppt"`
	// TglRekamTtrSppt    *time.Time      `json:"tglRekamTtrSppt"`
	// NipRekamTtrSppt    *string         `json:"nipRekamTtrSppt"`
}

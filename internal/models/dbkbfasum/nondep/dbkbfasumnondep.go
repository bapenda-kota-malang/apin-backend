package nondep

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/dbkbfasum"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DbkbFasumNonDep struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  string  `json:"provinsi_kode" gorm:"type:char(2)"`
	Daerah_Kode    string  `json:"daerah_kode" gorm:"type:char(2)"`
	Tahun          string  `json:"tahun" gorm:"type:char(4)"`
	Fasilitas_Kode string  `json:"fasilitas_kode" gorm:"type:char(2)"`
	Nilai          float64 `json:"nilai" gorm:"type:numeric(10,2)"`
	gormhelper.DateModel
	DbkbFasum *dbkbfasum.DbkbFasum `json:"dbkbFasum,omitempty" gorm:"foreignKey:Fasilitas_Kode;references:Kode"`
}

type CreateDto struct {
	Provinsi_Kode  string   `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode    string   `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Tahun          string   `json:"tahun" validate:"required;minLength=4;maxLength=4"`
	Fasilitas_Kode string   `json:"fasilitas_kode" validate:"required;minLength=2;maxLength=2"`
	Nilai          *float64 `json:"nilai" validate:"required"`
}

type FilterDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	Fasilitas_Kode *string  `json:"fasilitas_kode"`
	Tahun          *string  `json:"tahun"`
	Nilai          *float64 `json:"nilai"`
	Page           int      `json:"page"`
	PageSize       int      `json:"page_size"`
}

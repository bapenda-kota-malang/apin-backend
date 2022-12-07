package datanir

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DataNir struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  string  `json:"provinsi_kode" gorm:"type:char(2)"`
	Daerah_Kode    string  `json:"daerah_kode" gorm:"type:char(2)"`
	Kecamatan_Kode string  `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kelurahan_Kode string  `json:"kelurahan_kode" gorm:"type:char(3)"`
	Tahun          string  `json:"tahun" gorm:"type:char(4)"`
	NomerDokumen   string  `json:"nomerDokumen" gorm:"type:char(11)"`
	Kpbb_Id        *string `json:"kpbb_id" gorm:"type:char(2)"`
	Kanwil_Id      *string `json:"kanwil_id" gorm:"type:char(2)"`
	JenisDokumen   *string `json:"jenisDokumen" gorm:"type:char(1)"`
	Znt_Kode       string  `json:"znt_kode" gorm:"type:char(2)"`
	Nir            float64 `json:"nir" gorm:"type:decimal(12,2)"`
	gormhelper.DateModel
}

type CreateDto struct {
	Provinsi_Kode  string   `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode    string   `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Kecamatan_Kode string   `json:"kecamatan_kode" validate:"required;minLength=3;maxLength=3"`
	Kelurahan_Kode string   `json:"kelurahan_kode" validate:"required;minLength=3;maxLength=3"`
	Tahun          string   `json:"tahun" validate:"required;minLength=4;maxLength=4"`
	NomerDokumen   string   `json:"nomerDokumen" validate:"required;minLength=11;maxLength=11"`
	Kpbb_Id        *string  `json:"kpbb_id" validate:"required;minLength=2;maxLength=2"`
	Kanwil_Id      *string  `json:"kanwil_id" validate:"required;minLength=2;maxLength=2"`
	JenisDokumen   *string  `json:"jenisDokumen" validate:"required;minLength=1;maxLength=1"`
	Znt_Kode       string   `json:"znt_kode" validate:"required;minLength=2;maxLength=2"`
	Nir            *float64 `json:"nir" validate:"required"`
}

type DataDto struct {
	Znt_Kode string   `json:"znt_kode" validate:"required;minLength=2;maxLength=2"`
	Nir      *float64 `json:"nir" validate:"required"`
}

type CreateBulkDto struct {
	Provinsi_Kode  string    `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode    string    `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Kecamatan_Kode string    `json:"kecamatan_kode" validate:"required;minLength=3;maxLength=3"`
	Kelurahan_Kode string    `json:"kelurahan_kode" validate:"required;minLength=3;maxLength=3"`
	Tahun          string    `json:"tahun" validate:"required;minLength=4;maxLength=4"`
	NomerDokumen   string    `json:"nomerDokumen" validate:"required;minLength=11;maxLength=11"`
	Kpbb_Id        *string   `json:"kpbb_id" validate:"required;minLength=2;maxLength=2"`
	Kanwil_Id      *string   `json:"kanwil_id" validate:"required;minLength=2;maxLength=2"`
	JenisDokumen   *string   `json:"jenisDokumen" validate:"required;minLength=1;maxLength=1"`
	Datas          []DataDto `json:"datas" validate:"required"`
}

type FilterDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	Kecamatan_Kode *string  `json:"kecamatan_kode"`
	Kelurahan_Kode *string  `json:"kelurahan_kode"`
	Tahun          *string  `json:"tahun"`
	NomerDokumen   *string  `json:"nomerDokumen"`
	Kpbb_Id        *string  `json:"kpbb_id"`
	Kanwil_Id      *string  `json:"kanwil_id"`
	JenisDokumen   *string  `json:"jenisDokumen"`
	Znt_Kode       *string  `json:"znt_kode"`
	Nir            *float64 `json:"nir"`
	Page           int      `json:"page"`
	PageSize       int      `json:"page_size"`
	// NoPagination   bool    `json:"no_pagination"`
}

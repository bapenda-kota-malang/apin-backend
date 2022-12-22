package datapetaznt

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DataPetaZnt struct {
	Id             uint64 `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  string `json:"provinsi_kode" gorm:"type:char(2)"`
	Daerah_Kode    string `json:"daerah_kode" gorm:"type:char(2)"`
	Kecamatan_Kode string `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kelurahan_Kode string `json:"kelurahan_kode" gorm:"type:char(3)"`
	Blok_Kode      string `json:"blok_kode" gorm:"type:char(3)"`
	Znt_Kode       string `json:"znt_kode" gorm:"type:char(2)"`
	gormhelper.DateModel
}

type CreateDto struct {
	Provinsi_Kode  string `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode    string `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Kecamatan_Kode string `json:"kecamatan_kode" validate:"required;minLength=3;maxLength=3"`
	Kelurahan_Kode string `json:"kelurahan_kode" validate:"required;minLength=3;maxLength=3"`
	Blok_Kode      string `json:"blok_kode" validate:"required;minLength=3;maxLength=3"`
	Znt_Kode       string `json:"znt_kode" validate:"required;minLength=2;maxLength=2"`
}

type DataDto struct {
	Id        *int   `json:"id"`
	Blok_Kode string `json:"blok_kode" validate:"required;minLength=3;maxLength=3"`
	Znt_Kode  string `json:"znt_kode" validate:"required;minLength=2;maxLength=2"`
}

type CreateBulkDto struct {
	Provinsi_Kode  string    `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode    string    `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Kecamatan_Kode string    `json:"kecamatan_kode" validate:"required;minLength=3;maxLength=3"`
	Kelurahan_Kode string    `json:"kelurahan_kode" validate:"required;minLength=3;maxLength=3"`
	Datas          []DataDto `json:"datas" validate:"required"`
}

type FilterDto struct {
	Provinsi_Kode  *string `json:"provinsi_kode"`
	Daerah_Kode    *string `json:"daerah_kode"`
	Kecamatan_Kode *string `json:"kecamatan_kode"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Blok_Kode      *string `json:"blok_kode"`
	Znt_Kode       *string `json:"znt_kode"`
	Page           int     `json:"page"`
	PageSize       int     `json:"page_size"`
	NoPagination   bool    `json:"no_pagination"`
}

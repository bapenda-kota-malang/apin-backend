package hargareferensi

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type HargaReferensi struct {
	Id             uint64        `json:"id" gorm:"primaryKey"`
	Alamat         *string       `json:"alamat"`
	Kecamatan_Kode *string       `json:"kecamatan_kode" gorm:"type:varchar(15)"`
	Kelurahan_Kode *string       `json:"kelurahan_kode" gorm:"type:varchar(15)"`
	Harga          *int          `json:"harga"`
	Kelurahan      *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
	Kecamatan      *ad.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Kode;references:Kode"`
	gh.DateModel
}

type CreateDto struct {
	Alamat         *string `json:"alamat"`
	Kecamatan_Kode *string `json:"kecamatan_kode"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Harga          *int    `json:"harga"`
}

type UpdateDto struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Alamat         *string `json:"alamat"`
	Kecamatan_Kode *string `json:"kecamatan_kode"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Harga          *int    `json:"harga"`
}

type FilterDto struct {
	Alamat         *string `json:"alamat"`
	Kecamatan_Kode *string `json:"kecamatan_kode"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Harga          *int    `json:"harga"`
	Page           int     `json:"page"`
	PageSize       int     `json:"page_size"`
}

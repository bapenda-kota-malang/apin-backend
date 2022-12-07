package hargareferensi

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type HargaReferensi struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Alamat       *string       `json:"alamat"`
	Kecamatan_Id *uint64       `json:"kecamatan_id"`
	Kelurahan_Id *uint64       `json:"kelurahan_id"`
	Harga        *int          `json:"harga"`
	Kecamatan    *ad.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan    *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
}

type CreateDto struct {
	Alamat       *string `json:"alamat"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Harga        *int    `json:"harga"`
}

type UpdateDto struct {
	Id           uint64  `json:"id" gorm:"primaryKey"`
	Alamat       *string `json:"alamat"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Harga        *int    `json:"harga"`
}

type FilterDto struct {
	Alamat       *string `json:"alamat"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	Harga        *int    `json:"harga"`
	Page         int     `json:"page"`
	PageSize     int     `json:"page_size"`
}

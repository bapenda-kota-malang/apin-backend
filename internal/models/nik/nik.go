package nik

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Nik struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Nik          *string       `json:"nik"`
	Nama         *string       `json:"nama"`
	Alamat       *string       `json:"alamat"`
	Provinsi_Id  *uint64       `json:"provinsi_id"`
	Daerah_Id    *uint64       `json:"daerah_id"`
	Kecamatan_Id *uint64       `json:"kecamatan_id"`
	Kelurahan_Id *uint64       `json:"kelurahan_id"`
	RtRw         *string       `json:"rtRw"`
	KodePos      *string       `json:"kodePos"`
	Provinsi     *ad.Provinsi  `json:"provinsi,omitempty" gorm:"foreignKey:Provinsi_Id"`
	Daerah       *ad.Daerah    `json:"daerah,omitempty" gorm:"foreignKey:Daerah_Id"`
	Kecamatan    *ad.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan    *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
}

type CreateDto struct {
	Nik          *string `json:"nik" validate:"required;nik"`
	Nama         *string `json:"nama"`
	Alamat       *string `json:"alamat"`
	Provinsi_Id  *uint64 `json:"provinsi_id"`
	Daerah_Id    *uint64 `json:"daerah_id"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	RtRw         *string `json:"rtRw"`
	KodePos      *string `json:"kodePos"`
}

type UpdateDto struct {
	Id           *uint64 `json:"id"`
	Nik          *string `json:"nik" validate:"required;nik"`
	Nama         *string `json:"nama"`
	Alamat       *string `json:"alamat"`
	Provinsi_Id  *uint64 `json:"provinsi_id"`
	Daerah_Id    *uint64 `json:"daerah_id"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	RtRw         *string `json:"rtRw"`
	KodePos      *string `json:"kodePos"`
}

type FilterDto struct {
	Nik          *string `json:"nik"`
	Nama         *string `json:"nama" reffunc:"LOWER"`
	Nama_Opt     *string `json:"nama_opt"`
	Alamat       *string `json:"alamat"`
	Provinsi_Id  *uint64 `json:"provinsi_id"`
	Daerah_Id    *uint64 `json:"daerah_id"`
	Kecamatan_Id *uint64 `json:"kecamatan_id"`
	Kelurahan_Id *uint64 `json:"kelurahan_id"`
	RtRw         *string `json:"rtRw"`
	KodePos      *string `json:"kodePos"`
	Page         int     `json:"page"`
	PageSize     int     `json:"page_size"`
}

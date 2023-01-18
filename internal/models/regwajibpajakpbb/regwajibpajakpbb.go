package regwajibpajakpbb

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegWajibPajakPbb struct {
	Id             uint64        `json:"id" gorm:"primarykey"`
	Nik            *string       `json:"nik" gorm:"type:char(30)"`
	Nama           *string       `json:"nama" gorm:"type:varchar(30)"`
	Jalan          *string       `json:"jalan" gorm:"type:varchar(30)"`
	BlokKavNo      *string       `json:"blokKavNo" gorm:"type:varchar(15)"`
	Rw             *string       `json:"rw" gorm:"type:char(2)"`
	Rt             *string       `json:"rt" gorm:"type:char(2)"`
	Kelurahan_Kode *string       `json:"kelurahan_kode" gorm:"type:char(10)"`
	Kelurahan      *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
	Daerah_Kode    *string       `json:"daerah_kode" gorm:"type:char(4)"`
	Daerah         *ad.Daerah    `json:"kota,omitempty" gorm:"foreignKey:Daerah_Kode;references:Kode"`
	KodePos        *string       `json:"kodePos" gorm:"type:varchar(5)"`
	Telp           *string       `json:"telp" gorm:"type:varchar(20)"`
	Npwp           *string       `json:"npwp" gorm:"type:varchar(20)"`
	Pekerjaan      *string       `json:"pekerjaan" gorm:"type:varchar(100)"`
	User_ID        *uint64       `json:"user_ID"`

	gh.DateModel
}

type CreateDto struct {
	Nik            *string `json:"nik"`
	Nama           *string `json:"nama"`
	Jalan          *string `json:"jalan"`
	BlokKavNo      *string `json:"blokKavNo"`
	Rw             *string `json:"rw"`
	Rt             *string `json:"rt"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Daerah_Kode    *string `json:"daerah_kode"`
	KodePos        *string `json:"kodePos"`
	Telp           *string `json:"telp"`
	Npwp           *string `json:"npwp"`
	Pekerjaan      *string `json:"pekerjaan"`
	User_ID        *uint64 `json:"user_ID"`
}

type UpdateDto struct {
	Nik            *string `json:"nik"`
	Nama           *string `json:"nama"`
	Jalan          *string `json:"jalan"`
	BlokKavNo      *string `json:"blokKavNo"`
	Rw             *string `json:"rw"`
	Rt             *string `json:"rt"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Daerah_Kode    *string `json:"daerah_kode"`
	KodePos        *string `json:"kodePos"`
	Telp           *string `json:"telp"`
	Npwp           *string `json:"npwp"`
	Pekerjaan      *string `json:"pekerjaan"`
	User_ID        *uint64 `json:"user_ID"`
}

type FilterDto struct {
	Nik            *string `json:"nik"`
	Nama           *string `json:"nama"`
	Jalan          *string `json:"jalan"`
	BlokKavNo      *string `json:"blokKavNo"`
	Rw             *string `json:"rw"`
	Rt             *string `json:"rt"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Daerah_Kode    *string `json:"daerah_kode"`
	KodePos        *string `json:"kodePos"`
	Telp           *string `json:"telp"`
	Npwp           *string `json:"npwp"`
	Pekerjaan      *string `json:"pekerjaan"`
	Page           int     `json:"page"`
	PageSize       int     `json:"page_size"`
	User_ID        *uint64 `json:"user_ID"`
}

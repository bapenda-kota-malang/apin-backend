package wajibpajakpbb

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type WajibPajakPbb struct {
	Id             uint64        `json:"id" gorm:"primarykey"`
	Nik            string        `json:"nik" gorm:"type:char(30)"`
	Nama           string        `json:"nama" gorm:"type:varchar2(30)"`
	Jalan          string        `json:"jalan" gorm:"type:varchar2(30)"`
	BlokKavNo      string        `json:"blokKavNo" gorm:"type:varchar2(15)"`
	Rw             string        `json:"rw" gorm:"type:char(2)"`
	Rt             string        `json:"rt" gorm:"type:char(2)"`
	Kelurahan_Kode string        `json:"kelurahan_kode" gorm:"type:char(3)"`
	Kelurahan      *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
	Daerah_Kode    string        `json:"daerah_kode" gorm:"type:char(3)"`
	Daerah         *ad.Daerah    `json:"kota,omitempty" gorm:"foreignKey:Daerah_Kode;references:Kode"`
	KodePos        string        `json:"kodePos" gorm:"type:varchar2(5)"`
	Telp           string        `json:"telp" gorm:"type:varchar2(20)"`
	Npwp           string        `json:"npwp" gorm:"type:varchar(20)"`
	Pekerjaan      string        `json:"pekerjaan" gorm:"type:varchar(100)"`
	gh.DateModel
}

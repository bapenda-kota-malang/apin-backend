package wajibpajakpbb

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type WajibPajakPBB struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	Nik          string `json:"nik" gorm:"size:255;unique"`
	Nama         string `json:"nama" gorm:"size:100"`
	Alamat       string `json:"alamat"`
	BlokKavNo    string `json:"blokKavNo" gorm:"size:15"`
	Rw           string `json:"rw" gorm:"size:3"`
	Rt           string `json:"rt" gorm:"size:2"`
	Kelurahan_Id uint   `json:"kelurahan_id"`
	// Kelurahan    adm.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	// Kecamatan_Id uint `json:"kecamatan_id"`
	// Kecamatan    adm.Kecamatan `gorm:"foreignKey:Kecamatan_Id"`
	Kota_id uint `json:"kota_id"`
	// Daerah       adm.Daerah    `gorm:"foreignKey:Kota_Id"`
	// Provinsi_Id uint   `json:"provinsi_id"`
	// Provinsi     adm.Provinsi  `gorm:"foreignKey:Provinsi_Id"`
	KodePos   string `json:"kodePos" gorm:"size:50"`
	Telp      string `json:"telp" gorm:"size:20"`
	Npwp      string `json:"npwp" gorm:"size:20"`
	Pekerjaan string `json:"pekerjaan" gorm:"size:100"`
	gormhelper.DateModel
}

type RequestDto struct {
	Nik       string `json:"nik" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Alamat    string `json:"alamat" validate:"required"`
	BlokKavNo string `json:"blokKavNo"`
	Rt        string `json:"rt" validate:"required"`
	Rw        string `json:"rw" validate:"required"`
	// Provinsi_Id  uint      `json:"provinsi_id" validate:"required,min=1"`
	Kota_id uint `json:"kota_id" validate:"required,min=1"`
	// Kecamatan_Id uint      `json:"kecamatan_id" validate:"required,min=1"`
	Kelurahan_Id uint   `json:"kelurahan_id" validate:"required,min=1"`
	KodePos      string `json:"kodePos" validate:"required"`
	Telp         string `json:"telp" validate:"required"`
	Npwp         string `json:"npwp" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	// Email        string
}

type FilterDto struct {
	Nik       *string `json:"nik"`
	Nama      *string `json:"nama"`
	Alamat    *string `json:"alamat"`
	BlokKavNo *string `json:"blokKavNo"`
	Rt        *string `json:"rt"`
	Rw        *string `json:"rw"`
	// Provinsi_Id  *uint   `json:"provinsi_id"`
	Kota_id *uint `json:"kota_id"`
	// Kecamatan_Id *uint   `json:"kecamatan_id"`
	Kelurahan_Id *uint   `json:"kelurahan_id"`
	KodePos      *string `json:"kodePos"`
	Pekerjaan    *string `json:"pekerjaan"`
	Telp         *string `json:"telp"`
	Npwp         *string `json:"npwp"`
	// fixed fields
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

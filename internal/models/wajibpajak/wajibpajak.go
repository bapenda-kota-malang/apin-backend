package wajibpajak

import (
	"time"

	// adm "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type WajibPajak struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Nik         string `json:"nik" gorm:"size:255;unique"`
	Nama        string `json:"nama" gorm:"size:100"`
	Alamat      string `json:"alamat"`
	Provinsi_Id uint   `json:"provinsi_id"`
	// Provinsi     adm.Provinsi  `gorm:"foreignKey:Provinsi_Id"`
	Kota_id uint `json:"kota_id"`
	// Daerah       adm.Daerah    `gorm:"foreignKey:Kota_Id"`
	Kecamatan_Id uint `json:"kecamatan_id"`
	// Kecamatan    adm.Kecamatan `gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan_Id uint `json:"kelurahan_id"`
	// Kelurahan    adm.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
	Telp    string `json:"telp" gorm:"size:20"`
	KodePos string `json:"kodePos" gorm:"size:50"`
	// Email        string    `json:"email" gorm:"size:100;unique"`
	RtRw         string    `json:"rtRw" gorm:"size:10"`
	FotoKtp      string    `json:"fotoKtp" gorm:"size:255"`
	Status       uint8     `json:"status"`
	TempatLahir  string    `json:"tempatLahir" gorm:"size:50"`
	TanggalLahir time.Time `json:"tanggalLahir"`
	Gender       uint8     `json:"gender"`
	Pekerjaan    string    `json:"pekerjaan" gorm:"size:100"`
	gormhelper.DateModel
}

type CreateDto struct {
	Nik          string    `json:"nik" validate:"required"`
	Nama         string    `json:"nama" validate:"required"`
	Alamat       string    `json:"alamat" validate:"required"`
	Provinsi_Id  uint      `json:"provinsi_id" validate:"required,min=1"`
	Kota_id      uint      `json:"kota_id" validate:"required,min=1"`
	Kecamatan_Id uint      `json:"kecamatan_id" validate:"required,min=1"`
	Kelurahan_Id uint      `json:"kelurahan_id" validate:"required,min=1"`
	Telp         string    `json:"telp" validate:"required"`
	KodePos      string    `json:"kodePos" validate:"required"`
	RtRw         string    `json:"rtRw" validate:"required"`
	FotoKtp      string    `json:"fotoKtp" validate:"required"`
	TempatLahir  string    `json:"tempatLahir" validate:"required"`
	TanggalLahir time.Time `json:"tanggalLahir" validate:"required"`
	Gender       uint8     `json:"gender" validate:"required,min=1"`
	Pekerjaan    string    `json:"pekerjaan" validate:"required"`
	// Email        string
}
type RegistrasiWajibPajak struct {
	// Wajib Pajak
	WajibPajak CreateDto `json:"wajibPajak"`
	// User
	User user.RegisterDto `json:"user"`
}

type UpdateDto struct {
	Nama         *string    `json:"nama"`
	Alamat       *string    `json:"alamat"`
	Provinsi_Id  *uint      `json:"provinsi_id"`
	Kota_id      *uint      `json:"kota_id"`
	Kecamatan_Id *uint      `json:"kecamatan_id"`
	Kelurahan_Id *uint      `json:"kelurahan_id"`
	Telp         *string    `json:"telp"`
	KodePos      *string    `json:"kodePos"`
	RtRw         *string    `json:"rtRw"`
	FotoKtp      *string    `json:"fotoKtp"`
	TempatLahir  *string    `json:"tempatLahir"`
	TanggalLahir *time.Time `json:"tanggalLahir"`
	Gender       *uint8     `json:"gender"`
	Pekerjaan    *string    `json:"pekerjaan"`
	// User
	// Email *string `json:"email" validate:"validemail"`
}

type FilterDto struct {
	Nik          *string `json:"nik"`
	Nama         *string `json:"nama"`
	Alamat       *string `json:"alamat"`
	Provinsi_Id  *uint   `json:"provinsi_id"`
	Kota_id      *uint   `json:"kota_id"`
	Kecamatan_Id *uint   `json:"kecamatan_id"`
	Kelurahan_Id *uint   `json:"kelurahan_id"`
	TempatLahir  *string `json:"tempatLahir"`
	Gender       *uint8  `json:"gender"`
	Pekerjaan    *string `json:"pekerjaan"`
	Telp         *string `json:"telp"`
	KodePos      *string `json:"kodePos"`
	Status       *uint8  `json:"status"`
	// fixed fields
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type CheckerPOneDto struct {
	Nik          string    `json:"nik" validate:"required"`
	Nama         string    `json:"nama" validate:"required"`
	TempatLahir  string    `json:"tempatLahir" validate:"required"`
	TanggalLahir time.Time `json:"tanggalLahir" validate:"required"`
	Gender       uint8     `json:"gender" validate:"required,min=1"`
	Pekerjaan    string    `json:"pekerjaan" validate:"required"`
	Telp         string    `json:"telp" validate:"required"`
}

type CheckerPTwoDto struct {
	Alamat       string `json:"alamat" validate:"required"`
	Provinsi_Id  uint   `json:"provinsi_id" validate:"min=1"`
	Kota_id      uint   `json:"kota_id" validate:"min=1"`
	Kecamatan_Id uint   `json:"kecamatan_id" validate:"min=1"`
	Kelurahan_Id uint   `json:"kelurahan_id" validate:"min=1"`
	KodePos      string `json:"kodePos" validate:"required"`
	RtRw         string `json:"rtRw" validate:"required"`
}

type CheckerPFourDto struct {
	FotoKtp string `json:"fotoKtp" validate:"required"`
}

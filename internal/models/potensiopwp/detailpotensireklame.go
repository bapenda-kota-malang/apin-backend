package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiReklame struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Potensiop_Id  uuid.UUID `json:"potensiop_id" gorm:"type:uuid"`
	JudulReklame  string    `json:"judulReklame"`
	NamaProduk    string    `json:"namaProduk"`
	JumlahTahun   uint16    `json:"jumlahTahun"`
	JumlahBulan   uint16    `json:"jumlahBulan"`
	JumlahMinggu  uint16    `json:"jumlahMinggu"`
	JumlahHari    uint16    `json:"jumlahHari"`
	JenisReklame  string    `json:"jenisReklame"`
	Lokasi        string    `json:"lokasi"`
	NoPersil      string    `json:"noPersil"`
	JumlahReklame uint64    `json:"jumlahReklame"`
	JumlahSisi    uint64    `json:"jumlahSisi"`
	Panjang       *float64  `json:"panjang"`
	Lebar         *float64  `json:"lebar"`
	Diameter      *float64  `json:"diameter"`
	Luas          float64   `json:"luas"`
	gormhelper.DateModel
}

type CreateDtoDPReklame struct {
	Potensiop_Id  uuid.UUID `json:"-"`
	JudulReklame  string    `json:"judulReklame"`
	NamaProduk    string    `json:"namaProduk"`
	JumlahTahun   uint16    `json:"jumlahTahun"`
	JumlahBulan   uint16    `json:"jumlahBulan"`
	JumlahMinggu  uint16    `json:"jumlahMinggu"`
	JumlahHari    uint16    `json:"jumlahHari"`
	JenisReklame  string    `json:"jenisReklame"`
	Lokasi        string    `json:"lokasi"`
	NoPersil      string    `json:"noPersil"`
	JumlahReklame uint64    `json:"jumlahReklame"`
	JumlahSisi    uint64    `json:"jumlahSisi"`
	Panjang       *float64  `json:"panjang"`
	Lebar         *float64  `json:"lebar"`
	Diameter      *float64  `json:"diameter"`
	Luas          float64   `json:"luas"`
}

type CreateDtoReklame struct {
	CreateDto
	DetailPajakDtos []CreateDtoDPReklame `json:"detailPajaks"`
}

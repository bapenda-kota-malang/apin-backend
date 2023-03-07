package njoptkp

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type NJOPTKP struct {
	Id            uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode   string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	ThnAwal       string  `json:"thnAwal" gorm:"type:varchar(4)"`
	ThnAkhir      string  `json:"thnAkhir" gorm:"type:varchar(4)"`
	NilaiNJOPTKP  float64 `json:"nilaiNJOPTKP"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode string  `json:"provinsi_kode"`
	Daerah_Kode   string  `json:"daerah_kode"`
	ThnAwal       string  `json:"thnAwal"`
	ThnAkhir      string  `json:"thnAkhir"`
	NilaiNJOPTKP  float64 `json:"nilaiNJOPTKP"`
}

type UpdateDto struct {
	Id            uint64  `json:"id"`
	Provinsi_Kode string  `json:"provinsi_kode"`
	Daerah_Kode   string  `json:"daerah_kode"`
	ThnAwal       string  `json:"thnAwal"`
	ThnAkhir      string  `json:"thnAkhir"`
	NilaiNJOPTKP  float64 `json:"nilaiNJOPTKP"`
}

type FilterDto struct {
	ThnAwal  string `json:"thnAwal"`
	ThnAkhir string `json:"thnAkhir"`
}

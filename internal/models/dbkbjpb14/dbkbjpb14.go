package dbkbjpb14

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb14 struct {
	Id          uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode    *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb14 *string  `json:"tahunDbkbJpb14" gorm:"type:char(4)"`
	NilaiDbkbJpb14 *float64 `json:"nilaiDbkbJpb14"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb14 *string  `json:"tahunDbkbJpb14"`
	NilaiDbkbJpb14 *float64 `json:"nilaiDbkbJpb14"`
}

type UpdateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb14 *string  `json:"tahunDbkbJpb14"`
	NilaiDbkbJpb14 *float64 `json:"nilaiDbkbJpb14"`
}

type FilterDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb14 *string  `json:"tahunDbkbJpb14"`
	NilaiDbkbJpb14 *float64 `json:"nilaiDbkbJpb14"`
	Page           int      `json:"page"`
	PageSize       int      `json:"page_size"`
}

package dbkbjpb16

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb16 struct {
	Id             uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode    *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb16 *string  `json:"tahunDbkbJpb16" gorm:"type:char(4)"`
	KelasDbkbJpb16 *string  `json:"kelasDbkbJpb16" gorm:"type:char(1)"`
	LantaiMinJpb16 *int     `json:"lantaiMinJpb16"`
	LantaiMaxJpb16 *int     `json:"lantaiMaxJpb16"`
	NilaiDbkbJpb16 *float64 `json:"nilaiDbkbJpb16"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb16 *string  `json:"tahunDbkbJpb16"`
	KelasDbkbJpb16 *string  `json:"kelasDbkbJpb16"`
	LantaiMinJpb16 *int     `json:"lantaiMinJpb16"`
	LantaiMaxJpb16 *int     `json:"lantaiMaxJpb16"`
	NilaiDbkbJpb16 *float64 `json:"nilaiDbkbJpb16"`
}

type UpdateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb16 *string  `json:"tahunDbkbJpb16"`
	KelasDbkbJpb16 *string  `json:"kelasDbkbJpb16"`
	LantaiMinJpb16 *int     `json:"lantaiMinJpb16"`
	LantaiMaxJpb16 *int     `json:"lantaiMaxJpb16"`
	NilaiDbkbJpb16 *float64 `json:"nilaiDbkbJpb16"`
}

type FilterDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb16 *string  `json:"tahunDbkbJpb16"`
	KelasDbkbJpb16 *string  `json:"kelasDbkbJpb16"`
	LantaiMinJpb16 *int     `json:"lantaiMinJpb16"`
	LantaiMaxJpb16 *int     `json:"lantaiMaxJpb16"`
	NilaiDbkbJpb16 *float64 `json:"nilaiDbkbJpb16"`
	Page           int      `json:"page"`
	PageSize       int      `json:"page_size"`
}

package dbkbjpb13

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb13 struct {
	Id             uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode    *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb13 *string  `json:"tahunDbkbJpb13" gorm:"type:char(4)"`
	KelasDbkbJpb13 *string  `json:"kelasDbkbJpb13" gorm:"type:char(1)"`
	LantaiMinJpb13 *int     `json:"lantaiMinJpb13"`
	LantaiMaxJpb13 *int     `json:"lantaiMaxJpb13"`
	NilaiDbkbJpb13 *float64 `json:"nilaiDbkbJpb13"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb13 *string  `json:"tahunDbkbJpb13"`
	KelasDbkbJpb13 *string  `json:"kelasDbkbJpb13"`
	LantaiMinJpb13 *int     `json:"lantaiMinJpb13"`
	LantaiMaxJpb13 *int     `json:"lantaiMaxJpb13"`
	NilaiDbkbJpb13 *float64 `json:"nilaiDbkbJpb13"`
}

type UpdateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb13 *string  `json:"tahunDbkbJpb13"`
	KelasDbkbJpb13 *string  `json:"kelasDbkbJpb13"`
	LantaiMinJpb13 *int     `json:"lantaiMinJpb13"`
	LantaiMaxJpb13 *int     `json:"lantaiMaxJpb13"`
	NilaiDbkbJpb13 *float64 `json:"nilaiDbkbJpb13"`
}

type FilterDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb13 *string  `json:"tahunDbkbJpb13"`
	KelasDbkbJpb13 *string  `json:"kelasDbkbJpb13"`
	LantaiMinJpb13 *int     `json:"lantaiMinJpb13"`
	LantaiMaxJpb13 *int     `json:"lantaiMaxJpb13"`
	NilaiDbkbJpb13 *float64 `json:"nilaiDbkbJpb13"`
	Page           int      `json:"page"`
	PageSize       int      `json:"page_size"`
}

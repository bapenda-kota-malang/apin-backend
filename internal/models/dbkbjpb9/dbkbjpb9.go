package dbkbjpb9

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb9 struct {
	Id            uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode   *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb9 *string  `json:"tahunDbkbJpb9" gorm:"type:char(4)"`
	KelasDbkbJpb9 *string  `json:"kelasDbkbJpb9" gorm:"type:char(1)"`
	LantaiMinJpb9 *int     `json:"lantaiMinJpb9"`
	LantaiMaxJpb9 *int     `json:"lantaiMaxJpb9"`
	NilaiDbkbJpb9 *float64 `json:"nilaiDbkbJpb9"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb9 *string  `json:"tahunDbkbJpb9"`
	KelasDbkbJpb9 *string  `json:"kelasDbkbJpb9"`
	LantaiMinJpb9 *int     `json:"lantaiMinJpb9"`
	LantaiMaxJpb9 *int     `json:"lantaiMaxJpb9"`
	NilaiDbkbJpb9 *float64 `json:"nilaiDbkbJpb9"`
}

type UpdateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb9 *string  `json:"tahunDbkbJpb9"`
	KelasDbkbJpb9 *string  `json:"kelasDbkbJpb9"`
	LantaiMinJpb9 *int     `json:"lantaiMinJpb9"`
	LantaiMaxJpb9 *int     `json:"lantaiMaxJpb9"`
	NilaiDbkbJpb9 *float64 `json:"nilaiDbkbJpb9"`
}

type FilterDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb9 *string  `json:"tahunDbkbJpb9"`
	KelasDbkbJpb9 *string  `json:"kelasDbkbJpb9"`
	LantaiMinJpb9 *int     `json:"lantaiMinJpb9"`
	LantaiMaxJpb9 *int     `json:"lantaiMaxJpb9"`
	NilaiDbkbJpb9 *float64 `json:"nilaiDbkbJpb9"`
	Page          int      `json:"page"`
	PageSize      int      `json:"page_size"`
}

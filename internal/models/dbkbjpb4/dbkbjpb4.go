package dbkbjpb4

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb4 struct {
	Id            uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode   *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb4 *string  `json:"tahunDbkbJpb4" gorm:"type:char(4)"`
	KelasDbkbJpb4 *string  `json:"kelasDbkbJpb4" gorm:"type:char(1)"`
	LantaiMinJpb4 *int     `json:"lantaiMinJpb4"`
	LantaiMaxJpb4 *int     `json:"lantaiMaxJpb4"`
	NilaiDbkbJpb4 *float64 `json:"nilaiDbkbJpb4"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb4 *string  `json:"tahunDbkbJpb4"`
	KelasDbkbJpb4 *string  `json:"kelasDbkbJpb4"`
	LantaiMinJpb4 *int     `json:"lantaiMinJpb4"`
	LantaiMaxJpb4 *int     `json:"lantaiMaxJpb4"`
	NilaiDbkbJpb4 *float64 `json:"nilaiDbkbJpb4"`
}

type UpdateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb4 *string  `json:"tahunDbkbJpb4"`
	KelasDbkbJpb4 *string  `json:"kelasDbkbJpb4"`
	LantaiMinJpb4 *int     `json:"lantaiMinJpb4"`
	LantaiMaxJpb4 *int     `json:"lantaiMaxJpb4"`
	NilaiDbkbJpb4 *float64 `json:"nilaiDbkbJpb4"`
}

type FilterDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb4 *string  `json:"tahunDbkbJpb4"`
	KelasDbkbJpb4 *string  `json:"kelasDbkbJpb4"`
	LantaiMinJpb4 *int     `json:"lantaiMinJpb4"`
	LantaiMaxJpb4 *int     `json:"lantaiMaxJpb4"`
	NilaiDbkbJpb4 *float64 `json:"nilaiDbkbJpb4"`
	Page          int      `json:"page"`
	PageSize      int      `json:"page_size"`
}

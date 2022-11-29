package dbkbjpb5

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb5 struct {
	Id            uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode   *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb5 *string  `json:"tahunDbkbJpb5" gorm:"type:char(4)"`
	KelasDbkbJpb5 *string  `json:"kelasDbkbJpb5" gorm:"type:char(1)"`
	LantaiMinJpb5 *int     `json:"lantaiMinJpb5"`
	LantaiMaxJpb5 *int     `json:"lantaiMaxJpb5"`
	NilaiDbkbJpb5 *float64 `json:"nilaiDbkbJpb5"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb5 *string  `json:"tahunDbkbJpb5"`
	KelasDbkbJpb5 *string  `json:"kelasDbkbJpb5"`
	LantaiMinJpb5 *int     `json:"lantaiMinJpb5"`
	LantaiMaxJpb5 *int     `json:"lantaiMaxJpb5"`
	NilaiDbkbJpb5 *float64 `json:"nilaiDbkbJpb5"`
}

type UpdateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb5 *string  `json:"tahunDbkbJpb5"`
	KelasDbkbJpb5 *string  `json:"kelasDbkbJpb5"`
	LantaiMinJpb5 *int     `json:"lantaiMinJpb5"`
	LantaiMaxJpb5 *int     `json:"lantaiMaxJpb5"`
	NilaiDbkbJpb5 *float64 `json:"nilaiDbkbJpb5"`
}

type FilterDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb5 *string  `json:"tahunDbkbJpb5"`
	KelasDbkbJpb5 *string  `json:"kelasDbkbJpb5"`
	LantaiMinJpb5 *int     `json:"lantaiMinJpb5"`
	LantaiMaxJpb5 *int     `json:"lantaiMaxJpb5"`
	NilaiDbkbJpb5 *float64 `json:"nilaiDbkbJpb5"`
	Page          int      `json:"page"`
	PageSize      int      `json:"page_size"`
}

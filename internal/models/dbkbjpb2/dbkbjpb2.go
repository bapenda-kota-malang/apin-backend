package dbkbjpb2

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb2 struct {
	Id            uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode   *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb2 *string  `json:"tahunDbkbJpb2" gorm:"type:char(4)"`
	KelasDbkbJpb2 *string  `json:"kelasDbkbJpb2" gorm:"type:char(1)"`
	LantaiMinJpb2 *int     `json:"lantaiMinJpb2"`
	LantaiMaxJpb2 *int     `json:"lantaiMaxJpb2"`
	NilaiDbkbJpb2 *float64 `json:"nilaiDbkbJpb2"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb2 *string  `json:"tahunDbkbJpb2"`
	KelasDbkbJpb2 *string  `json:"kelasDbkbJpb2"`
	LantaiMinJpb2 *int     `json:"lantaiMinJpb2"`
	LantaiMaxJpb2 *int     `json:"lantaiMaxJpb2"`
	NilaiDbkbJpb2 *float64 `json:"nilaiDbkbJpb2"`
}

type UpdateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb2 *string  `json:"tahunDbkbJpb2"`
	KelasDbkbJpb2 *string  `json:"kelasDbkbJpb2"`
	LantaiMinJpb2 *int     `json:"lantaiMinJpb2"`
	LantaiMaxJpb2 *int     `json:"lantaiMaxJpb2"`
	NilaiDbkbJpb2 *float64 `json:"nilaiDbkbJpb2"`
}

type FilterDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb2 *string  `json:"tahunDbkbJpb2"`
	KelasDbkbJpb2 *string  `json:"kelasDbkbJpb2"`
	LantaiMinJpb2 *int     `json:"lantaiMinJpb2"`
	LantaiMaxJpb2 *int     `json:"lantaiMaxJpb2"`
	NilaiDbkbJpb2 *float64 `json:"nilaiDbkbJpb2"`
	Page          int      `json:"page"`
	PageSize      int      `json:"page_size"`
}

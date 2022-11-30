package dbkbjpb6

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb6 struct {
	Id            uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode   *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb6 *string  `json:"tahunDbkbJpb6" gorm:"type:char(4)"`
	KelasDbkbJpb6 *string  `json:"kelasDbkbJpb6" gorm:"type:char(1)"`
	NilaiDbkbJpb6 *float64 `json:"nilaiDbkbJpb6"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb6 *string  `json:"tahunDbkbJpb6"`
	KelasDbkbJpb6 *string  `json:"kelasDbkbJpb6"`
	NilaiDbkbJpb6 *float64 `json:"nilaiDbkbJpb6"`
}

type UpdateDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb6 *string  `json:"tahunDbkbJpb6"`
	KelasDbkbJpb6 *string  `json:"kelasDbkbJpb6"`
	NilaiDbkbJpb6 *float64 `json:"nilaiDbkbJpb6"`
}

type FilterDto struct {
	Provinsi_Kode *string  `json:"provinsi_kode"`
	Daerah_Kode   *string  `json:"daerah_kode"`
	TahunDbkbJpb6 *string  `json:"tahunDbkbJpb6"`
	KelasDbkbJpb6 *string  `json:"kelasDbkbJpb6"`
	NilaiDbkbJpb6 *float64 `json:"nilaiDbkbJpb6"`
	Page          int      `json:"page"`
	PageSize      int      `json:"page_size"`
}

package dbkbjpb12

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb12 struct {
	Id             uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode    *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb12 *string  `json:"tahunDbkbJpb12" gorm:"type:char(4)"`
	TipeDbkbJpb12  *string  `json:"tipeDbkbJpb12" gorm:"type:char(1)"`
	NilaiDbkbJpb12 *float64 `json:"nilaiDbkbJpb12"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb12 *string  `json:"tahunDbkbJpb12"`
	TipeDbkbJpb12  *string  `json:"tipeDbkbJpb12"`
	NilaiDbkbJpb12 *float64 `json:"nilaiDbkbJpb12"`
}

type UpdateDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb12 *string  `json:"tahunDbkbJpb12"`
	TipeDbkbJpb12  *string  `json:"tipeDbkbJpb12"`
	NilaiDbkbJpb12 *float64 `json:"nilaiDbkbJpb12"`
}

type FilterDto struct {
	Provinsi_Kode  *string  `json:"provinsi_kode"`
	Daerah_Kode    *string  `json:"daerah_kode"`
	TahunDbkbJpb12 *string  `json:"tahunDbkbJpb6"`
	TipeDbkbJpb12  *string  `json:"tipeDbkbJpb12"`
	NilaiDbkbJpb12 *float64 `json:"nilaiDbkbJpb12"`
	Page           int      `json:"page"`
	PageSize       int      `json:"page_size"`
}

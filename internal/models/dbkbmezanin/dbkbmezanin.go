package dbkbmezanin

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbMezanin struct {
	Id               uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode    *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode      *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbMezanin *string  `json:"tahunDbkbMezanin" gorm:"type:char(4)"`
	NilaiDbkbMezanin *float64 `json:"nilaiDbkbMezanin"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode    *string  `json:"provinsi_kode"`
	Daerah_Kode      *string  `json:"daerah_kode"`
	TahunDbkbMezanin *string  `json:"tahunDbkbMezanin"`
	NilaiDbkbMezanin *float64 `json:"nilaiDbkbMezanin"`
}

type UpdateDto struct {
	Provinsi_Kode    *string  `json:"provinsi_kode"`
	Daerah_Kode      *string  `json:"daerah_kode"`
	TahunDbkbMezanin *string  `json:"tahunDbkbMezanin"`
	NilaiDbkbMezanin *float64 `json:"nilaiDbkbMezanin"`
}

type FilterDto struct {
	Provinsi_Kode    *string  `json:"provinsi_kode"`
	Daerah_Kode      *string  `json:"daerah_kode"`
	TahunDbkbMezanin *string  `json:"tahunDbkbMezanin"`
	NilaiDbkbMezanin *float64 `json:"nilaiDbkbMezanin"`
	Page             int      `json:"page"`
	PageSize         int      `json:"page_size"`
}

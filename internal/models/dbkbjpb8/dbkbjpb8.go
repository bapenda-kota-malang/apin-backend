package dbkbjpb8

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb8 struct {
	Id                     uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode          *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode            *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb8          *string  `json:"tahunDbkbJpb8"`
	LebarBentukMinDbkbJpb8 *float64 `json:"lebarBentukMinDbkbJpb8"`
	LebarBentukMaxDbkbJpb8 *float64 `json:"lebarBentukMaxDbkbJpb8"`
	TinggiKolomMinDbkbJpb8 *float64 `json:"tinggiKolomMinDbkbJpb8"`
	TinggiKolomMaxDbkbJpb8 *float64 `json:"tinggiKolomMaxDbkbJpb8"`
	NilaiDbkbJpb8          *float64 `json:"nilaiDbkbJpb8" gorm:"type:decimal(12,2)"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode          *string  `json:"provinsi_kode"`
	Daerah_Kode            *string  `json:"daerah_kode"`
	TahunDbkbJpb8          *string  `json:"tahunDbkbJpb8"`
	LebarBentukMinDbkbJpb8 *float64 `json:"lebarBentukMinDbkbJpb8"`
	LebarBentukMaxDbkbJpb8 *float64 `json:"lebarBentukMaxDbkbJpb8"`
	TinggiKolomMinDbkbJpb8 *float64 `json:"tinggiKolomMinDbkbJpb8"`
	TinggiKolomMaxDbkbJpb8 *float64 `json:"tinggiKolomMaxDbkbJpb8"`
	NilaiDbkbJpb8          *float64 `json:"nilaiDbkbJpb8"`
}

type UpdateDto struct {
	Provinsi_Kode          *string  `json:"provinsi_kode"`
	Daerah_Kode            *string  `json:"daerah_kode"`
	TahunDbkbJpb8          *string  `json:"tahunDbkbJpb8"`
	LebarBentukMinDbkbJpb8 *float64 `json:"lebarBentukMinDbkbJpb8"`
	LebarBentukMaxDbkbJpb8 *float64 `json:"lebarBentukMaxDbkbJpb8"`
	TinggiKolomMinDbkbJpb8 *float64 `json:"tinggiKolomMinDbkbJpb8"`
	TinggiKolomMaxDbkbJpb8 *float64 `json:"tinggiKolomMaxDbkbJpb8"`
	NilaiDbkbJpb8          *float64 `json:"nilaiDbkbJpb8"`
}

type FilterDto struct {
	Provinsi_Kode          *string  `json:"provinsi_kode"`
	Daerah_Kode            *string  `json:"daerah_kode"`
	TahunDbkbJpb8          *string  `json:"tahunDbkbJpb8"`
	LebarBentukMinDbkbJpb8 *float64 `json:"lebarBentukMinDbkbJpb8"`
	LebarBentukMaxDbkbJpb8 *float64 `json:"lebarBentukMaxDbkbJpb8"`
	TinggiKolomMinDbkbJpb8 *float64 `json:"tinggiKolomMinDbkbJpb8"`
	TinggiKolomMaxDbkbJpb8 *float64 `json:"tinggiKolomMaxDbkbJpb8"`
	NilaiDbkbJpb8          *float64 `json:"nilaiDbkbJpb8"`
	Page                   int      `json:"page"`
	PageSize               int      `json:"page_size"`
}

package dbkbjpb3

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb3 struct {
	Id                     uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode          *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode            *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb3          *string  `json:"tahunDbkbJpb3"`
	LebarBentukMinDbkbJpb3 *float64 `json:"lebarBentukMinDbkbJpb3"`
	LebarBentukMaxDbkbJpb3 *float64 `json:"lebarBentukMaxDbkbJpb3"`
	TinggiKolomMinDbkbJpb3 *float64 `json:"tinggiKolomMinDbkbJpb3"`
	TinggiKolomMaxDbkbJpb3 *float64 `json:"tinggiKolomMaxDbkbJpb3"`
	NilaiDbkbJpb3          *float64 `json:"nilaiDbkbJpb3" gorm:"type:decimal(12,2)"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode          *string  `json:"provinsi_kode"`
	Daerah_Kode            *string  `json:"daerah_kode"`
	TahunDbkbJpb3          *string  `json:"tahunDbkbJpb3"`
	LebarBentukMinDbkbJpb3 *float64 `json:"lebarBentukMinDbkbJpb3"`
	LebarBentukMaxDbkbJpb3 *float64 `json:"lebarBentukMaxDbkbJpb3"`
	TinggiKolomMinDbkbJpb3 *float64 `json:"tinggiKolomMinDbkbJpb3"`
	TinggiKolomMaxDbkbJpb3 *float64 `json:"tinggiKolomMaxDbkbJpb3"`
	NilaiDbkbJpb3          *float64 `json:"nilaiDbkbJpb3"`
}

type UpdateDto struct {
	Provinsi_Kode          *string  `json:"provinsi_kode"`
	Daerah_Kode            *string  `json:"daerah_kode"`
	TahunDbkbJpb3          *string  `json:"tahunDbkbJpb3"`
	LebarBentukMinDbkbJpb3 *float64 `json:"lebarBentukMinDbkbJpb3"`
	LebarBentukMaxDbkbJpb3 *float64 `json:"lebarBentukMaxDbkbJpb3"`
	TinggiKolomMinDbkbJpb3 *float64 `json:"tinggiKolomMinDbkbJpb3"`
	TinggiKolomMaxDbkbJpb3 *float64 `json:"tinggiKolomMaxDbkbJpb3"`
	NilaiDbkbJpb3          *float64 `json:"nilaiDbkbJpb3"`
}

type FilterDto struct {
	Provinsi_Kode          *string  `json:"provinsi_kode"`
	Daerah_Kode            *string  `json:"daerah_kode"`
	TahunDbkbJpb3          *string  `json:"tahunDbkbJpb3"`
	LebarBentukMinDbkbJpb3 *float64 `json:"lebarBentukMinDbkbJpb3"`
	LebarBentukMaxDbkbJpb3 *float64 `json:"lebarBentukMaxDbkbJpb3"`
	TinggiKolomMinDbkbJpb3 *float64 `json:"tinggiKolomMinDbkbJpb3"`
	TinggiKolomMaxDbkbJpb3 *float64 `json:"tinggiKolomMaxDbkbJpb3"`
	NilaiDbkbJpb3          *float64 `json:"nilaiDbkbJpb3"`
	Page                   int      `json:"page"`
	PageSize               int      `json:"page_size"`
}

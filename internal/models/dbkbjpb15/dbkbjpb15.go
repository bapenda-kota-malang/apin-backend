package dbkbjpb15

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DbkbJpb15 struct {
	Id                    uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode         *string  `json:"provinsi_kode"  gorm:"type:char(2)"`
	Daerah_Kode           *string  `json:"daerah_kode"  gorm:"type:char(2)"`
	TahunDbkbJpb15        *string  `json:"tahunDbkbJpb15" gorm:"type:char(4)"`
	JenisTangkiDbkbJpb15  *string  `json:"jenisTangkiJpb15" gorm:"type:char(1)"`
	KapasitasMinDbkbJpb15 *float64 `json:"kapasitasMinDbkbJpb15" gorm:"type:decimal(9,4)"`
	KapasitasMaxDbkbJpb15 *float64 `json:"kapasitasMaxDbkbJpb15" gorm:"type:decimal(9,4)"`
	NilaiDbkbJpb15        *float64 `json:"nilaiDbkbJpb15"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode         *string  `json:"provinsi_kode"`
	Daerah_Kode           *string  `json:"daerah_kode"`
	TahunDbkbJpb15        *string  `json:"tahunDbkbJpb15"`
	JenisTangkiDbkbJpb15  *string  `json:"jenisTangkiJpb15"`
	KapasitasMinDbkbJpb15 *float64 `json:"kapasitasMinDbkbJpb15"`
	KapasitasMaxDbkbJpb15 *float64 `json:"kapasitasMaxDbkbJpb15"`
	NilaiDbkbJpb15        *float64 `json:"nilaiDbkbJpb15"`
}

type UpdateDto struct {
	Provinsi_Kode         *string  `json:"provinsi_kode"`
	Daerah_Kode           *string  `json:"daerah_kode"`
	TahunDbkbJpb15        *string  `json:"tahunDbkbJpb15"`
	JenisTangkiDbkbJpb15  *string  `json:"jenisTangkiJpb15"`
	KapasitasMinDbkbJpb15 *float64 `json:"kapasitasMinDbkbJpb15"`
	KapasitasMaxDbkbJpb15 *float64 `json:"kapasitasMaxDbkbJpb15"`
	NilaiDbkbJpb15        *float64 `json:"nilaiDbkbJpb15"`
}

type FilterDto struct {
	Provinsi_Kode         *string  `json:"provinsi_kode"`
	Daerah_Kode           *string  `json:"daerah_kode"`
	TahunDbkbJpb15        *string  `json:"tahunDbkbJpb15"`
	JenisTangkiDbkbJpb15  *string  `json:"jenisTangkiJpb15"`
	KapasitasMinDbkbJpb15 *float64 `json:"kapasitasMinDbkbJpb15"`
	KapasitasMaxDbkbJpb15 *float64 `json:"kapasitasMaxDbkbJpb15"`
	NilaiDbkbJpb15        *float64 `json:"nilaiDbkbJpb15"`
	Page                  int      `json:"page"`
	PageSize              int      `json:"page_size"`
}

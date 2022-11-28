package kelastanah

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type KelasTanah struct {
	Id                   uint64   `json:"id" gorm:"primaryKey"`
	NilaiMaxTanah        *float64 `json:"nilaiMaxTanah" gorm:"type:decimal(8,2)"`
	NilaiMinTanah        *float64 `json:"nilaiMinTanah" gorm:"type:decimal(8,2)"`
	NilaiPerM2Tanah      *float64 `json:"nilaiPerM2Tanah" gorm:"type:decimal(8,2)"`
	TahunAkhirKelasTanah *string  `json:"tahunAkhirKelasTanah" gorm:"type:char(4)"`
	TahunAwalKelasTanah  *string  `json:"tahunAwalKelasTanah" gorm:"type:char(4)"`
	gh.DateModel
}

type CreateDto struct {
	NilaiMaxTanah        *float64 `json:"nilaiMaxTanah"`
	NilaiMinTanah        *float64 `json:"nilaiMinTanah"`
	NilaiPerM2Tanah      *float64 `json:"nilaiPerM2Tanah"`
	TahunAkhirKelasTanah *string  `json:"tahunAkhirKelasTanah"`
	TahunAwalKelasTanah  *string  `json:"tahunAwalKelasTanah"`
}

type UpdateDto struct {
	NilaiMaxTanah        *float64 `json:"nilaiMaxTanah"`
	NilaiMinTanah        *float64 `json:"nilaiMinTanah"`
	NilaiPerM2Tanah      *float64 `json:"nilaiPerM2Tanah"`
	TahunAkhirKelasTanah *string  `json:"tahunAkhirKelasTanah"`
	TahunAwalKelasTanah  *string  `json:"tahunAwalKelasTanah"`
}

type FilterDto struct {
	NilaiMaxTanah        *float64 `json:"nilaiMaxTanah"`
	NilaiMinTanah        *float64 `json:"nilaiMinTanah"`
	NilaiPerM2Tanah      *float64 `json:"nilaiPerM2Tanah"`
	TahunAkhirKelasTanah *string  `json:"tahunAkhirKelasTanah"`
	TahunAwalKelasTanah  *string  `json:"tahunAwalKelasTanah"`
	Page                 int      `json:"page"`
	PageSize             int      `json:"page_size"`
}

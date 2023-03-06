package referensibuku

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type ReferensiBuku struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	Kode     string  `json:"kode" gorm:"type:varchar(1);not null"`
	ThnAwal  string  `json:"thnAwal" gorm:"type:varchar(4);not null"`
	ThnAkhir string  `json:"thnAkhir" gorm:"type:varchar(4);not null"`
	NilaiMin float64 `json:"nilaiMin" gorm:"size:15;not null"`
	NilaiMax float64 `json:"nilaiMax" gorm:"size:15;not null"`
	gormhelper.DateModel
}

type CreateDto struct {
	Kode     string  `json:"kode" validate:"required"`
	ThnAwal  string  `json:"thnAwal" validate:"required"`
	ThnAkhir string  `json:"thnAkhir" validate:"required"`
	NilaiMin float64 `json:"nilaiMin" validate:"required"`
	NilaiMax float64 `json:"nilaiMax" validate:"required"`
}

type FilterDto struct {
	Kode     *string  `json:"kode"`
	ThnAwal  *string  `json:"thnAwal"`
	ThnAkhir *string  `json:"thnAkhir"`
	NilaiMin *float64 `json:"nilaiMin"`
	NilaiMax *float64 `json:"nilaiMax"`
	Page     int      `json:"page"`
	PageSize int      `json:"page_size"`
}

type UpdateDto struct {
	Kode     *string  `json:"kode"`
	ThnAwal  *string  `json:"thnAwal"`
	ThnAkhir *string  `json:"thnAkhir"`
	NilaiMin *float64 `json:"nilaiMin"`
	NilaiMax *float64 `json:"nilaiMax"`
}

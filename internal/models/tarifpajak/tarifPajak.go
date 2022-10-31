package tarifpajak

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
)

type TarifPajak struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	Tahun       *uint64            `json:"tahun" gorm:"type:integer"`
	TarifRp     *float64           `json:"tarifRp" gorm:"type:decimal"`
	TarifPersen *float64           `json:"tarifPersen" gorm:"type:decimal"`
	OmsetAwal   *float64           `json:"omsetAwal" gorm:"type:decimal"`
	OmsetAkhir  *float64           `json:"omsetAkhir" gorm:"type:decimal"`
}

type CreateDto struct {
	Rekening_Id uint64  `json:"rekening_id"`
	Tahun       uint64  `json:"tahun" validate:"required"`
	TarifRp     float64 `json:"tarifRp" validate:"required"`
	TarifPersen float64 `json:"tarifPersen" validate:"required"`
	OmsetAwal   float64 `json:"omsetAwal" validate:"required"`
	OmsetAkhir  float64 `json:"omsetAkhir" validate:"required"`
}

type UpdateDto struct {
	Rekening_Id uint64  `json:"rekening_id"`
	Tahun       uint64  `json:"tahun"`
	TarifRp     float64 `json:"tarifRp"`
	TarifPersen float64 `json:"tarifPersen"`
	OmsetAwal   float64 `json:"omsetAwal"`
	OmsetAkhir  float64 `json:"omsetAkhir"`
}

type FilterDto struct {
	Rekening_Id *uint64  `json:"rekening_id"`
	Tahun       *uint64  `json:"tahun"`
	TarifRp     *float64 `json:"tarifRp"`
	TarifPersen *float64 `json:"tarifPersen"`
	OmsetAwal   *float64 `json:"omsetAwal"`
	OmsetAkhir  *float64 `json:"omsetAkhir"`
}

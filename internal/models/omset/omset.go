package omset

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
)

type Omset struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	OmsetMin    *float64           `json:"omsetMin" gorm:"type:decimal"`
	OmsetMax    *float64           `json:"omsetMax" gorm:"type:decimal"`
	TarifRp     *float64           `json:"tarifRp" gorm:"type:decimal"`
	TarifPersen *float64           `json:"tarifPersen" gorm:"type:decimal"`
	TahunPajak  *string            `json:"tahunPajak" gorm:"type:varchar(4)"`
}

type CreateDto struct {
	Rekening_Id uint64  `json:"rekening_id" validate:"required"`
	OmsetMin    float64 `json:"omsetMin" validate:"required"`
	OmsetMax    float64 `json:"omsetMax" validate:"required"`
	TarifRp     float64 `json:"tarifRp" validate:"required"`
	TarifPersen float64 `json:"tarifPersen" validate:"required"`
	TahunPajak  string  `json:"tahunPajak" validate:"required"`
}

type UpdateDto struct {
	Rekening_Id uint64  `json:"rekening_id"`
	OmsetMin    float64 `json:"omsetMin"`
	OmsetMax    float64 `json:"omsetMax"`
	TarifRp     float64 `json:"tarifRp"`
	TarifPersen float64 `json:"tarifPersen"`
	TahunPajak  string  `json:"tahunPajak"`
}

type FilterDto struct {
	Rekening_Id *uint64  `json:"rekening_id"`
	OmsetMin    *float64 `json:"omsetMin"`
	OmsetMax    *float64 `json:"omsetMax"`
	TarifRp     *float64 `json:"tarifRp"`
	TarifPersen *float64 `json:"tarifPersen"`
	TahunPajak  *string  `json:"tahunPajak"`
	Page        int      `json:"page"`
	PageSize    int      `json:"page_size"`
}

package tarifreklame

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/klasifikasijalan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/types"
)

type TarifReklame struct {
	Id                  uint64                             `json:"id" gorm:"primaryKey"`
	JenisMasa           types.JenisMasa                    `json:"jenisMasa" gorm:"type:integer"`
	JenisReklame        *string                            `json:"jenisReklame" gorm:"type:varchar(200)"`
	DasarPengenaan      *string                            `json:"dasarPengenaan" gorm:"type:varchar(100)"`
	KlasifikasiJalan_Id *string                            `json:"klasifikasiJalan_id" gorm:"type:varchar(3)"`
	KlasifikasiJalan    *klasifikasijalan.KlasifikasiJalan `json:"omitempty" gorm:"foreignKey:KlasifikasiJalan_Id"`
	MasaPajak           *string                            `json:"masaPajak,omitempty" gorm:"size:100"`
	Tarif               *float64                           `json:"tarif" gorm:"type:decimal"`
}

type CreateDto struct {
	JenisMasa           types.JenisMasa `json:"jenisMasa" validate:"required"`
	JenisReklame        string          `json:"jenisReklame" validate:"required"`
	DasarPengenaan      string          `json:"dasarPengenaan" validate:"required"`
	KlasifikasiJalan_Id string          `json:"klasifikasiJalan_id" validate:"required"`
	MasaPajak           *string         `json:"masaPajak,omitempty"`
	Tarif               float64         `json:"tarif" validate:"required"`
}

type UpdateDto struct {
	JenisMasa           types.JenisMasa `json:"jenisMasa"`
	JenisReklame        *string         `json:"jenisReklame"`
	DasarPengenaan      string          `json:"dasarPengenaan"`
	KlasifikasiJalan_Id string          `json:"klasifikasiJalan_id"`
	MasaPajak           *string         `json:"masaPajak,omitempty"`
	Tarif               float64         `json:"tarif"`
}

type FilterDto struct {
	JenisMasa           int16    `json:"jenisMasa"`
	JenisReklame        *string  `json:"jenisReklame"`
	DasarPengenaan      *string  `json:"dasarPengenaan"`
	KlasifikasiJalan_Id *string  `json:"klasifikasiJalan_id"`
	Tarif               *float64 `json:"tarif"`
	MasaPajak           *string  `json:"masaPajak,omitempty"`
	Page                int      `json:"page"`
	PageSize            int      `json:"page_size"`
}

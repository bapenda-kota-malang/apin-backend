package detailsptparkir

import (
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptParkir struct {
	Id             uint64            `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id         uuid.UUID         `json:"spt_id" gorm:"type:uuid"`
	JenisKendaraan mt.JenisKendaraan `json:"jenisKendaraan"`
	Kapasitas      *uint64           `json:"kapasitas"`
	Tarif          uint              `json:"tarif"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id         uuid.UUID
	JenisKendaraan mt.JenisKendaraan `json:"jenisKendaraan"  validate:"required"`
	Kapasitas      uint64            `json:"kapasitas"  validate:"required;min=1"`
	Tarif          uint              `json:"tarif" validate:"required;min=1"`
}

type UpdateDto struct {
	Id             uint   `json:"id"`
	JenisKendaraan *uint8 `json:"jenisKendaraan" validate:"min=1;max=2"`
	Kapasitas      *uint  `json:"kapasitas" validate:"min=1"`
	Tarif          *uint  `json:"tarif" validate:"min=1"`
}

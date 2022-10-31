package detailesptparkir

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptParkir struct {
	detailesptair.DetailEspt
	JenisKendaraan uint8 `json:"jenisKendaraan" gorm:"comment:roda dua roda empat"`
	Kapasitas      uint  `json:"kapasitas"`
	Tarif          uint  `json:"tarif"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id        uint
	JenisKendaraan uint8 `json:"jenisKendaraan" validate:"required;min=2;max=4"`
	Kapasitas      uint  `json:"kapasitas" validate:"required;min=1"`
	Tarif          uint  `json:"tarif" validate:"required;min=1"`
}

type UpdateDto struct {
	Id             uint   `json:"id"`
	JenisKendaraan *uint8 `json:"jenisKendaraan" validate:"min=2"`
	Kapasitas      *uint  `json:"kapasitas" validate:"min=1"`
	Tarif          *uint  `json:"tarif" validate:"min=1"`
}

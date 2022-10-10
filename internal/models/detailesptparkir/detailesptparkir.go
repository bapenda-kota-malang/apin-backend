package detailesptparkir

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptParkir struct {
	detailesptair.DetailEspt
	JenisKendaraan uint `json:"jenisKendaraan"`
	Kapasitas      uint `json:"kapasitas"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id        uint `json:"espt_id"`
	JenisKendaraan uint `json:"jenisKendaraan"  validate:"required"`
	Kapasitas      uint `json:"kapasitas"  validate:"required"`
}

type UpdateDto struct {
	Id             uint `json:"id"`
	Espt_Id        uint `json:"espt_id"`
	JenisKendaraan uint `json:"jenisKendaraan"`
	Kapasitas      uint `json:"kapasitas"`
}

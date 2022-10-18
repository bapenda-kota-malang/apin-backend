package detailesptparkir

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptParkir struct {
	detailesptair.DetailEspt
	JenisKendaraan JenisKendaraan `json:"jenisKendaraan"`
	Kapasitas      uint           `json:"kapasitas"`
	Tarif          uint           `json:"tarif"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id        uint
	JenisKendaraan JenisKendaraan `json:"jenisKendaraan"`
	Kapasitas      uint           `json:"kapasitas"`
	Tarif          uint           `json:"tarif"`
}

type UpdateDto struct {
	Id             uint           `json:"id"`
	JenisKendaraan JenisKendaraan `json:"jenisKendaraan"`
	Kapasitas      uint           `json:"kapasitas"`
	Tarif          uint           `json:"tarif"`
}

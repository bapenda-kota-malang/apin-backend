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

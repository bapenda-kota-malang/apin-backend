package detailesptparkir

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type CreateDetailReklameDto struct {
	detailesptair.CreateDto
	JenisKendaraan uint `json:"jenisKendaraan"  validate:"required"`
	Kapasitas      uint `json:"kapasitas"  validate:"required"`
}

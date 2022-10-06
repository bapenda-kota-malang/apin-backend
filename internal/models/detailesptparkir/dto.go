package detailesptparkir

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type CreateDto struct {
	detailesptair.EsptIdDto
	JenisKendaraan uint `json:"jenisKendaraan"  validate:"required"`
	Kapasitas      uint `json:"kapasitas"  validate:"required"`
}

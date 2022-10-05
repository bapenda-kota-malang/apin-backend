package detailesptresto

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type CreateDetailRestoDto struct {
	detailesptair.CreateDto
	JumlahMeja       uint    `json:"jumlahMeja"  validate:"required"`
	JumlahKursi      uint    `json:"jumlahKursi"  validate:"required"`
	TarifMinuman     float32 `json:"tarifMinuman"  validate:"required"`
	TarifMakanan     float32 `json:"tarifMakanan"  validate:"required"`
	JumlahPengunjung uint    `json:"jumlahPengunjung"  validate:"required"`
	LdBt             uint    `json:"ldBt"  validate:"required"`
}

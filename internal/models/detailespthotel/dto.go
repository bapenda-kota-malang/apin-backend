package detailespthotel

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type CreateDto struct {
	detailesptair.EsptIdDto
	GolonganKamar       string  `json:"golonganKamar"  validate:"required"`
	Tarif               float32 `json:"tarif"  validate:"required"`
	JumlahKamar         uint    `json:"jumlahKamar"  validate:"required"`
	JumlahKamarYangLaku uint    `json:"jumlahKamarYangLaku"  validate:"required"`
}

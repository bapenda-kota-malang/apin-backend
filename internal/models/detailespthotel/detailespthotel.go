package detailespthotel

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptHotel struct {
	detailesptair.DetailEspt
	GolonganKamar       string  `json:"golonganKamar" gorm:"size:20"`
	Tarif               float32 `json:"tarif"`
	JumlahKamar         uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku uint    `json:"jumlahKamarYangLaku"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id             uint    `json:"espt_id"`
	GolonganKamar       string  `json:"golonganKamar"  validate:"required"`
	Tarif               float32 `json:"tarif"  validate:"required"`
	JumlahKamar         uint    `json:"jumlahKamar"  validate:"required"`
	JumlahKamarYangLaku uint    `json:"jumlahKamarYangLaku"  validate:"required"`
}

type UpdateDto struct {
	Id                  uint    `json:"id"`
	Espt_Id             uint    `json:"espt_id"`
	GolonganKamar       string  `json:"golonganKamar"`
	Tarif               float32 `json:"tarif"`
	JumlahKamar         uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku uint    `json:"jumlahKamarYangLaku"`
}

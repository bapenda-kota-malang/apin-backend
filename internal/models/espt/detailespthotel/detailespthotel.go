package detailespthotel

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailEsptHotel struct {
	detailesptair.DetailEspt
	GolonganKamar       *string  `json:"golonganKamar" gorm:"size:20"`
	Tarif               *float32 `json:"tarif"`
	JumlahKamar         *uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint    `json:"jumlahKamarYangLaku"`
	KasRegister         bool     `json:"kasRegister"`
	Pembukuan           bool     `json:"pembukuan"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id             uuid.UUID
	GolonganKamar       *string  `json:"golonganKamar"`
	Tarif               *float32 `json:"tarif"`
	JumlahKamar         *uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint    `json:"jumlahKamarYangLaku"`
	KasRegister         bool     `json:"kasRegister" validate:"required"`
	Pembukuan           bool     `json:"pembukuan" validate:"required"`
}

type UpdateDto struct {
	Id                  uint     `json:"id"`
	GolonganKamar       *string  `json:"golonganKamar"`
	Tarif               *float32 `json:"tarif"`
	JumlahKamar         *uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint    `json:"jumlahKamarYangLaku"`
	KasRegister         bool     `json:"kasRegister"`
	Pembukuan           bool     `json:"pembukuan"`
}

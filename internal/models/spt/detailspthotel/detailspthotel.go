package detailspthotel

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptHotel struct {
	Id                  uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id              uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	GolonganKamar       *string   `json:"golonganKamar" gorm:"type:varchar(20)"`
	Tarif               *float64  `json:"tarif" gorm:"type:decimal"`
	JumlahKamar         *uint64   `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint64   `json:"jumlahKamarYangLaku"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id              uuid.UUID
	GolonganKamar       *string  `json:"golonganKamar"`
	Tarif               *float64 `json:"tarif"`
	JumlahKamar         *uint64  `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint64  `json:"jumlahKamarYangLaku"`
}

type UpdateDto struct {
	Id                  uint     `json:"id"`
	GolonganKamar       *string  `json:"golonganKamar"`
	Tarif               *float64 `json:"tarif"`
	JumlahKamar         *uint64  `json:"jumlahKamar"`
	JumlahKamarYangLaku *uint64  `json:"jumlahKamarYangLaku"`
}

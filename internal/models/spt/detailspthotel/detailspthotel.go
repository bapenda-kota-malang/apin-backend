package detailspthotel

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type DetailSptHotel struct {
	Id                  uint64           `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id              uuid.UUID        `json:"spt_id" gorm:"type:uuid"`
	GolonganKamar       *pq.StringArray  `json:"golonganKamar" gorm:"type:varchar[]"`
	Tarif               *pq.Float32Array `json:"tarif" gorm:"type:numeric[]"`
	JumlahKamar         *pq.Int32Array   `json:"jumlahKamar" gorm:"type:integer[]"`
	JumlahKamarYangLaku *pq.Int32Array   `json:"jumlahKamarYangLaku" gorm:"type:integer[]"`
	KasRegister         bool             `json:"kasRegister"`
	Pembukuan           bool             `json:"pembukuan"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id              uuid.UUID
	GolonganKamar       *pq.StringArray  `json:"golonganKamar"`
	Tarif               *pq.Float32Array `json:"tarif"`
	JumlahKamar         *pq.Int32Array   `json:"jumlahKamar"`
	JumlahKamarYangLaku *pq.Int32Array   `json:"jumlahKamarYangLaku"`
	KasRegister         bool             `json:"kasRegister" validate:"required"`
	Pembukuan           bool             `json:"pembukuan" validate:"required"`
}

type UpdateDto struct {
	Id                  uint             `json:"id"`
	GolonganKamar       *pq.StringArray  `json:"golonganKamar"`
	Tarif               *pq.Float32Array `json:"tarif"`
	JumlahKamar         *pq.Int32Array   `json:"jumlahKamar"`
	JumlahKamarYangLaku *pq.Int32Array   `json:"jumlahKamarYangLaku"`
	KasRegister         *bool            `json:"kasRegister"`
	Pembukuan           *bool            `json:"pembukuan"`
}

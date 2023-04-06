package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiHotel struct {
	Id              uint      `json:"id" gorm:"primaryKey"`
	Potensiop_Id    uuid.UUID `json:"potensiop_id" gorm:"type:uuid"`
	JenisFasilitas  string    `json:"jenisFasilitas"`
	JumlahFasilitas uint64    `json:"jumlahFasilitas"`
	TarifFasilitas  float64   `json:"tarifFasilitas"`
	Kapasitas       uint64    `json:"kapasitas"`
	gormhelper.DateModel
}

type CreateDtoDPHotel struct {
	Potensiop_Id    uuid.UUID `json:"-"`
	JenisFasilitas  string    `json:"jenisFasilitas"`
	JumlahFasilitas uint64    `json:"jumlahFasilitas"`
	TarifFasilitas  float64   `json:"tarifFasilitas"`
	Kapasitas       uint64    `json:"kapasitas"`
}

type CreateDtoHotel struct {
	CreateDto
	DetailPajakDtos []CreateDtoDPHotel `json:"detailPajaks"`
}

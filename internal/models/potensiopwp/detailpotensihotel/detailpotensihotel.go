package detailpotensihotel

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiHotel struct {
	Id              uint      `json:"id" gorm:"primaryKey"`
	Potensiop_Id    uuid.UUID `json:"potensiop_id" gorm:"type:uuid"`
	JenisFasilitas  string    `json:"jenisFasilitas"`
	JumlahFasilitas string    `json:"jumlahFasilitas"`
	TarifFasilitas  string    `json:"tarifFasilitas"`
	Kapasitas       string    `json:"kapasitas"`
	gormhelper.DateModel
}

type CreateDto struct {
	Potensiop_Id    uuid.UUID `json:"-"`
	JenisFasilitas  string    `json:"jenisFasilitas"`
	JumlahFasilitas string    `json:"jumlahFasilitas"`
	TarifFasilitas  string    `json:"tarifFasilitas"`
	Kapasitas       string    `json:"kapasitas"`
}

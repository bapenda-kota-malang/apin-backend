package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiHiburan struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Potensiop_Id   uuid.UUID `json:"potensiop_id" gorm:"type:uuid"`
	JenisFasilitas string    `json:"jenisFasilitas"`
	Jumlah         string    `json:"jumlah"`
	Tarif          string    `json:"tarif"`
	Kapasitas      int64     `json:"kapasitas"`
	gormhelper.DateModel
}

type CreateDtoDPHiburan struct {
	Potensiop_Id   uuid.UUID `json:"-"`
	JenisFasilitas string    `json:"jenisFasilitas"`
	Jumlah         string    `json:"jumlah"`
	Tarif          string    `json:"tarif"`
	Kapasitas      int64     `json:"kapasitas"`
}

type CreateDtoHiburan struct {
	CreateDto
	DetailPajakDtos CreateDtoDPHiburan `json:"detailPajaks"`
}

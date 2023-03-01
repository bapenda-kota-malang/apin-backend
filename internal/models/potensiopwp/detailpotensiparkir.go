package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiParkir struct {
	Id             uint                 `json:"id" gorm:"primaryKey"`
	Potensiop_Id   uuid.UUID            `json:"potensiop_id" gorm:"type:uuid"`
	JenisKendaraan types.JenisKendaraan `json:"jenisKendaraan"`
	JumlahPenuh    uint64               `json:"jumlahPenuh"`
	Tarif          float64              `json:"tarif"`
	gormhelper.DateModel
}

type CreateDtoDPParkir struct {
	Potensiop_Id   uuid.UUID            `json:"-"`
	JenisKendaraan types.JenisKendaraan `json:"jenisKendaraan"`
	JumlahPenuh    uint64               `json:"jumlahPenuh"`
	Tarif          float64              `json:"tarif"`
}

type CreateDtoParkir struct {
	CreateDto
	DetailPajakDtos []CreateDtoDPParkir `json:"detailPajaks"`
}

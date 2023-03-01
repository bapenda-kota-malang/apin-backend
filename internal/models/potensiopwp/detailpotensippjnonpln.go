package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiPPJNonPLN struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	Potensiop_Id  uuid.UUID `json:"potensiop_id" gorm:"type:uuid"`
	Klasifikasi   string    `json:"klasifikasi"`
	Jumlah        string    `json:"jumlah"`
	KapasitasDaya string    `json:"kapasitasDaya"`
	gormhelper.DateModel
}

type CreateDtoDPPPJNonPLN struct {
	Potensiop_Id  uuid.UUID `json:"-"`
	Klasifikasi   string    `json:"klasifikasi"`
	Jumlah        string    `json:"jumlah"`
	KapasitasDaya string    `json:"kapasitasDaya"`
}

type CreateDtoPPJNonPLN struct {
	CreateDto
	DetailPajakDtos []CreateDtoDPPPJNonPLN `json:"detailPajaks"`
}

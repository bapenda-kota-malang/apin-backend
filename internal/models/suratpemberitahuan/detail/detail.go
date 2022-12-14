package detail

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type SuratPemberitahuanDetail struct {
	Id                    uint64    `json:"id" gorm:"primaryKey"`
	SuratPemberitahuan_Id uuid.UUID `json:"suratPemberitahuan_id" gorm:"index;type:uuid"`
	Spt_Id                uuid.UUID `json:"spt_id" gorm:"index;type:uuid"`
	Denda                 float64   `json:"denda"`
	TelahDibayar          float64   `json:"telahDibayar"`
	SisaPajak             float64   `json:"sisaPajak"`
	gormhelper.DateModel
	Spt        *spt.Spt           `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	SspdDetail *[]sspd.SspdDetail `json:"sspdDetails,omitempty" gorm:"-"`
}

type CreateDto struct {
	SuratPemberitahuan_Id uuid.UUID `json:"suratPemberitahuan_id"`
	Spt_Id                uuid.UUID `json:"spt_id"`
	Denda                 float64   `json:"denda"`
	TelahDibayar          float64   `json:"telahDibayar"`
	SisaPajak             float64   `json:"sisaPajak"`
}

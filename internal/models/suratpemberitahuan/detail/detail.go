package detail

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type SuratPemberitahuanDetail struct {
	Id                    uint64    `json:"id" gorm:"primaryKey"`
	SuratPemberitahuan_Id uint64    `json:"suratPemberitahuan_id"`
	Spt_Id                uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	Spt                   *spt.Spt  `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	gormhelper.DateModel
}

type CreateDto struct {
	SuratPemberitahuan_Id uint64    `json:"suratPemberitahuan_id"`
	Spt_Id                uuid.UUID `json:"spt_id"`
}

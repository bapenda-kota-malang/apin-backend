package detailsptair

import (
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptAir struct {
	Id         uint64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id     uuid.UUID     `json:"spt_id" gorm:"type:uuid"`
	Peruntukan mt.Peruntukan `json:"peruntukan" gorm:"type:varchar(100)"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt" gorm:"type:varchar(20)"`
	Lokasi     *string       `json:"lokasi,omitempty" gorm:"size:50"`
	Pengenaan  *float64      `json:"pengenaan" gorm:"type:decimal"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id     uuid.UUID
	Peruntukan mt.Peruntukan `json:"peruntukan" validate:"required"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt" validate:"required"`
	Lokasi     *string       `json:"lokasi"`
	Pengenaan  float64       `json:"pengenaan" validate:"required"`
}

type UpdateDto struct {
	Id         uint          `json:"id"`
	Peruntukan mt.Peruntukan `json:"peruntukan"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt"`
	Lokasi     *string       `json:"lokasi"`
	Pengenaan  float64       `json:"pengenaan"`
}

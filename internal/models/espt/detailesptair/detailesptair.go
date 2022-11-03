package detailesptair

import (
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailEspt struct {
	Id      uint      `json:"id" gorm:"primarykey"`
	Espt_Id uuid.UUID `json:"espt_id"`
}

type DetailEsptAir struct {
	DetailEspt
	Peruntukan mt.Peruntukan `json:"peruntukan" gorm:"size:100"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt" gorm:"size:20"`
	Location   *string       `json:"location,omitempty" gorm:"size:50"`
	Pengenaan  float32       `json:"pengenaan"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id    uuid.UUID
	Peruntukan mt.Peruntukan `json:"peruntukan" validate:"required"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt"`
	Location   *string       `json:"location"`
	Pengenaan  float32       `json:"-"`
}

type UpdateDto struct {
	Id         uint          `json:"id"`
	Peruntukan mt.Peruntukan `json:"peruntukan"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt"`
	Lokasi     *string       `json:"lokasi"`
	Pengenaan  float32       `json:"pengenaan"`
}

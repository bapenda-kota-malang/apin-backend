package detailesptair

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEspt struct {
	Id      uint  `json:"id" gorm:"primarykey"`
	Espt_Id *uint `json:"espt_id"`
}

type DetailEsptAir struct {
	DetailEspt
	Peruntukan Peruntukan `json:"peruntukan" gorm:"size:100"`
	JenisAbt   JenisABT   `json:"jenisAbt" gorm:"20"`
	Pengenaan  float32    `json:"pengenaan"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id    uint
	Peruntukan Peruntukan `json:"peruntukan"  validate:"required"`
	JenisAbt   JenisABT   `json:"jenisAbt"  validate:"required"`
	Pengenaan  float32    `json:"pengenaan"  validate:"min=1"`
}

type UpdateDto struct {
	Id         uint       `json:"id"`
	Peruntukan Peruntukan `json:"peruntukan"`
	JenisAbt   JenisABT   `json:"jenisAbt"`
	Pengenaan  float32    `json:"pengenaan"`
}

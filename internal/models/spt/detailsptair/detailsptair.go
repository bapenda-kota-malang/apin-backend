package detailsptair

import (
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
)

type DetailSptAir struct {
	Id         uint64        `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id     uint64        `json:"spt_id"`
	Peruntukan mt.Peruntukan `json:"peruntukan" gorm:"type:varchar(100)"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt" gorm:"type:varchar(20)"`
	Pengenaan  *float64      `json:"pengenaan" gorm:"type:decimal"`
}

type CreateDto struct {
	Spt_Id     uint
	Peruntukan mt.Peruntukan `json:"peruntukan" validate:"required"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt" validate:"required"`
	Pengenaan  float64       `json:"pengenaan" validate:"required"`
}

type UpdateDto struct {
	Spt_Id     uint          `json:"spt_id"`
	Peruntukan mt.Peruntukan `json:"peruntukan"`
	JenisAbt   mt.JenisABT   `json:"jenisAbt"`
	Pengenaan  float64       `json:"pengenaan"`
}

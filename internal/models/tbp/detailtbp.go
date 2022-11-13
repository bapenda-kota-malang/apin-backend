package tbp

import (
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DetailTbp struct {
	Id     uint64     `json:"id" gorm:"primaryKey"`
	Tbp_Id *uint64    `json:"tbp_id"`
	Tbp    *Tbp       `json:"tbp" gorm:"foreignKey:Tbp_Id"`
	Spt_Id *uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	Spt    *ms.Spt    `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	// Teguran_Iduint64 `json:"teguran_id"`
	// Teguran *mt.Teguran `json:"teguran,omitempty" gorm:"foreignKey:Teguran_Id"`
	NominalBayar *float64 `json:"nominalBayar" gorm:"type:decimal"`
	// Sts_Id *uint64 `json:"sts_id"`
	// Sts *msts.Sts `json:"sts,omitempty" gorm:"foreignKey:Sts_Id"`
	AngsuranKe *int `json:"angsuranKe" gorm:"type:smallint"`
	Waktu_Detail_Tb *datatypes.Time `json:"waktu_detail_tb"`
}

type DetailTbpCreateDto struct {
	Tbp_Id *uint64    `json:"tbp_id"`
	Spt_Id *uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	// Teguran_Iduint64 `json:"teguran_id"`
	NominalBayar *float64 `json:"nominalBayar"`
	// Sts_Id *uint64 `json:"sts_id"`
	AngsuranKe *int `json:"angsuranKe"`
	Waktu_Detail_Tb *datatypes.Time `json:"waktu_detail_tb"`
}

type DetailTbpUpdateDto struct {
	NominalBayar *float64 `json:"nominalBayar"`
}

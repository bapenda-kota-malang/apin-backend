package tbp

import (
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type RincianTbp struct {
	Id     uint64     `json:"id" gorm:"primaryKey"`
	Tbp_Id *uint64    `json:"tbp_id"`
	Tbp    *Tbp       `json:"tbp" gorm:"foreignKey:Tbp_Id"`
	Spt_Id *uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	Spt    *ms.Spt    `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	// Teguran_Iduint64 `json:"teguran_id"`
	// Teguran *mt.Teguran `json:"teguran,omitempty" gorm:"foreignKey:Teguran_Id"`
	// Nominal      *float64 `json:"nominal" gorm:"type:decimal"`
	// Denda        *float64 `json:"denda" gorm:"type:decimal"`
	NominalBayar *float64 `json:"nominalBayar" gorm:"type:decimal"`
	// Sts_Id *uint64 `json:"sts_id"`
	// Sts *msts.Sts `json:"sts,omitempty" gorm:"foreignKey:Sts_Id"`
	AngsuranKe *int `json:"angsuranKe" gorm:"type:smallint"`
	// Bunga            *float64        `json:"bunga" gorm:"type:decimal"`
	Waktu_Rincian_Tb *datatypes.Time `json:"waktu_rincian_tb"`
}

type RincianTbpCreateDto struct {
	Tbp_Id *uint64    `json:"tbp_id"`
	Spt_Id *uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	// Teguran_Iduint64 `json:"teguran_id"`
	// Nominal      *float64 `json:"nominal"`
	// Denda        *float64 `json:"denda"`
	NominalBayar *float64 `json:"nominalBayar"`
	// Sts_Id *uint64 `json:"sts_id"`
	AngsuranKe *int `json:"angsuranKe"`
	// Bunga            *float64        `json:"bunga"`
	Waktu_Rincian_Tb *datatypes.Time `json:"waktu_rincian_tb"`
}

type RincianTbpUpdateDto struct {
	NominalBayar *float64 `json:"nominalBayar"`
}

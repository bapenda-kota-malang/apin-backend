package sspd

import (
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type SspdDetail struct {
	Id      uint64     `json:"id" gorm:"primaryKey"`
	Sspd_Id *uint64    `json:"sspd_id"`
	Sspd    *Sspd      `json:"sspd" gorm:"foreignKey:Sspd_Id"`
	Spt_Id  *uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	Spt     *ms.Spt    `json:"spt,omitempty" gorm:"foreignKey:Spt_Id"`
	// Teguran_Iduint64 `json:"teguran_id"`
	// Teguran *mt.Teguran `json:"teguran,omitempty" gorm:"foreignKey:Teguran_Id"`
	NominalBayar *float64 `json:"nominalBayar" gorm:"type:decimal"`
	// Sts_Id *uint64 `json:"sts_id"`
	// Sts *msts.Sts `json:"sts,omitempty" gorm:"foreignKey:Sts_Id"`
	AngsuranKe      *int            `json:"angsuranKe" gorm:"type:smallint"`
	WaktuSspdDetail *datatypes.Time `json:"waktuSspdDetail"`
}

type SspdDetailCreateDto struct {
	Sspd_Id *uint64    `json:"sspd_id"`
	Spt_Id  *uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	// Teguran_Iduint64 `json:"teguran_id"`
	NominalBayar *float64 `json:"nominalBayar"`
	// Sts_Id *uint64 `json:"sts_id"`
	AngsuranKe      *int            `json:"angsuranKe"`
	WaktuSspdDetail *datatypes.Time `json:"waktuSspdDetail"`
}

type SspdDetailUpdateDto struct {
	NominalBayar *float64 `json:"nominalBayar"`
}

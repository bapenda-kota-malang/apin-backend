package sts

import "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"

type StsDetail struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening" gorm:"foreignKey:Rekening_Id"`
	Nominal     *float64           `json:"nominal" gorm:"type:decimal"`
}

type StsDetailCreateDto struct {
	Rekening_Id *uint64  `json:"rekening_id"`
	Nominal     *float64 `json:"nominal"`
}

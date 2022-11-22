package sts

import "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"

type StsDetail struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	Sts_Id      *uint64            `json:"sts_id"`
	Sts         *Sts               `json:"sts,omitempty" gorm:"foreignKey:Sts_Id"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening" gorm:"foreignKey:Rekening_Id"`
	Nominal     *float64           `json:"nominal" gorm:"type:decimal"`
}

type StsDetailCreateDto struct {
	Sts_Id      uint64   `json:"sts_id"`
	Rekening_Id *uint64  `json:"rekening_id"`
	Nominal     *float64 `json:"nominal"`
}
type StsDetailUpdateDto struct {
	Id          *uint64  `json:"id"`
	Sts_Id      uint64   `json:"sts_id"`
	Rekening_Id *uint64  `json:"rekening_id"`
	Nominal     *float64 `json:"nominal"`
}

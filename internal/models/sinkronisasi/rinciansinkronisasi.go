package sinkronisasi

import (
	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RincianSinkronisasi struct {
	Id                 uint64         `json:"id" gorm:"primaryKey"`
	Npwpd_Id           *uint64        `json:"npwpd_id"`
	Npwpd              *mn.Npwpd      `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	Spt_No             *string        `json:"spt_no"`
	RincianTbp_No      *mt.RincianTbp `json:"rincianTbp_no,omitempty" gorm:"foreignKey:Spt_No"`
	Spt_Nominal        *string        `json:"spt_nominal"`
	RincianTbp_Nominal mt.RincianTbp  `json:"rincianTbp_nominal,omitempty" gorm:"foreignKey:Spt_Nominal;references:Nominal"`
	MasaPajak          *string        `json:"masaPajak"`
	NominalPembayaran  *string        `json:"nominalPembayaran"`
	IsSync             *bool          `json:"isSync"`
	Keterangan         *string        `json:"keterangan"`
	gh.DateModel
	Sinkronisasi_Id *uint64       `json:"sinkronisasi_id"`
	Sinkronisasi    *Sinkronisasi `json:"sinkronisasi,omitempty" gorm:"foreignKey:Sinkronisasi_Id"`
}

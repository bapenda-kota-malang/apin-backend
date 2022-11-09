package sinkronisasi

import (
	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/tbp"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RincianSinkronisasi struct {
	Id       uint64    `json:"id" gorm:"primaryKey"`
	Npwpd_Id *uint64   `json:"npwpd_id"`
	Npwpd    *mn.Npwpd `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	// Spt_No             *string        `json:"spt_no"`
	RincianTbp_Id *uint64        `json:"rincianTbp_id"`
	RincianTbp    *mt.RincianTbp `json:"rincianTbp_no,omitempty" gorm:"foreignKey:RincianTbp_Id"`
	// Spt_Nominal        *string        `json:"spt_nominal"`
	// RincianTbp_Nominal mt.RincianTbp  `json:"rincianTbp_nominal,omitempty" gorm:"foreignKey:Spt_Nominal;references:Nominal"`
	MasaPajak         *string `json:"masaPajak"`
	NominalPembayaran *string `json:"nominalPembayaran"`
	IsSync            *bool   `json:"isSync"`
	Keterangan        *string `json:"keterangan"`
	gh.DateModel
	Sinkronisasi_Id *uint64       `json:"sinkronisasi_id"`
	Sinkronisasi    *Sinkronisasi `json:"sinkronisasi,omitempty" gorm:"foreignKey:Sinkronisasi_Id"`
}

type RincianSinkronisasiCreateDto struct {
	Npwpd_Id *uint64 `json:"npwpd_id"`
	// Spt_No            *string `json:"spt_no"`
	// Spt_Nominal       *string `json:"spt_nominal"`
	RincianTbp_Id     *uint64 `json:"rincianTbp_Id"`
	MasaPajak         *string `json:"masaPajak"`
	NominalPembayaran *string `json:"nominalPembayaran"`
	IsSync            *bool   `json:"isSync"`
	Keterangan        *string `json:"keterangan"`
	gh.DateModel
	Sinkronisasi_Id *uint64 `json:"sinkronisasi_id"`
}

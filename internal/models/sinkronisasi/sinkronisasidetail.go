package sinkronisasi

import (
	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type SinkronisasiDetail struct {
	Id       uint64    `json:"id" gorm:"primaryKey"`
	Npwpd_Id *uint64   `json:"npwpd_id"`
	Npwpd    *mn.Npwpd `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	// Spt_No             *string        `json:"spt_no"`
	SspdDetail_Id *uint64        `json:"sspdDetail_id"`
	SspdDetail    *ms.SspdDetail `json:"sspdDetail,omitempty" gorm:"foreignKey:SspdDetail_Id"`
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

type SinkronisasiDetailCreateDto struct {
	Npwpd_Id *uint64 `json:"npwpd_id"`
	// Spt_No            *string `json:"spt_no"`
	// Spt_Nominal       *string `json:"spt_nominal"`
	SspdDetail_Id     *uint64 `json:"sspdDetail_Id"`
	MasaPajak         *string `json:"masaPajak"`
	NominalPembayaran *string `json:"nominalPembayaran"`
	IsSync            *bool   `json:"isSync"`
	Keterangan        *string `json:"keterangan"`
	gh.DateModel
	Sinkronisasi_Id *uint64 `json:"sinkronisasi_id"`
}

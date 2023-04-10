package sejarah

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type SejarahAnggotaObjekPajak struct {
	Id                       uint64        `json:"id" gorm:"primarykey"`
	Provinsi_Kode            *string       `json:"provinsi_kode" gorm:"type:char(2)"`
	Daerah_Kode              *string       `json:"daerah_kode" gorm:"type:char(2)"`
	Kecamatan_Kode           *string       `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kelurahan_Kode           *string       `json:"kelurahan_kode" gorm:"type:char(3)"`
	Blok_Kode                *string       `json:"blok_kode" gorm:"type:char(3)"`
	NoUrut                   *string       `json:"noUrut" gorm:"type:char(4)"`
	JenisOp                  *string       `json:"jenisOp" gorm:"type:char(1)"`
	Induk_Provinsi_Kode      *string       `json:"induk_provinsi_kode" gorm:"type:char(2)"`
	Induk_Daerah_Kode        *string       `json:"induk_daerah_kode" gorm:"type:char(2)"`
	Induk_Kecamatan_Kode     *string       `json:"induk_kecamatan_kode" gorm:"type:char(3)"`
	Induk_Kelurahan_Kode     *string       `json:"induk_kelurahan_kode" gorm:"type:char(3)"`
	Induk_Blok_Kode          *string       `json:"induk_blok_kode" gorm:"type:char(3)"`
	Induk_NoUrut             *string       `json:"induk_noUrut" gorm:"type:char(4)"`
	Induk_JenisOp            *string       `json:"induk_jenisOp" gorm:"type:char(1)"`
	Induk_Area_Kode          *string       `json:"induk_area_kode" gorm:"type:char(10)"`
	LuasBumiBeban            *int          `json:"luasBumiBeban"`
	LuasBangunanBeban        *int          `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int          `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int          `json:"nilaiSistemBangunanBeban"`
	Area_Kode                *string       `json:"area_kode" gorm:"char(10)"`
	Kelurahan                *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Area_Kode;references:Kode"`
	gh.DateModel
}

type SejarahOpAnggotaResponse struct {
	Nop                      *string `json:"nop"`
	LuasBumiBeban            *int    `json:"luasBumiBeban"`
	LuasBangunanBeban        *int    `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int    `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int    `json:"nilaiSistemBangunanBeban"`
}

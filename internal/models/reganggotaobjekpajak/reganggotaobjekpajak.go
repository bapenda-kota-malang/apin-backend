package reganggotaobjekpajak

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	i "github.com/bapenda-kota-malang/apin-backend/internal/models/regindukobjekpajak"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegAnggotaObjekPajak struct {
	Id                       uint64                `json:"id" gorm:"primarykey"`
	RegIndukObjekPajak_Id    *uint64               `json:"regIndukObjekPajak_id"`
	RegIndukObjekPajak       *i.RegIndukObjekPajak `json:"regIndukObjekPajak,omitempty" gorm:"foreignKey:RegIndukObjekPajak_Id;references:Id"`
	Provinsi_Kode            *string               `json:"provinsi_kode" gorm:"type:char(2)"`
	Daerah_Kode              *string               `json:"kota_kode" gorm:"type:char(2)"`
	Kecamatan_Kode           *string               `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kelurahan_Kode           *string               `json:"kelurahan_kode" gorm:"type:char(3)"`
	Blok_Kode                *string               `json:"blok_kode" gorm:"type:char(3)"`
	NoUrut                   *string               `json:"noUrut" gorm:"type:char(4)"`
	JenisOp                  *string               `json:"jenisOp" gorm:"type:char(1)"`
	LuasBumiBeban            *int                  `json:"luasBumiBeban"`
	LuasBangunanBeban        *int                  `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int                  `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int                  `json:"nilaiSistemBangunanBeban"`
	NjopBumiBeban            *int                  `json:"njopBumiBeban"`
	NjopBangunanBeban        *int                  `json:"njopBangunanBeban"`
	Area_Kode                *string               `json:"area_kode" gorm:"char(10)"`
	Kelurahan                *ad.Kelurahan         `json:"kelurahan,omitempty" gorm:"foreignKey:Area_Kode;references:Kode"`
	gh.DateModel
}

type CreateDto struct {
	RegIndukObjekPajak_Id    *uint64 `json:"regIndukObjekPajak_id"`
	Provinsi_Kode            *string `json:"provinsi_kode"`
	Daerah_Kode              *string `json:"daerah_kode"`
	Kecamatan_Kode           *string `json:"kecamatan_kode"`
	Kelurahan_Kode           *string `json:"kelurahan_kode"`
	Blok_Kode                *string `json:"blok_kode"`
	NoUrut                   *string `json:"noUrut"`
	JenisOp                  *string `json:"jenisOp"`
	LuasBumiBeban            *int    `json:"luasBumiBeban"`
	LuasBangunanBeban        *int    `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int    `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int    `json:"nilaiSistemBangunanBeban"`
	NjopBumiBeban            *int    `json:"njopBumiBeban"`
	NjopBangunanBeban        *int    `json:"njopBangunanBeban"`
	Area_Kode                *string `json:"area_kode"`
}

type UpdateDto struct {
	RegIndukObjekPajak_Id    *uint64 `json:"regIndukObjekPajak_id"`
	Provinsi_Kode            *string `json:"provinsi_kode"`
	Daerah_Kode              *string `json:"daerah_kode"`
	Kecamatan_Kode           *string `json:"kecamatan_kode"`
	Kelurahan_Kode           *string `json:"kelurahan_kode"`
	Blok_Kode                *string `json:"blok_kode"`
	NoUrut                   *string `json:"noUrut"`
	JenisOp                  *string `json:"jenisOp"`
	LuasBumiBeban            *int    `json:"luasBumiBeban"`
	LuasBangunanBeban        *int    `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int    `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int    `json:"nilaiSistemBangunanBeban"`
	NjopBumiBeban            *int    `json:"njopBumiBeban"`
	NjopBangunanBeban        *int    `json:"njopBangunanBeban"`
	Area_Kode                *string `json:"area_kode"`
}

type FilterDto struct {
	RegIndukObjekPajak_Id    *uint64 `json:"regIndukObjekPajak_id"`
	Provinsi_Kode            *string `json:"provinsi_kode"`
	Daerah_Kode              *string `json:"daerah_kode"`
	Kecamatan_Kode           *string `json:"kecamatan_kode"`
	Kelurahan_Kode           *string `json:"kelurahan_kode"`
	Blok_Kode                *string `json:"blok_kode"`
	NoUrut                   *string `json:"noUrut"`
	LuasBumiBeban            *int    `json:"luasBumiBeban"`
	LuasBangunanBeban        *int    `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int    `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int    `json:"nilaiSistemBangunanBeban"`
	NjopBumiBeban            *int    `json:"njopBumiBeban"`
	NjopBangunanBeban        *int    `json:"njopBangunanBeban"`
	Page                     int     `json:"page"`
	PageSize                 int     `json:"page_size"`
}

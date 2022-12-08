package anggotaobjekpajak

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	i "github.com/bapenda-kota-malang/apin-backend/internal/models/indukobjekpajak"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type AnggotaObjekPajak struct {
	Id                       uint64             `json:"id" gorm:"primarykey"`
	IndukObjekPajak_Id       *uint64            `json:"indukObjekPajak_id"`
	IndukObjekPajak          *i.IndukObjekPajak `json:"indukObjekPajak,omitempty" gorm:"foreignKey:IndukObjekPajak_Id;references:Id"`
	Provinsi_Kode            *string            `json:"provinsi_kode" gorm:"type:char(2)"`
	Kota_Kode                *string            `json:"kota_kode" gorm:"type:char(2)"`
	Kecamatan_Kode           *string            `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kelurahan_Kode           *string            `json:"kelurahan_kode" gorm:"type:char(3)"`
	Blok_Kode                *string            `json:"blok_kode" gorm:"type:char(3)"`
	NoUrut                   *string            `json:"noUrut" gorm:"type:char(4)"`
	JenisOp                  *string            `json:"jenisOp" gorm:"type:char(1)"`
	LuasBumiBeban            *int               `json:"luasBumiBeban"`
	LuasBangunanBeban        *int               `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     *int               `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban *int               `json:"nilaiSistemBangunanBeban"`
	NjopBumiBeban            *int               `json:"njopBumiBeban"`
	NjopBangunanBeban        *int               `json:"njopBangunanBeban"`
	Area_Kode                *string            `json:"area_kode" gorm:"char(10)"`
	Kelurahan                *ad.Kelurahan      `json:"kelurahan,omitempty" gorm:"foreignKey:Area_Kode;references:Kode"`
	gh.DateModel
}

type CreateDto struct {
	IndukObjekPajak_Id       *uint64 `json:"indukObjekPajak_id"`
	Provinsi_Kode            *string `json:"provinsi_kode"`
	Kota_Kode                *string `json:"kota_kode"`
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
	IndukObjekPajak_Id       *uint64 `json:"indukObjekPajak_id"`
	Provinsi_Kode            *string `json:"provinsi_kode"`
	Kota_Kode                *string `json:"kota_kode"`
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
	IndukObjekPajak_Id       *uint64 `json:"indukObjekPajak_id"`
	Provinsi_Kode            *string `json:"provinsi_kode"`
	Kota_Kode                *string `json:"kota_kode"`
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

package anggotaobjekpajak

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	i "github.com/bapenda-kota-malang/apin-backend/internal/models/indukobjekpajak"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type AnggotaObjekPajak struct {
	Id                 uint64            `json:"id" gorm:"primarykey"`
	IndukObjekPajak_Id uint64            `json:"indukObjekPajak_id"`
	IndukObjekPajak    i.IndukObjekPajak `json:"indukObjekPajak,omitempty" gorm:"foreignKey:IndukObjekPajak_Id;references:Id"`
	Provinsi_Kode      string            `json:"provinsi_kode" gorm:"type:char(2)"`
	Provinsi           *ad.Provinsi      `json:"provinsi,omitempty" gorm:"foreignKey:Provinsi_Kode;references:Kode"`
	Kota_Kode          string            `json:"kota_kode" gorm:"type:char(2)"`
	Kota               *ad.Daerah        `json:"kota,omitempty" gorm:"foreignKey:Kota_Kode;references:Kode"`
	Kecamatan_Kode     string            `json:"kecamatan_kode" gorm:"type:char(3)"`
	Kecamatan          *ad.Kecamatan     `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Kode;references:Kode"`
	Kelurahan_Kode     string            `json:"kelurahan_kode" gorm:"type:char(3)"`
	Kelurahan          *ad.Kelurahan     `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
	// Blok_Kode      string        `json:"blok_kode" gorm:"type:char(3)"`
	// Blok         *ad.Blok      `json:"blok,omitempty" gorm:"foreignKey:Blok_Kode;references:Kode"`
	NoUrut                   string    `json:"noUrut" gorm:"type:char(4)"`
	JenisOp                  t.JenisOp `json:"jenisOp" gorm:"type:char(1)"`
	LuasBumiBeban            int       `json:"luasBumiBeban"`
	LuasBangunanBeban        int       `json:"luasBangunanBeban"`
	NilaiSistemBumiBeban     int       `json:"nilaiSistemBumiBeban"`
	NilaiSistemBangunanBeban int       `json:"nilaiSistemBangunanBeban"`
	NjopBumiBeban            int       `json:"njopBumiBeban"`
	NjopBangunanBeban        int       `json:"njopBangunanBeban"`
	gh.DateModel
}

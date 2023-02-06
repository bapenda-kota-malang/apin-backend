package penilaianpbb

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RangePenyusutan struct {
	Id                    uint64   `json:"id" gorm:"primaryKey"`
	Range_Penyusutan_Kode string   `json:"range_penyusutan_kode" gorm:"type:varchar(1)"`
	NilaiMinPenyusutan    *float64 `json:"nilaiMinPenyusutan"`
	NilaiMaxPenyusutan    *float64 `json:"nilaiMaxPenyusutan"`
	gh.DateModel
}

type Penyusutan struct {
	Id                    uint64  `json:"id" gorm:"primaryKey"`
	Range_Penyusutan_Kode string  `json:"range_penyusutan_kode" gorm:"type:varchar(1)"`
	KondisiBangunanSusut  *string `json:"kondisiBangunanSusut" gorm:"type:varchar(1)"`
	UmurEfektif           *int    `json:"nilaiMinPenyusutan"`
	NilaiPenyusutan       *int    `json:"nilaiMaxPenyusutan"`
	gh.DateModel
}

type Fasilitas struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Fasilitas_Kode string  `json:"fasilitas_kode" gorm:"type:varchar(2)"`
	Nama           *string `json:"nama" gorm:"type:varchar(20)"`
	Satuan         *string `json:"satuan" gorm:"type:varchar(2)"`
	Status         *string `json:"status" gorm:"type:varchar(1)"`
	Ketergantungan *string `json:"ketergantungan" gorm:"type:varchar(1)"`
	gh.DateModel
}

type FasNonDep struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode    string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	Thn_Non_Dep    *string `json:"thn_non_dep" gorm:"type:varchar(4)"`
	Fasilitas_Kode *string `json:"fasilitas_kode" gorm:"type:varchar(2)"`
	gh.DateModel
}

type FasDepMinMax struct {
	Id             uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  string   `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode    string   `json:"daerah_kode" gorm:"type:varchar(2)"`
	TahunDepMinMax *string  `json:"tahunDepMinMax" gorm:"type:varchar(4)"`
	Fasilitas_Kode *string  `json:"fasilitas_kode" gorm:"type:varchar(2)"`
	KlsDepMin      *int     `json:"klsDepMin"`
	KlsDepMax      *int     `json:"klsDepMax"`
	NilaiDepMinMax *float64 `json:"nilaiDepMinMax"`
	gh.DateModel
}

type FasDepJpbKlsBintang struct {
	Id                       uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode            string   `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode              string   `json:"daerah_kode" gorm:"type:varchar(2)"`
	TahunDep                 *string  `json:"tahunDep" gorm:"type:varchar(4)"`
	Fasilitas_Kode           *string  `json:"fasilitas_kode" gorm:"type:varchar(2)"`
	Jpb_Kode                 *string  `json:"jpb_kode" gorm:"type:varchar(2)"`
	KlsBintang               *string  `json:"klsBintang" gorm:"type:varchar(4)"`
	NilaiFasilitasKlsBintang *float64 `json:"nilaiFasilitasKlsBintang"`
	gh.DateModel
}

type DbkbMaterial struct {
	Id                uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode     string   `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode       string   `json:"daerah_kode" gorm:"type:varchar(2)"`
	Thn_Dbkb_Material *string  `json:"thn_dbkb_material" gorm:"type:varchar(4)"`
	Pekerjaan_Kode    *string  `json:"pekerjaan_kode" gorm:"type:varchar(2)"`
	Kegiatan_Kode     *string  `json:"kegiatan_kode" gorm:"type:varchar(2)"`
	NilaiDbkbMaterial *float64 `json:"nilaiDbkbMaterial"`
	gh.DateModel
}

type BangunanLantai struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	Jpb_Kode        string  `json:"jpb_kode" gorm:"type:varchar(2)"`
	Tipe_Bng        string  `json:"tipe_bng" gorm:"type:varchar(5)"`
	Bng_Lantai_Kode *string `json:"bng_lantai_kode" gorm:"type:varchar(8)"`
	LantaiMinBng    *int    `json:"LantaiMinBng"`
	LantaiMaxBng    *int    `json:"LantaiMaxBng"`
	gh.DateModel
}

type TipeBangunan struct {
	Id                uint64   `json:"id" gorm:"primaryKey"`
	Bng_Tipe          string   `json:"bng_tipe" gorm:"type:varchar(5)"`
	NamaTipe          *string  `json:"namaTipe" gorm:"type:varchar(20)"`
	LuasMinTipe       *int     `json:"luasMinTipe"`
	LuasMaxTipe       *int     `json:"luasMaxTipe"`
	FaktorPembagiTipe *float64 `json:"faktorPembagiTipe"`
	gh.DateModel
}

type KayuUlin struct {
	Id                uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode     string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode       string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	ThnStatusKayuUlin *string `json:"thnStatusKayuUlin" gorm:"type:varchar(4)"`
	StatusKayuUlin    int     `json:"statusKayuUlin"`
	gh.DateModel
}

type DbkbStandard struct {
	Id                uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode     string   `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode       string   `json:"daerah_kode" gorm:"type:varchar(2)"`
	ThnDbkbStandard   *string  `json:"thnDbkbStandard" gorm:"type:varchar(4)"`
	Jpb_Kode          *string  `json:"jpb_kode" gorm:"type:varchar(2)"`
	TipeBng           *string  `json:"tipeBng" gorm:"type:varchar(5)"`
	Bng_Lantai_Kode   *string  `json:"bng_lantai_kode" gorm:"type:varchar(8)"`
	NilaiDbkbStandard *float64 `json:"nilaiDbkbStandard"`
	gh.DateModel
}

type DbkbDayaDukung struct {
	Id                  uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode       string   `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode         string   `json:"daerah_kode" gorm:"type:varchar(2)"`
	ThnDbkbDayaDukung   *string  `json:"thnDbkbDayaDukung" gorm:"type:varchar(4)"`
	TipeKonstruksi      *string  `json:"tipeKonstruksi" gorm:"type:varchar(1)"`
	NilaiDbkbDayaDukung *float64 `json:"nilaiDbkbDayaDukung"`
	gh.DateModel
}

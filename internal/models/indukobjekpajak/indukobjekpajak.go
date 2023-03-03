package indukobjekpajak

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type IndukObjekPajak struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Provinsi_Kode  string `json:"provinsi_kode" gorm:"type:varchar(15)"`
	Dati2_Kode     string `json:"dati2_kode" gorm:"type:varchar(15)"`
	Kecamatan_Kode string `json:"kecamatan_kode" gorm:"type:varchar(15)"`
	Kelurahan_Kode string `json:"kelurahan_kode" gorm:"type:varchar(15)"`
	Blok_Kode      string `json:"blok_kode" gorm:"type:varchar(15)"`
	NoUrut         string `json:"no_urut" gorm:"type:varchar(15)"`
	JenisOP_Kode   string `json:"jenis_op_kode" gorm:"type:varchar(15)"`
}

type CreateDto struct {
	Provinsi_Kode  string `json:"provinsi_kode"`
	Dati2_Kode     string `json:"dati2_kode"`
	Kecamatan_Kode string `json:"kecamatan_kode"`
	Kelurahan_Kode string `json:"kelurahan_kode"`
	Blok_Kode      string `json:"blok_kode"`
	NoUrut         string `json:"no_urut"`
	JenisOP_Kode   string `json:"jenis_op_kode"`
}

type UpdateDto struct {
	Id             *uint64 `json:"id"`
	Provinsi_Kode  string  `json:"provinsi_kode"`
	Dati2_Kode     string  `json:"dati2_kode"`
	Kecamatan_Kode string  `json:"kecamatan_kode"`
	Kelurahan_Kode string  `json:"kelurahan_kode"`
	Blok_Kode      string  `json:"blok_kode"`
	NoUrut         string  `json:"no_urut"`
	JenisOP_Kode   string  `json:"jenis_op_kode"`
}

type FilterDto struct {
	Provinsi_Kode  string `json:"provinsi_kode"`
	Dati2_Kode     string `json:"dati2_kode"`
	Kecamatan_Kode string `json:"kecamatan_kode"`
	Kelurahan_Kode string `json:"kelurahan_kode"`
	Blok_Kode      string `json:"blok_kode"`
	NoUrut         string `json:"no_urut"`
	JenisOP_Kode   string `json:"jenis_op_kode"`
	Page           int    `json:"page"`
	PageSize       int    `json:"page_size"`
}

type ListDto struct {
	AlamatObjekBersama *string         `json:"alamat_objek_bersama"`
	Kelurahan          string          `json:"kelurahan"`
	LuasBangunan       *int            `json:"luas_bangunan"`
	LuasBumi           *int            `json:"luas_bumi"`
	Details            []ListDetailDto `json:"details"`
}

type ListDetailDto struct {
	NOPAnggota        string  `json:"nop_anggota"`
	NamaWajibPajak    *string `json:"nama_wajib_pajak"`
	LuasBumiBeban     *int    `json:"luas_bumi_beban"`
	LuasBangunanBeban *int    `json:"luas_bangunan_beban"`
	NJOP              *int    `json:"njop"`
}

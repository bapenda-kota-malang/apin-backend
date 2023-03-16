package nilaiindividu

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type NilaiIndividu struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	nop.NopDetail
	NoBangunan             *int            `json:"noBangunan" gorm:"type:integer"`
	NoFormulir             *string         `json:"noFormulir" gorm:"type:varchar(11)"`
	NilaiIndividu          *float32        `json:"nilaiIndividu" gorm:"type:bigint"`
	TglPenilaianIndividu   *datatypes.Date `json:"tglPenilaianIndividu"`
	NipPenilaianIndividu   *string         `json:"nipPenilaianIndividu" gorm:"type:varchar(30)"`
	TglPemeriksaanIndividu *datatypes.Date `json:"tglPemeriksaanIndividu"`
	NipPemeriksaanIndividu *string         `json:"nipPemeriksaanIndividu" gorm:"type:varchar(30)"`
	NipCreate              *string         `json:"nipCreate" gorm:"type:varchar(30)"`
	gh.DateModel
}

type CreateDto struct {
	nop.NopDetail
	NoBangunan             *int            `json:"noBangunan"`
	NoFormulir             *string         `json:"noFormulir"`
	NilaiIndividu          *float32        `json:"nilaiIndividu"`
	TglPenilaianIndividu   *datatypes.Date `json:"tglPenilaianIndividu"`
	NipPenilaianIndividu   *string         `json:"nipPenilaianIndividu"`
	TglPemeriksaanIndividu *datatypes.Date `json:"tglPemeriksaanIndividu"`
	NipPemeriksaanIndividu *string         `json:"nipPemeriksaanIndividu"`
	NipCreate              *string         `json:"nipCreate"`
}

type UpdateDto struct {
	nop.NopDetail
	NoBangunan             *int            `json:"noBangunan"`
	NoFormulir             *string         `json:"noFormulir"`
	NilaiIndividu          *float32        `json:"nilaiIndividu"`
	TglPenilaianIndividu   *datatypes.Date `json:"tglPenilaianIndividu"`
	NipPenilaianIndividu   *string         `json:"nipPenilaianIndividu"`
	TglPemeriksaanIndividu *datatypes.Date `json:"tglPemeriksaanIndividu"`
	NipPemeriksaanIndividu *string         `json:"nipPemeriksaanIndividu"`
	NipCreate              *string         `json:"nipCreate"`
}

type FilterDto struct {
	Provinsi_Kode          *string         `json:"provinsi_kode"`
	Daerah_Kode            *string         `json:"daerah_kode"`
	Kecamatan_Kode         *string         `json:"kecamatan_kode"`
	Kelurahan_Kode         *string         `json:"kelurahan_kode"`
	Blok_Kode              *string         `json:"blok_kode"`
	NoUrut                 *string         `json:"noUrut"`
	JenisOp                *string         `json:"jenisOp"`
	Area_Kode              *string         `json:"area_kode"`
	NoBangunan             *int            `json:"noBangunan"`
	NoFormulir             *string         `json:"noFormulir"`
	NilaiIndividu          *float32        `json:"nilaiIndividu"`
	TglPenilaianIndividu   *datatypes.Date `json:"tglPenilaianIndividu"`
	NipPenilaianIndividu   *string         `json:"nipPenilaianIndividu"`
	TglPemeriksaanIndividu *datatypes.Date `json:"tglPemeriksaanIndividu"`
	NipPemeriksaanIndividu *string         `json:"nipPemeriksaanIndividu"`
	NipCreate              *string         `json:"nipCreate"`
	BlokNOP                *string         `json:"blok_nop"`
	Page                   int             `json:"page"`
	PageSize               int             `json:"page_size"`
}

// response
type ListResponse struct {
	BlokNOP         string   `json:"blok_nop"`
	LetakObjekPajak *string  `json:"letak_objek_pajak"`
	NamaWP          *string  `json:"nama_wp"`
	NoBng           *int     `json:"no_bng"`
	NilaiSistem     *int     `json:"nilai_sistem"`
	NilaiIndividu   *float32 `json:"nilai_individu"`
}

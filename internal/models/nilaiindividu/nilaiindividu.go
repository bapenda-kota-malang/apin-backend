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
	nop.NopDetail
	NoBangunan             *int            `json:"noBangunan"`
	NoFormulir             *string         `json:"noFormulir"`
	NilaiIndividu          *float32        `json:"nilaiIndividu"`
	TglPenilaianIndividu   *datatypes.Date `json:"tglPenilaianIndividu"`
	NipPenilaianIndividu   *string         `json:"nipPenilaianIndividu"`
	TglPemeriksaanIndividu *datatypes.Date `json:"tglPemeriksaanIndividu"`
	NipPemeriksaanIndividu *string         `json:"nipPemeriksaanIndividu"`
	NipCreate              *string         `json:"nipCreate"`
	Page                   int             `json:"page"`
	PageSize               int             `json:"page_size"`
}

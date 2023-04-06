package sejarah

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SejarahNilaiIndividu struct {
	nop.NopDetail
	NilaiIndividu_Id       *uint64         `json:"nilaiIndividu_id"`
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

type ListSejarahNilaiIndividu struct {
	NoBangunan             *int            `json:"noBangunan"`
	NoFormulir             *string         `json:"noFormulir"`
	NilaiIndividu          *float32        `json:"nilaiIndividu"`
	TglPenilaianIndividu   *datatypes.Date `json:"tglPenilaianIndividu"`
	NamaPenilai            *string         `json:"namaPenilai"`
	TglPemeriksaanIndividu *datatypes.Date `json:"tglPemeriksaanIndividu"`
	NamaPerekam            *string         `json:"namaPerekam"`
}

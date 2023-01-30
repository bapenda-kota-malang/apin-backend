package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type RegPembetulanSpptSKPSTP struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	PermohonanId    *uint64 `json:"permohonanId" gorm:"type:integer"`
	KanwilId        *string `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId         *string `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan  *string `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan *string `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan *string `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanNOP
	TahunPembetulan *datatypes.Date `json:"TahunPembetulan" gorm:"type:varchar(4)"`
	JnsSK           *string         `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK            *string         `json:"noSK" gorm:"type:varchar(30)"`
	NamaWP          *string         `json:"jnsKeputusan" gorm:"type:varchar(30)"`
	JalanWP         *string         `json:"klsTanah_Kode" gorm:"type:varchar(30)"`
	BlokKavNoWP     *string         `json:"klsTanah_TahunAwal" gorm:"type:varchar(15)"`
	LuasBumi        *int            `json:"luasBumi" gorm:"type:INTEGER"`
	LuasBangunan    *int            `json:"luasBangunan" gorm:"type:INTEGER"`
	NjopBumi        *int32          `json:"njopBumi" gorm:"type:BIGINT"`
	NjopBangunan    *int32          `json:"njopBangunan" gorm:"type:BIGINT"`
	PBB             *int32          `json:"pbb" gorm:"type:BIGINT"`
	JenisSurat      *string         `json:"jenisSurat" gorm:"type:varchar(2)"`
	NoSurat         *string         `json:"noSurat" gorm:"type:varchar(30)"`
	TanggalSurat    *datatypes.Date `json:"TanggalSurat"`
	gormhelper.DateModel
}

func (i RegPstPermohonan) SetPembetulanSpptSKPSTP(nop PermohonanNOP) *RegPembetulanSpptSKPSTP {
	return &RegPembetulanSpptSKPSTP{
		PermohonanId:    &i.Id,
		KanwilId:        i.KanwilId,
		KppbbId:         i.KppbbId,
		TahunPelayanan:  i.TahunPelayanan,
		BundelPelayanan: i.BundelPelayanan,
		NoUrutPelayanan: i.NoUrutPelayanan,
		PermohonanNOP:   nop,
	}
}

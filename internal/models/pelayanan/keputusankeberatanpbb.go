package pstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type KeputusanKeberatanPbb struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	PermohonanId    *uint64 `json:"permohonanId" gorm:"type:integer"`
	KanwilId        *string `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId         *string `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan  *string `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan *string `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan *string `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanNOP
	JnsSK                 *string         `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK                  *string         `json:"noSK" gorm:"type:varchar(30)"`
	JnsKeputusan          *string         `json:"jnsKeputusan" gorm:"type:varchar(1)"`
	KlsTanah_Kode         *string         `json:"klsTanah_Kode" gorm:"type:varchar(3)"`
	KlsTanah_TahunAwal    *string         `json:"klsTanah_TahunAwal" gorm:"type:varchar(4)"`
	KlsBangunan_Kode      *string         `json:"klsBangunan_Kode" gorm:"type:varchar(3)"`
	KlsBangunan_TahunAwal *string         `json:"klsBangunan_TahunAwal" gorm:"type:varchar(4)"`
	LuasBumi              *int            `json:"luasBumi" gorm:"type:INTEGER"`
	LuasBangunan          *int            `json:"luasBangunan" gorm:"type:INTEGER"`
	NjopBumi              *int32          `json:"njopBumi" gorm:"type:BIGINT"`
	NjopBangunan          *int32          `json:"njopBangunan" gorm:"type:BIGINT"`
	PBB                   *int32          `json:"pbb" gorm:"type:BIGINT"`
	TanggalCetak          *datatypes.Date `json:"tanggalCetak"`
	NIPPencetak           *string         `json:"nipPencetak" gorm:"type:varchar(9)"`
	gormhelper.DateModel
}

func (i PstPermohonan) SetKeputusanKeberatanPbb(nop PermohonanNOP) *KeputusanKeberatanPbb {
	return &KeputusanKeberatanPbb{
		PermohonanId:    &i.Id,
		KanwilId:        i.KanwilId,
		KppbbId:         i.KppbbId,
		TahunPelayanan:  i.TahunPelayanan,
		BundelPelayanan: i.BundelPelayanan,
		NoUrutPelayanan: i.NoUrutPelayanan,
		PermohonanNOP:   nop,
	}
}

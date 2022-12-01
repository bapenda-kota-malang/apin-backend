package sppt

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type SpptObjekBersama struct {
	Id                     uint64  `json:"id" gorm:"primaryKey"`
	Blok_Id                *string `json:"blok_Id" gorm:"type:varchar(3)"`
	Dati2_Id               *string `json:"dati2_Id" gorm:"type:varchar(2)"`
	JenisOP_Id             *string `json:"jenisOP_Id" gorm:"type:varchar(1)"`
	KanwilBank_Id          *string `json:"kanwilBank_Id" gorm:"type:varchar(2)"`
	Kecamatan_Id           *string `json:"kecamatan_Id" gorm:"type:varchar(3)"`
	Keluarahan_Id          *string `json:"keluarahan_Id" gorm:"type:varchar(3)"`
	KelasBangunan_Id       *string `json:"kelasBangunan_Id" gorm:"type:varchar(3)"`
	KelasTanah_Id          *string `json:"kelasTanah_Id" gorm:"type:varchar(3)"`
	Propinsi_Id            *string `json:"propinsi_Id" gorm:"type:varchar(2)"`
	LuasBangunanBeban_sppt *int    `json:"luasBangunanBeban_sppt" gorm:"type:number(12)"`
	LuasBumiBeban_sppt     *int    `json:"luasBumiBeban_sppt" gorm:"type:number(12)"`
	NJOPBangunanBeban_sppt *int    `json:"njopBangunanBeban_sppt" gorm:"type:number(15)"`
	NJOPBumiBeban_sppt     *int    `json:"njopBumiBeban_sppt" gorm:"type:number(15)"`
	NoUrut                 *string `json:"noUrut" gorm:"type:varchar(4)"`
	TahunAwalKelasBangunan *string `json:"tahunAwalKelasBangunan" gorm:"type:varchar(30)"`
	TahunAwalKelasTanah    *string `json:"tahunAwalKelasTanah" gorm:"type:varchar(30)"`
	TahunPajakskp_sppt     *string `json:"tahunPajakskp_sppt" gorm:"type:varchar(30)"`
	gormhelper.DateModel
}

type RequestSpptObjekBersamaDto struct {
	Blok_Id                *string `json:"blok_Id"`
	Dati2_Id               *string `json:"dati2_Id"`
	JenisOP_Id             *string `json:"jenisOP_Id"`
	KanwilBank_Id          *string `json:"kanwilBank_Id"`
	Kecamatan_Id           *string `json:"kecamatan_Id"`
	Keluarahan_Id          *string `json:"keluarahan_Id"`
	KelasBangunan_Id       *string `json:"kelasBangunan_Id"`
	KelasTanah_Id          *string `json:"kelasTanah_Id"`
	Propinsi_Id            *string `json:"propinsi_Id"`
	LuasBangunanBeban_sppt *int    `json:"luasBangunanBeban_sppt"`
	LuasBumiBeban_sppt     *int    `json:"luasBumiBeban_sppt"`
	NJOPBangunanBeban_sppt *int    `json:"njopBangunanBeban_sppt"`
	NJOPBumiBeban_sppt     *int    `json:"njopBumiBeban_sppt"`
	NoUrut                 *string `json:"noUrut"`
	TahunAwalKelasBangunan *string `json:"tahunAwalKelasBangunan"`
	TahunAwalKelasTanah    *string `json:"tahunAwalKelasTanah"`
	TahunPajakskp_sppt     *string `json:"tahunPajakskp_sppt"`
}

type FilterSpptObjekBersamaDto struct {
	Blok_Id                *string `json:"blok_Id"`
	Dati2_Id               *string `json:"dati2_Id"`
	JenisOP_Id             *string `json:"jenisOP_Id"`
	KanwilBank_Id          *string `json:"kanwilBank_Id"`
	Kecamatan_Id           *string `json:"kecamatan_Id"`
	Keluarahan_Id          *string `json:"keluarahan_Id"`
	KelasBangunan_Id       *string `json:"kelasBangunan_Id"`
	KelasTanah_Id          *string `json:"kelasTanah_Id"`
	Propinsi_Id            *string `json:"propinsi_Id"`
	LuasBangunanBeban_sppt *int    `json:"luasBangunanBeban_sppt"`
	LuasBumiBeban_sppt     *int    `json:"luasBumiBeban_sppt"`
	NJOPBangunanBeban_sppt *int    `json:"njopBangunanBeban_sppt"`
	NJOPBumiBeban_sppt     *int    `json:"njopBumiBeban_sppt"`
	NoUrut                 *string `json:"noUrut"`
	TahunAwalKelasBangunan *string `json:"tahunAwalKelasBangunan"`
	TahunAwalKelasTanah    *string `json:"tahunAwalKelasTanah"`
	TahunPajakskp_sppt     *string `json:"tahunPajakskp_sppt"`
	Page                   int     `json:"page"`
	PageSize               int     `json:"page_size"`
}

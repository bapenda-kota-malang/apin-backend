package keberatan

import (
	"time"
)

type PembetulanKeberatan struct {
	Id                          uint64     `json:"id" gorm:"primaryKey"`
	BundelPelayanan             *string    `json:"bundelPelayanan" gorm:"type:char(4)"`
	BundelPelayananKepKeberatan *string    `json:"bundelPelayananKepKeberatan" gorm:"type:char(4)"`
	JenisSK                     *string    `json:"jenisSK" gorm:"type:char(1)"`
	BlokPemohon_Kd              *string    `json:"blokPemohon_Kd"`
	DaerahPemohon_Kd            *string    `json:"daerahPemohon_Kd"`
	JenisOpPemohon_Kd           *string    `json:"jenisOpPemohon_Kd"`
	Kanwil_Kd                   *string    `json:"kanwil_Kd"`
	KecamatanPemohon_Kd         *string    `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd         *string    `json:"kelurahanPemohon_Kd"`
	KelasBangunan_Kd            *string    `json:"kelasBangunan"`
	KelasTanah_Kd               *string    `json:"kelasTanah"`
	Kppbb_Kd                    *string    `json:"kppbb_Kd"`
	ProvinsiPemohon_Kd          *string    `json:"provinsiPemohon_Kd"`
	LuasBangunanPembetulan      *float64   `json:"luasBangunanPembetulan"`
	LuasBumiPembetulan          *float64   `json:"luasBumiPembetulan"`
	NipPencetakPembetulan       *string    `json:"nipPencetakPembetulan"`
	NjopBangunanPembetulan      *float64   `json:"njopBangunanPembetulan"`
	NjopBumiPembetulan          *float64   `json:"njopBumiPembetulan"`
	NoSk                        *string    `json:"noSk" gorm:"type:char(30)"`
	NoUrutPelayanan             *string    `json:"noUrutPelayanan" gorm:"type:char(3)"`
	NoUrutPelayananKepKeberatan *string    `json:"noUrutPelayananKepKeberatan" gorm:"type:char(3)"`
	NoUrutPemohonan             *string    `json:"noUrutPemohonan" gorm:"type:char(4)"`
	PbbPembetulan               *float64   `json:"pbbPembetulan"`
	TanggalCetakPembetulan      *time.Time `json:"tanggalCetakPembetulan"`
	TahunAwalKelasBangunan      *string    `json:"tahunAwalKelasBangunan" gorm:"type:char(4)"`
	TahunAwalKelasTanah         *string    `json:"tahunAwalKelasTanah" gorm:"type:char(4)"`
	TahunPelayanan              *string    `json:"tahunPelayanan" gorm:"type:char(4)"`
	TahunPelayananKepKeberatan  *string    `json:"tahunPelayananKepKeberatan" gorm:"type:char(4)"`
}

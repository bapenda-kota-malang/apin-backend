package keberatan

import (
	"time"
)

type PembetulanKeberatan struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	// No Pelayanan
	Kanwil_Kd                   *string `json:"kanwil_Kd"`
	Kppbb_Kd                    *string `json:"kppbb_Kd"`
	TahunPelayanan              *string `json:"tahunPelayanan" gorm:"type:char(4)"`
	BundelPelayanan             *string `json:"bundelPelayanan" gorm:"type:char(4)"`
	NoUrutPelayanan             *string `json:"noUrutPelayanan" gorm:"type:char(3)"`
	TahunPelayananKepKeberatan  *string `json:"tahunPelayananKepKeberatan" gorm:"type:char(4)"`
	BundelPelayananKepKeberatan *string `json:"bundelPelayananKepKeberatan" gorm:"type:char(4)"`
	NoUrutPelayananKepKeberatan *string `json:"noUrutPelayananKepKeberatan" gorm:"type:char(3)"`
	// nop
	ProvinsiPemohon_Kd     *string    `json:"provinsiPemohon_Kd"`
	DaerahPemohon_Kd       *string    `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd    *string    `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd    *string    `json:"kelurahanPemohon_Kd"`
	BlokPemohon_Kd         *string    `json:"blokPemohon_Kd"`
	NoUrutPemohonan        *string    `json:"noUrutPemohonan" gorm:"type:char(4)"`
	JenisOpPemohon_Kd      *string    `json:"jenisOpPemohon_Kd"`
	JenisSK                *string    `json:"jenisSK" gorm:"type:char(1)"`
	NoSk                   *string    `json:"noSk" gorm:"type:char(30)"`
	KelasBangunan_Kd       *string    `json:"kelasBangunan"`
	KelasTanah_Kd          *string    `json:"kelasTanah"`
	LuasBangunanPembetulan *float64   `json:"luasBangunanPembetulan"`
	LuasBumiPembetulan     *float64   `json:"luasBumiPembetulan"`
	NjopBangunanPembetulan *float64   `json:"njopBangunanPembetulan"`
	NjopBumiPembetulan     *float64   `json:"njopBumiPembetulan"`
	PbbPembetulan          *float64   `json:"pbbPembetulan"`
	TahunAwalKelasBangunan *string    `json:"tahunAwalKelasBangunan" gorm:"type:char(4)"`
	TahunAwalKelasTanah    *string    `json:"tahunAwalKelasTanah" gorm:"type:char(4)"`
	NipPencetakPembetulan  *string    `json:"nipPencetakPembetulan"`
	TanggalCetakPembetulan *time.Time `json:"tanggalCetakPembetulan"`
}

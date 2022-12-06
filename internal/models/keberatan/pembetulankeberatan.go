package keberatan

import (
	"time"

	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
)

type PembetulanKeberatan struct {
	Id                          uint64     `json:"id" gorm:"primaryKey"`
	BundelPelayanan             *string    `json:"bundelPelayanan" gorm:"type:char(4)"`
	BundelPelayananKepKeberatan *string    `json:"bundelPelayananKepKeberatan" gorm:"type:char(4)"`
	JenisSK                     *string    `json:"jenisSK" gorm:"type:char(1)"`
	BlokPemohon_Id              *string    `json:"blokPemohon_id"`
	DaerahPemohon_Id            *string    `json:"daerahPemohon_id"`
	JenisOpPemohon_Id           *string    `json:"jenisOpPemohon_id"`
	Kanwil_Id                   *string    `json:"kanwil_id"`
	KecamatanPemohon_Id         *string    `json:"kecamatanPemohon_id"`
	KelurahanPemohon_Id         *string    `json:"kelurahanPemohon_id"`
	KelasBangunan_Id            *string    `json:"kelasBangunan"`
	KelasTanah_Id               *string    `json:"kelasTanah"`
	Kppbb_Id                    *string    `json:"kppbb_id"`
	ProvinsiPemohon_Id          *string    `json:"provinsiPemohon_id"`
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
	Nop_Id                      *uint64    `json:"nop_id"`
	Nop                         *mn.Nop    `json:"nop,omitempty" gorm:"foreignKey:Nop_Id"`
}

package keberatan

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sksk"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type PembetulanKeberatan struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	// No Pelayanan
	Kanwil_Kd                   string  `json:"kanwil_Kd" gorm:"not null"`
	Kppbb_Kd                    string  `json:"kppbb_Kd" gorm:"not null"`
	TahunPelayanan              string  `json:"tahunPelayanan" gorm:"type:char(4);not null"`
	BundelPelayanan             string  `json:"bundelPelayanan" gorm:"type:char(4);not null"`
	NoUrutPelayanan             string  `json:"noUrutPelayanan" gorm:"type:char(3);not null"`
	TahunPelayananKepKeberatan  *string `json:"tahunPelayananKepKeberatan" gorm:"type:char(4)"`
	BundelPelayananKepKeberatan *string `json:"bundelPelayananKepKeberatan" gorm:"type:char(4)"`
	NoUrutPelayananKepKeberatan *string `json:"noUrutPelayananKepKeberatan" gorm:"type:char(3)"`
	// nop
	ProvinsiPemohon_Kd     string                  `json:"provinsiPemohon_Kd" gorm:"not null"`
	Provinsi               *areadivision.Provinsi  `json:"provinsi" gorm:"foreignKey:ProvinsiPemohon_Kd;references:Kode"`
	DaerahPemohon_Kd       string                  `json:"daerahPemohon_Kd" gorm:"not null"`
	Daerah                 *areadivision.Daerah    `json:"daerah" gorm:"foreignKey:DaerahPemohon_Kd;references:Kode"`
	KecamatanPemohon_Kd    string                  `json:"kecamatanPemohon_Kd" gorm:"not null"`
	Kecamatan              *areadivision.Kecamatan `json:"kecamatan" gorm:"foreignKey:KecamatanPemohon_Kd;references:Kode"`
	KelurahanPemohon_Kd    string                  `json:"kelurahanPemohon_Kd" gorm:"not null"`
	Kelurahan              *areadivision.Kelurahan `json:"kelurahan" gorm:"foreignKey:KelurahanPemohon_Kd;references:Kode"`
	BlokPemohon_Kd         string                  `json:"blokPemohon_Kd" gorm:"not null"`
	Blok                   *areadivision.Blok      `json:"blok" gorm:"foreignKey:BlokPemohon_Kd;references:Kode"`
	NoUrutPemohonan        string                  `json:"noUrutPemohonan" gorm:"type:char(4);not null"`
	JenisOpPemohon_Kd      string                  `json:"jenisOpPemohon_Kd" gorm:"not null"`
	JenisSK                *string                 `json:"jenisSK" gorm:"type:char(1)"`
	NoSk                   *string                 `json:"noSk" gorm:"type:char(30)"`
	SkSk                   *sksk.SkSk              `json:"sk" gorm:"foreignKey:NoSk;references:NoSK"`
	KelasTanah_Kd          string                  `json:"kelasTanah" gorm:"default:'XXX'"`
	TahunAwalKelasTanah    string                  `json:"tahunAwalKelasTanah" gorm:"type:char(4);default:'1986'"`
	KelasBangunan_Kd       string                  `json:"kelasBangunan" gorm:"default:'XXX'"`
	TahunAwalKelasBangunan string                  `json:"tahunAwalKelasBangunan" gorm:"type:char(4);default:'1986'"`
	LuasBumiPembetulan     float64                 `json:"luasBumiPembetulan"`
	LuasBangunanPembetulan float64                 `json:"luasBangunanPembetulan"`
	NjopBumiPembetulan     float64                 `json:"njopBumiPembetulan"`
	NjopBangunanPembetulan float64                 `json:"njopBangunanPembetulan"`
	PbbPembetulan          float64                 `json:"pbbPembetulan"`
	NipPencetakPembetulan  *string                 `json:"nipPencetakPembetulan"`
	TanggalCetakPembetulan datatypes.Date          `json:"tanggalCetakPembetulan" gorm:"default:current_timestamp"`
	gormhelper.DateModel
}

type InputPembetulanSkKeberatan struct {
	TahunPelayanan  string          `json:"tahunPelayanan" validate:"required"`
	BundelPelayanan string          `json:"bundelPelayanan" validate:"required"`
	NoUrutPelayanan string          `json:"noUrutPelayanan" validate:"required"`
	JenisSKLama     *string         `json:"jenisSKLama" validate:"required"`
	NoSkLama        *string         `json:"noSkLama" validate:"required"`
	JenisSKBaru     *string         `json:"jenisSKBaru" validate:"required"`
	NoSkBaru        *string         `json:"noSkBaru" validate:"required"`
	TglSK           *datatypes.Date `json:"tglSK" validate:"required"`
	NoBaKantor      *string         `json:"noBaKantor"`
	TglBaKantor     *datatypes.Date `json:"tglBaKantor"`
	NoBaLapangan    *string         `json:"noBaLapangan"`
	TglBaLapangan   *datatypes.Date `json:"tglBaLapangan"`
}

type FilterDtoPemKeb struct {
	Kanwil_Kd                   *string         `json:"kanwil_Kd"`
	Kppbb_Kd                    *string         `json:"kppbb_Kd"`
	TahunPelayanan              *string         `json:"tahunPelayanan"`
	BundelPelayanan             *string         `json:"bundelPelayanan"`
	NoUrutPelayanan             *string         `json:"noUrutPelayanan"`
	TahunPelayananKepKeberatan  *string         `json:"tahunPelayananKepKeberatan"`
	BundelPelayananKepKeberatan *string         `json:"bundelPelayananKepKeberatan"`
	NoUrutPelayananKepKeberatan *string         `json:"noUrutPelayananKepKeberatan"`
	ProvinsiPemohon_Kd          *string         `json:"provinsiPemohon_Kd"`
	DaerahPemohon_Kd            *string         `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd         *string         `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd         *string         `json:"kelurahanPemohon_Kd"`
	BlokPemohon_Kd              *string         `json:"blokPemohon_Kd"`
	NoUrutPemohonan             *string         `json:"noUrutPemohonan"`
	JenisOpPemohon_Kd           *string         `json:"jenisOpPemohon_Kd"`
	JenisSK                     *string         `json:"jenisSK"`
	NoSk                        *string         `json:"noSk"`
	KelasTanah_Kd               *string         `json:"kelasTanah"`
	TahunAwalKelasTanah         *string         `json:"tahunAwalKelasTanah"`
	KelasBangunan_Kd            *string         `json:"kelasBangunan"`
	TahunAwalKelasBangunan      *string         `json:"tahunAwalKelasBangunan"`
	LuasBumiPembetulan          *float64        `json:"luasBumiPembetulan"`
	LuasBangunanPembetulan      *float64        `json:"luasBangunanPembetulan"`
	NjopBumiPembetulan          *float64        `json:"njopBumiPembetulan"`
	NjopBangunanPembetulan      *float64        `json:"njopBangunanPembetulan"`
	PbbPembetulan               *float64        `json:"pbbPembetulan"`
	NipPencetakPembetulan       *string         `json:"nipPencetakPembetulan"`
	TanggalCetakPembetulan      *datatypes.Date `json:"tanggalCetakPembetulan"`
}

type UpdateDtoPemKeb struct {
	Kanwil_Kd                   *string         `json:"kanwil_Kd"`
	Kppbb_Kd                    *string         `json:"kppbb_Kd"`
	TahunPelayanan              *string         `json:"tahunPelayanan"`
	BundelPelayanan             *string         `json:"bundelPelayanan"`
	NoUrutPelayanan             *string         `json:"noUrutPelayanan"`
	TahunPelayananKepKeberatan  *string         `json:"tahunPelayananKepKeberatan"`
	BundelPelayananKepKeberatan *string         `json:"bundelPelayananKepKeberatan"`
	NoUrutPelayananKepKeberatan *string         `json:"noUrutPelayananKepKeberatan"`
	ProvinsiPemohon_Kd          *string         `json:"provinsiPemohon_Kd"`
	DaerahPemohon_Kd            *string         `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd         *string         `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd         *string         `json:"kelurahanPemohon_Kd"`
	BlokPemohon_Kd              *string         `json:"blokPemohon_Kd"`
	NoUrutPemohonan             *string         `json:"noUrutPemohonan"`
	JenisOpPemohon_Kd           *string         `json:"jenisOpPemohon_Kd"`
	JenisSK                     *string         `json:"jenisSK"`
	NoSk                        *string         `json:"noSk"`
	KelasTanah_Kd               *string         `json:"kelasTanah"`
	TahunAwalKelasTanah         *string         `json:"tahunAwalKelasTanah"`
	KelasBangunan_Kd            *string         `json:"kelasBangunan"`
	TahunAwalKelasBangunan      *string         `json:"tahunAwalKelasBangunan"`
	LuasBumiPembetulan          *float64        `json:"luasBumiPembetulan"`
	LuasBangunanPembetulan      *float64        `json:"luasBangunanPembetulan"`
	NjopBumiPembetulan          *float64        `json:"njopBumiPembetulan"`
	NjopBangunanPembetulan      *float64        `json:"njopBangunanPembetulan"`
	PbbPembetulan               *float64        `json:"pbbPembetulan"`
	NipPencetakPembetulan       *string         `json:"nipPencetakPembetulan"`
	TanggalCetakPembetulan      *datatypes.Date `json:"tanggalCetakPembetulan"`
}

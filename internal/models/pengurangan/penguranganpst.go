package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganPST struct {
	BaseModel
	ThnPengPST        *string            `json:"thnPengPST" gorm:"type:char(4)"`
	StatusSK          *StatusPengurangan `json:"statusSK" gorm:"type:char(1)"`
	PCTPenguranganPST *float64           `json:"PCTPenguranganPST" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type CreateDtoPST struct {
	Kanwil_Kode            *string            `json:"kanwil_kode"`
	KPPBB_Kode             *string            `json:"kppbb_kode"`
	ThnPelayanan           *string            `json:"thnPelayanan"`
	BundelPelayanan        *string            `json:"bundelPelayanan"`
	NoUrutPelayanan        *string            `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string            `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string            `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string            `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string            `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string            `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string            `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string            `json:"jenisOp_pemohon"`
	JenisSk                *string            `json:"jenisSk"`
	NoSk                   *string            `json:"noSk"`
	ThnPengPST             *string            `json:"thnPengPST"`
	StatusSK               *StatusPengurangan `json:"statusSK"`
	PCTPenguranganPST      *float64           `json:"PCTPenguranganPST"`
}

type FilterDtoPST struct {
	Kanwil_Kode            *string            `json:"kanwil_kode"`
	KPPBB_Kode             *string            `json:"kppbb_kode"`
	ThnPelayanan           *string            `json:"thnPelayanan"`
	BundelPelayanan        *string            `json:"bundelPelayanan"`
	NoUrutPelayanan        *string            `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string            `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string            `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string            `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string            `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string            `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string            `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string            `json:"jenisOp_pemohon"`
	JenisSk                *string            `json:"jenisSk"`
	NoSk                   *string            `json:"noSk"`
	ThnPengPST             *string            `json:"thnPengPST"`
	StatusSK               *StatusPengurangan `json:"statusSK"`
	PCTPenguranganPST      *float64           `json:"PCTPenguranganPST"`
	Page                   int                `json:"page"`
	PageSize               int                `json:"page_size"`
}

type UpdateDtoPST struct {
	Kanwil_Kd           *string            `json:"kanwil_Kd"`
	Kppbb_Kd            *string            `json:"kppbb_Kd"`
	TahunPelayanan      *string            `json:"tahunPelayanan"`
	BundelPelayanan     *string            `json:"bundelPelayanan"`
	NoUrutPelayanan     *string            `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd  *string            `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd    *string            `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd *string            `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd *string            `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd      *string            `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd    *string            `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd   *string            `json:"jenisOpPemohon_Kd"`
	JenisSk             *string            `json:"jenisSk"`
	NoSk                *string            `json:"noSk"`
	ThnPengPST          *string            `json:"thnPengPST"`
	StatusSK            *StatusPengurangan `json:"statusSK"`
	PCTPenguranganPST   *float64           `json:"PCTPenguranganPST"`
}

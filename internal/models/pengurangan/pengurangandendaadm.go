package pengurangan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PenguranganDendaADM struct {
	BaseModel
	TahunPengurangan *string            `json:"tahunPengurangan" gorm:"type:char(4)"`
	StatusPenguragan *StatusPengurangan `json:"statusPenguragan" gorm:"type:char(1)"`
	PCTPengurangan   *float64           `json:"pCTPengurangan" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type DataDendaAdmDto struct {
	DataBaseDto
	TahunPengurangan *string            `json:"tahunPengurangan"`
	StatusPenguragan *StatusPengurangan `json:"statusPenguragan"`
	PCTPengurangan   *float64           `json:"PCTPengurangan"`
}

type CreateDtoDendaADM struct {
	Kanwil_Kode     string            `json:"kanwil_kode"`
	Kppbb_Kode      string            `json:"kppbb_kode"`
	ThnPelayanan    string            `json:"thnPelayanan" validate:"required"`
	BundelPelayanan string            `json:"bundelPelayanan" validate:"required"`
	NoUrutPelayanan string            `json:"noUrutPelayanan" validate:"required"`
	SkSk            CreateSkskDto     `json:"sksk" validate:"required"`
	Data            []DataDendaAdmDto `json:"data" validate:"required"`
}

type FilterDtoDendaADM struct {
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
	TahunPengurangan    *string            `json:"tahunPengurangan"`
	StatusPenguragan    *StatusPengurangan `json:"statusPenguragan"`
	PCTPengurangan      *float64           `json:"pCTPengurangan"`
	Page                int                `json:"page"`
	PageSize            int                `json:"page_size"`
}

type UpdateDtoDendaADM struct {
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
	TahunPengurangan    *string            `json:"tahunPengurangan"`
	StatusPenguragan    *StatusPengurangan `json:"statusPenguragan"`
	PCTPengurangan      *float64           `json:"pCTPengurangan"`
}

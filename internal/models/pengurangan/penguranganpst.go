package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganPST struct {
	BaseModel
	ThnPengPST        *string            `json:"thnPengPST" gorm:"type:char(4)"`
	StatusSK          *StatusPengurangan `json:"statusSK" gorm:"type:char(1)"`
	PCTPenguranganPST *float64           `json:"PCTPenguranganPST" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type DataPstDto struct {
	DataBaseDto
	ThnPengPST        string            `json:"thnPengPST"`
	StatusSK          StatusPengurangan `json:"statusSK"`
	PCTPenguranganPST float64           `json:"PCTPenguranganPST"`
}

type CreateDtoPST struct {
	Kanwil_Kode     string        `json:"kanwil_kode"`
	Kppbb_Kode      string        `json:"kppbb_kode"`
	ThnPelayanan    string        `json:"thnPelayanan" validate:"required"`
	BundelPelayanan string        `json:"bundelPelayanan" validate:"required"`
	NoUrutPelayanan string        `json:"noUrutPelayanan" validate:"required"`
	SkSk            CreateSkskDto `json:"sksk" validate:"required"`
	Data            []DataPstDto  `json:"data" validate:"required"`
}

type FilterDtoPST struct {
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
	Page                int                `json:"page"`
	PageSize            int                `json:"page_size"`
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

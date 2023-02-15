package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganDendaADM struct {
	BaseModel
	TahunPenguranganDendaADM *string            `json:"tahunPenguranganDendaADM" gorm:"type:char(4)"`
	StatusPenguraganADM      *StatusPengurangan `json:"statusPenguraganADM" gorm:"type:char(1)"`
	PctPenguranganDendaADM   *float64           `json:"pctPenguranganDendaADM" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type CreateDtoDendaADM struct {
	Kanwil_Kd                *string            `json:"kanwil_Kd"`
	Kppbb_Kd                 *string            `json:"kppbb_Kd"`
	TahunPelayanan           *string            `json:"tahunPelayanan"`
	BundelPelayanan          *string            `json:"bundelPelayanan"`
	NoUrutPelayanan          *string            `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd       *string            `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd         *string            `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd      *string            `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd      *string            `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd           *string            `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd         *string            `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd        *string            `json:"jenisOpPemohon_Kd"`
	JenisSk                  *string            `json:"jenisSk"`
	NoSk                     *string            `json:"noSk"`
	TahunPenguranganDendaADM *string            `json:"tahunPenguranganDendaADM"`
	StatusPenguraganADM      *StatusPengurangan `json:"statusPenguraganADM"`
	PctPenguranganDendaADM   *float64           `json:"pctPenguranganDendaADM"`
}

type FilterDtoDendaADM struct {
	Kanwil_Kd                *string  `json:"kanwil_Kd"`
	Kppbb_Kd                 *string  `json:"kppbb_Kd"`
	TahunPelayanan           *string  `json:"tahunPelayanan"`
	BundelPelayanan          *string  `json:"bundelPelayanan"`
	NoUrutPelayanan          *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd       *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd         *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd      *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd      *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd           *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd         *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd        *string  `json:"jenisOpPemohon_Kd"`
	JenisSk                  *string  `json:"jenisSk"`
	NoSk                     *string  `json:"noSk"`
	TahunPenguranganDendaADM *string  `json:"tahunPenguranganDendaADM"`
	StatusPenguraganADM      *string  `json:"statusPenguraganADM"`
	PctPenguranganDendaADM   *float64 `json:"pctPenguranganDendaADM"`
	Page                     int      `json:"page"`
	PageSize                 int      `json:"page_size"`
}

type UpdateDtoDendaADM struct {
	Kanwil_Kd                *string  `json:"kanwil_Kd"`
	Kppbb_Kd                 *string  `json:"kppbb_Kd"`
	TahunPelayanan           *string  `json:"tahunPelayanan"`
	BundelPelayanan          *string  `json:"bundelPelayanan"`
	NoUrutPelayanan          *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd       *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd         *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd      *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd      *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd           *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd         *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd        *string  `json:"jenisOpPemohon_Kd"`
	JenisSk                  *string  `json:"jenisSk"`
	NoSk                     *string  `json:"noSk"`
	TahunPenguranganDendaADM *string  `json:"tahunPenguranganDendaADM"`
	StatusPenguraganADM      *string  `json:"statusPenguraganADM"`
	PctPenguranganDendaADM   *float64 `json:"pctPenguranganDendaADM"`
}

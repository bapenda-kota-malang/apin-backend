package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganPermanen struct {
	BaseModel
	TahunPenguranganPermanenAwal  *string            `json:"tahunPenguranganPermanenAwal" gorm:"type:char(4)"`
	TahunPenguranganPermanenAkhir *string            `json:"tahunPenguranganPermanenAkhir" gorm:"type:char(4)"`
	StatusSkPenguraganPermanen    *StatusPengurangan `json:"statusSkPenguranganPermanen" gorm:"type:char(1)"`
	PctPenguranganPermanen        *float64           `json:"pctPenguranganPermanen" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type CreateDtoPermanen struct {
	Kanwil_Kd                     *string  `json:"kanwil_Kd"`
	Kppbb_Kd                      *string  `json:"kppbb_Kd"`
	TahunPelayanan                *string  `json:"tahunPelayanan"`
	BundelPelayanan               *string  `json:"bundelPelayanan"`
	NoUrutPelayanan               *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd            *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd              *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd           *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd           *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd                *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd              *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd             *string  `json:"jenisOpPemohon_Kd"`
	JenisSk                       *string  `json:"jenisSk"`
	NoSk                          *string  `json:"noSk"`
	TahunPenguranganPermanenAwal  *string  `json:"tahunPenguranganPermanenAwal"`
	TahunPenguranganPermanenAkhir *string  `json:"tahunPenguranganPermanenAkhir"`
	StatusSkPenguraganPermanen    *string  `json:"statusSkPenguranganPermanen"`
	PctPenguranganPermanen        *float64 `json:"pctPenguranganPermanen"`
}

type FilterDtoPermanen struct {
	Kanwil_Kd                     *string  `json:"kanwil_Kd"`
	Kppbb_Kd                      *string  `json:"kppbb_Kd"`
	TahunPelayanan                *string  `json:"tahunPelayanan"`
	BundelPelayanan               *string  `json:"bundelPelayanan"`
	NoUrutPelayanan               *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd            *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd              *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd           *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd           *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd                *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd              *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd             *string  `json:"jenisOpPemohon_Kd"`
	JenisSk                       *string  `json:"jenisSk"`
	NoSk                          *string  `json:"noSk"`
	TahunPenguranganPermanenAwal  *string  `json:"tahunPenguranganPermanenAwal"`
	TahunPenguranganPermanenAkhir *string  `json:"tahunPenguranganPermanenAkhir"`
	StatusSkPenguraganPermanen    *string  `json:"statusSkPenguranganPermanen"`
	PctPenguranganPermanen        *float64 `json:"pctPenguranganPermanen"`
	Page                          int      `json:"page"`
	PageSize                      int      `json:"page_size"`
}

type UpdateDtoPermanen struct {
	Kanwil_Kd                     *string  `json:"kanwil_Kd"`
	Kppbb_Kd                      *string  `json:"kppbb_Kd"`
	TahunPelayanan                *string  `json:"tahunPelayanan"`
	BundelPelayanan               *string  `json:"bundelPelayanan"`
	NoUrutPelayanan               *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd            *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd              *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd           *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd           *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd                *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd              *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd             *string  `json:"jenisOpPemohon_Kd"`
	JenisSk                       *string  `json:"jenisSk"`
	NoSk                          *string  `json:"noSk"`
	TahunPenguranganPermanenAwal  *string  `json:"tahunPenguranganPermanenAwal"`
	TahunPenguranganPermanenAkhir *string  `json:"tahunPenguranganPermanenAkhir"`
	StatusSkPenguraganPermanen    *string  `json:"statusSkPenguranganPermanen"`
	PctPenguranganPermanen        *float64 `json:"pctPenguranganPermanen"`
}

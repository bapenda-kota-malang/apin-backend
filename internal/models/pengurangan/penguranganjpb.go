package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganJPB struct {
	BaseModel
	TahunPengenaanJPB *string  `json:"tahunPengenaanJPB" gorm:"type:char(4)"`
	PctPenguranganJPB *float64 `json:"pctPenguranganJPB" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type CreateDtoJPB struct {
	Kanwil_Kd           *string  `json:"kanwil_Kd"`
	Kppbb_Kd            *string  `json:"kppbb_Kd"`
	TahunPelayanan      *string  `json:"tahunPelayanan"`
	BundelPelayanan     *string  `json:"bundelPelayanan"`
	NoUrutPelayanan     *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd  *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd    *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd      *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd    *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd   *string  `json:"jenisOpPemohon_Kd"`
	JenisSk             *string  `json:"jenisSk"`
	NoSk                *string  `json:"noSk"`
	TahunPengenaanJPB   *string  `json:"tahunPengenaanJPB"`
	PctPenguranganJPB   *float64 `json:"pctPenguranganJPB"`
}

type FilterDtoJPB struct {
	Kanwil_Kd           *string  `json:"kanwil_Kd"`
	Kppbb_Kd            *string  `json:"kppbb_Kd"`
	TahunPelayanan      *string  `json:"tahunPelayanan"`
	BundelPelayanan     *string  `json:"bundelPelayanan"`
	NoUrutPelayanan     *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd  *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd    *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd      *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd    *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd   *string  `json:"jenisOpPemohon_Kd"`
	JenisSk             *string  `json:"jenisSk"`
	NoSk                *string  `json:"noSk"`
	TahunPengenaanJPB   *string  `json:"tahunPengenaanJPB"`
	PctPenguranganJPB   *float64 `json:"pctPenguranganJPB"`
	Page                int      `json:"page"`
	PageSize            int      `json:"page_size"`
}

type UpdateDtoJPB struct {
	Kanwil_Kd           *string  `json:"kanwil_Kd"`
	Kppbb_Kd            *string  `json:"kppbb_Kd"`
	TahunPelayanan      *string  `json:"tahunPelayanan"`
	BundelPelayanan     *string  `json:"bundelPelayanan"`
	NoUrutPelayanan     *string  `json:"noUrutPelayanan"`
	ProvinsiPemohon_Kd  *string  `json:"ProvinsiPemohon_Kd"`
	DaerahPemohon_Kd    *string  `json:"daerahPemohon_Kd"`
	KecamatanPemohon_Kd *string  `json:"kecamatanPemohon_Kd"`
	KelurahanPemohon_Kd *string  `json:"KelurahanPemohon_Kd"`
	BlokPemohon_Kd      *string  `json:"blokPemohon_Kd"`
	NoUrutPemohon_Kd    *string  `json:"noUrutPemohon_Kd"`
	JenisOpPemohon_Kd   *string  `json:"jenisOpPemohon_Kd"`
	JenisSk             *string  `json:"jenisSk"`
	NoSk                *string  `json:"noSk"`
	TahunPengenaanJPB   *string  `json:"tahunPengenaanJPB"`
	PctPenguranganJPB   *float64 `json:"pctPenguranganJPB"`
}

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
	Kanwil_Kode              *string            `json:"kanwil_kode"`
	KPPBB_Kode               *string            `json:"kppbb_kode"`
	ThnPelayanan             *string            `json:"thnPelayanan"`
	BundelPelayanan          *string            `json:"bundelPelayanan"`
	NoUrutPelayanan          *string            `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon    *string            `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon      *string            `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon   *string            `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon   *string            `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon        *string            `json:"blok_kode_pemohon"`
	NoUrut_Pemohon           *string            `json:"noUrut_pemohon"`
	JenisOp_Pemohon          *string            `json:"jenisOp_pemohon"`
	JenisSk                  *string            `json:"jenisSk"`
	NoSk                     *string            `json:"noSk"`
	TahunPenguranganDendaADM *string            `json:"tahunPenguranganDendaADM"`
	StatusPenguraganADM      *StatusPengurangan `json:"statusPenguraganADM"`
	PctPenguranganDendaADM   *float64           `json:"pctPenguranganDendaADM"`
}

type FilterDtoDendaADM struct {
	Kanwil_Kode              *string  `json:"kanwil_kode"`
	KPPBB_Kode               *string  `json:"kppbb_kode"`
	ThnPelayanan             *string  `json:"thnPelayanan"`
	BundelPelayanan          *string  `json:"bundelPelayanan"`
	NoUrutPelayanan          *string  `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon    *string  `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon      *string  `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon   *string  `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon   *string  `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon        *string  `json:"blok_kode_pemohon"`
	NoUrut_Pemohon           *string  `json:"noUrut_pemohon"`
	JenisOp_Pemohon          *string  `json:"jenisOp_pemohon"`
	JenisSk                  *string  `json:"jenisSk"`
	NoSk                     *string  `json:"noSk"`
	TahunPenguranganDendaADM *string  `json:"tahunPenguranganDendaADM"`
	StatusPenguraganADM      *string  `json:"statusPenguraganADM"`
	PctPenguranganDendaADM   *float64 `json:"pctPenguranganDendaADM"`
	Page                     int      `json:"page"`
	PageSize                 int      `json:"page_size"`
}

type UpdateDtoDendaADM struct {
	Kanwil_Kode              *string  `json:"kanwil_kode"`
	KPPBB_Kode               *string  `json:"kppbb_kode"`
	ThnPelayanan             *string  `json:"thnPelayanan"`
	BundelPelayanan          *string  `json:"bundelPelayanan"`
	NoUrutPelayanan          *string  `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon    *string  `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon      *string  `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon   *string  `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon   *string  `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon        *string  `json:"blok_kode_pemohon"`
	NoUrut_Pemohon           *string  `json:"noUrut_pemohon"`
	JenisOp_Pemohon          *string  `json:"jenisOp_pemohon"`
	JenisSk                  *string  `json:"jenisSk"`
	NoSk                     *string  `json:"noSk"`
	TahunPenguranganDendaADM *string  `json:"tahunPenguranganDendaADM"`
	StatusPenguraganADM      *string  `json:"statusPenguraganADM"`
	PctPenguranganDendaADM   *float64 `json:"pctPenguranganDendaADM"`
}

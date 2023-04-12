package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganPermanen struct {
	BaseModel
	ThnAwal                *string            `json:"thnAwal" gorm:"type:char(4)"`
	ThnAkhir               *string            `json:"thnAkhir" gorm:"type:char(4)"`
	StatusSK               *StatusPengurangan `json:"statusSk" gorm:"type:char(1)"`
	PCTPenguranganPermanen *float64           `json:"PCTPenguranganPermanen" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type DataPermanenDto struct {
	DataBaseDto
	ThnAwal                string            `json:"thnAwal"`
	ThnAkhir               string            `json:"thnAkhir"`
	StatusSK               StatusPengurangan `json:"statusSk"`
	PCTPenguranganPermanen float64           `json:"PCTPenguranganPermanen"`
}

type CreateDtoPermanen struct {
	Kanwil_Kode     string            `json:"kanwil_kode"`
	Kppbb_Kode      string            `json:"kppbb_kode"`
	ThnPelayanan    string            `json:"thnPelayanan" validate:"required"`
	BundelPelayanan string            `json:"bundelPelayanan" validate:"required"`
	NoUrutPelayanan string            `json:"noUrutPelayanan" validate:"required"`
	SkSk            CreateSkskDto     `json:"sksk" validate:"required"`
	Data            []DataPermanenDto `json:"data" validate:"required"`
}

type FilterDtoPermanen struct {
	Kanwil_Kode            *string            `json:"kanwil_kode"`
	Kppbb_Kode             *string            `json:"kppbb_kode"`
	ThnPelayanan           *string            `json:"thnPelayanan"`
	BundelPelayanan        *string            `json:"bundelPelayanan"`
	NoUrutPelayanan        *string            `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string            `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string            `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string            `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string            `json:"kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string            `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string            `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string            `json:"jenisOp_pemohon"`
	JenisSk                *string            `json:"jenisSk"`
	NoSk                   *string            `json:"noSk"`
	ThnAwal                *string            `json:"thnAwal"`
	ThnAkhir               *string            `json:"thnAkhir"`
	StatusSK               *StatusPengurangan `json:"statusSk"`
	PCTPenguranganPermanen *float64           `json:"PCTPenguranganPermanen"`
	Page                   int                `json:"page"`
	PageSize               int                `json:"page_size"`
}

type UpdateDtoPermanen struct {
	Kanwil_Kode            *string            `json:"kanwil_kode"`
	Kppbb_Kode             *string            `json:"kppbb_kode"`
	ThnPelayanan           *string            `json:"thnPelayanan"`
	BundelPelayanan        *string            `json:"bundelPelayanan"`
	NoUrutPelayanan        *string            `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string            `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string            `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string            `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string            `json:"kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string            `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string            `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string            `json:"jenisOp_pemohon"`
	JenisSk                *string            `json:"jenisSk"`
	NoSk                   *string            `json:"noSk"`
	ThnAwal                *string            `json:"thnAwal"`
	ThnAkhir               *string            `json:"thnAkhir"`
	StatusSK               *StatusPengurangan `json:"statusSk"`
	PCTPenguranganPermanen *float64           `json:"PCTPenguranganPermanen"`
}

package pengurangan

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type PenguranganJPB struct {
	BaseModel
	ThnPengenaan      *string  `json:"thnPengenaan" gorm:"type:char(4)"`
	PCTPenguranganJPB *float64 `json:"PCTPenguranganJPB" gorm:"type:decimal(5,2)"`
	gormhelper.DateModel
}

type CreateDtoJPB struct {
	Kanwil_Kode            *string  `json:"kanwil_kode"`
	KPPBB_Kode             *string  `json:"kppbb_kode"`
	ThnPelayanan           *string  `json:"thnPelayanan"`
	BundelPelayanan        *string  `json:"bundelPelayanan"`
	NoUrutPelayanan        *string  `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string  `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string  `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string  `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string  `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string  `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string  `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string  `json:"jenisOp_pemohon"`
	JenisSk                *string  `json:"jenisSk"`
	NoSk                   *string  `json:"noSk"`
	ThnPengenaan           *string  `json:"thnPengenaan"`
	PCTPenguranganJPB      *float64 `json:"PCTPenguranganJPB"`
}

type FilterDtoJPB struct {
	Kanwil_Kode            *string  `json:"kanwil_kode"`
	KPPBB_Kode             *string  `json:"kppbb_kode"`
	ThnPelayanan           *string  `json:"thnPelayanan"`
	BundelPelayanan        *string  `json:"bundelPelayanan"`
	NoUrutPelayanan        *string  `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string  `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string  `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string  `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string  `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string  `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string  `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string  `json:"jenisOp_pemohon"`
	JenisSk                *string  `json:"jenisSk"`
	NoSk                   *string  `json:"noSk"`
	ThnPengenaan           *string  `json:"thnPengenaan"`
	PCTPenguranganJPB      *float64 `json:"PCTPenguranganJPB"`
	Page                   int      `json:"page"`
	PageSize               int      `json:"page_size"`
}

type UpdateDtoJPB struct {
	Kanwil_Kode            *string  `json:"kanwil_kode"`
	KPPBB_Kode             *string  `json:"kppbb_kode"`
	ThnPelayanan           *string  `json:"thnPelayanan"`
	BundelPelayanan        *string  `json:"bundelPelayanan"`
	NoUrutPelayanan        *string  `json:"noUrutPelayanan"`
	Provinsi_Kode_Pemohon  *string  `json:"provinsi_kode_pemohon"`
	Daerah_Kode_Pemohon    *string  `json:"daerah_kode_pemohon"`
	Kecamatan_Kode_Pemohon *string  `json:"kecamatan_kode_pemohon"`
	Kelurahan_Kode_Pemohon *string  `json:"Kelurahan_kode_pemohon"`
	Blok_Kode_Pemohon      *string  `json:"blok_kode_pemohon"`
	NoUrut_Pemohon         *string  `json:"noUrut_pemohon"`
	JenisOp_Pemohon        *string  `json:"jenisOp_pemohon"`
	JenisSk                *string  `json:"jenisSk"`
	NoSk                   *string  `json:"noSk"`
	ThnPengenaan           *string  `json:"thnPengenaan"`
	PCTPenguranganJPB      *float64 `json:"PCTPenguranganJPB"`
}

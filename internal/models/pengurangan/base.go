package pengurangan

import "gorm.io/datatypes"

// this struct base for:
// PenguranganDendaADM, PenguranganJPB, PengurangaPermanen, PenguranganPST
type BaseModel struct {
	Id                     uint64 `json:"id" gorm:"primaryKey"`
	Kanwil_Kode            string `json:"kanwil_kode" gorm:"type:char(2);not null"`
	Kppbb_Kode             string `json:"kppbb_kode" gorm:"type:char(2);not null"`
	ThnPelayanan           string `json:"thnPelayanan" gorm:"type:char(4);not null"`
	BundelPelayanan        string `json:"bundelPelayanan" gorm:"type:char(1);not null"`
	NoUrutPelayanan        string `json:"noUrutPelayanan" gorm:"type:char(3);not null"`
	Provinsi_Kode_Pemohon  string `json:"provinsi_kode_pemohon" gorm:"type:char(2);not null"`
	Daerah_Kode_Pemohon    string `json:"daerah_kode_pemohon" gorm:"type:char(2);not null"`
	Kecamatan_Kode_Pemohon string `json:"kecamatan_kode_pemohon" gorm:"type:char(3);not null"`
	Kelurahan_Kode_Pemohon string `json:"kelurahan_kode_pemohon" gorm:"type:char(3);not null"`
	Blok_Kode_Pemohon      string `json:"blok_kode_pemohon" gorm:"type:char(3);not null"`
	NoUrut_Pemohon         string `json:"noUrut_pemohon" gorm:"type:char(4);not null"`
	JenisOp_Pemohon        string `json:"jenisOp_pemohon" gorm:"type:char(1);not null"`
	JnsSk                  string `json:"jnsSk" gorm:"type:char(1)"`
	NoSk                   string `json:"noSk" gorm:"type:char(30)"`
}

type CreateSkskDto struct {
	JnsSK  string          `json:"jnsSK" validate:"required"`
	NoSK   string          `json:"noSK" validate:"required"`
	TglSK  *datatypes.Date `json:"tglSK" validate:"required"`
	NoLhp  string          `json:"noLhp" validate:"required"`
	TglLhp *datatypes.Date `json:"tglLhp" validate:"required"`
}

type DataBaseDto struct {
	Provinsi_Kode_Pemohon  string `json:"provinsi_kode_pemohon" validate:"required"`
	Daerah_Kode_Pemohon    string `json:"daerah_kode_pemohon" validate:"required"`
	Kecamatan_Kode_Pemohon string `json:"kecamatan_kode_pemohon" validate:"required"`
	Kelurahan_Kode_Pemohon string `json:"kelurahan_kode_pemohon" validate:"required"`
	Blok_Kode_Pemohon      string `json:"blok_kode_pemohon" validate:"required"`
	NoUrut_Pemohon         string `json:"noUrut_pemohon" validate:"required"`
	JenisOp_Pemohon        string `json:"jenisOp_pemohon" validate:"required"`
}

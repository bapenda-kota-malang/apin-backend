package pengurangan

// this struct base for:
// PenguranganDendaADM, PenguranganJPB, PengurangaPermanen, PenguranganPST
type BaseModel struct {
	Id                     uint64  `json:"id" gorm:"primaryKey"`
	Kanwil_Kode            *string `json:"kanwil_kode" gorm:"type:char(2);not null"`
	KPPBB_Kode             *string `json:"kppbb_kode" gorm:"type:char(2);not null"`
	ThnPelayanan           *string `json:"thnPelayanan" gorm:"type:char(4);not null"`
	BundelPelayanan        *string `json:"bundelPelayanan" gorm:"type:char(1);not null"`
	NoUrutPelayanan        *string `json:"noUrutPelayanan" gorm:"type:char(3);not null"`
	Provinsi_Kode_Pemohon  *string `json:"provinsi_kode_pemohon" gorm:"type:char(2);not null"`
	Daerah_Kode_Pemohon    *string `json:"daerah_kode_pemohon" gorm:"type:char(2);not null"`
	Kecamatan_Kode_Pemohon *string `json:"kecamatan_kode_pemohon" gorm:"type:char(3);not null"`
	Kelurahan_Kode_Pemohon *string `json:"Kelurahan_kode_pemohon" gorm:"type:char(3);not null"`
	Blok_Kode_Pemohon      *string `json:"blok_kode_pemohon" gorm:"type:char(3);not null"`
	NoUrut_Pemohon         *string `json:"noUrut_pemohon" gorm:"type:char(4);not null"`
	JenisOp_Pemohon        *string `json:"jenisOp_pemohon" gorm:"type:char(1);not null"`
	JenisSk                *string `json:"jenisSk" gorm:"type:char(1)"`
	NoSk                   *string `json:"noSk" gorm:"type:char(30)"`
}

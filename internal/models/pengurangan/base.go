package pengurangan

// this struct base for:
// PenguranganDendaADM, PenguranganJPB, PengurangaPermanen, PenguranganPST
type BaseModel struct {
	Id                  uint64  `json:"id" gorm:"primaryKey"`
	Kanwil_Kd           *string `json:"kanwil_Kd" gorm:"type:char(2);not null"`
	Kppbb_Kd            *string `json:"kppbb_Kd" gorm:"type:char(2);not null"`
	TahunPelayanan      *string `json:"tahunPelayanan" gorm:"type:char(4);not null"`
	BundelPelayanan     *string `json:"bundelPelayanan" gorm:"type:char(1);not null"`
	NoUrutPelayanan     *string `json:"noUrutPelayanan" gorm:"type:char(3);not null"`
	ProvinsiPemohon_Kd  *string `json:"provinsiPemohon_Kd" gorm:"type:char(2);not null"`
	DaerahPemohon_Kd    *string `json:"daerahPemohon_Kd" gorm:"type:char(2);not null"`
	KecamatanPemohon_Kd *string `json:"kecamatanPemohon_Kd" gorm:"type:char(3);not null"`
	KelurahanPemohon_Kd *string `json:"KelurahanPemohon_Kd" gorm:"type:char(3);not null"`
	BlokPemohon_Kd      *string `json:"blokPemohon_Kd" gorm:"type:char(3);not null"`
	NoUrutPemohon_Kd    *string `json:"noUrutPemohon_Kd" gorm:"type:char(4);not null"`
	JenisOpPemohon_Kd   *string `json:"jenisOpPemohon_Kd" gorm:"type:char(1);not null"`
	JenisSk             *string `json:"jenisSk" gorm:"type:char(1)"`
	NoSk                *string `json:"noSk" gorm:"type:char(30)"`
}

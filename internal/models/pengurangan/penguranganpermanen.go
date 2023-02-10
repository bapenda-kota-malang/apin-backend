package pengurangan

type PenguranganaPermanen struct {
	Id                     uint64  `json:"id" gorm:"primaryKey"`
	BundelPelayanan        *string `json:"bundelPelayanan" gorm:"type:char(1)"`
	JenisSk                *string `json:"jenisSk" gorm:"type:char(4)"`
	BlokPemohon_Kd         *string `json:"blokPemohon_Kd" gorm:"type:char(3)"`
	DaerahPemohon_Kd       *string `json:"daerahPemohon_Kd" gorm:"type:char(2)"`
	JenisOpPemohon         *string `json:"jenisOpPemohon" gorm:"type:char(1)"`
	Kanwil_Kd              *string `json:"kanwil_Kd" gorm:"type:char(2)"`
	KecamatanPemohon_Kd    *string `json:"kecamatanPemohon_Kd" gorm:"type:char(3)"`
	KelurahanPemohon_Kd    *string `json:"KelurahanPemohon_Kd" gorm:"type:char(3)"`
	Kppbb_Kd               *string `json:"kppbb_Kd" gorm:"type:char(2)"`
	PropinsiPemohon_Kd     *string `json:"PropinsiPemohon_Kd" gorm:"type:char(1)"`
	NoSk                   *string `json:"noSk" gorm:"type:char(30)"`
	NoUrutPelayanan        *string `json:"noUrutPelayanan" gorm:"type:char(3)"`
	NoUrutPemohon          *string `json:"noUrutPemohon" gorm:"type:char(4)"`
	PctPenguranganPermanen *string `json:"pctPenguranganPermanen" gorm:"type:decimal(5,2)"`
}

package areadivision

type Daerah struct {
	ID           uint64       `json:"id" gorm:"primaryKey"`
	ProvinsiKode string       `json:"provinsi_kode"`
	Provinsi     *Provinsi    `json:"provinsi,omitempty" gorm:"foreignKey:ProvinsiKode;references:Kode"`
	Kode         string       `json:"kode" gorm:"unique"`
	Nama         string       `json:"nama"`
	Kecamatan    []*Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:DaerahKode;references:Kode"`
}

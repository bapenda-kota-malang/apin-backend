package areadivision

type Kecamatan struct {
	ID         uint64       `json:"id" gorm:"primaryKey"`
	DaerahKode string       `json:"daerah_kode"`
	Daerah     *Daerah      `json:"daerah,omitempty" gorm:"foreignKey:DaerahKode;references:Kode"`
	Kode       string       `json:"kode" gorm:"unique"`
	Nama       string       `json:"nama"`
	Kelurahan  []*Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:KecamatanKode;references:Kode"`
}

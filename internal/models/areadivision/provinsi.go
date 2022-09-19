package areadivision

type Provinsi struct {
	ID     uint64    `json:"id" gorm:"primaryKey"`
	Kode   string    `json:"kode" gorm:"unique"`
	Nama   string    `json:"nama"`
	Daerah []*Daerah `json:"daerah,omitempty" gorm:"foreignKey:ProvinsiKode;references:Kode"`
}

package areadivision

type Kelurahan struct {
	ID            uint64     `json:"id" gorm:"primaryKey"`
	KecamatanKode string     `json:"kecamatan_kode"`
	Kecamatan     *Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:KecamatanKode;references:Kode"`
	Kode          string     `json:"kode" gorm:"unique"`
	Nama          string     `json:"nama"`
}

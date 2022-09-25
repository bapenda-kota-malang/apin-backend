package areadivision

type Kecamatan struct {
	ID         uint64       `json:"id" gorm:"primaryKey"`
	DaerahKode string       `json:"daerah_kode"`
	Daerah     *Daerah      `json:"daerah,omitempty" gorm:"foreignKey:DaerahKode;references:Kode"`
	Kode       string       `json:"kode" gorm:"unique"`
	Nama       string       `json:"nama"`
	Kelurahan  []*Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:KecamatanKode;references:Kode"`
}

type KecamatanCreateDto struct {
	DaerahKode string `json:"daerah_kode" validate:"requred;min=1"`
	Kode       string `json:"kode" validate:"requred;min=1"`
	Nama       string `json:"nama" validate:"required;maxLength=100"`
}

type KecamatanUpdateDto struct {
	DaerahKode string `json:"daerah_kode" validate:"requred;min=1"`
	Kode       string `json:"kode" validate:"requred;min=1"`
	Nama       string `json:"nama" validate:"required;maxLength=100"`
}

type KecamatanFilterDto struct {
	// DaerahKode *string `json:"daerah_kode"`
	// Kode       *string `json:"kode"`
	// Nama       string  `json:"nama"`
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

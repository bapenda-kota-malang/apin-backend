package areadivision

type Kecamatan struct {
	ID          uint64       `json:"id" gorm:"primaryKey"`
	Daerah_Kode string       `json:"daerah_kode" gorm:"size:4"`
	Daerah      *Daerah      `json:"daerah,omitempty" gorm:"foreignKey:Daerah_Kode;references:Kode"`
	Kode        string       `json:"kode" gorm:"unique;size:7"`
	Nama        string       `json:"nama" gorm:"size:100"`
	Kelurahan   []*Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kecamatan_Kode;references:Kode"`
}

type KecamatanCreateDto struct {
	Daerah_Kode string `json:"daerah_kode" validate:"requred;min=1"`
	Kode        string `json:"kode" validate:"requred;min=1"`
	Nama        string `json:"nama" validate:"required;maxLength=100"`
}

type KecamatanUpdateDto struct {
	Daerah_Kode string `json:"daerah_kode" validate:"requred;min=1"`
	Kode        string `json:"kode" validate:"requred;min=1"`
	Nama        string `json:"nama" validate:"required;maxLength=100"`
}

type KecamatanFilterDto struct {
	// DaerahKode *string `json:"daerah_kode"`
	// Kode       *string `json:"kode"`
	// Nama       string  `json:"nama"`
	Page     int   `json:"page"`
	PageSize int64 `json:"pageSize"`
}

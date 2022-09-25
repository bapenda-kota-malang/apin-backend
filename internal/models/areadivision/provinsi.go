package areadivision

type Provinsi struct {
	ID     uint64    `json:"id" gorm:"primaryKey"`
	Kode   string    `json:"kode" gorm:"unique"`
	Nama   string    `json:"nama"`
	Daerah []*Daerah `json:"daerah,omitempty" gorm:"foreignKey:ProvinsiKode;references:Kode"`
}

type ProvinsiCreateDto struct {
	Kode string `json:"kode" validate:"requred;min=1"`
	Nama string `json:"nama" validate:"required;maxLength=100"`
}

type ProvinsiUpdateDto struct {
	Kode string `json:"kode" validate:"requred;min=1"`
	Nama string `json:"nama" validate:"required;maxLength=100"`
}

type ProvinsiFilterDto struct {
	// ID       *uint64 `json:"id"`
	// Kode     *string `json:"kode"`
	// Nama     *string `json:"nama"`
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

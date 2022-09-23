package areadivision

type Daerah struct {
	ID           uint64       `json:"id" gorm:"primaryKey"`
	ProvinsiKode string       `json:"provinsi_kode"`
	Provinsi     *Provinsi    `json:"provinsi,omitempty" gorm:"foreignKey:ProvinsiKode;references:Kode"`
	Kode         string       `json:"kode" gorm:"unique"`
	Nama         string       `json:"nama"`
	Kecamatan    []*Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:DaerahKode;references:Kode"`
}

type DaerahCreateDto struct {
	ProvinsiKode string `json:"provinsi_kode" validate:"requred;min=1"`
	Kode         string `json:"kode" validate:"requred;min=1"`
	Nama         string `json:"nama" validate:"required;maxLength=100"`
}

type DaerahUpdateDto struct {
	ProvinsiKode string `json:"provinsi_kode" validate:"requred;min=1"`
	Kode         string `json:"kode" validate:"requred;min=1"`
	Nama         string `json:"nama" validate:"required;maxLength=100"`
}

type DaerahFilterDto struct {
	ProvinsiKode *string `json:"provinsi_kode"`
	// Kode         *string `json:"kode"`
	// Nama         *string `json:"nama"`
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

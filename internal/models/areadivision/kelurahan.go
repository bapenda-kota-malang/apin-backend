package areadivision

type Kelurahan struct {
	ID            uint64     `json:"id" gorm:"primaryKey"`
	KecamatanKode string     `json:"kecamatan_kode"`
	Kecamatan     *Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:KecamatanKode;references:Kode"`
	Kode          string     `json:"kode" gorm:"unique"`
	Nama          string     `json:"nama"`
}
type KelurahanCreateDto struct {
	KecamatanKode string `json:"kecamatan_kode" validate:"requred;min=1"`
	Kode          string `json:"kode" validate:"requred;min=1"`
	Nama          string `json:"nama" validate:"required;maxLength=100"`
}

type KelurahanUpdateDto struct {
	KecamatanKode string `json:"kecamatan_kode" validate:"requred;min=1"`
	Kode          string `json:"kode" validate:"requred;min=1"`
	Nama          string `json:"nama" validate:"required;maxLength=100"`
}

type KelurahanFilterDto struct {
	// KecamatanKode *string `json:"kecamatan_kode"`
	// Kode          *string `json:"kode"`
	// Nama          *string `json:"nama"`
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

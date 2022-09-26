package areadivision

type Daerah struct {
	ID            uint64       `json:"id" gorm:"primaryKey"`
	Provinsi_Kode string       `json:"provinsi_kode" gorm:"size:2"`
	Provinsi      *Provinsi    `json:"provinsi,omitempty" gorm:"foreignKey:Provinsi_Kode;references:Kode"`
	Kode          string       `json:"kode" gorm:"unique;size:4"`
	Nama          string       `json:"nama" gorm:"size:100"`
	Kecamatan     []*Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Daerah_Kode;references:Kode"`
}

type DaerahCreateDto struct {
	Provinsi_Kode string `json:"provinsi_kode" validate:"requred;min=1"`
	Kode          string `json:"kode" validate:"requred;min=1"`
	Nama          string `json:"nama" validate:"required;maxLength=100"`
}

type DaerahUpdateDto struct {
	Provinsi_Kode string `json:"provinsi_kode" validate:"requred;min=1"`
	Kode          string `json:"kode" validate:"requred;min=1"`
	Nama          string `json:"nama" validate:"required;maxLength=100"`
}

type DaerahFilterDto struct {
	Provinsi_Kode *string `json:"provinsi_kode"`
	// Kode         *string `json:"kode"`
	// Nama         *string `json:"nama"`
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

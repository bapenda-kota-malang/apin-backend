package areadivision

type Kelurahan struct {
	Id             uint64     `json:"id" gorm:"primaryKey"`
	Kecamatan_Kode string     `json:"kecamatan_kode" gorn:"size:7"`
	Kecamatan      *Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Kode;references:Kode"`
	Kode           string     `json:"kode" gorm:"unique;size:10"`
	Nop            string     `json:"nop" gorm:"unique;size:10"`
	Nama           string     `json:"nama" gorm:"size:100"`
}
type KelurahanCreateDto struct {
	Kecamatan_Kode string `json:"kecamatan_kode" validate:"requred;min=1"`
	Kode           string `json:"kode" validate:"requred;min=1"`
	Nama           string `json:"nama" validate:"required;maxLength=100"`
}

type KelurahanUpdateDto struct {
	Kecamatan_Kode string `json:"kecamatan_kode" validate:"requred;min=1"`
	Kode           string `json:"kode" validate:"requred;min=1"`
	Nama           string `json:"nama" validate:"required;maxLength=100"`
}

type KelurahanFilterDto struct {
	Kecamatan_Kode *string `json:"kecamatan_kode"`
	Kode           *string `json:"kode"`
	Kode_Opt       *string `json:"kode_opt"`
	Nama           *string `json:"nama"`
	Page           int     `json:"page"`
	PageSize       int64   `json:"page_size"`
	NoPagination   bool    `json:"no_pagination"`
}

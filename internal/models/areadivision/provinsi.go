package areadivision

type Provinsi struct {
	Id     uint64    `json:"id" gorm:"primaryKey"`
	Kode   string    `json:"kode" gorm:"unique;size:2"`
	Nama   string    `json:"nama" gorm:"size:100"`
	Daerah []*Daerah `json:"daerah,omitempty" gorm:"foreignKey:Provinsi_Kode;references:Kode"`
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
	Kode         *string `json:"kode"`
	Nama         *string `json:"nama"`
	Page         int     `json:"page"`
	PageSize     int64   `json:"page_size"`
	NoPagination bool    `json:"no_pagination"`
}

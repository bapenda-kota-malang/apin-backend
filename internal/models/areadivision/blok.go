package areadivision

type Blok struct {
	Kode string `json:"kode" validate:"requred;min=1"`
	Nama string `json:"nama" validate:"required;maxLength=100"`
}

type BlokCreateDto struct {
	Kode string `json:"kode" validate:"requred;min=1"`
	Nama string `json:"nama" validate:"required;maxLength=100"`
}

type BlokUpdateDto struct {
	Kode string `json:"kode" validate:"requred;min=1"`
	Nama string `json:"nama" validate:"required;maxLength=100"`
}

type BlokFilterDto struct {
	Kode         *string `json:"kode"`
	Nama         *string `json:"nama"`
	Page         int     `json:"page"`
	PageSize     int64   `json:"page_size"`
	NoPagination bool    `json:"no_pagination"`
}

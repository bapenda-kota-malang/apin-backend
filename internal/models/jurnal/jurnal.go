package jurnal

type Jurnal struct {
	Id   uint64  `json:"id" gorm:"primaryKey"`
	Kode *string `json:"kode" gorm:"type:varchar(100)"`
	Nama *string `json:"nama" gorm:"type:varchar(50)"`
}

type CreateDto struct {
	Kode string `json:"kode" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

type UpdateDto struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Nama     *string `json:"nama"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

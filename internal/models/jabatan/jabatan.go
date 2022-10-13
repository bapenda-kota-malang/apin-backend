package jabatan

type Jabatan struct {
	Id   uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Kode *string `json:"kode" gorm:"size:20"`
	Nama *string `json:"nama" gorm:"size:100"`
}

type CreateDto struct {
	Kode *string `json:"kode" validate:"required"`
	Nama *string `json:"nama" validate:"required"`
}

type UpdateDto struct {
	Kode *string `json:"kode" validate:"required"`
	Nama *string `json:"nama" validate:"required"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Nama     *string `json:"nama"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

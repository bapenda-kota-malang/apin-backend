package pangkat

type Pangkat struct {
	Id       uint64  `json:"id" gorm:"primaryKey"`
	Kode     *string `json:"kode" gorm:"type:varchar(50)"`
	Nama     *string `json:"nama" gorm:"type:varchar(100)"`
	Golongan *string `json:"golongan" gorm:"type:char(1)"`
	Ruang    *string `json:"ruang" gorm:"type:char(1)"`
}

type CreateDto struct {
	Kode     string `json:"kode" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Golongan string `json:"golongan" validate:"required"`
	Ruang    string `json:"ruang" validate:"required, gte=1, lte=9"`
}

type UpdateDto struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	Golongan string `json:"golongan"`
	Ruang    string `json:"ruang"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Nama     *string `json:"nama"`
	Golongan *string `json:"golongan"`
	Ruang    *string `json:"ruang"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

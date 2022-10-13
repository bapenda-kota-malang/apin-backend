package klasifikasijalan

type KlasifikasiJalan struct {
	Id   uint64  `json:"id" gorm:"primaryKey"`
	Kode *string `json:"kode" gorm:"type:varchar(3)"`
}

type CreateDto struct {
	Kode string `json:"kode" validate:"required"`
}

type UpdateDto struct {
	Kode string `json:"kode"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

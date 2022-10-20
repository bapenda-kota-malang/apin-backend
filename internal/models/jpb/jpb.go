package jpb

type Jpb struct {
	Id   uint64 `json:"id" gorm:"primarykey"`
	Kode string `json:"kode" gorm:"type:char(2)"`
	Nama string `json:"nama" gorm:"type:varchar(30)"`
}

type CreateDto struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type UpdateDto struct {
	Id   string `json:"id"`
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Nama     *string `json:"nama"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

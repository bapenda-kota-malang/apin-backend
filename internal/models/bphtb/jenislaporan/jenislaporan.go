package jenislaporan

type BphtbJenisLaporan struct {
	Id   uint64 `json:"id" gorm:"primaryKey"`
	Kode string `json:"kode" gorm:"type:varchar(2)"`
	Nama string `json:"nama" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	Kode string `json:"kode" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

type FilterDto struct {
	Kode *string `json:"kode"`
	Nama *string `json:"nama"`
}

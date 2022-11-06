package tarifjambong

type TarifJambong struct {
	Id           uint64   `json:"id" gorm:"primaryKey"`
	JenisReklame *string  `json:"jenisReklame" gorm:"type:varchar(100)"`
	Nominal      *float64 `json:"nominal" gorm:"type:decimal"`
}

type CreateDto struct {
	JenisReklame string  `json:"jenisReklame" validate:"required"`
	Nominal      float64 `json:"nominal" validate:"required"`
}

type UpdateDto struct {
	JenisReklame string  `json:"jenisReklame"`
	Nominal      float64 `json:"nominal"`
}

type FilterDto struct {
	JenisReklame *string  `json:"jenisReklame"`
	Nominal      *float64 `json:"nominal"`
	Page         int      `json:"page"`
	PageSize     int      `json:"page_size"`
}

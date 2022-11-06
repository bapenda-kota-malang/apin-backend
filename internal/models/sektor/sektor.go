package sektor

type Sektor struct {
	Id         uint64  `json:"id" gorm:"primaryKey"`
	KodeSektor *string `json:"kodeSektor" gorm:"type:char(2)"`
	NamaSektor *string `json:"namaSektor" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	KodeSektor string `json:"kodeSektor" validate:"required"`
	NamaSektor string `json:"namaSektor" validate:"required"`
}

type UpdateDto struct {
	KodeSektor string `json:"kodeSektor"`
	NamaSektor string `json:"namaSektor"`
}

type FilterDto struct {
	KodeSektor *string `json:"kodeSektor"`
	NamaSektor *string `json:"namaSektor"`
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
}

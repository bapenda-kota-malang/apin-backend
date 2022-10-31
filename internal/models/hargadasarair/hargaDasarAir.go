package hargadasarair

type HargaDasarAir struct {
	Id                uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Peruntukan        *string  `json:"peruntukan" gorm:"type:varchar(100)"`
	BatasBawah        *string  `json:"batasBawah" gorm:"type:varchar(10)"`
	BatasAtas         *string  `json:"batasAtas" gorm:"type:varchar(10)"`
	TarifMataAir      *float64 `json:"tarifMataAir" gorm:"type:decimal"`
	TarifBukanMataAir *float64 `json:"tarifBukanMataAir" gorm:"type:decimal"`
}

type CreateDto struct {
	Peruntukan        string  `json:"peruntukan" validate:"required"`
	BatasBawah        string  `json:"batasBawah" validate:"required"`
	BatasAtas         string  `json:"batasAtas" validate:"required"`
	TarifMataAir      float64 `json:"tarifMataAir" validate:"required"`
	TarifBukanMataAir float64 `json:"tarifBukanMataAir" validate:"required"`
}

type UpdateDto struct {
	Peruntukan        string  `json:"peruntukan"`
	BatasBawah        string  `json:"batasBawah"`
	BatasAtas         string  `json:"batasAtas"`
	TarifMataAir      float64 `json:"tarifMataAir"`
	TarifBukanMataAir float64 `json:"tarifBukanMataAir"`
}

type FilterDto struct {
	Peruntukan        *string  `json:"peruntukan"`
	BatasBawah        *string  `json:"batasBawah"`
	BatasAtas         *string  `json:"batasAtas"`
	TarifMataAir      *float64 `json:"tarifMataAir"`
	TarifBukanMataAir *float64 `json:"tarifBukanMataAir"`
	Page              int      `json:"page"`
	PageSize          int      `json:"page_size"`
}

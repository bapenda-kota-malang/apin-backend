package hargadasarair

type HargaDasarAir struct {
	Id                uint64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Peruntukan        *string  `json:"peruntukan" gorm:"type:varchar(100)"`
	BatasBawah        *float64 `json:"batasBawah"`
	BatasAtas         *float64 `json:"batasAtas"`
	TarifMataAir      *float64 `json:"tarifMataAir" gorm:"type:decimal"`
	TarifBukanMataAir *float64 `json:"tarifBukanMataAir" gorm:"type:decimal"`
}

type CreateDto struct {
	Peruntukan        string  `json:"peruntukan" validate:"required"`
	BatasBawah        float64 `json:"batasBawah" validate:"required"`
	BatasAtas         float64 `json:"batasAtas" validate:"required"`
	TarifMataAir      float64 `json:"tarifMataAir" validate:"required"`
	TarifBukanMataAir float64 `json:"tarifBukanMataAir" validate:"required"`
}

type UpdateDto struct {
	Peruntukan        string  `json:"peruntukan"`
	BatasBawah        float64 `json:"batasBawah"`
	BatasAtas         float64 `json:"batasAtas"`
	TarifMataAir      float64 `json:"tarifMataAir"`
	TarifBukanMataAir float64 `json:"tarifBukanMataAir"`
}

type FilterDto struct {
	Peruntukan        *string  `json:"peruntukan"`
	BatasBawah        *float64 `json:"batasBawah"`
	BatasBawah_Opt    *string  `json:"batasBawah_Opt"`
	BatasAtas         *float64 `json:"batasAtas"`
	BatasAtas_Opt     *string  `json:"batasAtas_Opt"`
	TarifMataAir      *float64 `json:"tarifMataAir"`
	TarifBukanMataAir *float64 `json:"tarifBukanMataAir"`
	Page              int      `json:"page"`
	PageSize          int      `json:"page_size"`
}

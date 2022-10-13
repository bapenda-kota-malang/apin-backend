package tarifjambongrek

type TarifJambongRek struct {
	Id         uint64   `json:"id" gorm:"primaryKey"`
	Tipe       *string  `json:"tipe" gorm:"type:varchar(50)"`
	BatasBawah *string  `json:"batasBawah" gorm:"type:varchar(10)"`
	BatasAtas  *string  `json:"batasAtas" gorm:"type:varchar(10)"`
	Nominal    *float64 `json:"nominal" gorm:"type:decimal"`
}

type CreateDto struct {
	Tipe       string  `json:"tipe" validate:"required"`
	BatasBawah string  `json:"batasBawah" validate:"required"`
	BatasAtas  string  `json:"batasAtas" validate:"required"`
	Nominal    float64 `json:"nominal" validate:"required"`
}

type UpdateDto struct {
	Tipe       string  `json:"tipe"`
	BatasBawah string  `json:"batasBawah"`
	BatasAtas  string  `json:"batasAtas"`
	Nominal    float64 `json:"nominal"`
}
type FilterDto struct {
	Tipe       *string  `json:"tipe"`
	BatasBawah *string  `json:"batasBawah"`
	BatasAtas  *string  `json:"batasAtas"`
	Nominal    *float64 `json:"nominal"`
	Page       int      `json:"page"`
	PageSize   int      `json:"page_size"`
}

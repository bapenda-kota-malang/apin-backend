package sts

type SumberDanaSts struct {
	Id           uint64   `json:"id" gorm:"primaryKey"`
	Sts_Id       uint64   `json:"sts_id"`
	Sts          *Sts     `json:"sts,omitempty" gorm:"foreignKey:Sts_Id"`
	IsSumberDana bool     `json:"isSumberDana"`
	Nominal      *float64 `json:"nominal" gorm:"type:decimal"`
}

type SumberDanaStsCreateDto struct {
	Sts_Id       uint64   `json:"sts_id"`
	IsSumberDana bool     `json:"isSumberDana"`
	Nominal      *float64 `json:"nominal"`
}
type SumberDanaStsUpdateDto struct {
	Id           *uint64  `json:"id"`
	Sts_Id       uint64   `json:"sts_id"`
	IsSumberDana bool     `json:"isSumberDana"`
	Nominal      *float64 `json:"nominal"`
}

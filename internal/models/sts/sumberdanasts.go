package sts

type SumberDanaSts struct {
	Id           uint64   `json:"id" gorm:"primaryKey"`
	IsSumberDana *bool    `json:"isSumberDana"`
	Nominal      *float64 `json:"nominal" gorm:"type:decimal"`
}

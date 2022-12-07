package kppbb

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kanwil"
)

type kppbb struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	AlKppbb         *string `json:"alKppbb" gorm:"type:varchar(50)"`
	KdKanwil        *string `json:"kdKanwil,omitempty" gorm:"type:varchar(2);foreignKey:Rekening_Id"`
	KdKppbb         *string `json:"kdKppbb,omitempty" gorm:"type:varchar(2)"`
	KotaTerbitKppbb *string `json:"kotaTerbitKppbb" gorm:"type:varchar(30)"`
	NmKppbb         *string `json:"nmKppbb" gorm:"type:varchar(30)"`
	NoFaksimili     *string `json:"noFaksimili" gorm:"type:varchar(50)"`
	NoTelpon        *string `json:"noTelpon" gorm:"type:varchar(50)"`

	Kanwil *kanwil.Kanwil
}

package kppbb

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/kanwil"
)

type Kppbb struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	AlKppbb         *string `json:"alKppbb" gorm:"type:varchar(50)"`
	KanwilId        *uint64  `json:"kanwilId" gorm:"type:bigint"`
	KdKppbb         *string `json:"kdKppbb,omitempty" gorm:"type:varchar(2)"`
	KotaTerbitKppbb *string `json:"kotaTerbitKppbb" gorm:"type:varchar(30)"`
	NmKppbb         *string `json:"nmKppbb" gorm:"type:varchar(30)"`
	NoFaksimili     *string `json:"noFaksimili" gorm:"type:varchar(50)"`
	NoTelpon        *string `json:"noTelpon" gorm:"type:varchar(50)"`

	Kanwil *kanwil.Kanwil `json:"kanwil" gorm:"foreignKey:KanwilId;references:Id"`
}

type CreateDto struct {
	AlKppbb         *string `json:"alKppbb"`
	KanwilId        *uint64 `json:"kanwilId"`
	KdKppbb         *string `json:"kdKppbb"`
	KotaTerbitKppbb *string `json:"kotaTerbitKppbb"`
	NmKppbb         *string `json:"nmKppbb"`
	NoFaksimili     *string `json:"noFaksimili"`
	NoTelpon        *string `json:"noTelpon"`
}

type UpdateDto struct {
	Id               *uint64 `json:"id"`
	AlKppbb         *string `json:"alKppbb"`
	KanwilId        *uint64 `json:"kanwilId"`
	KdKppbb         *string `json:"kdKppbb"`
	KotaTerbitKppbb *string `json:"kotaTerbitKppbb"`
	NmKppbb         *string `json:"nmKppbb"`
	NoFaksimili     *string `json:"noFaksimili"`
	NoTelpon        *string `json:"noTelpon"`
}

type FilterDto struct {
	AlKppbb         *string `json:"alKppbb"`
	KanwilId        *uint64 `json:"kanwilId"`
	KdKppbb         *string `json:"kdKppbb"`
	KotaTerbitKppbb *string `json:"kotaTerbitKppbb"`
	NmKppbb         *string `json:"nmKppbb"`
	NoFaksimili     *string `json:"noFaksimili"`
	NoTelpon        *string `json:"noTelpon"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

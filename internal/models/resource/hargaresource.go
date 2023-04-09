package hargaresource

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type HargaResource struct {
	Id                 uint64       `json:"id" gorm:"primaryKey"`
	Provisi_Kode       string       `json:"provisi_kode" gorm:"type:varchar(2)"`
	Dati2_Kode         string       `json:"dati2_kode" gorm:"type:varchar(2)"`
	ThnHrgResource     string       `json:"thnHrgResource" gorm:"type:varchar(4)"`
	GroupResource_Kode string       `json:"groupResource_Kode" gorm:"type:varchar(2)"`
	Resource_Kode      string       `json:"resource_Kode" gorm:"type:varchar(2)"`
	KdKanwil           *string      `json:"kdKanwil" gorm:"type:varchar(2)"`
	KdKppbb            *string      `json:"kdKppbb" gorm:"type:varchar(2)"`
	JnsDokumen         *string      `json:"jnsDokumen" gorm:"type:varchar(1)"`
	NoDokumen          *string      `json:"noDokumen" gorm:"type:varchar(2)"`
	HrgResource        *float64     `json:"hrgResource" gorm:"type:varchar(2)"`
	Provinsi           *ad.Provinsi `json:"provinsi,omitempty" gorm:"foreignKey:Kelurahan_Kode;references:Kode"`
	Dati2              *ad.Daerah   `json:"dati2,omitempty" gorm:"foreignKey:Kecamatan_Kode;references:Kode"`
	gh.DateModel
}

type CreateDto struct {
	Provisi_Kode       string   `json:"provisi_kode"`
	Dati2_Kode         string   `json:"dati2_kode"`
	ThnHrgResource     string   `json:"thnHrgResource"`
	GroupResource_Kode string   `json:"groupResource_Kode"`
	Resource_Kode      string   `json:"resource_Kode"`
	KdKanwil           *string  `json:"kdKanwil"`
	KdKppbb            *string  `json:"kdKppbb"`
	JnsDokumen         *string  `json:"jnsDokumen"`
	NoDokumen          *string  `json:"noDokumen"`
	HrgResource        *float64 `json:"hrgResource"`
}

type UpdateDto struct {
	Provisi_Kode       string   `json:"provisi_kode"`
	Dati2_Kode         string   `json:"dati2_kode"`
	ThnHrgResource     string   `json:"thnHrgResource"`
	GroupResource_Kode string   `json:"groupResource_Kode"`
	Resource_Kode      string   `json:"resource_Kode"`
	KdKanwil           *string  `json:"kdKanwil"`
	KdKppbb            *string  `json:"kdKppbb"`
	JnsDokumen         *string  `json:"jnsDokumen"`
	NoDokumen          *string  `json:"noDokumen"`
	HrgResource        *float64 `json:"hrgResource"`
}

type FilterDto struct {
	Provisi_Kode       string   `json:"provisi_kode"`
	Dati2_Kode         string   `json:"dati2_kode"`
	ThnHrgResource     string   `json:"thnHrgResource"`
	GroupResource_Kode string   `json:"groupResource_Kode"`
	Resource_Kode      string   `json:"resource_Kode"`
	KdKanwil           *string  `json:"kdKanwil"`
	KdKppbb            *string  `json:"kdKppbb"`
	JnsDokumen         *string  `json:"jnsDokumen"`
	NoDokumen          *string  `json:"noDokumen"`
	HrgResource        *float64 `json:"hrgResource"`
	Page               int      `json:"page"`
	PageSize           int      `json:"page_size"`
}

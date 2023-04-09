package hargaresource

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type ItemResource struct {
	Id                 uint64 `json:"id" gorm:"primaryKey"`
	GroupResource_Kode string `json:"groupResource_Kode" gorm:"type:varchar(2)"`
	Resource_Kode      string `json:"resource_kode" gorm:"type:varchar(2)"`
	Resource_Name      string `json:"resource_name" gorm:"type:varchar(45)"`
	SatuanResource     string `json:"satuanResource" gorm:"type:varchar(15)"`
	gh.DateModel
}

type CreateItemDto struct {
	GroupResource_Kode string `json:"groupResource_Kode"`
	Resource_Kode      string `json:"resource_kode"`
	Resource_Name      string `json:"resource_name"`
	SatuanResource     string `json:"satuanResource"`
}

type UpdateItemDto struct {
	GroupResource_Kode string `json:"groupResource_Kode"`
	Resource_Kode      string `json:"resource_kode"`
	Resource_Name      string `json:"resource_name"`
	SatuanResource     string `json:"satuanResource"`
}

type FilterItemDto struct {
	GroupResource_Kode string `json:"groupResource_Kode"`
	Resource_Kode      string `json:"resource_kode"`
	Resource_Name      string `json:"resource_name"`
	SatuanResource     string `json:"satuanResource"`
	Page               int    `json:"page"`
	PageSize           int    `json:"page_size"`
}

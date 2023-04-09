package hargaresource

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type GroupResource struct {
	Id         uint64 `json:"id" gorm:"primaryKey"`
	Group_Kode string `json:"group_kode" gorm:"type:varchar(2)"`
	Name       string `json:"name" gorm:"type:varchar(45)"`
	gh.DateModel
}

type CreateGroupDto struct {
	Group_Kode string `json:"group_kode"`
	Name       string `json:"name"`
}

type UpdateGroupDto struct {
	Group_Kode string `json:"group_kode"`
	Name       string `json:"name"`
}

type FilterGroupDto struct {
	Group_Kode string `json:"group_kode"`
	Name       string `json:"name"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

package prosesjambong

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type ProsesJambong struct {
	Id                uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	JaminanBongkar_Id uint64          `json:"jaminanBongkar_Id"`
	Status            Status          `json:"status"`
	Date              *datatypes.Date `json:"date"`
	Note              *string         `json:"note"`
	Nominal           *float64        `json:"nominal"`
	Location          *string         `json:"location"`
	CreateBy_User_Id  uint            `json:"createBy_user_id"`
	gormhelper.DateModel
	CreateBy *user.User `json:"createBy,omitempty" gorm:"foreignKey:CreateBy_User_Id"`
}

type CreateDto struct {
	JaminanBongkar_Id *uint64         `json:"jaminanBongkar_Id" validate:"required"`
	Status            Status          `json:"status" validate:"required;min=1;max=2"`
	Date              *datatypes.Date `json:"date" validate:"required"`
	Note              *string         `json:"note"`
	Nominal           *float64        `json:"nominal"`
	Location          *string         `json:"location"`
}

type UpdateDto struct {
	Status           *Status         `json:"status" validate:"min=1;max=2"`
	Date             *datatypes.Date `json:"date"`
	Note             *string         `json:"note"`
	Nominal          *float64        `json:"nominal"`
	Location         *string         `json:"location"`
	CreateBy_User_Id *uint           `json:"createBy_user_id"`
}

type FilterDto struct {
	JaminanBongkar_Id *uint64         `json:"jaminanBongkar_Id"`
	Status            *Status         `json:"status" validate:"required;min=1;max=2"`
	Date              *datatypes.Date `json:"date"`
	Note              *string         `json:"note"`
	Nominal           *float64        `json:"nominal"`
	Location          *string         `json:"location"`
	CreateBy_User_Id  *uint           `json:"createBy_user_id"`
	Page              int             `json:"page"`
	PageSize          int             `json:"page_size"`
}

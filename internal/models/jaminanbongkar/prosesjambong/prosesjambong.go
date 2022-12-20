package prosesjambong

import (
	"os/user"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/jaminanbongkar"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Prosesjambong struct {
	Id                uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	JaminanBongkar_Id uint64          `json:"jaminanBongkar_Id"`
	Status            Status          `json:"status"`
	Date              *datatypes.Date `json:"date"`
	Note              string          `json:"note"`
	Nominal           *float64        `json:"nominal"`
	Location          string          `json:"location"`
	CreateBy_User_Id  uint            `json:"createBy_user_id"`
	gormhelper.DateModel
	JaminanBongkar *jaminanbongkar.JaminanBongkar `json:"jaminanBongkar,omitempty" gorm:"foreignKey:JaminanBongkar_Id"`
	CreateBy       *user.User                     `json:"createBy,omitempty" gorm:"foreignKey:CreateBy_User_Id"`
}

type CreateDto struct {
	JaminanBongkar_Id uint64          `json:"jaminanBongkar_Id"`
	Status            Status          `json:"status" validate:"required;min=1;max=2"`
	Date              *datatypes.Date `json:"date"`
	Note              string          `json:"note"`
	Nominal           *float64        `json:"nominal"`
	Location          string          `json:"location"`
	CreateBy_User_Id  uint            `json:"createBy_user_id"`
}

type UpdateDto struct {
	Status           *Status         `json:"status" validate:"min=1;max=2"`
	Date             *datatypes.Date `json:"date"`
	Note             *string         `json:"note"`
	Nominal          *float64        `json:"nominal"`
	Location         *string         `json:"location"`
	CreateBy_User_Id *uint           `json:"createBy_user_id"`
}

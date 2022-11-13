package prosesjambong

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Prosesjambong struct {
	Id                uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	JaminanBongkar_Id uint64          `json:"jaminanBongkar_Id"`
	Status            uint8           `json:"status"`
	Date              *datatypes.Date `json:"date"`
	Note              string          `json:"note"`
	Nominal           *float64        `json:"nominal"`
	Location          string          `json:"location"`
	CreateBy_User_Id  uint            `json:"createBy_user_id"`
	gormhelper.DateModel
}

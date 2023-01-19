package bphtb

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type LogApproval struct {
	Id       uint64    `json:"id" gorm:"primaryKey"`
	Sptpd_Id uuid.UUID `json:"sptpd_id" gorm:"type:uuid"`
	User_id  string    `json:"user_id" gorm:"type:varchar(50)"`
	Status   string    `json:"status" gorm:"type:varchar(50)"`
	gormhelper.DateModel
}

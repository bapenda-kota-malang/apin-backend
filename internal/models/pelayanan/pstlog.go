package pstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PstLogApproval struct {
	Id            uint64 `json:"id" gorm:"primaryKey"`
	Permohonan_Id uint64 `json:"permohonan_id"`
	User_id       string `json:"user_id" gorm:"type:varchar(50)"`
	Catatan       string `json:"catatan" gorm:"type:varchar(255)"`
	Status        string `json:"status" gorm:"type:varchar(50)"`
	gormhelper.DateModel
}

type RequestPSTLogApproval struct {
	Permohonan_Id uint64 `json:"permohonan_id"`
	User_id       string `json:"user_id"`
	Catatan       string `json:"catatan"`
	Status        string `json:"status"`
	gormhelper.DateModel
}

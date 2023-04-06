package logbankjatim

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"gorm.io/datatypes"
)

type LogBankJatim struct {
	Id           uint64         `json:"id" gorm:"primaryKey"`
	Path         string         `json:"path" gorm:"not null"`
	Content      datatypes.JSON `json:"content"  gorm:"not null"`
	Response     datatypes.JSON `json:"response"  gorm:"not null"`
	ResponseCode string         `json:"responseCode" gorm:"size:2;not null"`
	User_Id      uint64         `json:"userId"`
	Timestamp    time.Time      `json:"timestamp" gorm:"index"`
	User         *user.User     `json:"user,omitempty" gorm:"foreignKey:User_Id"`
}

type CreateDto struct {
	Path         string         `json:"path"`
	Content      datatypes.JSON `json:"content"`
	Response     datatypes.JSON `json:"response"`
	ResponseCode string         `json:"responseCode"`
	User_Id      uint64         `json:"userId"`
}

type CreateDynamicDto struct {
	Path         string `json:"path"`
	Content      any    `json:"content"`
	Response     any    `json:"response"`
	ResponseCode string `json:"responseCode"`
	User_Id      uint64 `json:"userId"`
}

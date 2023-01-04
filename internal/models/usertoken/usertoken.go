package usertoken

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

// type JenisResetPass string

type UserToken struct {
	Id         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	User_Email string     `json:"user_email"`
	Jenis      string     `json:"jenis" gorm:"size:20"`
	Token      uuid.UUID  `json:"token" gorm:"size:255"`
	ExpiredAt  *time.Time `json:"expiredAt"`
	gormhelper.DateModel
}

type CreateDto struct {
	User_Email string `json:"user_id" validate:"required;email"`
}

type UpdateDto struct {
	User_Email string `json:"user_id" validate:"required;email"`
	Jenis      string `json:"jenis" validate:"required"`
	Token      int16  `json:"position" validate:"required"`
}

const JenisConfirmByEmail = "ConfirmByEmail"
const JenisResetPass = "ResetPass"

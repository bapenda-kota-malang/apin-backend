package bapenagihandetail

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaPenagihanDetail struct {
	Id              uint64    `json:"id" gorm:"primaryKey"`
	BaPenagihan_Id  uuid.UUID `json:"baPenagihan_id"`
	Petugas_User_Id uint64    `json:"petugas_user_id"`

	BaPenagihan *bapenagihan.BaPenagihan `json:"baPenagihan,omitempty" gorm:"foreignKey:BaPenagihan_Id"`
	Petugas     *user.User               `json:"user,omitempty" gorm:"foreignKey:Petugas_User_Id"`

	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

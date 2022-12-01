package detail

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type BaPenagihanDetail struct {
	Id              uint64    `json:"id" gorm:"primaryKey"`
	BaPenagihan_Id  uuid.UUID `json:"baPenagihan_id" gorm:"type:uuid"`
	Petugas_User_Id uint64    `json:"petugas_user_id"`
	gormhelper.DateModel
	Petugas *user.User `json:"petugas,omitempty" gorm:"foreignKey:Petugas_User_Id"`
}

type CreateDto struct {
	BaPenagihan_Id  uuid.UUID `json:"-"`
	Petugas_User_Id uint64    `json:"petugas_user_id"`
}

type UpdateDto struct {
	BaPenagihan_Id  uuid.UUID `json:"-"`
	Petugas_User_Id uint64    `json:"petugas_user_id" validate:"required"`
	Deleted         bool      `json:"deleted"`
}

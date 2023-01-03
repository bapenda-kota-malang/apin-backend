package user

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type User struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string     `json:"name" gorm:"size:100;unique"`
	Salt        *string    `json:"salt,omitempty" gorm:"size:100"`
	Password    *string    `json:"password,omitempty" gorm:"size:200"`
	Position    int16      `json:"position"`
	Ref_Id      int        `json:"ref_id" gorm:"index"`
	Group_Id    int        `json:"group_id" gorm:"index"`
	Email       string     `json:"email" gorm:"size:100;unique"`
	Notes       string     `json:"notes"`
	SysAdmin    bool       `json:"sysAdmin"`
	ValidPeriod time.Time  `json:"validPeriod"`
	VeifyStatus int16      `json:"verifyStatus"`
	RegMode     int16      `json:"regMode"`
	Status      int16      `json:"status"`
	ApprovedAt  *time.Time `json:"approvedAt"`
	InactiveAt  *time.Time `json:"inactiveAt"`
	gormhelper.DateModel
}

type CreateDto struct {
	Name        string    `json:"name" validate:"required"`
	Password    *string   `json:"password" validate:"required"`
	Position    int16     `json:"position" validate:"required"`
	Status      int16     `json:"status"`
	Ref_Id      int       `json:"ref_id"`
	Group_Id    int       `json:"group_id" validate:"min=1"`
	RegMode     int16     `json:"regMode"`
	Email       string    `json:"email" validate:"validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type UpdateDto struct {
	Name     string `json:"name"`
	Position int16  `json:"position"`
	Group_Id int    `json:"group_id"`
	Email    string `json:"email"`
	Notes    string `json:"notes"`
	SysAdmin bool   `json:"sysAdmin"`
}

type RegisterDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"validemail"`
}

type LoginDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Position int16  `json:"position"`
	LongTerm bool   `json:"long"`
}

type FilterDto struct {
	Name     *string `json:"name"`
	Position *int16  `json:"position"`
	Email    *string `json:"email"`
	Status   *int16  `json:"status"`
	Ref_Id   *int    `json:"ref_id"`

	// fixed fields
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type VerifikasiDto struct {
	Status int16 `json:"status" validate:"required"`
}

type CheckerPThreeDto struct {
	Name       string  `json:"name" validate:"required"`
	Password   *string `json:"password,omitempty" validate:"minLength=6"`
	Email      string  `json:"email" validate:"validemail"`
	RePassword *string `json:"rePassword,omitempty" validate:"minLength=6"`
}

type ChangePassDto struct {
	OldPassword *string `json:"oldPassword,omitempty" validate:"minLength=6"`
	RePassword  *string `json:"rePassword,omitempty" validate:"minLength=6"`
	NewPassword *string `json:"newPassword,omitempty" validate:"minLength=6"`
}

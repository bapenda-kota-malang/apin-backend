package user

import "time"

type User struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:100;unique"`
	Salt        string    `json:"salt" gorm:"size:100"`
	Password    string    `json:"password" gorm:"size:200"`
	Position    int16     `json:"position"`
	Ref_Id      int       `json:"ref_id" gorm:"index"`
	Group_Id    int       `json:"group_id" gorm:"index"`
	Email       string    `json:"email" gorm:"size:100;unique"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
	VeifyStatus int16     `json:"verifyStatus"`
	RegMode     int16     `json:"regMode"`
	Status      int16     `json:"status"`
	ApprovedAt  time.Time `json:"approvedAt"`
	InactiveAt  time.Time `json:"inactiveAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type CreateDto struct {
	Name        string    `json:"name" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Position    int16     `json:"position" validate:"required"`
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
	Position string `json:"position"`
	Group_Id int    `json:"group_id"`
	Email    string `json:"email"`
	Notes    string `json:"notes"`
	SysAdmin bool   `json:"sysAdmin"`
}

type RegisterDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
}

type LoginDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Position int16  `json:"position"`
}

type FilterDto struct {
	Name     *string `json:"name"`
	Position *int16  `json:"position"`
	Email    *string `json:"email"`
	Status   *int16  `json:"status"`

	// fixed fields
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

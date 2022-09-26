package user

import "time"

type User struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"size:100"`
	Salt        string    `json:"salt" gorm:"size:100"`
	Password    string    `json:"password" gorm:"size:200"`
	Position    int16     `json:"position"`
	Ref_Id      int       `json:"ref_id" gorm:"index"`
	Group_Id    int       `json:"group_id" gorm:"index"`
	Email       string    `json:"email" gorm:"size:100;index"`
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

type Create struct {
	Name        string    `json:"name" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Position    string    `json:"position" validate:"required"`
	Group_Id    int       `json:"group_id" validate:"min=1"`
	Email       string    `json:"email" validate:"validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type Update struct {
	Name     string `json:"name"`
	Position string `json:"position" validate:"required"`
	Group_Id int    `json:"group_id" validate:"required"`
	Email    string `json:"email"`
	Notes    string `json:"notes"`
	SysAdmin bool   `json:"sysAdmin"`
}

type Register struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
}

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Filter struct {
	Name     string `json:"name"`
	Position int16  `json:"position"`
	Email    string `json:"email"`
	Status   int16  `json:"status"`
}

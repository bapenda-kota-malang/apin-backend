package user

import "time"

type User struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	Group_id     int       `json:"group"`
	Name         string    `json:"name"`
	Salt         string    `json:"salt"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	Position     string    `json:"position"`
	Ref_id       int       `json:"ref_id"`
	ValidPeriod  time.Time `json:"validPeriod"`
	Notes        string    `json:"notes"`
	SysAdmin     int       `json:"sysAdmin"`
	Status       int       `json:"status"`
	VerifyStatus int       `json:"verifyStatus"`
	InactiveAt   time.Time `json:"inactiveAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
	ApprovalAt   time.Time `json:"approvedAt"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// self creation
type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// by operator creation
type UserCreate struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Position string `json:"position" validate:"required"`
}

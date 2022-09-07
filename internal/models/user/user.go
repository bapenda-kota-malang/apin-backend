package user

import "time"

type User struct {
	Name        string    `json:"name"`
	Salt        string    `json:"salt"`
	Password    string    `json:"password"`
	Position    string    `json:"position"`
	Ref_id      int       `json:"ref_id"`
	Group_ID    int       `json:"group_id"`
	Email       string    `json:"email"`
	Notes       string    `json:"notes"`
	SysAdmin    int       `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
	VeifyStatus int       `json:"verifyStatus"`
	Status      int       `json:"status"`
	ApprovedAt  time.Time `json:"approvedAt"`
	InactiveAt  time.Time `json:"inactiveAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// self creation
type UserCreate struct {
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	Position     string    `json:"position" validate:"required"`
	Group_ID     int       `json:"group_id" validate:"required"`
	Email        string    `json:"email"`
	Notes        string    `json:"notes"`
	SysAdmin     bool      `json:"sysAdmin"`
	ValidPeriod  time.Time `json:"validPeriod"`
	VerifyStatus int       `json:"verifyStatus"`
}

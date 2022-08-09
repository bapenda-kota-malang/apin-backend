package usermodel

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Position string
	Ref_id   int64
	Salt     string
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

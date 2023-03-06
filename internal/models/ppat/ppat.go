package ppat

import "time"

type Ppat struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama   string `json:"nama" gorm:"size(100)"`
	Alamat string `json:"alamat" gorm:"size(200)"`
	Nik    string `json:"nik"`
}

type OutputDto struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	Nik        string `json:"nik"`
	User_Name  string `json:"user_name"`
	User_Email string `json:"user_email"`
	User_Notes string `json:"user_notes"`
}

// create ppat data alongside with user data
type CreateDto struct {
	// for ppat
	Nama   string `json:"nama" validate:"required;maxLength=100;alphabetspace"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required"`

	// for user
	User_Name     string    `json:"user_name" validate:"required;alphanumeric"`
	User_Password string    `json:"user_password" validate:"required"`
	User_Email    string    `json:"user_email" validate:"required;validemail"`
	User_Notes    string    `json:"user_notes"`
	ValidPeriod   time.Time `json:"validPeriod"`
}

type UpdateDto struct {
	// for ppat
	Nama   string `json:"nama" validate:"required;alphabetspace"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required"`

	// for user
	User_Email  string    `json:"user_email" validate:"required; validemail"`
	User_Notes  string    `json:"user_notes"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type FilterDto struct {
	Nama string `json:"nama"`
	Nik  string `json:"nip"`
}

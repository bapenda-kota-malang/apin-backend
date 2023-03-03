package ppat

import "time"

type Ppat struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama   string `json:"nama" gorm:"size:100"`
	Alamat string `json:"alamat" gorm:"size:200"`
	Nik    string `json:"nik"`
}

// create ppat data alongside with user data
type Create struct {
	// for ppat
	Nama   string `json:"nama" validate:"required;maxLength=100;alphabetspace"`
	Alamat string `json:"alamat" validate:"required;maxLength=200"`
	Nik    string `json:"nik" validate:"required;nik"`

	// for user
	Name     string `json:"name" validate:"required;alphanumeric"`
	Password string `json:"password" validate:"required"`
	// Position    int16     `json:"position"`
	// Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email" validate:"required;validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type Update struct {
	// for ppat
	Nama   string `json:"nama" validate:"required';alphabetspace"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required;nik"`

	// for user
	// Position    int16     `json:"position"`
	// Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email" validate:"required; validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type Filter struct {
	Nama string `json:"nama"`
	Nik  string `json:"nip"`
}

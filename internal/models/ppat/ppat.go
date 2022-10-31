package ppat

import "time"

type Ppat struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required"`
}

// create ppat data alongside with user data
type Create struct {
	// for ppat
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required"`

	// for user
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	// Position    int16     `json:"position"`
	// Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email" validate:"required; validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type Update struct {
	// for ppat
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required"`

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

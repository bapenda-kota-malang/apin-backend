package ppat

import "time"

type Ppat struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaLengkap string `json:"namaLengkap" validate:"required"`
	Alamat      string `json:"alamat" validate:"required"`
	Nik         string `json:"nik" validate:"required"`
}

// create ppat data alongside with user data
type CreateByUser struct {
	Name        string    `json:"name" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Position    int16     `json:"position" validate:"required"`
	Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
	NamaLengkap string    `json:"namaLengkap" validate:"required"`
	Alamat      string    `json:"alamat" validate:"required"`
	Nik         string    `json:"nik" validate:"required"`
}

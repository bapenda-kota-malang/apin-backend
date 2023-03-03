// model
package pegawai

import (
	"time"
)

type Pegawai struct {
	Id         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama       string    `json:"nama" gorm:"size:100"`
	Nip        string    `json:"nip" gorm:"size:30;unique"`
	Jabatan_Id int       `json:"jabatan_id" gorm:"index"`
	Pangkat_Id int       `json:"pangkat_id" gorm:"index"`
	Skpd_Id    int       `json:"skpd_id"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	Aktif      int16     `json:"aktif"`
}

// Creates pegawai data alongside with user data
type Create struct {
	// for pegawai
	Nama       string    `json:"nama" validate:"required;maxLength=100"`
	Nip        string    `json:"nip" validate:"required;maxLength=30"`
	Jabatan_Id int       `json:"jabatan_id" validate:"required;min=1"`
	Pangkat_Id int       `json:"pangkat_id" validate:"required;min=1"`
	Skpd_Id    int       `json:"skpd_id" validate:"required;min=1"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`

	// user.Create	// using embed composition is not good since structvalidation inlcudes the parent namespace
	// for user
	Name     string `json:"name" validate:"required;alphanumeric"`
	Password string `json:"password" validate:"required"`
	// Position    int16     `json:"position"`
	Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email" validate:"required; validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

// Creates pegawai data alongside with user data
type Update struct {
	// for pegawai
	Nama       string    `json:"nama" validate:"required;maxLength=100"`
	Nip        string    `json:"nip" validate:"required;maxLength=30"`
	Jabatan_Id int       `json:"jabatan_id" validate:"required;min=1"`
	Pangkat_Id int       `json:"pangkat_id" validate:"required;min=1"`
	Skpd_Id    int       `json:"skpd_id" validate:"required;min=1"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`

	// user.Create	// using composition is not good since structvalidation inlcudes the parent namespace
	// for user
	// Position    int16     `json:"position"`
	Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email" validate:"required; validemail"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
}

type Filter struct {
	// Nama string `json:"nama"`
	// Nip         string `json:"nip"`
	// Jabatan_Id  int    `json:"jabatan_id"`
	// Pangkat_Id  int    `json:"pangkat_id"`
	// Skpd_Id     int    `json:"skpd_id"`
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

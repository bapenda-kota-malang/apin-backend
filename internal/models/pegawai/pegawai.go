package pegawai

import (
	"time"
)

type Pegawai struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaLengkap string    `json:"namaLengkap" gorm:"size:100"`
	Nip         string    `json:"nip" gorm:"size:30;index"`
	Jabatan_Id  int       `json:"jabatan_id"`
	Pangkat_Id  int       `json:"pangkat_id"`
	Aktif       string    `json:"aktif"`
	Skpd_Id     string    `json:"skpd_id"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

// by operator creation
type CreateByUser struct {
	// user.Create	// using composition is not good since validation inlcudes the parent namespace
	// from user.Create
	Name        string    `json:"name" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Position    int16     `json:"position" validate:"required"`
	Group_Id    int       `json:"group_id" validate:"required;min=1"`
	Email       string    `json:"email"`
	Notes       string    `json:"notes"`
	SysAdmin    bool      `json:"sysAdmin"`
	ValidPeriod time.Time `json:"validPeriod"`
	//
	NamaLengkap string    `json:"namaLengkap" validate:"required"`
	Nip         string    `json:"nip" validate:"required"`
	Jabatan_Id  int       `json:"jabatan_id" validate:"required"`
	Pangkat_Id  int       `json:"pangkat_id" validate:"required"`
	Skpd_ID     int       `json:"skpd_id" validate:"required"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
}

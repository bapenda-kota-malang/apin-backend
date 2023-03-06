// model
package pegawai

import (
	"time"
)

type Pegawai struct {
	Id               int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama             string    `json:"nama" gorm:"size:100"`
	Nip              string    `json:"nip" gorm:"size:30;unique"`
	Jabatan_Id       int       `json:"jabatan_id" gorm:"index"`
	BidangKerja_Kode string    `json:"bidangKerja_kode" gorm:"index"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
	Aktif            int16     `json:"aktif"`
	// SatuanKerja_Id   int       `json:"satuanKerja_id"`
	// Pangkat_Id       int       `json:"pangkat_id" gorm:"index"`
}

type OutputDto struct {
	Id               int       `json:"id"`
	Nama             string    `json:"nama"`
	Nip              string    `json:"nip"`
	Jabatan_Id       int       `json:"jabatan_id"`
	BidangKerja_Kode string    `json:"bidangKerja_kode"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
	// SatuanKerja_Id   int       `json:"satuanKerja_id"`
	// Pangkat_Id       int       `json:"pangkat_id" gorm:"index"`
	Aktif           int16  `json:"aktif"`
	User_Group_Id   int    `json:"user_group_id"`
	User_Group_Name string `json:"user_group_name"`
	User_SysAdmin   bool   `json:"user_sysAdmin"`
	User_Notes      string `json:"user_notes"`
	User_Name       string `json:"user_name"`
	User_Email      string `json:"user_email"`
}

// Creates pegawai data alongside with user data
type CreateDto struct {
	// for pegawai
	Nama             string    `json:"nama" validate:"required;maxLength=100;alphabetspace"`
	Nip              string    `json:"nip" validate:"required;maxLength=30"`
	Jabatan_Id       int       `json:"jabatan_id" validate:"required;min=1"`
	BidangKerja_Kode string    `json:"bidangKerja_kode"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
	// SatuanKerja_Id   int       `json:"satuanKerja_id"`
	// Pangkat_Id       int       `json:"pangkat_id"`

	// user.Create	// using embed composition is not good since structvalidation inlcudes the parent namespace
	// for user
	User_Name        string    `json:"user_name" validate:"required;alphanumeric"`
	User_Password    string    `json:"user_password" validate:"required"`
	User_Group_Id    int       `json:"user_group_id" validate:"required;min=1"`
	User_Email       string    `json:"user_email" validate:"required; validemail"`
	User_SysAdmin    bool      `json:"user_sysAdmin"`
	User_Notes       string    `json:"user_notes"`
	User_ValidPeriod time.Time `json:"user_validPeriod"`
}

// Creates pegawai data alongside with user data
type UpdateDto struct {
	// for pegawai
	Nama             string    `json:"nama" validate:"required;maxLength=100;alphabetspace"`
	Nip              string    `json:"nip" validate:"required;maxLength=30"`
	Jabatan_Id       int       `json:"jabatan_id" validate:"required;min=1"`
	BidangKerja_Kode string    `json:"bidangKerja_kode" validate:"required"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
	// SatuanKerja_Id int       `json:"satuanKerja_id" validate:"required;min=1"`
	// Pangkat_Id     int       `json:"pangkat_id" validate:"required;min=1"`

	// user.Create	// using composition is not good since structvalidation inlcudes the parent namespace
	// for user
	// Position    int16     `json:"position"`
	User_Group_Id    int       `json:"user_group_id" validate:"required;min=1"`
	User_Email       string    `json:"user_email" validate:"required; validemail"`
	User_Notes       string    `json:"user_notes"`
	User_SysAdmin    bool      `json:"user_sysAdmin"`
	User_ValidPeriod time.Time `json:"user_validPeriod"`
}

type FilterDto struct {
	Nama             string  `json:"nama"`
	Nip              *string `json:"nip"`
	Jabatan_Id       *int    `json:"jabatan_id"`
	BidangKerja_Kode *int    `json:"bidangKerja_kode"`
	Page             int     `json:"page"`
	PageSize         int64   `json:"page_size"`
}

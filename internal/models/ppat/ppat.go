package ppat

import "github.com/bapenda-kota-malang/apin-backend/internal/models/user"

type Ppat struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaLengkap string `json:"namaLengkap" validate:"required"`
	Alamat      string `json:"alamat" validate:"required"`
	Nik         string `json:"nik" validate:"required"`
}

type CreateByUser struct {
	user.Create
	NamaLengkap string `json:"namaLengkap" validate:"required"`
	Alamat      string `json:"alamat" validate:"required"`
	Nik         string `json:"nik" validate:"required"`
}

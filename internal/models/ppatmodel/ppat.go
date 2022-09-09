package ppat

import "github.com/bapenda-kota-malang/apin-backend/internal/models/user"

type PpatCreate struct {
	user.UserCreate
	Nama   string `json:"nama" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	Nik    string `json:"nik" validate:"required"`
}

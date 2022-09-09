package pegawai

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

type Pegawai struct {
	Nama string `json:"string"`
}

// by operator creation
type PegawaiCreate struct {
	user.UserCreate
	Nama       string    `json:"nama" validate:"required"`
	Nip        string    `json:"nip" validate:"required"`
	Jabatan_ID string    `json:"jabatan_id" validate:"required"`
	Pangkat_ID string    `json:"pangkat_id" validate:"required"`
	Aktif      string    `json:"aktif" validate:"required"`
	Skpd_ID    string    `json:"skpd_id" validate:"required"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
}

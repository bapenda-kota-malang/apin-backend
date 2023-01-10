package kunjungan

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

// untuk spop
type Kunjungan struct {
	Id               uint64          `json:"id" gorm:"primaryKey"`
	TypeNo           bool            `json:"typeNo" gorm:"type:boolean"`
	NoKunjungan      string          `json:"noKunjungan" gorm:"type:varchar(50)"`
	JenisOP          *string         `json:"jenisOP" gorm:"type:varchar(50)"`
	NamaOP           *string         `json:"namaOP" gorm:"type:varchar(50)"`
	NamaWP           *string         `json:"namaWP" gorm:"type:varchar(50)"`
	AlamatWP         *string         `json:"alamatWP" gorm:"type:TEXT"`
	Agenda           *string         `json:"agenda" gorm:"type:TEXT"`
	TanggalKunjungan *datatypes.Date `json:"tanggalKunjungan"`
	gh.DateModel

	// Pegawais *[]KunjunganDetail `json:"pegawais,omitempty" gorm:"foreignKey:Id;references:Kunjungan_Id"`
}

type CreateDto struct {
	TypeNo           bool            `json:"typeNo"`
	NoKunjungan      string          `json:"noKunjungan"`
	JenisOP          *string         `json:"jenisOP"`
	NamaOP           *string         `json:"namaOP"`
	NamaWP           *string         `json:"namaWP"`
	AlamatWP         *string         `json:"alamatWP"`
	Agenda           *string         `json:"agenda"`
	TanggalKunjungan *datatypes.Date `json:"tanggalKunjungan"`

	Pegawais *[]KunjunganDetail `json:"pegawais"`
}

type UpdateDto struct {
	TypeNo           bool            `json:"typeNo"`
	NoKunjungan      string          `json:"noKunjungan"`
	JenisOP          *string         `json:"jenisOP"`
	NamaOP           *string         `json:"namaOP"`
	NamaWP           *string         `json:"namaWP"`
	AlamatWP         *string         `json:"alamatWP"`
	Agenda           *string         `json:"agenda"`
	TanggalKunjungan *datatypes.Date `json:"tanggalKunjungan"`

	Pegawais *[]KunjunganDetail `json:"pegawais"`
}

type FilterDto struct {
	TypeNo           bool            `json:"typeNo"`
	NoKunjungan      string          `json:"noKunjungan"`
	JenisOP          *string         `json:"jenisOP"`
	NamaOP           *string         `json:"namaOP"`
	NamaWP           *string         `json:"namaWP"`
	AlamatWP         *string         `json:"alamatWP"`
	Agenda           *string         `json:"agenda"`
	TanggalKunjungan *datatypes.Date `json:"tanggalKunjungan"`
	Page             int             `json:"page"`
	PageSize         int             `json:"page_size"`
}

package sinkronisasi

import (
	"gorm.io/datatypes"
)

type ExcelPbbBillerAgregat struct {
	Nop     string
	Nama    string
	Tahun   string
	Pokok   float64
	Denda   float64
	Nominal float64
	Tanggal datatypes.Date
	Jam     datatypes.Time
	Biller  string
}

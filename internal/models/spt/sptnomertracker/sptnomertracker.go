package sptnomertracker

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type SptNomerTracker struct {
	Id             uint64 `json:"id" gorm:"primaryKey"`
	Tahun          uint   `json:"tahun"`
	KodeJenisPajak string `json:"kodeJenisPajak" gorm:"size:5"`
	LastNumber     uint32 `json:"lastNumber"`
	gormhelper.DateModel
}

type Dto struct {
	Tahun          *uint   `json:"tahun"`
	KodeJenisPajak *string `json:"kodeJenisPajak"`
	LastNumber     *uint32 `json:"lastNumber"`
}

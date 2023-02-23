package tarif

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type Tarif struct {
	Id            uint64 `json:"id" gorm:"primaryKey"`
	Provinsi_Kode string `json:"provinsi_kd" gorm:"type:varchar(2);not null"`
	Daerah_Kode   string `json:"daerah_kd" gorm:"type:varchar(2);not null"`
	TahunAwal     string `json:"tahunAwal" gorm:"size:4;not null"`
	TahunAkhir    string `json:"tahunAkhir" gorm:"size:4"`
	NjopMin       int64  `json:"njopMin" gorm:"size:15"`
	NjopMax       int64  `json:"njopMax" gorm:"size:15"`
	NilaiTarif    float64
	gormhelper.DateModel
}

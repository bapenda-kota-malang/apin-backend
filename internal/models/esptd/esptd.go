package esptd

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Espt struct {
	Id               uint           `json:"id" gorm:"primarykey"`
	Npwpd_Id         uint           `json:"npwp_id"`
	ObjekPajak_Id    uint           `json:"objekPajak_id"`
	Rekening_Id      uint           `json:"rekening_id"`
	Location         string         `gorm:"size:50"`
	Description      string         `json:"description"`
	PeriodeAwal      datatypes.Date `json:"periodeAwal"`
	PeriodeAkhir     datatypes.Date `json:"periodeAkhir"`
	TarifPajak_id    uint           `json:"tarifPajak_id"`
	EtaxSubTotal     string         `gorm:"size:100"`
	EtaxTotal        string         `gorm:"size:100"`
	EtaxJumlahPajak  string         `gorm:"size:100"`
	Omset            float64        `json:"omset"`
	JumlahPajak      float32        `json:"jumlahPajak"`
	Attachment       string         `json:"attachment"`
	JatuhTempo       datatypes.Date `json:"jatuhTempo"`
	Sunset           datatypes.Date `json:"sunset"`
	LaporBy_User_Id  uint           `json:"laporBy_user_id"`
	VerifyBy_User_Id uint           `json:"verifyBy_user_id"`
	VerifyStatus     string         `json:"verifyStatus" gorm:"size:20"`
	gormhelper.DateModel
}

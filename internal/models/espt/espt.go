package espt

import (
	mdea "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	mdeh "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	mdep "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	mderek "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptreklame"
	mderes "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Espt struct {
	Id                uint                        `json:"id" gorm:"primarykey"`
	Npwpd_Id          uint                        `json:"npwp_id"`
	ObjekPajak_Id     uint                        `json:"objekPajak_id"`
	Rekening_Id       uint                        `json:"rekening_id"`
	Location          string                      `json:"location" gorm:"size:50"`
	Description       *string                     `json:"description"`
	PeriodeAwal       datatypes.Date              `json:"periodeAwal"`
	PeriodeAkhir      datatypes.Date              `json:"periodeAkhir"`
	TarifPajak_id     uint                        `json:"tarifPajak_id"`
	EtaxSubTotal      string                      `json:"etaxSubTotal" gorm:"size:100"`
	EtaxTotal         string                      `json:"etaxTotal" gorm:"size:100"`
	EtaxJumlahPajak   string                      `json:"etaxJumlahPajak" gorm:"size:100"`
	Omset             float64                     `json:"omset"`
	JumlahPajak       float32                     `json:"jumlahPajak"`
	Attachment        string                      `json:"attachment"`
	JatuhTempo        datatypes.Date              `json:"jatuhTempo"`
	Sunset            datatypes.Date              `json:"sunset"`
	LaporBy_User_Id   uint                        `json:"laporBy_user_id"`
	VerifyBy_User_Id  uint                        `json:"verifyBy_user_id"`
	VerifyStatus      string                      `json:"verifyStatus" gorm:"size:20"`
	DetailEsptAir     *[]mdea.DetailEsptAir       `json:"detailEsptAir,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptHotel   *[]mdeh.DetailEsptHotel     `json:"detailEsptHotel,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptParkir  *[]mdep.DetailEsptParkir    `json:"detailEsptParkir,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptReklame *[]mderek.DetailEsptReklame `json:"detailEsptReklame,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptResto   *[]mderes.DetailEsptResto   `json:"detailEsptResto,omitempty" gorm:"foreignKey:Espt_Id"`
	gormhelper.DateModel
}

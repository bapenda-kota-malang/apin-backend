package espt

import (
	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	mdhib "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthiburan"
	mdhot "github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	mdpar "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptparkir"
	mdrek "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptreklame"
	mdres "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptresto"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Espt struct {
	Id               uint            `json:"id" gorm:"primarykey"`
	Npwpd_Id         uint            `json:"npwp_id"`
	ObjekPajak_Id    uint            `json:"objekPajak_id"`
	Rekening_Id      uint            `json:"rekening_id"`
	Location         *string         `json:"location,omitempty" gorm:"size:50"`
	LuasLokasi       *uint           `json:"luasLokasi,omitempty" gorm:"comment:untuk pajak parkir"`
	Description      *string         `json:"description,omitempty"`
	PeriodeAwal      datatypes.Date  `json:"periodeAwal"`
	PeriodeAkhir     datatypes.Date  `json:"periodeAkhir"`
	TarifPajak_id    uint            `json:"tarifPajak_id"`
	EtaxSubTotal     *string         `json:"etaxSubTotal,omitempty" gorm:"size:100"`
	EtaxTotal        *string         `json:"etaxTotal,omitempty" gorm:"size:100"`
	EtaxJumlahPajak  *string         `json:"etaxJumlahPajak,omitempty" gorm:"size:100"`
	Omset            float64         `json:"omset"`
	JumlahPajak      float32         `json:"jumlahPajak"`
	Attachment       string          `json:"attachment"`
	JatuhTempo       datatypes.Date  `json:"jatuhTempo"`
	Sunset           *datatypes.Date `json:"sunset,omitempty"`
	LaporBy_User_Id  uint            `json:"laporBy_user_id"`
	VerifyBy_User_Id *uint           `json:"verifyBy_user_id"`
	VerifyStatus     Status          `json:"verifyStatus" gorm:"size:20"`
	// Rekening         *rekening.Rekening        `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	LaporUser         *user.User                 `json:"laporUser,omitempty" gorm:"foreignKey:LaporBy_User_Id"`
	VerifyUser        *user.User                 `json:"verifyUser,omitempty" gorm:"foreignKey:VerifyBy_User_Id"`
	DetailEsptAir     *[]mdair.DetailEsptAir     `json:"detailEsptAir,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptHotel   *[]mdhot.DetailEsptHotel   `json:"detailEsptHotel,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptHiburan *[]mdhib.DetailEsptHiburan `json:"detailEsptHiburan,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptParkir  *[]mdpar.DetailEsptParkir  `json:"detailEsptParkir,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptReklame *[]mdrek.DetailEsptReklame `json:"detailEsptReklame,omitempty" gorm:"foreignKey:Espt_Id"`
	DetailEsptResto   *[]mdres.DetailEsptResto   `json:"detailEsptResto,omitempty" gorm:"foreignKey:Espt_Id"`
	gormhelper.DateModel
}

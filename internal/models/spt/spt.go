package spt

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Spt struct {
	Id               uint64          `json:"id" gorm:"primaryKey"`
	Npwpd_Id         *uint64         `json:"npwpd_Id"`
	ObjekPajak_Id    *uint64         `json:"objekPajak_Id"`
	Rekening_Id      *uint64         `json:"rekening_Id"`
	LuasLokasi       *string         `json:"luasLokasi" gorm:"type:varchar(50)"`
	Description      *string         `json:"description" gorm:"type:varchar(255)"`
	PeriodeAwal      *datatypes.Date `json:"periodeAwal,omitempty"`
	PeriodeAkhir     *datatypes.Date `json:"periodeAkhir,omitempty"`
	TarifPajak_Id    uint            `json:"tarifPajak_id"`
	EtaxSubTotal     *string         `json:"etaxSubTotal,omitempty" gorm:"size:100"`
	EtaxTotal        *string         `json:"etaxTotal,omitempty" gorm:"size:100"`
	EtaxJumlahPajak  *string         `json:"etaxJumlahPajak,omitempty" gorm:"size:100"`
	Omset            float64         `json:"omset"`
	JumlahPajak      float32         `json:"jumlahPajak"`
	Lampiran         string          `json:"lampiran"`
	JatuhTempo       datatypes.Date  `json:"jatuhTempo"`
	Sunset           *datatypes.Date `json:"sunset,omitempty"`
	CreateBy_User_Id uint            `json:"createBy_user_id"`
	TanggalSpt       time.Time       `json:"tanggalSpt"`
	NomorSpt         *string         `json:"NomorSpt" gorm:"type:varchar(20)"`
	KodeBilling      *string         `json:"kodeBilling" gorm:"varchar(30)"`
	Type             *nt.JenisPajak  `json:"type" gorm:"type:varchar(2)"`
	Status           *SptStatus      `json:"status" gorm:"type:varchar(5)"`
	TanggalLunas     *time.Time      `json:"tanggalLunas"`
	BatalPenetapan   *string         `json:"batalPenetapan" gorm:"type:char(1)"`
	IdBT             *uint64         `json:"idBT"`
	JumlahTahun      *float64        `json:"jumlahTahun" gorm:"type:decimal"`
	JumlahBulan      *float64        `json:"jumlahBulan" gorm:"type:decimal"`
	JumlahMinggu     *float64        `json:"jumlahMinggu" gorm:"type:decimal"`
	JumlahHari       *float64        `json:"jumlahHari" gorm:"type:decimal"`
	Gambar           *string         `json:"gambar" gorm:"type:varchar(255)"`
	Keterangan       *string         `json:"keterangan" gorm:"type:varchar(255)"`
	KoefisienPajak   *uint64         `json:"koefisienPajak"`
	NamaProduk       *string         `json:"productName" gorm:"type:varchar(200)"`
	NomorRegister    *string         `json:"registerNumber" gorm:"type:varchar(100)"`
	VaJatim          *string         `json:"vaJatim" gorm:"type:varchar(20)"`
	JenisKetetapan   *string         `json:"jenisKetetapan"`
	Kenaikan         *float64        `json:"kenaikan" gorm:"type:decimal"`
	Bunga            *float64        `json:"bunga" gorm:"type:decimal"`
	Denda            *float64        `json:"denda" gorm:"type:decimal"`
	Pengurangan      *float64        `json:"pengurangan" gorm:"type:decimal"`
	Total            *float64        `json:"total" gorm:"type:decimal"`
	Ref_Spt_Id       *uint64         `json:"ref_spt_id"`
	BillingPenetapan *string         `json:"billingPenetapan" gorm:"type:varchar(20)"`
	Teguran_Id       *uint64         `json:"teguran_id"`
	IsTeguran        bool            `json:"isTeguran"`
	// IsCancelled         bool            `json:"isCancelled"`
	KeteranganPenetapan *string `json:"keteranganPenetapan" gorm:"type:varchar(100)"`
	Kasubid_User_Id     *string `json:"kasubid_user_id"`
	Kabid_User_Id       *string `json:"kabid_user_id"`
	gormhelper.DateModel
	CancelledAt           *time.Time                      `json:"cancelledAt"`
	Npwpd                 *npwpd.Npwpd                    `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	ObjekPajak            *op.ObjekPajak                  `json:"objekPajak,omitempty" gorm:"foreignKey:ObjekPajak_Id"`
	Rekening              *rekening.Rekening              `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	CreateUser            *user.User                      `json:"createBy,omitempty" gorm:"foreignKey:CreateBy_User_Id"`
	KasubidUser           *user.User                      `json:"kasubid,omitempty" gorm:"foreignKey:Kasubid_User_Id"`
	KabidUser             *user.User                      `json:"kabid,omitempty" gorm:"foreignKey:Kabid_User_Id"`
	DetailSptAir          *[]mdsa.DetailSptAir            `json:"detailSptAir,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptHotel        *[]mdsh.DetailSptHotel          `json:"detailSptHotel,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptParkir       *[]mdsp.DetailSptParkir         `json:"detailSptParkir,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptReklame      *[]mdsrek.DetailSptReklame      `json:"detailSptReklame,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptResto        *[]mdsres.DetailSptResto        `json:"detailSptResto,omitempty" gorm:"foreignKey:Spt_Id"`
	JaminanBongkarReklame *[]mdsjbr.JaminanBongkarReklame `json:"jaminanBongkarReklame,omitempty" gorm:"foreignKey:Spt_Id"`
	// Teguran     *teguran.Teguran `json:"teguran,omitempty" gorm:"foreignKey:Teguran_Id"`
}

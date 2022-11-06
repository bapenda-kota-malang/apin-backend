package spt

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdhiburan "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthiburan"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdnonpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjnonpln"
	mdpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjpln"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Spt struct {
	Id                  uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Npwpd_Id            uint64          `json:"npwpd_Id"`
	ObjekPajak_Id       uint64          `json:"objekPajak_Id"`
	Rekening_Id         uint64          `json:"rekening_Id"`
	LuasLokasi          *string         `json:"luasLokasi,omitempty" gorm:"type:varchar(50)"`
	Description         *string         `json:"description,omitempty" gorm:"type:varchar(255)"`
	PeriodeAwal         datatypes.Date  `json:"periodeAwal"`
	PeriodeAkhir        datatypes.Date  `json:"periodeAkhir"`
	TarifPajak_Id       uint64          `json:"tarifPajak_id"`
	EtaxSubTotal        *string         `json:"etaxSubTotal,omitempty" gorm:"size:100"`
	EtaxTotal           *string         `json:"etaxTotal,omitempty" gorm:"size:100"`
	EtaxJumlahPajak     *string         `json:"etaxJumlahPajak,omitempty" gorm:"size:100"`
	Omset               float64         `json:"omset"`
	JumlahPajak         float32         `json:"jumlahPajak"`
	Lampiran            string          `json:"lampiran"`
	JatuhTempo          datatypes.Date  `json:"jatuhTempo"`
	Sunset              *datatypes.Date `json:"sunset,omitempty"`
	CreateBy_User_Id    uint            `json:"createBy_user_id"`
	TanggalSpt          time.Time       `json:"tanggalSpt"`
	NomorSpt            string          `json:"NomorSpt" gorm:"type:varchar(20)"`
	KodeBilling         string          `json:"kodeBilling" gorm:"varchar(30)"`
	Type                mt.JenisPajak   `json:"type" gorm:"type:varchar(2)"`
	Status              SptStatus       `json:"status" gorm:"type:varchar(5)"`
	TanggalLunas        *time.Time      `json:"tanggalLunas,omitempty"`
	BatalPenetapan      bool            `json:"batalPenetapan,omitempty"`
	IdBT                *uint64         `json:"idBT,omitempty"`
	JumlahTahun         *float64        `json:"jumlahTahun,omitempty" gorm:"type:decimal"`
	JumlahBulan         *float64        `json:"jumlahBulan,omitempty" gorm:"type:decimal"`
	JumlahMinggu        *float64        `json:"jumlahMinggu,omitempty" gorm:"type:decimal"`
	JumlahHari          *float64        `json:"jumlahHari,omitempty" gorm:"type:decimal"`
	Gambar              *string         `json:"gambar,omitempty" gorm:"type:varchar(255)"`
	Keterangan          *string         `json:"keterangan,omitempty" gorm:"type:varchar(255)"`
	KoefisienPajak      *uint64         `json:"koefisienPajak,omitempty"`
	NamaProduk          *string         `json:"productName,omitempty" gorm:"type:varchar(200)"`
	NomorRegister       *string         `json:"registerNumber,omitempty" gorm:"type:varchar(100)"`
	VaJatim             *string         `json:"vaJatim,omitempty" gorm:"type:varchar(20)"`
	JenisKetetapan      *JenisKetetapan `json:"jenisKetetapan,omitempty"`
	Kenaikan            *float64        `json:"kenaikan,omitempty" gorm:"type:decimal"`
	Bunga               *float64        `json:"bunga,omitempty" gorm:"type:decimal"`
	Denda               *float64        `json:"denda,omitempty" gorm:"type:decimal"`
	Pengurangan         *float64        `json:"pengurangan,omitempty" gorm:"type:decimal"`
	Total               *float64        `json:"total,omitempty" gorm:"type:decimal"`
	Ref_Spt_Id          *uint64         `json:"ref_spt_id,omitempty"`
	BillingPenetapan    *string         `json:"billingPenetapan,omitempty" gorm:"type:varchar(20)"`
	Teguran_Id          *uint64         `json:"teguran_id,omitempty"`
	IsTeguran           bool            `json:"isTeguran,omitempty"`
	KeteranganPenetapan *string         `json:"keteranganPenetapan,omitempty" gorm:"type:varchar(100)"`
	Kasubid_User_Id     *string         `json:"kasubid_user_id,omitempty"`
	Kabid_User_Id       *string         `json:"kabid_user_id,omitempty"`
	gormhelper.DateModel
	CancelledAt *time.Time   `json:"cancelledAt,omitempty"`
	Npwpd       *npwpd.Npwpd `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	// ObjekPajak            *op.ObjekPajak                  `json:"objekPajak,omitempty" gorm:"foreignKey:ObjekPajak_Id"`
	Rekening              *rekening.Rekening              `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	CreateUser            *user.User                      `json:"createBy,omitempty" gorm:"foreignKey:CreateBy_User_Id"`
	KasubidUser           *user.User                      `json:"kasubid,omitempty" gorm:"foreignKey:Kasubid_User_Id"`
	KabidUser             *user.User                      `json:"kabid,omitempty" gorm:"foreignKey:Kabid_User_Id"`
	DetailSptAir          *mdsa.DetailSptAir              `json:"detailSptAir,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptHiburan      *mdhiburan.DetailSptHiburan     `json:"detailSptHiburan,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptHotel        *[]mdsh.DetailSptHotel          `json:"detailSptHotel,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptParkir       *[]mdsp.DetailSptParkir         `json:"detailSptParkir,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptNonPln       *mdnonpln.DetailSptPpjNonPln    `json:"detailSptPpjNonPln,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptPln          *[]mdpln.DetailSptPpjPln        `json:"detailSptPpjPln,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptReklame      *[]mdsrek.DetailSptReklame      `json:"detailSptReklame,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptResto        *mdsres.DetailSptResto          `json:"detailSptResto,omitempty" gorm:"foreignKey:Spt_Id"`
	JaminanBongkarReklame *[]mdsjbr.JaminanBongkarReklame `json:"jaminanBongkarReklame,omitempty" gorm:"foreignKey:Spt_Id"`
	// Teguran     *teguran.Teguran `json:"teguran,omitempty" gorm:"foreignKey:Teguran_Id"`
}

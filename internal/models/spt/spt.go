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
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Spt struct {
	Id                uint64             `json:"id" gorm:"primaryKey"`
	ObjekPajak_Id     *uint64            `json:"objekPajak_id"`
	ObjekPajak        *op.ObjekPajak     `json:"objekPajak,omitempty" gorm:"foreignKey:ObjekPajak_Id"`
	Rekening_Id       *uint64            `json:"rekening_id"`
	Rekening          *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	SptDate           time.Time          `json:"sptDate"`
	SptNumber         *string            `json:"sptNumber" gorm:"type:varchar(20)"`
	BillingNumber     *uint64            `json:"billingNumber" gorm:"type:numeric"`
	PaymentCode       *string            `json:"paymentCode" gorm:"varchar(30)"`
	Type              *nt.JenisPajak     `json:"type" gorm:"type:varchar(2)"`
	Status            *mt.SptStatus      `json:"status" gorm:"type:varchar(5)"`
	Nama_objekPajak   *string            `json:"nama_objekPajak" gorm:"type:varchar(50)"`
	Alamat_objekPajak *string            `json:"alamat_objekPajak" gorm:"type:varchar(200)"`
	Location          *string            `json:"location" gorm:"type:varchar(50)"`
	Description       *string            `json:"description" gorm:"type:varchar(255)"`
	StartDate         datatypes.Date     `json:"startDate"`
	EndDate           datatypes.Date     `json:"endDate"`
	DueDate           datatypes.Date     `json:"dueDate"`
	Jumlah            *float64           `json:"jumlah" gorm:"type:decimal"`
	TarifRp           *float64           `json:"tarifRp" gorm:"type:decimal"`
	TarifPersen       *float64           `json:"tarifPersen" gorm:"type:decimal"`
	JumlahPajak       *float64           `json:"jumlahPajak" gorm:"type:decimal"`
	TanggalLunas      *time.Time         `json:"tanggalLunas"`
	Npwpd             *string            `json:"npwpd" gorm:"type:varchar(20)"`
	Npwpd_Id          *uint64            `json:"npwpd_id,omitempty" gorm:"type:varchar(20)"`
	NpwpdR            *npwpd.Npwpd       `json:"npwpdr,omitempty" gorm:"foreignKey:Npwpd_Id"`
	BatalPenetapan    *string            `json:"batalPenetapan" gorm:"type:char(1)"`
	IdBT              *uint64            `json:"idBT"`
	User_Name         *string            `json:"user_name" gorm:"type:varchar(100)"`
	User_Id           *uint64            `json:"user_id"`
	User              *user.User         `gorm:"foreignKey:User_Id"`
	JumlahTahun       *float64           `json:"jumlahTahun" gorm:"type:decimal"`
	JumlahBulan       *float64           `json:"jumlahBulan" gorm:"type:decimal"`
	JumlahMinggu      *float64           `json:"jumlahMinggu" gorm:"type:decimal"`
	Picture           *string            `json:"picture" gorm:"type:varchar(255)"`
	Notes             *string            `json:"notes" gorm:"type:varchar(255)"`
	KoefisienPajak    *uint64            `json:"koefisienPajak"`
	ProductName       *string            `json:"productName" gorm:"type:varchar(200)"`
	RegisterNumber    *string            `json:"registerNumber" gorm:"type:varchar(100)"`
	VaJatim           *string            `json:"vaJatim" gorm:"type:varchar(20)"`
	EtaxSubtotal      *string            `json:"etaxSubtotal" gorm:"type:varchar(100)"`
	EtaxTotal         *string            `json:"etaxTotal" gorm:"type:varchar(100)"`
	EtaxJumlahPajak   *string            `json:"etaxJumlahPajak" gorm:"type:varchar(100)"`
	JenisKetetapan    *string            `json:"jenisKetetapan"`
	Kenaikan          *float64           `json:"kenaikan" gorm:"type:decimal"`
	Bunga             *float64           `json:"bunga" gorm:"type:decimal"`
	Denda             *float64           `json:"denda" gorm:"type:decimal"`
	Pengurangan       *float64           `json:"pengurangan" gorm:"type:decimal"`
	Total             *float64           `json:"total" gorm:"type:decimal"`
	Ref_Spt_Id        *uint64            `json:"ref_spt_id"`
	BillingPenetapan  *string            `json:"billingPenetapan" gorm:"type:varchar(20)"`
	// Teguran_Id *uint64 `json:"teguran_id"`
	// Teguran *teguran.Teguran `json:"teguran,omitempty" gorm:"foreignKey:Teguran_Id"`
	IsTeguran   *string `json:"isTeguran" gorm:"type:char(1)"`
	IsCancelled *string `json:"isCancelled" gorm:"type:char(1)"`
	Note        *string `json:"note" gorm:"type:varchar(100)"`
	gormhelper.DateModel
	CancelledAt           *time.Time                      `json:"cancelledAt"`
	Kasubid_User_Id       *string                         `json:"kasubid_user_id" gorm:"type:varchar(100)"`
	Kabid_User_Id         *string                         `json:"kabid_user_id" gorm:"type:varchar(100)"`
	DetailSptAir          *[]mdsa.DetailSptAir            `json:"detailSptAir,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptHotel        *[]mdsh.DetailSptHotel          `json:"detailSptHotel,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptParkir       *[]mdsp.DetailSptParkir         `json:"detailSptParkir,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptReklame      *[]mdsrek.DetailSptReklame      `json:"detailSptReklame,omitempty" gorm:"foreignKey:Spt_Id"`
	DetailSptResto        *[]mdsres.DetailSptResto        `json:"detailSptResto,omitempty" gorm:"foreignKey:Spt_Id"`
	JaminanBongkarReklame *[]mdsjbr.JaminanBongkarReklame `json:"jaminanBongkarReklame,omitempty" gorm:"foreignKey:Spt_Id"`
}

type UpdateDto struct {
	ObjekPajak_Id     *uint64        `json:"objekPajak_id"`
	Rekening_Id       *uint64        `json:"rekening_id"`
	SptDate           string         `json:"sptDate"`
	SptNumber         *string        `json:"sptNumber"`
	BillingNumber     *uint64        `json:"billingNumber"`
	PaymentCode       *string        `json:"payment"`
	Type              *nt.JenisPajak `json:"type"`
	Status            *mt.SptStatus  `json:"status"`
	Nama_objekPajak   *string        `json:"nama_objekPajak"`
	Alamat_objekPajak *string        `json:"alamat_objekPajak"`
	Location          *string        `json:"location"`
	Description       *string        `json:"description"`
	StartDate         string         `json:"startDate"`
	EndDate           string         `json:"endDate"`
	DueDate           string         `json:"dueDate"`
	Jumlah            *float64       `json:"jumlah"`
	TarifRp           *float64       `json:"tarifRp"`
	TarifPersen       *float64       `json:"tarifPersen"`
	JumlahPajak       *float64       `json:"jumlahPajak"`
	TanggalLunas      string         `json:"tanggalLunas"`
	Npwpd             *string        `json:"npwpd"`
	Npwpd_Id          *uint64        `json:"npwpd_id"`
	BatalPenetapan    *string        `json:"batalPenetapan"`
	IdBT              *uint64        `json:"idBT"`
	User_Name         *string        `json:"user_name"`
	User_Id           *uint64        `json:"user_id"`
	JumlahTahun       *float64       `json:"jumlahTahun"`
	JumlahBulan       *float64       `json:"jumlahBulan"`
	JumlahMinggu      *float64       `json:"jumlahMinggu"`
	Picture           *string        `json:"picture"`
	Notes             *string        `json:"notes"`
	KoefisienPajak    *uint64        `json:"koefisienPajak"`
	ProductName       *string        `json:"productName"`
	RegisterNumber    *string        `json:"registerNumber"`
	VaJatim           *string        `json:"vaJatim"`
	EtaxSubtotal      *string        `json:"etaxSubtotal"`
	EtaxTotal         *string        `json:"etaxTotal"`
	EtaxJumlahPajak   *string        `json:"etaxJumlahPajak"`
	JenisKetetapan    *string        `json:"jenisKetetapan"`
	Kenaikan          *float64       `json:"kenaikan"`
	Bunga             *float64       `json:"bunga"`
	Denda             *float64       `json:"denda"`
	Pengurangan       *float64       `json:"pengurangan"`
	Total             *float64       `json:"total"`
	Ref_Spt_Id        *uint64        `json:"ref_spt_id"`
	BillingPenetapan  *string        `json:"billingPenetapan"`
	// Teguran_Id *uint64 `json:"teguran_id"`
	IsTeguran   *string `json:"isTeguran"`
	IsCancelled *string `json:"isCancelled"`
	Note        *string `json:"note"`
	gormhelper.DateModel
	CancelledAt     string  `json:"cancelledAt"`
	Kasubid_User_Id *string `json:"kasubid_user_id"`
	Kabid_User_Id   *string `json:"kabid_user_id"`
}

type FilterDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

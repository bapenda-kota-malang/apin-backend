package spt

import (
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	mdsa "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdsh "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailspthotel"
	mdsp "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptparkir"
	mdsrek "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mdsres "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptresto"
	mdsjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/types"
)

type CreateAirDto struct {
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
	IsTeguran       *string `json:"isTeguran"`
	IsCancelled     *string `json:"isCancelled"`
	Note            *string `json:"note"`
	CancelledAt     string  `json:"cancelledAt"`
	Kasubid_User_Id *string `json:"kasubid_user_id"`
	Kabid_User_Id   *string `json:"kabid_user_id"`
	mdsa.CreateDto
}

type CreateHotelDto struct {
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
	IsTeguran       *string          `json:"isTeguran"`
	IsCancelled     *string          `json:"isCancelled"`
	Note            *string          `json:"note"`
	CancelledAt     string           `json:"cancelledAt"`
	Kasubid_User_Id *string          `json:"kasubid_user_id"`
	Kabid_User_Id   *string          `json:"kabid_user_id"`
	DetailSpt       []mdsh.CreateDto `json:"detailSpt"`
}

type CreateParkirDto struct {
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
	IsTeguran       *string          `json:"isTeguran"`
	IsCancelled     *string          `json:"isCancelled"`
	Note            *string          `json:"note"`
	CancelledAt     string           `json:"cancelledAt"`
	Kasubid_User_Id *string          `json:"kasubid_user_id"`
	Kabid_User_Id   *string          `json:"kabid_user_id"`
	DetailSpt       []mdsp.CreateDto `json:"detailSpt"`
}

type CreateReklameDto struct {
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
	IsTeguran             *string          `json:"isTeguran"`
	IsCancelled           *string          `json:"isCancelled"`
	Note                  *string          `json:"note"`
	CancelledAt           string           `json:"cancelledAt"`
	Kasubid_User_Id       *string          `json:"kasubid_user_id"`
	Kabid_User_Id         *string          `json:"kabid_user_id"`
	DetailSptReklame      mdsrek.CreateDto `json:"detailSptReklame"`
	JaminanBongkarReklame mdsjbr.CreateDto `json:"jaminanBongkarReklame"`
}

type CreateRestoDto struct {
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
	IsTeguran       *string `json:"isTeguran"`
	IsCancelled     *string `json:"isCancelled"`
	Note            *string `json:"note"`
	CancelledAt     string  `json:"cancelledAt"`
	Kasubid_User_Id *string `json:"kasubid_user_id"`
	Kabid_User_Id   *string `json:"kabid_user_id"`
	mdsres.CreateDto
}

// type CreateJaminanDto struct {
// 	ObjekPajak_Id     *uint64        `json:"objekPajak_id"`
// 	Rekening_Id       *uint64        `json:"rekening_id"`
// 	SptDate           string         `json:"sptDate"`
// 	SptNumber         *string        `json:"sptNumber"`
// 	BillingNumber     *uint64        `json:"billingNumber"`
// 	PaymentCode       *string        `json:"payment"`
// 	Type              *nt.JenisPajak `json:"type"`
// 	Status            *mt.SptStatus  `json:"status"`
// 	Nama_objekPajak   *string        `json:"nama_objekPajak"`
// 	Alamat_objekPajak *string        `json:"alamat_objekPajak"`
// 	Location          *string        `json:"location"`
// 	Description       *string        `json:"description"`
// 	StartDate         string         `json:"startDate"`
// 	EndDate           string         `json:"endDate"`
// 	DueDate           string         `json:"dueDate"`
// 	Jumlah            *float64       `json:"jumlah"`
// 	TarifRp           *float64       `json:"tarifRp"`
// 	TarifPersen       *float64       `json:"tarifPersen"`
// 	JumlahPajak       *float64       `json:"jumlahPajak"`
// 	TanggalLunas      string         `json:"tanggalLunas"`
// 	Npwpd             *string        `json:"npwpd"`
// 	Npwpd_Id          *uint64        `json:"npwpd_id"`
// 	BatalPenetapan    *string        `json:"batalPenetapan"`
// 	IdBT              *uint64        `json:"idBT"`
// 	User_Name         *string        `json:"user_name"`
// 	User_Id           *uint64        `json:"user_id"`
// 	JumlahTahun       *float64       `json:"jumlahTahun"`
// 	JumlahBulan       *float64       `json:"jumlahBulan"`
// 	JumlahMinggu      *float64       `json:"jumlahMinggu"`
// 	Picture           *string        `json:"picture"`
// 	Notes             *string        `json:"notes"`
// 	KoefisienPajak    *uint64        `json:"koefisienPajak"`
// 	ProductName       *string        `json:"productName"`
// 	RegisterNumber    *string        `json:"registerNumber"`
// 	VaJatim           *string        `json:"vaJatim"`
// 	EtaxSubtotal      *string        `json:"etaxSubtotal"`
// 	EtaxTotal         *string        `json:"etaxTotal"`
// 	EtaxJumlahPajak   *string        `json:"etaxJumlahPajak"`
// 	JenisKetetapan    *string        `json:"jenisKetetapan"`
// 	Kenaikan          *float64       `json:"kenaikan"`
// 	Bunga             *float64       `json:"bunga"`
// 	Denda             *float64       `json:"denda"`
// 	Pengurangan       *float64       `json:"pengurangan"`
// 	Total             *float64       `json:"total"`
// 	Ref_Spt_Id        *uint64        `json:"ref_spt_id"`
// 	BillingPenetapan  *string        `json:"billingPenetapan"`
// 	// Teguran_Id *uint64 `json:"teguran_id"`
// 	IsTeguran       *string `json:"isTeguran"`
// 	IsCancelled     *string `json:"isCancelled"`
// 	Note            *string `json:"note"`
// 	CancelledAt     string  `json:"cancelledAt"`
// 	Kasubid_User_Id *string `json:"kasubid_user_id"`
// 	Kabid_User_Id   *string `json:"kabid_user_id"`
// 	mdsjbr.CreateDto
// }

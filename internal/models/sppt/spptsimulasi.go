package sppt

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SpptSimulasi struct {
	Id                     uint64          `json:"id" gorm:"primaryKey"`
	BlokKavNoWP_sppt       *string         `json:"blokKavNoWP_sppt" gorm:"type:varchar(15)"`
	Faktorpengurangan_sppt *int            `json:"faktorpengurangan_sppt"`
	JalanWPskp_sppt        *string         `json:"jalanWPskp_sppt" gorm:"type:varchar(30)"`
	BankPersepsi_Id        *string         `json:"bankPersepsi_Id" gorm:"type:varchar(2)"`
	BankTunggal_Id         *string         `json:"bankTunggal_Id" gorm:"type:varchar(2)"`
	Blok_Id                *string         `json:"blok_Id" gorm:"type:varchar(3)"`
	Dati2_Id               *string         `json:"dati2_Id" gorm:"type:varchar(2)"`
	JenisOP_Id             *string         `json:"jenisOP_Id" gorm:"type:varchar(1)"`
	KanwilBank_Id          *string         `json:"kanwilBank_Id" gorm:"type:varchar(2)"`
	Kecamatan_Id           *string         `json:"kecamatan_Id" gorm:"type:varchar(3)"`
	Keluarahan_Id          *string         `json:"keluarahan_Id" gorm:"type:varchar(3)"`
	KelasBangunan_Id       *string         `json:"kelasBangunan_Id" gorm:"type:varchar(3)"`
	KelasTanah_Id          *string         `json:"kelasTanah_Id" gorm:"type:varchar(3)"`
	KPPBBbank_Id           *string         `json:"kppbbBank_Id" gorm:"type:varchar(2)"`
	PosWPsppt_Id           *string         `json:"posWPsppt_Id" gorm:"type:varchar(5)"`
	Propinsi_Id            *string         `json:"propinsi_Id" gorm:"type:varchar(2)"`
	TP_Id                  *string         `json:"tp_Id" gorm:"type:varchar(2)"`
	KelurahanWP_sppt       *string         `json:"kelurahanWP_sppt" gorm:"type:varchar(30)"`
	KotaWP_sppt            *string         `json:"kotaWP_sppt" gorm:"type:varchar(30)"`
	LuasBangunan_sppt      *int            `json:"luasBangunan_sppt"`
	LuasBumi_sppt          *int            `json:"luasBumi_sppt"`
	NIPPencetakan_sppt     *string         `json:"nipPencetakan_sppt" gorm:"type:varchar(9)"`
	NJKPskp_sppt           *float32        `json:"njKPskp_sppt" gorm:"type:decimal"`
	NJOPTKP_sppt           *int            `json:"njopTKP_sppt"`
	NJOPBangunan_sppt      *int            `json:"njopBangunan_sppt"`
	NJOPBumi_sppt          *int            `json:"njopBumi_sppt"`
	NJOP_sppt              *int            `json:"njop_sppt"`
	NamaWP_sppt            *string         `json:"namaWP_sppt" gorm:"type:varchar(30)"`
	NoPersil_sppt          *string         `json:"noPersil_sppt" gorm:"type:varchar(5)"`
	NoUrut                 *string         `json:"noUrut" gorm:"type:varchar(4)"`
	NoVirtualAccount       *string         `json:"noVirtualAccount" gorm:"type:varchar(20)"`
	Npwp_sppt              *string         `json:"npwp_sppt" gorm:"type:varchar(15)"`
	PBBterhutang_sppt      *int            `json:"pbbTerhutang_sppt"`
	PBBygHarusDibayar_sppt *int            `json:"pbbYgHarusDibayar_sppt"`
	Pemutihan              *datatypes.Date `json:"pemutihan"`
	QRCode                 *string         `json:"qrCode" gorm:"type:TEXT"`
	QRInvoiceNumber        *string         `json:"qrInvoiceNumber" gorm:"type:varchar(25)"`
	RtWP_sppt              *string         `json:"rtWP_sppt" gorm:"type:varchar(3)"`
	RwWP_sppt              *string         `json:"rwWP_sppt" gorm:"type:varchar(2)"`
	Siklus_sppt            *int            `json:"siklus_sppt" gorm:"type:decimal"`
	StatusCetak_sppt       *string         `json:"statusCetak_sppt" gorm:"type:varchar(1)"`
	StatusPembayaran_sppt  *string         `json:"statusPembayaran_sppt" gorm:"type:varchar(1)"`
	StatusTagihan_sppt     *string         `json:"statusTagihan_sppt" gorm:"type:varchar(1)"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan" gorm:"type:varchar(30)"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah" gorm:"type:varchar(30)"`
	TahunPajakskp_sppt     *string         `json:"tahunPajakskp_sppt" gorm:"type:varchar(30)"`
	VirtualAccoountJatim   *int            `json:"virtualAccoountJatim" gorm:"type:varchar(30)"`
	gormhelper.DateModel
}

type SimulasiNOP struct {
	ProvinsiID     string            `json:"provinsiID"`
	KotaID         string            `json:"kotaID"`
	KecamatanID    string            `json:"kecamatanID"`
	KelurahanID    string            `json:"kelurahanID"`
	BlokID         *string           `json:"blokID"`
	NoUrut         *string           `json:"noUrut"`
	NoUrutExt      *string           `json:"noUrutExt"`
	JenisOpId      *string           `json:"jenisOpId"`
	TahunPajak     *string           `json:"tahunPajak"`
	BukuJatuhTempo []*datatypes.Date `json:"bukuJatuhTempo"`
	BukuTerbit     []*datatypes.Date `json:"bukuTerbit"`
}

type RequestSimulasiDto struct {
	BlokKavNoWP_sppt       *string         `json:"blokKavNoWP_sppt"`
	Faktorpengurangan_sppt *int            `json:"faktorpengurangan_sppt"`
	JalanWPskp_sppt        *string         `json:"jalanWPskp_sppt"`
	BankPersepsi_Id        *string         `json:"bankPersepsi_Id"`
	BankTunggal_Id         *string         `json:"bankTunggal_Id"`
	Blok_Id                *string         `json:"blok_Id"`
	Dati2_Id               *string         `json:"dati2_Id"`
	JenisOP_Id             *string         `json:"jenisOP_Id"`
	KanwilBank_Id          *string         `json:"kanwilBank_Id"`
	Kecamatan_Id           *string         `json:"kecamatan_Id"`
	Keluarahan_Id          *string         `json:"keluarahan_Id"`
	KelasBangunan_Id       *string         `json:"kelasBangunan_Id"`
	KelasTanah_Id          *string         `json:"kelasTanah_Id"`
	KPPBBbank_Id           *string         `json:"kppbbBank_Id"`
	PosWPsppt_Id           *string         `json:"posWPsppt_Id"`
	Propinsi_Id            *string         `json:"propinsi_Id"`
	TP_Id                  *string         `json:"tp_Id"`
	KelurahanWP_sppt       *string         `json:"kelurahanWP_sppt"`
	KotaWP_sppt            *string         `json:"kotaWP_sppt"`
	LuasBangunan_sppt      *int            `json:"luasBangunan_sppt"`
	LuasBumi_sppt          *int            `json:"luasBumi_sppt"`
	NIPPencetakan_sppt     *string         `json:"nipPencetakan_sppt"`
	NJKPskp_sppt           *float32        `json:"njKPskp_sppt"`
	NJOPTKP_sppt           *int            `json:"njopTKP_sppt"`
	NJOPBangunan_sppt      *int            `json:"njopBangunan_sppt"`
	NJOPBumi_sppt          *int            `json:"njopBumi_sppt"`
	NJOP_sppt              *int            `json:"njop_sppt"`
	NamaWP_sppt            *string         `json:"namaWP_sppt"`
	NoPersil_sppt          *string         `json:"noPersil_sppt"`
	NoUrut                 *string         `json:"noUrut"`
	NoVirtualAccount       *string         `json:"noVirtualAccount"`
	Npwp_sppt              *string         `json:"npwp_sppt"`
	PBBterhutang_sppt      *int            `json:"pbbTerhutang_sppt"`
	PBBygHarusDibayar_sppt *int            `json:"pbbYgHarusDibayar_sppt"`
	Pemutihan              *datatypes.Date `json:"pemutihan"`
	QRCode                 *string         `json:"qrCode"`
	QRInvoiceNumber        *string         `json:"qrInvoiceNumber"`
	RtWP_sppt              *string         `json:"rtWP_sppt"`
	RwWP_sppt              *string         `json:"rwWP_sppt"`
	Siklus_sppt            *int            `json:"siklus_sppt"`
	StatusCetak_sppt       *string         `json:"statusCetak_sppt"`
	StatusPembayaran_sppt  *string         `json:"statusPembayaran_sppt"`
	StatusTagihan_sppt     *string         `json:"statusTagihan_sppt"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah"`
	TahunPajakskp_sppt     *string         `json:"tahunPajakskp_sppt"`
	VirtualAccoountJatim   *int            `json:"virtualAccoountJatim"`
}

type FilterSimulasiDto struct {
	BlokKavNoWP_sppt       *string         `json:"blokKavNoWP_sppt"`
	Faktorpengurangan_sppt *int            `json:"faktorpengurangan_sppt"`
	JalanWPskp_sppt        *string         `json:"jalanWPskp_sppt"`
	BankPersepsi_Id        *string         `json:"bankPersepsi_Id"`
	BankTunggal_Id         *string         `json:"bankTunggal_Id"`
	Blok_Id                *string         `json:"blok_Id"`
	Dati2_Id               *string         `json:"dati2_Id"`
	JenisOP_Id             *string         `json:"jenisOP_Id"`
	KanwilBank_Id          *string         `json:"kanwilBank_Id"`
	Kecamatan_Id           *string         `json:"kecamatan_Id"`
	Keluarahan_Id          *string         `json:"keluarahan_Id"`
	KelasBangunan_Id       *string         `json:"kelasBangunan_Id"`
	KelasTanah_Id          *string         `json:"kelasTanah_Id"`
	KPPBBbank_Id           *string         `json:"kppbbBank_Id"`
	PosWPsppt_Id           *string         `json:"posWPsppt_Id"`
	Propinsi_Id            *string         `json:"propinsi_Id"`
	TP_Id                  *string         `json:"tp_Id"`
	KelurahanWP_sppt       *string         `json:"kelurahanWP_sppt"`
	KotaWP_sppt            *string         `json:"kotaWP_sppt"`
	LuasBangunan_sppt      *int            `json:"luasBangunan_sppt"`
	LuasBumi_sppt          *int            `json:"luasBumi_sppt"`
	NIPPencetakan_sppt     *string         `json:"nipPencetakan_sppt"`
	NJKPskp_sppt           *float32        `json:"njKPskp_sppt"`
	NJOPTKP_sppt           *int            `json:"njopTKP_sppt"`
	NJOPBangunan_sppt      *int            `json:"njopBangunan_sppt"`
	NJOPBumi_sppt          *int            `json:"njopBumi_sppt"`
	NJOP_sppt              *int            `json:"njop_sppt"`
	NamaWP_sppt            *string         `json:"namaWP_sppt"`
	NoPersil_sppt          *string         `json:"noPersil_sppt"`
	NoUrut                 *string         `json:"noUrut"`
	NoVirtualAccount       *string         `json:"noVirtualAccount"`
	Npwp_sppt              *string         `json:"npwp_sppt"`
	PBBterhutang_sppt      *int            `json:"pbbTerhutang_sppt"`
	PBBygHarusDibayar_sppt *int            `json:"pbbYgHarusDibayar_sppt"`
	Pemutihan              *datatypes.Date `json:"pemutihan"`
	QRCode                 *string         `json:"qrCode"`
	QRInvoiceNumber        *string         `json:"qrInvoiceNumber"`
	RtWP_sppt              *string         `json:"rtWP_sppt"`
	RwWP_sppt              *string         `json:"rwWP_sppt"`
	Siklus_sppt            *int            `json:"siklus_sppt"`
	StatusCetak_sppt       *string         `json:"statusCetak_sppt"`
	StatusPembayaran_sppt  *string         `json:"statusPembayaran_sppt"`
	StatusTagihan_sppt     *string         `json:"statusTagihan_sppt"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah"`
	TahunPajakskp_sppt     *string         `json:"tahunPajakskp_sppt"`
	VirtualAccoountJatim   *int            `json:"virtualAccoountJatim"`
	Page                   int             `json:"page"`
	PageSize               int             `json:"page_size"`
}

func (i SpptSimulasi) SetDataPermohonanTransformer(req Sppt) SpptSimulasi {
	return SpptSimulasi{
		Id:                     req.Id,
		BlokKavNoWP_sppt:       req.BlokKavNoWP_sppt,
		Faktorpengurangan_sppt: req.Faktorpengurangan_sppt,
		JalanWPskp_sppt:        req.JalanWPskp_sppt,
		BankPersepsi_Id:        req.BankPersepsi_Id,
		BankTunggal_Id:         req.BankTunggal_Id,
		Blok_Id:                req.Blok_Id,
		Dati2_Id:               req.Dati2_Id,
		JenisOP_Id:             req.JenisOP_Id,
		KanwilBank_Id:          req.KanwilBank_Id,
		Kecamatan_Id:           req.Kecamatan_Id,
		Keluarahan_Id:          req.Keluarahan_Id,
		KelasBangunan_Id:       req.KelasBangunan_Id,
		KelasTanah_Id:          req.KelasTanah_Id,
		KPPBBbank_Id:           req.KPPBBbank_Id,
		PosWPsppt_Id:           req.PosWPsppt_Id,
		Propinsi_Id:            req.Propinsi_Id,
		TP_Id:                  req.TP_Id,
		KelurahanWP_sppt:       req.KelurahanWP_sppt,
		KotaWP_sppt:            req.KotaWP_sppt,
		LuasBangunan_sppt:      req.LuasBangunan_sppt,
		LuasBumi_sppt:          req.LuasBumi_sppt,
		NIPPencetakan_sppt:     req.NIPPencetakan_sppt,
		NJKPskp_sppt:           req.NJKPskp_sppt,
		NJOPTKP_sppt:           req.NJOPTKP_sppt,
		NJOPBangunan_sppt:      req.NJOPBangunan_sppt,
		NJOPBumi_sppt:          req.NJOPBumi_sppt,
		NJOP_sppt:              req.NJOP_sppt,
		NamaWP_sppt:            req.NamaWP_sppt,
		NoPersil_sppt:          req.NoPersil_sppt,
		NoUrut:                 req.NoUrut,
		NoVirtualAccount:       req.NoVirtualAccount,
		Npwp_sppt:              req.Npwp_sppt,
		PBBterhutang_sppt:      req.PBBterhutang_sppt,
		PBBygHarusDibayar_sppt: req.PBBterhutang_sppt,
		Pemutihan:              req.Pemutihan,
		QRCode:                 req.QRCode,
		QRInvoiceNumber:        req.QRInvoiceNumber,
		RtWP_sppt:              req.RtWP_sppt,
		RwWP_sppt:              req.RwWP_sppt,
		Siklus_sppt:            req.Siklus_sppt,
		StatusCetak_sppt:       req.StatusCetak_sppt,
		StatusPembayaran_sppt:  req.StatusPembayaran_sppt,
		StatusTagihan_sppt:     req.StatusTagihan_sppt,
		TanggalCetak_sppt:      req.TanggalCetak_sppt,
		TanggalJatuhTempo_sppt: req.TanggalJatuhTempo_sppt,
		TanggalTerbit_sppt:     req.TanggalTerbit_sppt,
		TahunAwalKelasBangunan: req.TahunAwalKelasBangunan,
		TahunAwalKelasTanah:    req.TahunAwalKelasTanah,
		TahunPajakskp_sppt:     req.TahunPajakskp_sppt,
		VirtualAccoountJatim:   req.VirtualAccoountJatim,
	}
}

type Buku struct {
	Id  int
	Min float64
	Max float64
}

func (i Buku) GetBukus() []Buku {
	var Bukus []Buku
	tempBuku := Buku{
		Id:  1,
		Min: 0,
		Max: 100000,
	}
	Bukus = append(Bukus, tempBuku)

	tempBuku = Buku{
		Id:  2,
		Min: 100001,
		Max: 500000,
	}
	Bukus = append(Bukus, tempBuku)

	tempBuku = Buku{
		Id:  3,
		Min: 500001,
		Max: 2000000,
	}
	Bukus = append(Bukus, tempBuku)

	tempBuku = Buku{
		Id:  4,
		Min: 2000001,
		Max: 5000000,
	}
	Bukus = append(Bukus, tempBuku)

	tempBuku = Buku{
		Id:  5,
		Min: 5000001,
		Max: 999999999999999,
	}
	Bukus = append(Bukus, tempBuku)
	return Bukus
}

package sppt

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

const (
	NjopTkpYear = 2012
	NjopTkpOld  = 6000000
	NjopTkpNew  = 10000000
)

type Sppt struct {
	Id                     uint64          `json:"id" gorm:"primaryKey"`
	Propinsi_Id            *string         `json:"propinsi_Id" gorm:"type:varchar(2)"`
	Dati2_Id               *string         `json:"dati2_Id" gorm:"type:varchar(2)"`
	Kecamatan_Id           *string         `json:"kecamatan_Id" gorm:"type:varchar(3)"`
	Keluarahan_Id          *string         `json:"keluarahan_Id" gorm:"type:varchar(3)"`
	Blok_Id                *string         `json:"blok_Id" gorm:"type:varchar(3)"`
	NoUrut                 *string         `json:"noUrut" gorm:"type:varchar(4)"`
	JenisOP_Id             *string         `json:"jenisOP_Id" gorm:"type:varchar(1)"`
	BlokKavNoWP_sppt       *string         `json:"blokKavNoWP_sppt" gorm:"type:varchar(15)"`
	Faktorpengurangan_sppt *int            `json:"faktorpengurangan_sppt"`
	JalanWPskp_sppt        *string         `json:"jalanWPskp_sppt" gorm:"type:varchar(30)"`
	BankPersepsi_Id        *string         `json:"bankPersepsi_Id" gorm:"type:varchar(2)"`
	BankTunggal_Id         *string         `json:"bankTunggal_Id" gorm:"type:varchar(2)"`
	KanwilBank_Id          *string         `json:"kanwilBank_Id" gorm:"type:varchar(2)"`
	KelasBangunan_Id       *string         `json:"kelasBangunan_Id" gorm:"type:varchar(3)"`
	KelasTanah_Id          *string         `json:"kelasTanah_Id" gorm:"type:varchar(3)"`
	KPPBBbank_Id           *string         `json:"kppbbBank_Id" gorm:"type:varchar(2)"`
	PosWPsppt_Id           *string         `json:"posWPsppt_Id" gorm:"type:varchar(5)"`
	TP_Id                  *string         `json:"tp_Id" gorm:"type:varchar(2)"`
	KelurahanWP_sppt       *string         `json:"kelurahanWP_sppt" gorm:"type:varchar(30)"`
	KotaWP_sppt            *string         `json:"kotaWP_sppt" gorm:"type:varchar(30)"`
	LuasBangunan_sppt      *int            `json:"luasBangunan_sppt"`
	LuasBumi_sppt          *int            `json:"luasBumi_sppt"`
	NIPPencetakan_sppt     *string         `json:"nipPencetakan_sppt" gorm:"type:varchar(9)"`
	NJKPskp_sppt           *float32        `json:"njKPskp_sppt" gorm:"type:decimal(5,2)"`
	NJOPTKP_sppt           *int            `json:"njopTKP_sppt"`
	NJOPBangunan_sppt      *int            `json:"njopBangunan_sppt"`
	NJOPBumi_sppt          *int            `json:"njopBumi_sppt"`
	NJOP_sppt              *int            `json:"njop_sppt"`
	NamaWP_sppt            *string         `json:"namaWP_sppt" gorm:"type:varchar(30)"`
	NoPersil_sppt          *string         `json:"noPersil_sppt" gorm:"type:varchar(5)"`
	NoVirtualAccount       *string         `json:"noVirtualAccount" gorm:"type:varchar(20)"`
	Npwp_sppt              *string         `json:"npwp_sppt" gorm:"type:varchar(15)"`
	PBBterhutang_sppt      *int            `json:"pbbTerhutang_sppt"`
	PBBygHarusDibayar_sppt *int            `json:"pbbYgHarusDibayar_sppt"`
	Pemutihan              *datatypes.Date `json:"pemutihan"`
	QRCode                 *string         `json:"qrCode" gorm:"type:TEXT"`
	QRInvoiceNumber        *string         `json:"qrInvoiceNumber" gorm:"type:varchar(25)"`
	RtWP_sppt              *string         `json:"rtWP_sppt" gorm:"type:varchar(3)"`
	RwWP_sppt              *string         `json:"rwWP_sppt" gorm:"type:varchar(2)"`
	Siklus_sppt            *int            `json:"siklus_sppt"`
	StatusCetak_sppt       *string         `json:"statusCetak_sppt" gorm:"type:varchar(1)"`
	StatusPembayaran_sppt  *string         `json:"statusPembayaran_sppt" gorm:"type:varchar(1)"`
	StatusTagihan_sppt     *string         `json:"statusTagihan_sppt" gorm:"type:varchar(1)"`
	Sunset                 *datatypes.Date `json:"sunset"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	TanggalCreateQRCode    *datatypes.Date `json:"tanggalCreateQRCode"`
	TanggalCreateVA_Jatim  *datatypes.Date `json:"tanggalCreateVA_Jatim"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan" gorm:"type:varchar(4)"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah" gorm:"type:varchar(4)"`
	TahunPajakskp_sppt     *string         `json:"tahunPajakskp_sppt" gorm:"type:varchar(4)"`
	VirtualAccoountJatim   *int            `json:"virtualAccoountJatim" gorm:"type:varchar(30)"`
	User_ID                *uint64         `json:"user_id"`
	gormhelper.DateModel
}

type RequestDto struct {
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
	Sunset                 *datatypes.Date `json:"sunset"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	TanggalCreateQRCode    *datatypes.Date `json:"tanggalCreateQRCode"`
	TanggalCreateVA_Jatim  *datatypes.Date `json:"tanggalCreateVA_Jatim"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah"`
	TahunPajakskp_sppt     *string         `json:"tahunPajakskp_sppt"`
	User_ID                *uint64         `json:"user_id"`
	VirtualAccoountJatim   *int            `json:"virtualAccoountJatim"`
}

type FilterDto struct {
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
	Sunset                 *datatypes.Date `json:"sunset"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	TanggalCreateQRCode    *datatypes.Date `json:"tanggalCreateQRCode"`
	TanggalCreateVA_Jatim  *datatypes.Date `json:"tanggalCreateVA_Jatim"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah"`
	TahunPajakskp_sppt     *string         `json:"tahunPajakskp_sppt"`
	VirtualAccoountJatim   *int            `json:"virtualAccoountJatim"`
	Page                   int             `json:"page"`
	PageSize               int             `json:"page_size"`
}

type PenilaianDto struct {
	Provinsi_Kode  *string `json:"provinsi_kode"`
	Daerah_Kode    *string `json:"daerah_kode"`
	Kecamatan_Kode *string `json:"kecamatan_kode"`
	Kelurahan_Kode *string `json:"kelurahan_kode"`
	Tahun          string  `json:"tahun"`
}

type GetDaftarTagihan struct {
	Provinsi_Kode     string  `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode       string  `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Kecamatan_Kode    string  `json:"kecamatan_kode" validate:"required;minLength=3;maxLength=3"`
	NamaPropinsi      *string `json:"namaPropinsi"`
	NamaDati2         *string `json:"namaDati2"`
	NamaKecamatan     *string `json:"namaKecamatan"`
	TahunPajakAwal    *string `json:"tahunPajakAwal"`
	TahunPajakAkhir   *string `json:"tahunPajakAkhir"`
	KetetapanPBBAwal  *string `json:"ketetapanPBBAwal"`
	KetetapanPBBAkhir *string `json:"ketetapanPBBAkhir"`
}

type PenetapanMassalDto struct {
	Provinsi_Kode  *string   `json:"provinsi_kode"`
	Daerah_Kode    *string   `json:"daerah_kode"`
	Kecamatan_Kode *string   `json:"kecamatan_kode"`
	Kelurahan_Kode *string   `json:"kelurahan_kode"`
	Tahun          string    `json:"tahun"`
	TglJatuhTempo1 string    `json:"tglJatuhTempo1"`
	TglJatuhTempo2 string    `json:"tglJatuhTempo2"`
	TglJatuhTempo3 string    `json:"tglJatuhTempo3"`
	TglJatuhTempo4 string    `json:"tglJatuhTempo4"`
	TglJatuhTempo5 string    `json:"tglJatuhTempo5"`
	TglTerbit1     string    `json:"tglTerbit1"`
	TglTerbit2     string    `json:"tglTerbit2"`
	TglTerbit3     string    `json:"tglTerbit3"`
	TglTerbit4     string    `json:"tglTerbit4"`
	TglTerbit5     string    `json:"tglTerbit5"`
	BukuMin        []float64 `json:"bukuMin"`
	BukuMax        []float64 `json:"bukuMax"`
}

package sppt

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SpptSejarah struct {
	Id            uint64  `json:"id" gorm:"primaryKey"`
	Propinsi_Id   *string `json:"propinsi_Id" gorm:"type:varchar(2)"`
	Dati2_Id      *string `json:"dati2_Id" gorm:"type:varchar(2)"`
	Kecamatan_Id  *string `json:"kecamatan_Id" gorm:"type:varchar(3)"`
	Keluarahan_Id *string `json:"keluarahan_Id" gorm:"type:varchar(3)"`
	Blok_Id       *string `json:"blok_Id" gorm:"type:varchar(3)"`
	NoUrut        *string `json:"noUrut" gorm:"type:varchar(4)"`
	JenisOP_Id    *string `json:"jenisOP_Id" gorm:"type:varchar(1)"`

	NamaWP_sppt      *string `json:"namaWP_sppt" gorm:"type:varchar(30)"`
	JalanWPskp_sppt  *string `json:"jalanWPskp_sppt" gorm:"type:varchar(30)"`
	BlokKavNoWP_sppt *string `json:"blokKavNoWP_sppt" gorm:"type:varchar(15)"`
	RtWP_sppt        *string `json:"rtWP_sppt" gorm:"type:varchar(3)"`
	RwWP_sppt        *string `json:"rwWP_sppt" gorm:"type:varchar(2)"`
	KelurahanWP_sppt *string `json:"kelurahanWP_sppt" gorm:"type:varchar(30)"`
	KotaWP_sppt      *string `json:"kotaWP_sppt" gorm:"type:varchar(30)"`
	PosWPsppt_Id     *string `json:"posWPsppt_Id" gorm:"type:varchar(5)"`

	TahunPajakskp_sppt *string `json:"tahunPajakskp_sppt" gorm:"type:varchar(4)"`
	Siklus_sppt        *int    `json:"siklus_sppt"`

	Npwp_sppt     *string `json:"npwp_sppt" gorm:"type:varchar(15)"`
	NoPersil_sppt *string `json:"noPersil_sppt" gorm:"type:varchar(5)"`
	KanwilBank_Id *string `json:"kanwilBank_Id" gorm:"type:varchar(2)"`
	KPPBBbank_Id  *string `json:"kppbbBank_Id" gorm:"type:varchar(2)"`
	JnsSK         *string `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK          *string `json:"noSK" gorm:"type:varchar(30)"`
	TP_Id         *string `json:"tp_Id" gorm:"type:varchar(2)"`

	KelasTanah_Id          *string         `json:"kelasTanah_Id" gorm:"type:varchar(3)"`
	TahunAwalKelasTanah    *string         `json:"tahunAwalKelasTanah" gorm:"type:varchar(4)"`
	KelasBangunan_Id       *string         `json:"kelasBangunan_Id" gorm:"type:varchar(3)"`
	TahunAwalKelasBangunan *string         `json:"tahunAwalKelasBangunan" gorm:"type:varchar(4)"`
	LuasBumi_sppt          *int            `json:"luasBumi_sppt"`
	LuasBangunan_sppt      *int            `json:"luasBangunan_sppt"`
	NJOPBumi_sppt          *int            `json:"njopBumi_sppt"`
	NJOPBangunan_sppt      *int            `json:"njopBangunan_sppt"`
	NJOP_sppt              *int            `json:"njop_sppt"`
	NJOPTKP_sppt           *int            `json:"njopTKP_sppt"`
	NJKPskp_sppt           *float32        `json:"njKPskp_sppt" gorm:"type:decimal(5,2)"`
	PBBterhutang_sppt      *int            `json:"pbbTerhutang_sppt"`
	Faktorpengurangan_sppt *int            `json:"faktorpengurangan_sppt"`
	PBBygHarusDibayar_sppt *int            `json:"pbbYgHarusDibayar_sppt"`
	StatusPembayaran_sppt  *string         `json:"statusPembayaran_sppt" gorm:"type:varchar(1)"`
	StatusTagihan_sppt     *string         `json:"statusTagihan_sppt" gorm:"type:varchar(1)"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	NIPPencetakan_sppt     *string         `json:"nipPencetakan_sppt" gorm:"type:varchar(9)"`
	gormhelper.DateModel
}

type ListSejarahSpptResponse struct {
	NamaWajibPajak  *string               `json:"namaWajibPajak"`
	JalanObjekPajak *string               `json:"jalanObjekPajak"`
	JalanWajibPajak *string               `json:"jalanWajibPajak"`
	List            []SejarahSpptResponse `json:"list"`
}

type SejarahSpptResponse struct {
	NamaWP_sppt            *string         `json:"namaWP_sppt"`
	JalanWPskp_sppt        *string         `json:"jalanWPskp_sppt"`
	BlokKavNoWP_sppt       *string         `json:"blokKavNoWP_sppt"`
	RtWP_sppt              *string         `json:"rtWP_sppt"`
	RwWP_sppt              *string         `json:"rwWP_sppt"`
	KelurahanWP_sppt       *string         `json:"kelurahanWP_sppt"`
	KotaWP_sppt            *string         `json:"kotaWP_sppt"`
	PosWPsppt_Id           *string         `json:"posWPsppt_Id"`
	Npwp_sppt              *string         `json:"npwp_sppt"`
	NoPersil_sppt          *string         `json:"noPersil_sppt"`
	JnsSK                  *string         `json:"jnsSK"`
	NoSK                   *string         `json:"noSK"`
	TP_Id                  *string         `json:"tp_Id"`
	NamaTp                 *string         `json:"namaTp" `
	KelasBangunan_Id       *string         `json:"kelasBangunan_Id"`
	KelasTanah_Id          *string         `json:"kelasTanah_Id"`
	TanggalJatuhTempo_sppt *datatypes.Date `json:"tanggalJatuhTempo_sppt"`
	LuasBumi_sppt          *int            `json:"luasBumi_sppt"`
	LuasBangunan_sppt      *int            `json:"luasBangunan_sppt"`
	NJOPBumi_sppt          *int            `json:"njopBumi_sppt"`
	NJOPBangunan_sppt      *int            `json:"njopBangunan_sppt"`
	NJOP_sppt              *int            `json:"njop_sppt"`
	NJOPTKP_sppt           *int            `json:"njopTKP_sppt"`
	NJKPskp_sppt           *float32        `json:"njKPskp_sppt"`
	PBBterhutang_sppt      *int            `json:"pbbTerhutang_sppt"`
	Faktorpengurangan_sppt *int            `json:"faktorpengurangan_sppt"`
	PBBygHarusDibayar_sppt *int            `json:"pbbYgHarusDibayar_sppt"`
	StatusPembayaran_sppt  *string         `json:"statusPembayaran_sppt"`
	TanggalTerbit_sppt     *datatypes.Date `json:"tanggalTerbit_sppt"`
	TanggalCetak_sppt      *datatypes.Date `json:"tanggalCetak_sppt"`
	NIPPencetakan_sppt     *string         `json:"nipPencetakan_sppt"`
}

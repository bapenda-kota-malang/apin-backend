package bphtb

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type BphtbSptpd struct {
	Id                        uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Sptpd_Id                  *string         `json:"sptpd_id" gorm:"type:varchar(50)"`
	Ppat_Id                   *string         `json:"ppat_id" gorm:"type:varchar(50)"`
	WajibPajak_NIK            *string         `json:"wajibPajak_NIK" gorm:"type:varchar(50)"`
	PermohonanProvinsiID      *string         `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID          *string         `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID     *string         `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID     *string         `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID          *string         `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon             *string         `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID          *string         `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	NOPAlamat                 *string         `json:"opAlamat" gorm:"type:varchar(50)"`
	OPProvinsi                *string         `json:"opProvinsi" gorm:"type:varchar(50)"`
	OPKota                    *string         `json:"opKota" gorm:"type:varchar(50)"`
	OPKecamatan               *string         `json:"opKecamatan" gorm:"type:varchar(50)"`
	OPKelurahan               *string         `json:"opKelurahan" gorm:"type:varchar(50)"`
	OP_RtRW                   *string         `json:"op_RtRW" gorm:"type:varchar(50)"`
	OPLuasTanah               float64         `json:"opLuasTanah" gorm:"type:decimal(10,2)"`
	OPLuasBangunan            float64         `json:"opLuasBangunan" gorm:"type:decimal(10,2)"`
	OPLuasTanahBersama        float64         `json:"opLuasTanahBersama" gorm:"type:decimal(10,2)"`
	OPLuasBangunanBersama     float64         `json:"opLuasBangunanBersama" gorm:"type:decimal(10,2)"`
	NjopLuasTanah             float64         `json:"njopLuasTanah" gorm:"type:decimal(10,2)"`
	NjopLuasBangunan          float64         `json:"njopLuasBangunan" gorm:"type:decimal(10,2)"`
	NjopTanahBersama          float64         `json:"njopTanahBersama" gorm:"type:decimal(20,2)"`
	NjopBangunanBersama       float64         `json:"njopBangunanBersama" gorm:"type:decimal(20,2)"`
	NilaiOp                   float64         `json:"nilaiOp" gorm:"type:decimal(20,2)"`
	JenisPerolehanOp          *string         `json:"jenisPerolehanOp" gorm:"type:varchar(50)"`
	NoSertifikatOp            *string         `json:"noSertifikatOp" gorm:"type:varchar(50)"`
	NjopPbbOp                 float64         `json:"njopPbbOp" gorm:"type:decimal(20,2)"`
	LokasiOp                  *string         `json:"lokasiOp" gorm:"type:varchar(50)"`
	TahunPajakSppt            *string         `json:"tahunPajakSppt" gorm:"type:varchar(50)"`
	Npop                      float64         `json:"npop" gorm:"type:decimal(50,2)"`
	Npoptkp                   float64         `json:"npoptkp" gorm:"type:decimal(50,2)"`
	JenisSetoran              *string         `json:"jenisSetoran" gorm:"type:varchar(50)"`
	JenisSetoranNomor         *string         `json:"jenisSetoranNomor" gorm:"type:varchar(50)"`
	JenisSetoranTanggal       *datatypes.Date `json:"jenisSetoranTanggal"`
	JenisSetoranHitungSendiri *string         `json:"jenisSetoranHitungSendiri" gorm:"type:varchar(50)"`
	JenisSetoranCustom        *string         `json:"jenisSetoranCustom" gorm:"type:varchar(50)"`
	JumlahSetor               float64         `json:"jumlahSetor" gorm:"type:decimal(50,2)"`
	NominalSPT                float64         `json:"nominalSPT" gorm:"type:decimal(50,2)"`
	Tanggal                   *datatypes.Date `json:"tanggal"`
	NoDokumen                 *string         `json:"noDokumen" gorm:"type:varchar(50)"`
	NopPbbBaru                *string         `json:"nopPbbBaru" gorm:"type:varchar(50)"`
	User_id                   *string         `json:"user_id" gorm:"type:varchar(50)"`
	Id_pp                     *string         `json:"id_pp" gorm:"type:varchar(50)"`
	JenisSetoranKeterangan    *string         `json:"jenisSetoranKeterangan" gorm:"type:varchar(50)"`
	NilaiPasar                float64         `json:"nilaiPasar" gorm:"type:decimal(50,2)"`
	JenisPerolehan_id         *string         `json:"jenisPerolehan_id" gorm:"type:varchar(50)"`
	InpAphb1                  *string         `json:"inpAphb1" gorm:"type:varchar(50)"`
	InpAphb2                  *string         `json:"inpAphb2" gorm:"type:varchar(50)"`
	InpAphb3                  *string         `json:"inpAphb3" gorm:"type:varchar(50)"`
	KodeValidasi              *string         `json:"kodeValidasi" gorm:"type:varchar(50)"`
	GambarOp                  *string         `json:"gambarOp" gorm:"type:varchar(100)"`
	IsLunas                   *string         `json:"isLunas" gorm:"type:varchar(50)"`
	ValidasiDisependa         *string         `json:"validasiDisependa" gorm:"type:varchar(50)"`
	ValidasiBank              *string         `json:"validasiBank" gorm:"type:varchar(50)"`
	VerifikasiPpatAt          *time.Time      `json:"verifikasiPpatAt"`
	TglValidasiBank           *datatypes.Date `json:"tglValidasiBank"`
	TglValidasiDispenda       *datatypes.Date `json:"tglValidasiDispenda"`
	Dispenda_User_id          *string         `json:"pegawai_User_id" gorm:"type:varchar(50)"`
	AlasanReject              *string         `json:"alasanReject" gorm:"type:text"`
	Bank_Id                   *string         `json:"bank_Id" gorm:"type:varchar(50)"`
	GambarInt                 *string         `json:"gambarInt" gorm:"type:text"`
	FlagDispenda              *int            `json:"flagDispenda" gorm:"type:int"`
	FlagPPAT                  *int            `json:"flagPPAT" gorm:"type:int"`
	TanahInpAphb1             *float64        `json:"tanahInpAphb1" gorm:"type:decimal(20,2)"`
	TanahInpAphb2             *float64        `json:"tanahInpAphb2" gorm:"type:decimal(20,2)"`
	TanahInpAphb3             *float64        `json:"tanahInpAphb3" gorm:"type:decimal(20,2)"`
	BangunanInpAphb1          *float64        `json:"bangunanInpAphb1" gorm:"type:decimal(20,2)"`
	BangunanInpAphb2          *float64        `json:"bangunanInpAphb2" gorm:"type:decimal(20,2)"`
	BangunanInpAphb3          *float64        `json:"bangunanInpAphb3" gorm:"type:decimal(20,2)"`
	TanahBersamaInpAphb1      *float64        `json:"tanahBersamaInpAphb1" gorm:"type:decimal(20,2)"`
	TanahBersamaInpAphb2      *float64        `json:"tanahBersamaInpAphb2" gorm:"type:decimal(20,2)"`
	TanahBersamaInpAphb3      *float64        `json:"tanahBersamaInpAphb3" gorm:"type:decimal(20,2)"`
	BangunanBersamaInpAphb1   *float64        `json:"bangunanBersamaInpAphb1" gorm:"type:decimal(20,2)"`
	BangunanBersamaInpAphb2   *float64        `json:"bangunanBersamaInpAphb2" gorm:"type:decimal(20,2)"`
	BangunanBersamaInpAphb3   *float64        `json:"bangunanBersamaInpAphb3" gorm:"type:decimal(20,2)"`
	IsKurangBayar             *string         `json:"isKurangBayar" gorm:"type:varchar(2)"`
	KurangBayar               *float64        `json:"kurangBayar" gorm:"type:decimal(20,2)"`
	SSPDLama                  *int            `json:"sspdLama" gorm:"type:int"`
	NoReff                    *string         `json:"noReff" gorm:"type:varchar(50)"`
	IdBilling                 *string         `json:"idBilling" gorm:"type:varchar(50)"`
	Proses                    *string         `json:"proses" gorm:"type:varchar(2)"`
	NamaWp                    *string         `json:"namaWp" gorm:"type:varchar(50)"`
	Alamat                    *string         `json:"alamat" gorm:"type:varchar(50)"`
	Provinsi_Id               *string         `json:"provinsi_id" gorm:"type:varchar(2)"`
	Kabupaten_id              *string         `json:"kabupaten_id" gorm:"type:varchar(2)"`
	Kecamatan_Id              *string         `json:"kecamatan_id" gorm:"type:varchar(3)"`
	Kelurahan_Id              *string         `json:"kelurahan_id" gorm:"type:varchar(3)"`
	KodePos                   *string         `json:"kodePos" gorm:"type:varchar(50)"`
	RtRw                      *string         `json:"rtRw" gorm:"type:varchar(50)"`
	WajibPajak_id             *uint64         `json:"wajibPajak_id" gorm:"type:int"`
	ApprovePPAT               *string         `json:"approvePPAT" gorm:"type:varchar(2)"`
	NoPelayanan               *string         `json:"noPelayanan" gorm:"type:varchar(50)"`
	NamaPetugasLapangan       *string         `json:"namaPetugasLapangan" gorm:"type:varchar(50)"`
	NamaStaff                 *string         `json:"namaStaff" gorm:"type:varchar(50)"`
	NoAkta                    *string         `json:"noAkta" gorm:"type:varchar(50)"`
	TglAkta                   *datatypes.Date `json:"tglAkta"`
	PihakMengalihkan          *string         `json:"pihakMengalihkan" gorm:"type:varchar(50)"`
	Tgl                       *datatypes.Date `json:"tgl"`
	Rp                        *float64        `json:"rp" gorm:"type:decimal(15,2)"`
	Batas                     *int            `json:"batas" gorm:"type:int"`
	Pengurangan               *float64        `json:"pengurangan" gorm:"type:decimal(15,2)"`
	SkPengurangan             *string         `json:"skPengurangan" gorm:"type:varchar(50)"`
	TglSkPengurangan          *datatypes.Date `json:"tglSkPengurangan"`
	WajibPajak                *string         `json:"wajibPajak" gorm:"type:varchar(50)"`
	TglExpBilling             *datatypes.Date `json:"tglExpBilling"`
	JenisTransaksiSpop        *int            `json:"jenisTransaksiSpop" gorm:"type:int"`
	Pekerjaan                 *int            `json:"pekerjaan" gorm:"type:int"`
	JenisTanah                *int            `json:"jenisTanah" gorm:"type:int"`
	JumlahBangunan            *int            `json:"jumlahBangunan" gorm:"type:int"`
	BlokKavNo                 *string         `json:"blokKavNo" gorm:"type:varchar(50)"`
	RT                        *string         `json:"rT" gorm:"type:varchar(50)"`
	Rw                        *string         `json:"rw" gorm:"type:varchar(50)"`
	Status                    string          `json:"status" gorm:"type:varchar(50)"`
	Cs_Nama                   *pq.StringArray `json:"cs_nama" gorm:"type:varchar[]"`
	Cs_Nik                    *pq.StringArray `json:"cs_nik" gorm:"type:varchar[]"`
	Qq_Nama                   *string         `json:"qq_nama"`
	Qq_Nib                    *string         `json:"qq_nib"`
	gormhelper.DateModel
	Lampiran *Lampiran `json:"lampiran,omitempty" gorm:"foreignKey:BphtbSptpd_Id"`
}

type ResponseSptpd struct {
	BphtbSptpd
	Provinsi_wp  *string `json:"provinsi_wp"`
	Kabupaten_wp *string `json:"kabupaten_wp"`
	Kecamatan_wp *string `json:"kecamatan_wp"`
	Kelurahan_wp *string `json:"kelurahan_wp"`
}

type RequestApprovalSptpd struct {
	NoDokumen           *string  `json:"noDokumen"`
	User_id             *string  `json:"user_id"`
	Id_pp               *string  `json:"id_pp"`
	KodeValidasi        *string  `json:"kodeValidasi"`
	IsLunas             *string  `json:"isLunas"`
	ValidasiDisependa   *string  `json:"validasiDisependa"`
	ValidasiBank        *string  `json:"validasiBank"`
	Dispenda_User_id    *string  `json:"pegawai_User_id"`
	AlasanReject        *string  `json:"alasanReject"`
	Bank_Id             *string  `json:"bank_Id"`
	FlagDispenda        *int     `json:"flagDispenda"`
	FlagPPAT            *int     `json:"flagPPAT"`
	IsKurangBayar       *string  `json:"isKurangBayar"`
	KurangBayar         *float64 `json:"kurangBayar"`
	NoReff              *string  `json:"noReff"`
	IdBilling           *string  `json:"idBilling"`
	Proses              *string  `json:"proses"`
	ApprovePPAT         *string  `json:"approvePPAT"`
	NoPelayanan         *string  `json:"noPelayanan"`
	NamaPetugasLapangan *string  `json:"namaPetugasLapangan"`
	NamaStaff           *string  `json:"namaStaff"`
	Status              *string  `json:"status"`
	gormhelper.DateModel
}

func (req RequestApprovalSptpd) SetDataApproval(i *BphtbSptpd) *BphtbSptpd {
	if req.User_id != nil {
		i.User_id = req.User_id
	}
	if req.Id_pp != nil {
		i.Id_pp = req.Id_pp
	}
	if req.KodeValidasi != nil {
		i.KodeValidasi = req.KodeValidasi
	}
	if req.IsLunas != nil {
		i.IsLunas = req.IsLunas
	}
	if req.ValidasiDisependa != nil {
		i.ValidasiDisependa = req.ValidasiDisependa
	}
	if req.ValidasiBank != nil {
		i.ValidasiBank = req.ValidasiBank
	}
	if req.Dispenda_User_id != nil {
		i.Dispenda_User_id = req.Dispenda_User_id
	}
	if req.AlasanReject != nil {
		i.AlasanReject = req.AlasanReject
	}
	if req.Bank_Id != nil {
		i.Bank_Id = req.Bank_Id
	}
	if req.FlagDispenda != nil {
		i.FlagDispenda = req.FlagDispenda
	}
	if req.FlagPPAT != nil {
		i.FlagPPAT = req.FlagPPAT
	}
	if req.IsKurangBayar != nil {
		i.IsKurangBayar = req.IsKurangBayar
	}
	if req.KurangBayar != nil {
		i.KurangBayar = req.KurangBayar
	}
	if req.NoReff != nil {
		i.NoReff = req.NoReff
	}
	if req.IdBilling != nil {
		i.IdBilling = req.IdBilling
	}
	if req.Proses != nil {
		i.Proses = req.Proses
	}
	if req.ApprovePPAT != nil {
		i.ApprovePPAT = req.ApprovePPAT
	}
	if req.NoPelayanan != nil {
		i.NoPelayanan = req.NoPelayanan
	}
	if req.NamaPetugasLapangan != nil {
		i.NamaPetugasLapangan = req.NamaPetugasLapangan
	}
	if req.NamaStaff != nil {
		i.NamaStaff = req.NamaStaff
	}
	i.Status = *req.Status
	return i
}

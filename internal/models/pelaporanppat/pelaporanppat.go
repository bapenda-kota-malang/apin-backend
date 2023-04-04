package pelaporanppat

import (
	"time"

	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type PelaporanPpat struct {
	Id                 uint64          `json:"id" gorm:"primaryKey"`
	Sptpd_Id           *string         `json:"sptpd_id" gorm:"type:varchar(50)"`
	Ppat_Id            *string         `json:"ppat_id" gorm:"type:varchar(50)"`
	NoAkta             *string         `json:"noAkta" gorm:"type:varchar(50)"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal" gorm:"type:varchar(2)"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal" gorm:"type:varchar(4)"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir" gorm:"type:varchar(2)"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir" gorm:"type:varchar(4)"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan" gorm:"type:varchar(50)"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
	NominalSSP         *float64        `json:"nominalSSP"`
	gh.DateModel
}

type CreateDto struct {
	Sptpd_Id           *string         `json:"sptpd_id"`
	Ppat_Id            *string         `json:"ppat_id"`
	NoAkta             *string         `json:"noAkta"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
}

type UpdateDto struct {
	Id                 *uint64         `json:"id"`
	Sptpd_Id           *string         `json:"sptpd_id"`
	Ppat_Id            *string         `json:"ppat_id"`
	NoAkta             *string         `json:"noAkta"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
}

type FilterDto struct {
	Sptpd_Id           *string         `json:"sptpd_id"`
	Ppat_Id            *string         `json:"ppat_id"`
	NoAkta             *string         `json:"noAkta"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
	Page               int             `json:"page"`
	PageSize           int             `json:"page_size"`
}

type ResponsePelaporanPpat struct {
	Ppat_Id     *string `json:"ppat_id"`
	Ppat_Name   *string `json:"ppat_name"`
	TglLapor    *string `json:"tglLapor"`
	Sptpd_Id    *string `json:"sptpd_Id"`
	NilaiOp     *string `json:"nilaiOp"`
	JumlahSetor *string `json:"jumlahSetor"`
	Status      *string `json:"status"`
}

type ResponseDetilPelaporanPpat struct {
	Ppat_Id                   *string         `json:"ppat_id"`
	Ppat_Name                 *string         `json:"ppat_name"`
	TglLapor                  *string         `json:"tglLapor"`
	PihakYgMengalihkan        *string         `json:"pihakYgMengalihkan"`
	TglSSP                    *datatypes.Date `json:"tglSSP"`
	NominalSSP                *float64        `json:"nominalSSP"`
	Sptpd_Id                  *string         `json:"sptpd_Id"`
	Status                    *string         `json:"status"`
	WajibPajak_NIK            *string         `json:"wajibPajak_NIK"`
	PermohonanProvinsiID      *string         `json:"permohonanProvinsiID"`
	PermohonanKotaID          *string         `json:"permohonanKotaID"`
	PermohonanKecamatanID     *string         `json:"permohonanKecamatanID"`
	PermohonanKelurahanID     *string         `json:"permohonanKelurahanID"`
	PermohonanBlokID          *string         `json:"permohonanBlokID"`
	NoUrutPemohon             *string         `json:"noUrutPemohon"`
	PemohonJenisOPID          *string         `json:"pemohonJenisOPID"`
	NOPAlamat                 *string         `json:"opAlamat"`
	OPProvinsi                *string         `json:"opProvinsi"`
	OPKota                    *string         `json:"opKota"`
	OPKecamatan               *string         `json:"opKecamatan"`
	OPKelurahan               *string         `json:"opKelurahan"`
	OP_RtRW                   *string         `json:"op_RtRW"`
	OPLuasTanah               float64         `json:"opLuasTanah"`
	OPLuasBangunan            float64         `json:"opLuasBangunan"`
	OPLuasTanahBersama        float64         `json:"opLuasTanahBersama"`
	OPLuasBangunanBersama     float64         `json:"opLuasBangunanBersama"`
	NjopLuasTanah             float64         `json:"njopLuasTanah"`
	NjopLuasBangunan          float64         `json:"njopLuasBangunan"`
	NjopTanahBersama          float64         `json:"njopTanahBersama"`
	NjopBangunanBersama       float64         `json:"njopBangunanBersama"`
	NilaiOp                   float64         `json:"nilaiOp"`
	JenisPerolehanOp          *string         `json:"jenisPerolehanOp"`
	NoSertifikatOp            *string         `json:"noSertifikatOp"`
	NjopPbbOp                 float64         `json:"njopPbbOp"`
	LokasiOp                  *string         `json:"lokasiOp"`
	TahunPajakSppt            *string         `json:"tahunPajakSppt"`
	Npop                      float64         `json:"npop"`
	Npoptkp                   float64         `json:"npoptkp"`
	JenisSetoran              *string         `json:"jenisSetoran"`
	JenisSetoranNomor         *string         `json:"jenisSetoranNomor"`
	JenisSetoranTanggal       *datatypes.Date `json:"jenisSetoranTanggal"`
	JenisSetoranHitungSendiri *string         `json:"jenisSetoranHitungSendiri"`
	JenisSetoranCustom        *string         `json:"jenisSetoranCustom"`
	JumlahSetor               float64         `json:"jumlahSetor"`
	NominalSPT                float64         `json:"nominalSPT"`
	Tanggal                   *datatypes.Date `json:"tanggal"`
	NoDokumen                 *string         `json:"noDokumen"`
	NopPbbBaru                *string         `json:"nopPbbBaru"`
	User_id                   *string         `json:"user_id"`
	Id_pp                     *string         `json:"id_pp"`
	JenisSetoranKeterangan    *string         `json:"jenisSetoranKeterangan"`
	NilaiPasar                float64         `json:"nilaiPasar"`
	JenisPerolehan_id         *string         `json:"jenisPerolehan_id"`
	InpAphb1                  *string         `json:"inpAphb1"`
	InpAphb2                  *string         `json:"inpAphb2"`
	InpAphb3                  *string         `json:"inpAphb3"`
	KodeValidasi              *string         `json:"kodeValidasi"`
	GambarOp                  *string         `json:"gambarOp"`
	IsLunas                   *string         `json:"isLunas"`
	ValidasiDisependa         *string         `json:"validasiDisependa"`
	ValidasiBank              *string         `json:"validasiBank"`
	VerifikasiPpatAt          *time.Time      `json:"verifikasiPpatAt"`
	TglValidasiBank           *datatypes.Date `json:"tglValidasiBank"`
	TglValidasiDispenda       *datatypes.Date `json:"tglValidasiDispenda"`
	Dispenda_User_id          *string         `json:"pegawai_User_id"`
	AlasanReject              *string         `json:"alasanReject"`
	Bank_Id                   *string         `json:"bank_Id"`
	GambarInt                 *string         `json:"gambarInt"`
	FlagDispenda              *int            `json:"flagDispenda"`
	FlagPPAT                  *int            `json:"flagPPAT"`
	TanahInpAphb1             *float64        `json:"tanahInpAphb1"`
	TanahInpAphb2             *float64        `json:"tanahInpAphb2"`
	TanahInpAphb3             *float64        `json:"tanahInpAphb3"`
	BangunanInpAphb1          *float64        `json:"bangunanInpAphb1"`
	BangunanInpAphb2          *float64        `json:"bangunanInpAphb2"`
	BangunanInpAphb3          *float64        `json:"bangunanInpAphb3"`
	TanahBersamaInpAphb1      *float64        `json:"tanahBersamaInpAphb1"`
	TanahBersamaInpAphb2      *float64        `json:"tanahBersamaInpAphb2"`
	TanahBersamaInpAphb3      *float64        `json:"tanahBersamaInpAphb3"`
	BangunanBersamaInpAphb1   *float64        `json:"bangunanBersamaInpAphb1"`
	BangunanBersamaInpAphb2   *float64        `json:"bangunanBersamaInpAphb2"`
	BangunanBersamaInpAphb3   *float64        `json:"bangunanBersamaInpAphb3"`
	IsKurangBayar             *string         `json:"isKurangBayar"`
	KurangBayar               *float64        `json:"kurangBayar"`
	SSPDLama                  *int            `json:"sspdLama"`
	NoReff                    *string         `json:"noReff"`
	IdBilling                 *string         `json:"idBilling"`
	Proses                    *string         `json:"proses"`
	NamaWp                    *string         `json:"namaWp"`
	Alamat                    *string         `json:"alamat"`
	Provinsi_Id               *string         `json:"provinsi_id"`
	Kabupaten_id              *string         `json:"kabupaten_id"`
	Kecamatan_Id              *string         `json:"kecamatan_id"`
	Kelurahan_Id              *string         `json:"kelurahan_id"`
	KodePos                   *string         `json:"kodePos"`
	RtRw                      *string         `json:"rtRw"`
	WajibPajak_id             *uint64         `json:"wajibPajak_id"`
	ApprovePPAT               *string         `json:"approvePPAT"`
	NoPelayanan               *string         `json:"noPelayanan"`
	NamaPetugasLapangan       *string         `json:"namaPetugasLapangan"`
	NamaStaff                 *string         `json:"namaStaff"`
	NoAkta                    *string         `json:"noAkta"`
	TglAkta                   *datatypes.Date `json:"tglAkta"`
	PihakMengalihkan          *string         `json:"pihakMengalihkan"`
	Tgl                       *datatypes.Date `json:"tgl"`
	Rp                        *float64        `json:"rp"`
	Batas                     *int            `json:"batas"`
	Pengurangan               *float64        `json:"pengurangan"`
	SkPengurangan             *string         `json:"skPengurangan"`
	TglSkPengurangan          *datatypes.Date `json:"tglSkPengurangan"`
	WajibPajak                *string         `json:"wajibPajak"`
	TglExpBilling             *datatypes.Date `json:"tglExpBilling"`
	JenisTransaksiSpop        *int            `json:"jenisTransaksiSpop"`
	Pekerjaan                 *int            `json:"pekerjaan"`
	JenisTanah                *int            `json:"jenisTanah"`
	JumlahBangunan            *int            `json:"jumlahBangunan"`
	BlokKavNo                 *string         `json:"blokKavNo"`
	RT                        *string         `json:"rT"`
	Rw                        *string         `json:"rw"`
}

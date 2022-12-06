package bphtb

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Sptpd struct {
	Id                        uint            `json:"id" gorm:"primarykey"`
	Sptpd_Id                  *string         `json:"sptpd_Id" gorm:"type:varchar(50)"`
	Ppat_Id                   *string         `json:"ppat_Id" gorm:"type:varchar(50)"`
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
	OPLuasTanah               *float64        `json:"opLuasTanah" gorm:"type:number(10,2)"`
	OPLuasBangunan            *float64        `json:"opLuasBangunan" gorm:"type:number(10,2)"`
	OPLuasTanahBersama        *float64        `json:"opLuasTanahBersama" gorm:"type:number(10,2)"`
	OPLuasBangunanBersama     *float64        `json:"opLuasBangunanBersama" gorm:"type:number(10,2)"`
	NjopLuasTanah             *float64        `json:"njopLuasTanah" gorm:"type:number(10,2)"`
	NjopLuasBangunan          *float64        `json:"njopLuasBangunan" gorm:"type:number(10,2)"`
	NjopTanahBersama          *float64        `json:"njopTanahBersama" gorm:"type:number(20,2)"`
	NjopBangunanBersama       *float64        `json:"njopBangunanBersama" gorm:"type:number(20,2)"`
	NilaiOp                   *float64        `json:"nilaiOp" gorm:"type:number(20,2)"`
	JenisPerolehanOp          *string         `json:"jenisPerolehanOp" gorm:"type:varchar(50)"`
	NoSertifikatOp            *string         `json:"noSertifikatOp" gorm:"type:varchar(50)"`
	NjopPbbOp                 *float64        `json:"njopPbbOp" gorm:"type:number(20,2)"`
	LokasiOp                  *string         `json:"lokasiOp" gorm:"type:varchar(50)"`
	TahunPajakSppt            *string         `json:"tahunPajakSppt" gorm:"type:varchar(50)"`
	Npop                      *float64        `json:"npop" gorm:"type:number(50,2)"`
	Npoptkp                   *float64        `json:"npoptkp" gorm:"type:number(50,2)"`
	JenisSetoran              *string         `json:"jenisSetoran" gorm:"type:varchar(50)"`
	JenisSetoranNomor         *string         `json:"jenisSetoranNomor" gorm:"type:varchar(50)"`
	JenisSetoranTanggal       *datatypes.Date `json:"jenisSetoranTanggal"`
	JenisSetoranHitungSendiri *string         `json:"jenisSetoranHitungSendiri" gorm:"type:varchar(50)"`
	JenisSetoranCustom        *string         `json:"jenisSetoranCustom" gorm:"type:varchar(50)"`
	JumlahSetor               *float64        `json:"jumlahSetor" gorm:"type:number(50,2)"`
	NominalSPT                *float64        `json:"nominalSPT" gorm:"type:number(50,2)"`
	Tanggal                   *datatypes.Date `json:"tanggal"`
	NoDokumen                 *string         `json:"noDokumen" gorm:"type:varchar(50)"`
	NopPbbBaru                *string         `json:"nopPbbBaru" gorm:"type:varchar(50)"`
	User_id                   *string         `json:"user_id" gorm:"type:varchar(50)"`
	Id_pp                     *string         `json:"id_pp" gorm:"type:varchar(50)"`
	JenisSetoranKeterangan    *string         `json:"jenisSetoranKeterangan" gorm:"type:varchar(50)"`
	NilaiPasar                *float64        `json:"nilaiPasar" gorm:"type:number(50,2)"`
	JenisPerolehan_id         *string         `json:"jenisPerolehan_id" gorm:"type:varchar(50)"`
	InpAphb1                  *string         `json:"inpAphb1" gorm:"type:varchar(50)"`
	InpAphb2                  *string         `json:"inpAphb2" gorm:"type:varchar(50)"`
	InpAphb3                  *string         `json:"inpAphb3" gorm:"type:varchar(50)"`
	KodeValidasi              *string         `json:"kodeValidasi" gorm:"type:varchar(50)"`
	GambarOp                  *string         `json:"GambarOp" gorm:"type:varchar(50)"`
	IsLunas                   *string         `json:"isLunas" gorm:"type:varchar(50)"`
	ValidasiDisependa         *string         `json:"validasiDisependa" gorm:"type:varchar(50)"`
	ValidasiBank              *string         `json:"validasiBank" gorm:"type:varchar(50)"`
	TglValidasiDispenda       *datatypes.Date `json:"tglValidasiDispenda"`
	TglValidasiBank           *datatypes.Date `json:"tglValidasiBank"`
	Dispenda_User_id          *string         `json:"pegawai_User_id" gorm:"type:varchar(50)"`
	AlasanReject              *string         `json:"alasanReject" gorm:"type:text"`
	Bank_Id                   *string         `json:"bank_Id" gorm:"type:varchar(50)"`
	GambarInt                 *string         `json:"gambarInt" gorm:"type:text"`
	FlagDispenda              *int            `json:"flagDispenda" gorm:"type:int"`
	FlagPPAT                  *int            `json:"flagPPAT" gorm:"type:int"`
	TanahInpAphb1             *float64        `json:"tanahInpAphb1" gorm:"type:number(20,2)"`
	TanahInpAphb2             *float64        `json:"tanahInpAphb2" gorm:"type:number(20,2)"`
	TanahInpAphb3             *float64        `json:"tanahInpAphb3" gorm:"type:number(20,2)"`
	BangunanInpAphb1          *float64        `json:"bangunanInpAphb1" gorm:"type:number(20,2)"`
	BangunanInpAphb2          *float64        `json:"bangunanInpAphb2" gorm:"type:number(20,2)"`
	BangunanInpAphb3          *float64        `json:"bangunanInpAphb3" gorm:"type:number(20,2)"`
	TanahBersamaInpAphb1      *float64        `json:"tanahBersamaInpAphb1" gorm:"type:number(20,2)"`
	TanahBersamaInpAphb2      *float64        `json:"tanahBersamaInpAphb2" gorm:"type:number(20,2)"`
	TanahBersamaInpAphb3      *float64        `json:"tanahBersamaInpAphb3" gorm:"type:number(20,2)"`
	BangunanBersamaInpAphb1   *float64        `json:"bangunanBersamaInpAphb1" gorm:"type:number(20,2)"`
	BangunanBersamaInpAphb2   *float64        `json:"bangunanBersamaInpAphb2" gorm:"type:number(20,2)"`
	BangunanBersamaInpAphb3   *float64        `json:"bangunanBersamaInpAphb3" gorm:"type:number(20,2)"`
	IsKurangBayar             *interface{}    `json:"isKurangBayar" gorm:"type:varchar(50)"`
	KurangBayar               *float64        `json:"kurangBayar" gorm:"type:number(20,2)"`
	SSPDLama                  *int            `json:"sspdLama" gorm:"type:int"`
	NoReff                    *string         `json:"noReff" gorm:"type:varchar(50)"`
	IdBilling                 *string         `json:"idBilling" gorm:"type:varchar(50)"`
	Proses                    *interface{}    `json:"proses" gorm:"type:varchar(50)"`
	NamaWp                    *string         `json:"namaWp" gorm:"type:varchar(50)"`
	Alamat                    *string         `json:"alamat" gorm:"type:varchar(50)"`
	Provinsi_Id               *uint64         `json:"provinsi_Id" gorm:"type:int"`
	Kabupaten_id              *uint64         `json:"kabupaten_id" gorm:"type:int"`
	Kecamatan_Id              *uint64         `json:"kecamatan_Id" gorm:"type:int"`
	Kelurahan_Id              *uint64         `json:"kelurahan_Id" gorm:"type:int"`
	KodePos                   *string         `json:"kodePos" gorm:"type:varchar(50)"`
	RtRw                      *string         `json:"rtRw" gorm:"type:varchar(50)"`
	WajibPajak_id             *uint64         `json:"wajibPajak_id" gorm:"type:int"`
	ApprovePPAT               *interface{}    `json:"approvePPAT" gorm:"type:varchar(50)"`
	NoPelayanan               *string         `json:"noPelayanan" gorm:"type:varchar(50)"`
	NamaPetugasLapangan       *string         `json:"namaPetugasLapangan" gorm:"type:varchar(50)"`
	NamaStaff                 *string         `json:"namaStaff" gorm:"type:varchar(50)"`
	NoAkta                    *string         `json:"noAkta" gorm:"type:varchar(50)"`
	TglAkta                   *datatypes.Date `json:"tglAkta"`
	PihakMengalihkan          *string         `json:"pihakMengalihkan" gorm:"type:varchar(50)"`
	Tgl                       *datatypes.Date `json:"tgl"`
	Rp                        *float64        `json:"rp" gorm:"type:number(15,2)"`
	Batas                     *int            `json:"batas" gorm:"type:int"`
	Pengurangan               *float64        `json:"pengurangan" gorm:"type:number(15,2)"`
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
	gormhelper.DateModel
}

type RequestSptpd struct {
	Sptpd_Id                  *string         `json:"sptpd_Id"`
	Ppat_Id                   *string         `json:"ppat_Id"`
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
	OPLuasTanah               *float64        `json:"opLuasTanah"`
	OPLuasBangunan            *float64        `json:"opLuasBangunan"`
	OPLuasTanahBersama        *float64        `json:"opLuasTanahBersama"`
	OPLuasBangunanBersama     *float64        `json:"opLuasBangunanBersama"`
	NjopLuasTanah             *float64        `json:"njopLuasTanah"`
	NjopLuasBangunan          *float64        `json:"njopLuasBangunan"`
	NjopTanahBersama          *float64        `json:"njopTanahBersama"`
	NjopBangunanBersama       *float64        `json:"njopBangunanBersama"`
	NilaiOp                   *float64        `json:"nilaiOp"`
	JenisPerolehanOp          *string         `json:"jenisPerolehanOp"`
	NoSertifikatOp            *string         `json:"noSertifikatOp"`
	NjopPbbOp                 *float64        `json:"njopPbbOp"`
	LokasiOp                  *string         `json:"lokasiOp"`
	TahunPajakSppt            *string         `json:"tahunPajakSppt"`
	Npop                      *float64        `json:"npop"`
	Npoptkp                   *float64        `json:"npoptkp"`
	JenisSetoran              *string         `json:"jenisSetoran"`
	JenisSetoranNomor         *string         `json:"jenisSetoranNomor"`
	JenisSetoranTanggal       *datatypes.Date `json:"jenisSetoranTanggal"`
	JenisSetoranHitungSendiri *string         `json:"jenisSetoranHitungSendiri"`
	JenisSetoranCustom        *string         `json:"jenisSetoranCustom"`
	JumlahSetor               *float64        `json:"jumlahSetor"`
	NominalSPT                *float64        `json:"nominalSPT"`
	Tanggal                   *datatypes.Date `json:"tanggal"`
	NoDokumen                 *string         `json:"noDokumen"`
	NopPbbBaru                *string         `json:"nopPbbBaru"`
	User_id                   *string         `json:"user_id"`
	Id_pp                     *string         `json:"id_pp"`
	JenisSetoranKeterangan    *string         `json:"jenisSetoranKeterangan"`
	NilaiPasar                *float64        `json:"nilaiPasar"`
	JenisPerolehan_id         *string         `json:"jenisPerolehan_id"`
	InpAphb1                  *string         `json:"inpAphb1"`
	InpAphb2                  *string         `json:"inpAphb2"`
	InpAphb3                  *string         `json:"inpAphb3"`
	KodeValidasi              *string         `json:"kodeValidasi"`
	GambarOp                  *string         `json:"GambarOp"`
	IsLunas                   *string         `json:"isLunas"`
	ValidasiDisependa         *string         `json:"validasiDisependa"`
	ValidasiBank              *string         `json:"validasiBank"`
	TglValidasiDispenda       *datatypes.Date `json:"tglValidasiDispenda"`
	TglValidasiBank           *datatypes.Date `json:"tglValidasiBank"`
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
	IsKurangBayar             *interface{}    `json:"isKurangBayar"`
	KurangBayar               *float64        `json:"kurangBayar"`
	SSPDLama                  *int            `json:"sspdLama"`
	NoReff                    *string         `json:"noReff"`
	IdBilling                 *string         `json:"idBilling"`
	Proses                    *interface{}    `json:"proses"`
	NamaWp                    *string         `json:"namaWp"`
	Alamat                    *string         `json:"alamat"`
	Provinsi_Id               *uint64         `json:"provinsi_Id"`
	Kabupaten_id              *uint64         `json:"kabupaten_id"`
	Kecamatan_Id              *uint64         `json:"kecamatan_Id"`
	Kelurahan_Id              *uint64         `json:"kelurahan_Id"`
	KodePos                   *string         `json:"kodePos"`
	RtRw                      *string         `json:"rtRw"`
	WajibPajak_id             *uint64         `json:"wajibPajak_id"`
	ApprovePPAT               *interface{}    `json:"approvePPAT"`
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
	gormhelper.DateModel
}

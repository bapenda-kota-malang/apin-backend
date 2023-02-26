package penetapan

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PenerimaKompensasi struct {
	Id                        uint64   `json:"id" gorm:"primaryKey"`
	Kanwil_Kode               string   `json:"kanwil_kode" gorm:"type:varchar(2)"`
	KPPBB_Kode                string   `json:"kppbb_kode" gorm:"type:varchar(2)"`
	ThnPelayanan              string   `json:"thnPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan           string   `json:"bundelPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan           string   `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	Provinsi_Kode_Pemohon     string   `json:"provinsi_kode_pemohon" gorm:"type:varchar(2)"`
	Daerah_Kode_Pemohon       string   `json:"daerah_kode_pemohon" gorm:"type:varchar(2)"`
	Kecamatan_Kode_Pemohon    string   `json:"kecamatan_kode_pemohon" gorm:"type:varchar(3)"`
	Kelurahan_Kode_Pemohon    string   `json:"kelurahan_kode_pemohon" gorm:"type:varchar(3)"`
	Blok_Kode_Pemohon         string   `json:"blok_kode_pemohon" gorm:"type:varchar(3)"`
	NoUrut_Pemohon            string   `json:"noUrut_pemohon" gorm:"type:varchar(4)"`
	JenisOp_Pemohon           string   `json:"jenisOp_pemohon" gorm:"type:varchar(1)"`
	NoUrutPenerima            *uint64  `json:"noUrutPenerima"`
	Provinsi_Kode_Kompensasi  *string  `json:"provinsi_kode_kompensasi" gorm:"type:varchar(2)"`
	Daerah_Kode_Kompensasi    *string  `json:"daerah_kode_kompensasi" gorm:"type:varchar(2)"`
	Kecamatan_Kode_Kompensasi *string  `json:"kecamatan_kode_kompensasi" gorm:"type:varchar(3)"`
	Kelurahan_Kode_Kompensasi *string  `json:"kelurahan_kode_kompensasi" gorm:"type:varchar(3)"`
	Blok_Kode_Kompensasi      *string  `json:"blok_kode_kompensasi" gorm:"type:varchar(3)"`
	NoUrut_Kompensasi         *string  `json:"noUrut_kompensasi" gorm:"type:varchar(4)"`
	JenisOp_Kompensasi        *string  `json:"jenisOp_kompensasi" gorm:"type:varchar(1)"`
	TahunPajakKompensasi      *string  `json:"tahunPajakKompensasi" gorm:"type:varchar(4)"`
	NilaiYangDikompensasi     *float64 `json:"nilaiYangDikompensasi"`
	gh.DateModel
}

type SubjekPajakNJOPTKP struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode  *string `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode    *string `json:"daerah_kode" gorm:"type:varchar(2)"`
	Kecamatan_Kode *string `json:"kecamatan_kode" gorm:"type:varchar(3)"`
	Kelurahan_Kode *string `json:"kelurahan_kode" gorm:"type:varchar(3)"`
	Blok_Kode      *string `json:"blok_kode" gorm:"type:varchar(3)"`
	NoUrut         *string `json:"noUrut" gorm:"type:varchar(4)"`
	JenisOp        *string `json:"jenisOp" gorm:"type:varchar(1)"`
	TahunPajak     *string `json:"tahunPajak" gorm:"type:varchar(4)"`
	gh.DateModel
}

type Tarif struct {
	Id            uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode *string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode   *string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	TahunAwal     *string  `json:"tahunAwal" gorm:"type:varchar(4)"`
	TahunAkhir    *string  `json:"tahunAkhir" gorm:"type:varchar(4)"`
	NJOPMix       *float64 `json:"NJOPMix"`
	NJOPMax       *float64 `json:"NJOPMax"`
	NilaiTarif    *float64 `json:"nilaiTarif"`
	gh.DateModel
}

type PenguranganPengenaanJPB struct {
	Id                     uint64   `json:"id" gorm:"primaryKey"`
	Kanwil_Kode            string   `json:"kanwil_kode" gorm:"type:varchar(2)"`
	KPPBB_Kode             string   `json:"kppbb_kode" gorm:"type:varchar(2)"`
	ThnPelayanan           string   `json:"thnPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan        string   `json:"bundelPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan        string   `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	Provinsi_Kode_Pemohon  string   `json:"provinsi_kode_pemohon" gorm:"type:varchar(2)"`
	Daerah_Kode_Pemohon    string   `json:"daerah_kode_pemohon" gorm:"type:varchar(2)"`
	Kecamatan_Kode_Pemohon string   `json:"kecamatan_kode_pemohon" gorm:"type:varchar(3)"`
	Kelurahan_Kode_Pemohon string   `json:"kelurahan_kode_pemohon" gorm:"type:varchar(3)"`
	Blok_Kode_Pemohon      string   `json:"blok_kode_pemohon" gorm:"type:varchar(3)"`
	NoUrut_Pemohon         string   `json:"noUrut_pemohon" gorm:"type:varchar(4)"`
	JenisOp_Pemohon        string   `json:"jenisOp_pemohon" gorm:"type:varchar(1)"`
	ThnPengenaan           *string  `json:"thnPengenaan" gorm:"type:varchar(4)"`
	JnsSK                  *string  `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK                   *string  `json:"noSK" gorm:"type:varchar(30)"`
	PCTPenguranganJPB      *float64 `json:"PCTPenguranganJPB"`
	gh.DateModel
}

type PenguranganPermanen struct {
	Id                     uint64   `json:"id" gorm:"primaryKey"`
	Kanwil_Kode            string   `json:"kanwil_kode" gorm:"type:varchar(2)"`
	KPPBB_Kode             string   `json:"kppbb_kode" gorm:"type:varchar(2)"`
	ThnPelayanan           string   `json:"thnPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan        string   `json:"bundelPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan        string   `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	Provinsi_Kode_Pemohon  string   `json:"provinsi_kode_pemohon" gorm:"type:varchar(2)"`
	Daerah_Kode_Pemohon    string   `json:"daerah_kode_pemohon" gorm:"type:varchar(2)"`
	Kecamatan_Kode_Pemohon string   `json:"kecamatan_kode_pemohon" gorm:"type:varchar(3)"`
	Kelurahan_Kode_Pemohon string   `json:"kelurahan_kode_pemohon" gorm:"type:varchar(3)"`
	Blok_Kode_Pemohon      string   `json:"blok_kode_pemohon" gorm:"type:varchar(3)"`
	NoUrut_Pemohon         string   `json:"noUrut_pemohon" gorm:"type:varchar(4)"`
	JenisOp_Pemohon        string   `json:"jenisOp_pemohon" gorm:"type:varchar(1)"`
	ThnAwal                *string  `json:"thnAwal" gorm:"type:varchar(4)"`
	ThnAkhir               *string  `json:"thnAkhir" gorm:"type:varchar(4)"`
	JnsSK                  *string  `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK                   *string  `json:"noSK" gorm:"type:varchar(30)"`
	StatusSK               *string  `json:"statusSK" gorm:"type:varchar(1)"`
	PCTPenguranganPermanen *float64 `json:"PCTPenguranganPermanen"`
	gh.DateModel
}

type PenguranganPST struct {
	Id                     uint64   `json:"id" gorm:"primaryKey"`
	Kanwil_Kode            string   `json:"kanwil_kode" gorm:"type:varchar(2)"`
	KPPBB_Kode             string   `json:"kppbb_kode" gorm:"type:varchar(2)"`
	ThnPelayanan           string   `json:"thnPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan        string   `json:"bundelPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan        string   `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	Provinsi_Kode_Pemohon  string   `json:"provinsi_kode_pemohon" gorm:"type:varchar(2)"`
	Daerah_Kode_Pemohon    string   `json:"daerah_kode_pemohon" gorm:"type:varchar(2)"`
	Kecamatan_Kode_Pemohon string   `json:"kecamatan_kode_pemohon" gorm:"type:varchar(3)"`
	Kelurahan_Kode_Pemohon string   `json:"kelurahan_kode_pemohon" gorm:"type:varchar(3)"`
	Blok_Kode_Pemohon      string   `json:"blok_kode_pemohon" gorm:"type:varchar(3)"`
	NoUrut_Pemohon         string   `json:"noUrut_pemohon" gorm:"type:varchar(4)"`
	JenisOp_Pemohon        string   `json:"jenisOp_pemohon" gorm:"type:varchar(1)"`
	ThnPengPST             *string  `json:"thnPengPST" gorm:"type:varchar(4)"`
	JnsSK                  *string  `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK                   *string  `json:"noSK" gorm:"type:varchar(30)"`
	StatusSK               *string  `json:"statusSK" gorm:"type:varchar(1)"`
	PCTPenguranganPST      *float64 `json:"PCTPenguranganPST"`
	gh.DateModel
}

type SubjekPajak struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	Nama_WP         string  `json:"nama_WP" gorm:"type:varchar(30)"`
	Jalan_WP        string  `json:"jalan_WP" gorm:"type:varchar(30)"`
	Blok_Kav_No     string  `json:"blok_Kav_No" gorm:"type:varchar(15)"`
	RT              *string `json:"rt" gorm:"type:varchar(3)"`
	RW              *string `json:"rw" gorm:"type:varchar(2)"`
	Kelurahan_Kode  *string `json:"Kelurahan_Kode" gorm:"type:varchar(2)"`
	Kota_Kode       *string `json:"Kota_Kode" gorm:"type:varchar(3)"`
	Kode_Pos        *string `json:"Kode_Pos" gorm:"type:varchar(5)"`
	NoUrut_Pemohon  *string `json:"noUrut_pemohon" gorm:"type:varchar(4)"`
	NoTelp          *string `json:"NoTelp" gorm:"type:varchar(20)"`
	NPWP            *string `json:"NPWP" gorm:"type:varchar(12)"`
	StatusPekerjaan *string `json:"statusPekerjaan" gorm:"type:varchar(1)"`
	gh.DateModel
}

type DafnomOP struct {
	Id              uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode   *string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode     *string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	Kecamatan_Kode  *string  `json:"kecamatan_kode" gorm:"type:varchar(3)"`
	Kelurahan_Kode  *string  `json:"kelurahan_kode" gorm:"type:varchar(3)"`
	Blok_Kode       *string  `json:"blok_kode" gorm:"type:varchar(3)"`
	NoUrut          *string  `json:"noUrut" gorm:"type:varchar(4)"`
	JenisOp         *string  `json:"jenisOp" gorm:"type:varchar(1)"`
	Blok_Kav_No     *string  `json:"Blok_Kav_No" gorm:"type:varchar(15)"`
	RT              *string  `json:"rt" gorm:"type:varchar(3)"`
	RW              *string  `json:"rw" gorm:"type:varchar(2)"`
	Jns_Bumi        *string  `json:"Jns_Bumi" gorm:"type:varchar(1)"`
	JPB_Kode        *string  `json:"JPB_Kode" gorm:"type:varchar(2)"`
	Kategori_OP     *string  `json:"Kategori_OP" gorm:"type:varchar(1)"`
	Keterangan      *string  `json:"keterangan" gorm:"type:varchar(255)"`
	NoFormulir      *string  `json:"formuliar" gorm:"type:varchar(30)"`
	ThnPembentukan  *string  `json:"ThnPembentukan" gorm:"type:varchar(30)"`
	TglPemutakhiran *string  `json:"TglPemutakhiran" gorm:"type:varchar(9)"`
	Status_WP_Kode  *float64 `json:"status_WP_kode"`
}

type ReferensiBuku struct {
	Id                uint64   `json:"id" gorm:"primaryKey"`
	Kode              *string  `json:"kode" gorm:"type:varchar(1)"`
	ThnAwal           *string  `json:"thnAwal" gorm:"type:varchar(4)"`
	ThnAkhir          *string  `json:"thnAkhir" gorm:"type:varchar(4)"`
	NilaiMin          *float64 `json:"nilaiMin"`
	NilaiMax          *float64 `json:"nilaiMax"`
	LuasMinTipe       *int     `json:"luasMinTipe"`
	LuasMaxTipe       *int     `json:"luasMaxTipe"`
	FaktorPembagiTipe *float64 `json:"faktorPembagiTipe"`
	Status            *float64 `json:"STATUS"`
	gh.DateModel
}

type NJOPTKP struct {
	Id            uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Kode string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode   string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	ThnAwal       string  `json:"thnAwal" gorm:"type:varchar(4)"`
	ThnAkhir      string  `json:"thnAkhir" gorm:"type:varchar(4)"`
	NilaiNJOPTKP  float64 `json:"nilaiNJOPTKP"`
	gh.DateModel
}

type NJKP struct {
	Id             uint64  `json:"id" gorm:"primaryKey"`
	IndeksRange    string  `json:"indeksRange" gorm:"type:varchar(2)"`
	JNS_Range_Kode string  `json:"jns_range_kode" gorm:"type:varchar(1)"`
	Provinsi_Kode  string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode    string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	JPB_JPT_Kode   string  `json:"jpb_jpt_kode" gorm:"type:varchar(2)"`
	UrutanNJKP     int     `json:"urutanNJKP"`
	NJOPMin        float64 `json:"NJOPMin"`
	NJOPMax        float64 `json:"NJOPMax"`
	NilaiNJKP      float64 `json:"nilaiNJKP"`
	gh.DateModel
}

type ParameterSPPTSTTSDHKP struct {
	Id               uint64  `json:"id" gorm:"primaryKey"`
	SPPTTerpisah     int     `json:"SPPTTerpisah"`
	STTSTerpisah     int     `json:"STTSTerpisah"`
	DHKPTerpisah     int     `json:"DHKPTerpisah"`
	KelompokBuku     *string `json:"KelompokBuku" gorm:"type:varchar(2)"`
	TahunPajakSPPT   int     `json:"TahunPajakSPPT"`
	TahunPajakSTTS   int     `json:"TahunPajakSTTS"`
	NamaKANWIL       *string `json:"NamaKANWIL" gorm:"type:varchar(30)"`
	NamaKPPBB        *string `json:"NamaKPPBB" gorm:"type:varchar(30)"`
	KotaTerbit       *string `json:"KotaTerbit" gorm:"type:varchar(30)"`
	NamaKepalaKPPBB  *string `json:"NamaKepalaKPPBB" gorm:"type:varchar(30)"`
	BarcodeSPPT      *string `json:"BarcodeSPPT" gorm:"type:varchar(50)"`
	JnsBarcodeSPPT   *string `json:"JnsBarcodeSPPT" gorm:"type:varchar(2)"`
	BarcodeSTTS      *string `json:"BarcodeSTTS" gorm:"type:varchar(50)"`
	JnsBarcodeSTTS   *string `json:"JnsBarcodeSTTS" gorm:"type:varchar(2)"`
	SPPTFormBaru     *string `json:"SPPTFormBaru" gorm:"type:varchar(2)"`
	STTSFormBaru     *string `json:"STTSFormBaru" gorm:"type:varchar(2)"`
	SPPTFormLama     *string `json:"SPPTFormLama" gorm:"type:varchar(2)"`
	STTSFormLama     *string `json:"STTSFormLama" gorm:"type:varchar(2)"`
	SuratHimbauan    *string `json:"SuratHimbauan" gorm:"type:varchar(2)"`
	BukuHimbauan     *string `json:"BukuHimbauan" gorm:"type:varchar(2)"`
	TunggakanTHNLalu int     `json:"TunggakanTHNLalu"`
	SPPTPBBNOL       int     `json:"SPPTPBBNOL"`
	PBBMinimal       int     `json:"PBBMinimal"`
	TeksKANWIL       int     `json:"TeksKANWIL"`
	TeksKPPBB        int     `json:"TeksKPPBB"`
	TeksSPPT         int     `json:"TeksSPPT"`
	TeksSTTS         int     `json:"TeksSTTS"`
	TeksPratama      int     `json:"TeksPratama"`
	Kakap            *string `json:"Kakap" gorm:"type:varchar(4)"`
	gh.DateModel
}

type RefThnNJKPNJOPTKPTarif struct {
	Id                uint64  `json:"id" gorm:"primaryKey"`
	IndeksRange       string  `json:"indeksRange" gorm:"type:varchar(2)"`
	JNS_Range_Kode    string  `json:"jns_range_kode" gorm:"type:varchar(1)"`
	Provinsi_Kode     string  `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode       string  `json:"daerah_kode" gorm:"type:varchar(2)"`
	SK_NJOP_NJKP_Kode int     `json:"SK_NJOP_NJKP_Kode"`
	RangeThnAwal      *string `json:"rangeThnAwal" gorm:"type:varchar(4)"`
	RangeThnAkhir     *string `json:"rangeThnAkhir" gorm:"type:varchar(4)"`
	gh.DateModel
}

type ReportDataTable struct {
	No           uint16
	MasaPajak    string
	JatuhTempo   string
	Skpd         string
	Ketetapan    string
	Denda        string
	TelahDibayar string
	SisaPajak    string
}

type Pdf interface {
	GeneratePdf(outFile string) error
	AppendContent(ReportDataTable)
	SetTotal(map[string]float64)
}

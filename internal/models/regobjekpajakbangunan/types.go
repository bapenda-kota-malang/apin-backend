package regobjekpajakbangunan

import (
	mrfb "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
)

type Kondisi string
type JenisKonstruksi string
type JenisAtap string
type JenisDinding string
type JenisLantai string
type JenisLangitLangit string
type JenisTransaksi string
type KelasBangunan string
type JenisHotel string
type JumlahBintang string
type TipeBangunan string
type LetakTanki string
type VerifyStatus int

const (
	// kondisi bangunan
	KondisiSangatBaik Kondisi = "1" // sangat baik
	KondisiBaik       Kondisi = "2" // baik
	KondisiSedang     Kondisi = "3" // sedang
	KondisiJelek      Kondisi = "4" // jelek

	// jenis konstruksi
	JenisKonstruksiBaja     JenisKonstruksi = "1" // baja
	JenisKonstruksiBeton    JenisKonstruksi = "2" // beton
	JenisKonstruksiBatuBata JenisKonstruksi = "3" // batu bata
	JenisKonstruksiKayu     JenisKonstruksi = "4" // kayu

	// jenis atap
	JenisAtapDecrabon JenisAtap = "1" // decrabon/beton genteng glazur
	JenisAtapBeton    JenisAtap = "2" // beton
	JenisAtapBatuBata JenisAtap = "3" // batu bata
	JenisAtapKayu     JenisAtap = "4" // kayu

	// jenis dinding
	KodeDindingKaca            JenisDinding = "1" // kaca/aluminium
	KodeDindingBeton           JenisDinding = "2" // beton
	KodeDndingBatuBata         JenisDinding = "3" // batu bata/coblonk
	KodeDindingKayu            JenisDinding = "4" // kayu
	KodeDindingSeng            JenisDinding = "5" // seng
	KodeDindingTidakAdaDinding JenisDinding = "6" // tidak ada dinding

	// jenis lantai
	KodeLantaiMarmer        JenisLantai = "1" // marmer
	KodeLantaiMarmerKeramik JenisLantai = "2" // keramik
	KodeLantaiTeraso        JenisLantai = "3" // teraso
	KodeLantaiUbinPc        JenisLantai = "4" // ubin pc
	KodeLantaiSemen         JenisLantai = "5" // semen

	// jenis langit langit
	KodeLangitJati     JenisLangitLangit = "1" // akustik/jati
	KodeLangitTripleks JenisLangitLangit = "2" // tripleks/asbes bambu
	KodeLangitTidakAda JenisLangitLangit = "3" // tidak ada

	// jenis transaksi
	JenisTransaksiPerekam           JenisTransaksi = "21" // perekam data bangunan
	JenisTransaksiPemutakhiran      JenisTransaksi = "22" // pemutakhiran data bangunan
	JenisTransaksiPenghapusan       JenisTransaksi = "23" // penghapusan data bangunan
	JenisTransaksiPenilaianIndividu JenisTransaksi = "24" // penilaian individu

	KelasBangunan1 KelasBangunan = "1" // kelas 1
	KelasBangunan2 KelasBangunan = "2" // kelas 2
	KelasBangunan3 KelasBangunan = "3" // kelas 3
	KelasBangunan4 KelasBangunan = "4" // kelas 4

	// jenis hotel
	JenisHotelNonResort JenisHotel = "1" // non-resort
	JenisHotelResort    JenisHotel = "2" // resort

	// jumlah bintang
	JumlahBintang5          JumlahBintang = "1" // 5
	JumlahBintang4          JumlahBintang = "2" // 4
	JumlahBintang3          JumlahBintang = "3" // 3
	JumlahBintang2          JumlahBintang = "4" // 2
	JumlahBintangNonBintang JumlahBintang = "5" // non-bintang

	TipeBangunan4 TipeBangunan = "1" // 4
	TipeBangunan3 TipeBangunan = "2" // 3
	TipeBangunan2 TipeBangunan = "3" // 2
	TipeBangunan1 TipeBangunan = "4" // 1

	LetakTankiDiatasTanah  LetakTanki = "1" // di atas tanah
	LetakTankiDibawahTanah LetakTanki = "2" // di bawah tanah

	StatusBaru      VerifyStatus = 0
	StatusDisetujui VerifyStatus = 1
	StatusDitolak   VerifyStatus = 2
)

type RegInput interface {
	GetFasilitasBangunan() *mrfb.CreateDto
	GetNop() *string
	GetTanggalPendataan() *string
	GetTanggalPemeriksaan() *string
	GetTanggalPerekaman() *string
	GetObjekPajakBangunan() (CreateDto, error)
}

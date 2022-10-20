package objekpajakbangunan

type Kondisi string
type JenisKonstruksi string
type JenisAtap string
type JenisDinding string
type JenisLantai string
type JenisLangitLangit string
type JenisTransaksi string

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
)

package types

type Peruntukan string
type JenisABT string
type JenisKendaraan int8
type JenisPajak string
type Golongan int16
type Status int16 // StatusNpwdp
type StatusBL int16
type StatusVerifikasi uint8
type JenisMasa int16

const (
	PeruntukanNonNiaga    Peruntukan = "NON NIAGA"
	PeruntukanNiaga       Peruntukan = "NIAGA"
	PeruntukanIndustriAir Peruntukan = "INDUSTRI" // Industri dengan bahan baku air
	PeruntukanPdam        Peruntukan = "PDAM"

	JenisABTMataAir    JenisABT = "Mata Air"
	JenisABTNonMataAir JenisABT = "Bukan Mata Air"

	KendaraanRodaDua   JenisKendaraan = 1 //roda dua
	KendaraanRodaEmpat JenisKendaraan = 2 //roda empat

	JenisPajakSA JenisPajak = "SA" //self_assessment
	JenisPajakOA JenisPajak = "OA" //official_assessment

	GolonganOrangPribadi Golongan = 1 //orang pribadi
	GolonganBadan        Golongan = 2 //badan

	StatusAktif          Status = 1 //aktif
	StatusTutupSementara Status = 2 //tutup sementara
	StatusTutup          Status = 3 //tutup

	StatusBaru StatusBL = 0 //baru
	StatusLama StatusBL = 1 //lama

	StatusVerifikasiBaru             StatusVerifikasi = 0 // baru
	StatusVerifikasiDisetujuiKasubid StatusVerifikasi = 1 // DisetujuiKasubid
	StatusVerifikasiDisetujuiKabid   StatusVerifikasi = 2 // DisetujuiKabid
	StatusVerifikasiDitolakKasubid   StatusVerifikasi = 3 // DitolakKasubid
	StatusVerifikasiDitolakKabid     StatusVerifikasi = 4 // DitolakKabid

	MasaPajakTahun         JenisMasa = 1 //masa pajak tetap 1 tahun
	MasaPajakBulan         JenisMasa = 2 //masa pajak insidentil 1 bulan
	MasaPajakHari          JenisMasa = 3 //masa pajak insidentil 1 hari
	MasaPajakPenyelenggara JenisMasa = 4 //masa pajak insidentil 1 kali penyelenggaraan
)

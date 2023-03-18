package pengurangan

type StatusPengurangan string
type JenisPengurangan uint8
type PosisiVerifikasi uint8
type StatusVerifikasi uint8
type Status uint8

const (
	// 0 diterima
	StatusPenguranganDiterima StatusPengurangan = "0"
	// 1 diterima sebagian
	StatusPenguranganDiterimaSebagian StatusPengurangan = "1"
	// 2 ditolak
	StatusPenguranganDitolak StatusPengurangan = "2"

	// 1 Jumlah Pajak
	JenisPenguranganJumlahPajak JenisPengurangan = 1
	// 2 Denda
	JenisPenguranganDenda JenisPengurangan = 2

	PosisiVerifikasiStaff   PosisiVerifikasi = 0
	PosisiVerifikasiAnalis  PosisiVerifikasi = 1
	PosisiVerifikasiKasubid PosisiVerifikasi = 2
	PosisiVerifikasikabid   PosisiVerifikasi = 3
	PosisiVerifikasiSekban  PosisiVerifikasi = 4
	PosisiVerifikasiKaban   PosisiVerifikasi = 5

	StatusVerifikasiDisetujui StatusVerifikasi = 0
	StatusVerifikasiDitolak   StatusVerifikasi = 1

	StatusDiproses Status = 0
	StatusDiterima Status = 1
	StatusDitolak  Status = 2
)

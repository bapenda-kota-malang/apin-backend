package pengurangan

type StatusPengurangan string

const (
	// 0 diterima
	StatusPenguranganDiterima StatusPengurangan = "0"
	// 1 diterima sebagian
	StatusPenguranganDiterimaSebagian StatusPengurangan = "1"
	// 2 ditolak
	StatusPenguranganDitolak StatusPengurangan = "2"
)

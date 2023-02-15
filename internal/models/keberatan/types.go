package keberatan

type JnsKeputusan string

const (
	// Diterima
	JnsKeputusanDiterima JnsKeputusan = "0"
	// Diterima Sebagian
	JnsKeputusanDiterimaSebagian JnsKeputusan = "1"
	// Ditolak
	JnsKeputusanDitolak JnsKeputusan = "2"
	// Menambah Besarnya Pajak
	JnsKeputusanMenambahBesarnyaPajak JnsKeputusan = "3"
)

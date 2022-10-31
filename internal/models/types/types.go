package types

type Peruntukan string
type JenisABT string
type JenisKendaraan int8

const (
	PeruntukanNonNiaga    Peruntukan = "NON NIAGA"
	PeruntukanNiaga       Peruntukan = "NIAGA"
	PeruntukanIndustriAir Peruntukan = "INDUSTRI" // Industri dengan bahan baku air
	PeruntukanPdam        Peruntukan = "PDAM"

	JenisABTMataAir    JenisABT = "Mata Air"
	JenisABTNonMataAir JenisABT = "Bukan Mata Air"

	KendaraanRodaDua   JenisKendaraan = 1 //roda dua
	KendaraanRodaEmpat JenisKendaraan = 2 //roda empat
)

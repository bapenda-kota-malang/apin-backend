package detailesptair

type Peruntukan string
type JenisABT string

const (
	PeruntukanNonNiaga    Peruntukan = "Non Niaga"
	PeruntukanNiaga       Peruntukan = "Niaga"
	PeruntukanIndustriAir Peruntukan = "Industri dengan bahan baku air"
	PeruntukanPdam        Peruntukan = "PDAM"

	JenisABTMataAir    JenisABT = "Mata Air"
	JenisABTNonMataAir JenisABT = "Bukan Mata Air"
)

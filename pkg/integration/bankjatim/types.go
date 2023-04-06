package bankjatim

type FlagProses uint8
type StatusCode string

const (
	// Flag Create Process
	ProsesCreate FlagProses = 1
	// Flag Update Process
	ProsesUpdate FlagProses = 2
	// Flag Delete Process
	ProsesDelete FlagProses = 3

	// Code Sukses
	CodeSukses StatusCode = "00"
	// Code Data Tagihan Tidak Ditemukan
	CodeTidakDitemukan StatusCode = "10"
	// Code Data Tagihan Telah Lunas
	CodeLunas StatusCode = "13"
	// Code Data Tagihan Tidak Sesuai
	CodeTidakSesuai StatusCode = "14"
	// Code Data Reversal Tidak Ditemukan
	CodeReversalTidakDitemukan StatusCode = "34"
	// Code Data Reversal Telah Dibatalkan
	CodeReversalDibatalkan StatusCode = "36"
	// Code System Error
	CodeError StatusCode = "99"

	DescSukses         = "Success"
	DescLunas          = "Tagihan Lunas"
	DescTidakSesuai    = "Jumlah tagihan yang dibayarkan tidak sesuai"
	DescTidakDitemukan = "Tagihan Tidak Ditemukan"
	DescError          = "System Bank Error: "
)

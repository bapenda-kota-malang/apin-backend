package espt

type Status string

const (
	StatusBaru      Status = "baru"
	StatusDisetujui Status = "disetujui"
	StatusDitolak   Status = "ditolak"
)

type CreateInput interface {
	GetEspt() CreateDto
	CalculateTax(taxPercentage *float64)
	GetDetails() interface{}
	LenDetails() int
	ChangeDetails(newDetail interface{})
	ReplaceEsptId(id uint)
	ReplaceTarifPajakId(id uint)
}

type UpdateInput interface {
	GetEspt() UpdateDto
	CalculateTax(taxPercentage *float64)
	GetDetails() interface{}
	LenDetails() int
}

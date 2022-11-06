package espt

import "github.com/google/uuid"

type Status string

const (
	StatusBaru      Status = "baru"
	StatusDisetujui Status = "disetujui"
	StatusDitolak   Status = "ditolak"
)

type Input interface {
	GetEspt() interface{}
	CalculateTax(taxPercentage *float64)
	GetDetails() interface{}
	LenDetails() int
	ChangeDetails(newDetail interface{})
	ReplaceEsptId(id uuid.UUID)
	ReplaceTarifPajakId(id uint)
}

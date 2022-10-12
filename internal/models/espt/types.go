package espt

type Status string

const (
	StatusBaru      Status = "baru"
	StatusDisetujui Status = "disetujui"
	StatusDitolak   Status = "ditolak"
)

type CreateInput interface {
	GetEspt() CreateDto
	GetDetails() interface{}
	LenDetails() int
	ReplaceEsptId(id uint)
}

type UpdateInput interface {
	GetEspt() UpdateDto
	GetDetails() interface{}
	LenDetails() int
}

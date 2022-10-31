package spt

type SptStatus string

const (
	StatusBelumLunas          SptStatus = "00" //belum lunas
	StatusLunas               SptStatus = "11" //lunas penuh
	StatusKurangBayar         SptStatus = "12" //kurang bayar
	StatusKurangBayarAngsuran SptStatus = "13" //kurang bayar angsuran
	StatusLebihBayar          SptStatus = "14" //lebih bayar
	StatusSalahPenetapan      SptStatus = "21" //salah penetapan
	StatusDouble              SptStatus = "22" //double
	StatusRestitusiPenuh      SptStatus = "31" //restitusi penuh
	StatusRestitusiSebagian   SptStatus = "32" //restitusi sebagian
)

type CreateInput interface {
	GetSpt() CreateDto
	GetDetails() interface{}
	LenDetails() int
	ReplaceSptId(id uint)
}

type UpdateInput interface {
	GetSpt() CreateDto
	GetDetails() interface{}
	LenDetails() int
	ReplaceSptId(id uint)
}

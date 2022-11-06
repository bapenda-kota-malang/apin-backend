package spt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/google/uuid"
)

type SptStatus string
type JenisKetetapan string

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

	JenisKetetapanSkpd    JenisKetetapan = "skpd"
	JenisKetetapanSkpdkb  JenisKetetapan = "skpdkb"
	JenisKetetapanSkpdkbt JenisKetetapan = "skpdkbt"
)

type Input interface {
	GetSpt(baseUri string) interface{}
	ReplaceTarifPajakId(id uint)
	CalculateTax(taxPercentage *float64)
	ReplaceSptId(id uuid.UUID)
	GetDetails() interface{}
	LenDetails() int
	ChangeDetails(newDetail interface{})
	DuplicateEspt(esptDetail *espt.Espt) error
}

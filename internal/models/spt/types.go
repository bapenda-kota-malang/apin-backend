package spt

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/google/uuid"
)

type SptStatus string
type JenisKetetapan string
type StatusPenetapan uint8

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
	StatusPenyetoran          SptStatus = "40" //penyetoran

	JenisKetetapanSkpd    JenisKetetapan = "skpd"
	JenisKetetapanSkpdkb  JenisKetetapan = "skpdkb"
	JenisKetetapanSkpdkbt JenisKetetapan = "skpdkbt"

	StatusPenetapanBaru             StatusPenetapan = 0 // baru
	StatusPenetapanDisetujuiKasubid StatusPenetapan = 1 // DisetujuiKasubid
	StatusPenetapanDisetujuiKabid   StatusPenetapan = 2 // DisetujuiKabid
	StatusPenetapanDitolakKasubid   StatusPenetapan = 3 // DitolakKasubid
	StatusPenetapanDitolakKabid     StatusPenetapan = 4 // DitolakKabid

	TbpStatusFilterBaru       uint8 = 1 // baru
	TbpStatusFilterPembayaran uint8 = 2 // Pembayaran
	TbpStatusFilterPenyetoran uint8 = 3 // Penyetoran
	TbpStatusFilterLunas      uint8 = 4 // Lunas
	TbpStatusFilterJatuhTempo uint8 = 5 // Jatuh Tempo
	TbpStatusFilterPenetapan  uint8 = 6 // Penetapan
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
	SkpdkbDuplicate(sptDetail *Spt, skpdkb *SkpdkbExisting) error
	CalculateSkpdkb()
}

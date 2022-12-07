package objekpajakpbb

type JenisTransaksi string
type StatusWp int

const (
	JenisTransaksiPerekaman          JenisTransaksi = "1" // perekaman data op
	JenisTransaksiPemutakhiran       JenisTransaksi = "2" // pemutakhiran data op
	JenisTransaksiPenghapusan        JenisTransaksi = "3" // penghapusan op
	JenisTransaksiPenghapusanBersama JenisTransaksi = "4" // penghapusan status op bersamax

	StatusWpPemilik   StatusWp = 1
	StatusWpPenyewa   StatusWp = 2
	StatusWpPengelola StatusWp = 3
	StatusWpPemakai   StatusWp = 4
	StatusWpSengketa  StatusWp = 5
)

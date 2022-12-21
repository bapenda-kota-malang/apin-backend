package objekpajakpbb

type JenisTransaksi string
type StatusWp string
type JenisOp string

const (
	JenisTransaksiPerekaman          JenisTransaksi = "11" // perekaman data op
	JenisTransaksiPemutakhiran       JenisTransaksi = "12" // pemutakhiran data op
	JenisTransaksiPenghapusan        JenisTransaksi = "13" // penghapusan op
	JenisTransaksiPenghapusanBersama JenisTransaksi = "14" // penghapusan status op bersama

	StatusWpPemilik   StatusWp = "1" // pemilik
	StatusWpPenyewa   StatusWp = "2" // penyewa
	StatusWpPengelola StatusWp = "3" // pengelola
	StatusWpPemakai   StatusWp = "4" // pemakai
	StatusWpSengketa  StatusWp = "5" // sengketa

	JenisOpBiasa   JenisOp = "0" // objek biasa
	JenisOpBersama JenisOp = "9" // objek bersama
)

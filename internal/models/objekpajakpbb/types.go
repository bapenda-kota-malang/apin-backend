package objekpajakpbb

type JenisTransaksi string

const (
	JenisTransaksiPerekaman          JenisTransaksi = "1" // perekaman data op
	JenisTransaksiPemutakhiran       JenisTransaksi = "2" // pemutakhiran data op
	JenisTransaksiPenghapusan        JenisTransaksi = "3" // penghapusan op
	JenisTransaksiPenghapusanBersama JenisTransaksi = "4" // penghapusan status op bersama
)

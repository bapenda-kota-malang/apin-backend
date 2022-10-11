package npwpd

type Golongan int16
type Status int16 // StatusNpwdp
type JenisPajak string
type JalurRegistrasi int16

type StatusBL int16
type GolonganObjekPajak string

const (
	GolonganOrangPribadi Golongan = 1 //orang pribadi
	GolonganBadan        Golongan = 2 //badan

	StatusAktif          Status = 1 //aktif
	StatusTutupSementara Status = 2 //tutup sementara
	StatusTutup          Status = 3 //tutup

	JenisPajakSA JenisPajak = "SA" //self_assessment
	JenisPajakOA JenisPajak = "OA" //official_assessment

	StatusBaru StatusBL = 0 //baru
	StatusLama StatusBL = 1 //lama

	JalurRegistrasiOperator JalurRegistrasi = 1 //operator
	JalurRegistrasiMandiri  JalurRegistrasi = 2 //mandiri
)

package regobjekpajakbumi

type JenisBumi string

const (
	JenisBumiTanah         JenisBumi = "1" // tanah + bangunan
	JenisBumiKavling       JenisBumi = "2" // kavling siap bangun
	JenisBumiTanahKosong   JenisBumi = "3" // tanah kosong
	JenisBumiFasilitasUmum JenisBumi = "4" // fasilitas umum
)

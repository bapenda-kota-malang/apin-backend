package jpb7

type JenisHotel string
type JumlahBintang string

const (
	// jenis hotel
	JenisHotelNonResort JenisHotel = "1" // non-resort
	JenisHotelResort    JenisHotel = "2" // resort

	// jumlah bintang
	JumlahBintang5          JumlahBintang = "1" // 5
	JumlahBintang4          JumlahBintang = "2" // 4
	JumlahBintang3          JumlahBintang = "3" // 3
	JumlahBintang2          JumlahBintang = "4" // 2
	JumlahBintangNonBintang JumlahBintang = "5" // non-bintang
)

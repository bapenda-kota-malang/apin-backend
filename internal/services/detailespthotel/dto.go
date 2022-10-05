package detailespthotel

type CreateDto struct {
	GolonganKamar       string  `json:"golonganKamar" size:"20"`
	Tarif               float32 `json:"tarif"`
	JumlahKamar         uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku uint    `json:"jumlahKamarYangLaku"`
}

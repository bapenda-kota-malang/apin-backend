package detailesptd

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DetailEspt struct {
	Id      uint `json:"id" gorm:"primarykey"`
	Espt_Id uint `json:"espt_id"`
}

type DetailEsptAir struct {
	DetailEspt
	Peruntukan string  `json:"peruntukan" gorm:"size:100"`
	JenisAbt   string  `json:"jenisAbt" gorm:"20"`
	Pengenaan  float32 `json:"pengenaan"`
	gormhelper.DateModel
}

type DetailEsptHotel struct {
	DetailEspt
	GolonganKamar       string  `json:"golonganKamar" size:"20"`
	Tarif               float32 `json:"tarif"`
	JumlahKamar         uint    `json:"jumlahKamar"`
	JumlahKamarYangLaku uint    `json:"jumlahKamarYangLaku"`
	gormhelper.DateModel
}

type DetailEsptParkir struct {
	DetailEspt
	JenisKendaraan uint `json:"jenisKendaraan"`
	Kapasitas      uint `json:"kapasitas"`
	gormhelper.DateModel
}

type DetailEsptResto struct {
	DetailEspt
	JumlahMeja       uint    `json:"jumlahMeja"`
	JumlahKursi      uint    `json:"jumlahKursi"`
	TarifMinuman     float32 `json:"tarifMinuman"`
	TarifMakanan     float32 `json:"tarifMakanan"`
	JumlahPengunjung uint    `json:"jumlahPengunjung"`
	LdBt             uint    `json:"ldBt"`
	gormhelper.DateModel
}

type DetailEsptReklame struct {
	DetailEspt
	TarifReklame_Id uint    `json:"tarifReklame_id"`
	Jumlah          uint    `json:"jumlah"`
	Sisi            uint    `json:"sisi"`
	Panjang         float32 `json:"panjang"`
	Lebar           float32 `json:"lebar"`
	Diameter        float32 `json:"diameter"`
	Diskon          float32 `json:"diskon"`
	Lokasi          string  `json:"lokasi" gorm:"size:200"`
	TarifHari       float32 `json:"tarifHari"`
	TarifMinggu     float32 `json:"tarifMinggu"`
	TarifBulan      float32 `json:"tarifBulan"`
	TarifTahun      float32 `json:"tarifTahun"`
	gormhelper.DateModel
}

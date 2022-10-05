package detailesptreklame

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptReklame struct {
	detailesptair.DetailEspt
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

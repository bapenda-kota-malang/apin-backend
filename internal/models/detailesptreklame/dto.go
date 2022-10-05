package detailesptreklame

import "github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"

type CreateDetailReklameDto struct {
	detailesptair.CreateDto
	TarifReklame_Id uint    `json:"tarifReklame_id"  validate:"required"`
	Jumlah          uint    `json:"jumlah"  validate:"required"`
	Sisi            uint    `json:"sisi"  validate:"required"`
	Panjang         float32 `json:"panjang"  validate:"required"`
	Lebar           float32 `json:"lebar"  validate:"required"`
	Diameter        float32 `json:"diameter"  validate:"required"`
	Diskon          float32 `json:"diskon"  validate:"required"`
	Lokasi          string  `json:"lokasi"  validate:"required"`
	TarifHari       float32 `json:"tarifHari"  validate:"required"`
	TarifMinggu     float32 `json:"tarifMinggu"  validate:"required"`
	TarifBulan      float32 `json:"tarifBulan"  validate:"required"`
	TarifTahun      float32 `json:"tarifTahun"  validate:"required"`
}

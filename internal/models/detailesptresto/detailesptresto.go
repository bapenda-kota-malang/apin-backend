package detailesptresto

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptResto struct {
	detailesptair.DetailEspt
	JumlahMeja       uint     `json:"jumlahMeja"`
	JumlahKursi      uint     `json:"jumlahKursi"`
	TarifMinuman     *float32 `json:"tarifMinuman"`
	TarifMakanan     *float32 `json:"tarifMakanan"`
	JumlahPengunjung uint     `json:"jumlahPengunjung"`
	IdBt             *uint    `json:"idBt"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id          uint
	JumlahMeja       uint     `json:"jumlahMeja" validate:"min=1"`
	JumlahKursi      uint     `json:"jumlahKursi" validate:"min=1"`
	TarifMinuman     *float32 `json:"tarifMinuman"`
	TarifMakanan     *float32 `json:"tarifMakanan"`
	JumlahPengunjung uint     `json:"jumlahPengunjung" validate:"min=1"`
	IdBt             *uint    `json:"idBt" validate:"min=1"`
}

type UpdateDto struct {
	Id               uint     `json:"id"`
	JumlahMeja       *uint    `json:"jumlahMeja"`
	JumlahKursi      *uint    `json:"jumlahKursi"`
	TarifMinuman     *float32 `json:"tarifMinuman"`
	TarifMakanan     *float32 `json:"tarifMakanan"`
	JumlahPengunjung *uint    `json:"jumlahPengunjung"`
	IdBt             *uint    `json:"idBt"`
}

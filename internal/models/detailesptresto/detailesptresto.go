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
	JumlahMeja       uint     `json:"jumlahMeja"  validate:"required"`
	JumlahKursi      uint     `json:"jumlahKursi"  validate:"required"`
	TarifMinuman     *float32 `json:"tarifMinuman"  validate:"required"`
	TarifMakanan     *float32 `json:"tarifMakanan"  validate:"required"`
	JumlahPengunjung uint     `json:"jumlahPengunjung"  validate:"required"`
	IdBt             *uint    `json:"idBt"  validate:"required"`
}

type UpdateDto struct {
	Id               uint     `json:"id"`
	JumlahMeja       uint     `json:"jumlahMeja"`
	JumlahKursi      uint     `json:"jumlahKursi"`
	TarifMinuman     *float32 `json:"tarifMinuman"`
	TarifMakanan     *float32 `json:"tarifMakanan"`
	JumlahPengunjung uint     `json:"jumlahPengunjung"`
	IdBt             *uint    `json:"idBt"`
}

package detailesptresto

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptResto struct {
	detailesptair.DetailEspt
	JumlahMeja       uint    `json:"jumlahMeja"`
	JumlahKursi      uint    `json:"jumlahKursi"`
	TarifMinuman     float32 `json:"tarifMinuman"`
	TarifMakanan     float32 `json:"tarifMakanan"`
	JumlahPengunjung uint    `json:"jumlahPengunjung"`
	LdBt             uint    `json:"ldBt"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id          uint    `json:"espt_id"`
	JumlahMeja       uint    `json:"jumlahMeja"  validate:"required"`
	JumlahKursi      uint    `json:"jumlahKursi"  validate:"required"`
	TarifMinuman     float32 `json:"tarifMinuman"  validate:"required"`
	TarifMakanan     float32 `json:"tarifMakanan"  validate:"required"`
	JumlahPengunjung uint    `json:"jumlahPengunjung"  validate:"required"`
	LdBt             uint    `json:"ldBt"  validate:"required"`
}

type UpdateDto struct {
	Id               uint    `json:"id"`
	Espt_Id          uint    `json:"espt_id"`
	JumlahMeja       uint    `json:"jumlahMeja"`
	JumlahKursi      uint    `json:"jumlahKursi"`
	TarifMinuman     float32 `json:"tarifMinuman"`
	TarifMakanan     float32 `json:"tarifMakanan"`
	JumlahPengunjung uint    `json:"jumlahPengunjung"`
	LdBt             uint    `json:"ldBt"`
}

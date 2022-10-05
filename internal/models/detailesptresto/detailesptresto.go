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

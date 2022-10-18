package detailesptppjpln

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptPpjNonPln struct {
	detailesptair.DetailEspt
	JumlahPelanggan string `json:"jenisMesinPenggerak" gorm:""`
	JumlahRekening  string `json:"tahunMesin"`
	gormhelper.DateModel
}

package detailesptppjpln

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DetailEsptPpjPln struct {
	detailesptair.DetailEspt
	JenisPPJ_Id     uint `json:"jenisPpj_Id"`
	JumlahPelanggan uint `json:"jumlahPelanggan"`
	JumlahRekening  uint `json:"jumlahRekening"`
	gormhelper.DateModel
}

type CreateDto struct {
	Espt_Id         uint
	JenisPPJ_Id     uint `json:"jenisPpj_Id"`
	JumlahPelanggan uint `json:"jumlahPelanggan"`
	JumlahRekening  uint `json:"jumlahRekening"`
	JenisPPJ        *jenisppj.JenisPPJ
}

type UpdateDto struct {
	Id              uint  `json:"id"`
	JenisPPJ_Id     *uint `json:"jenisPpj_Id"`
	JumlahPelanggan *uint `json:"jumlahPelanggan"`
	JumlahRekening  *uint `json:"jumlahRekening"`
}

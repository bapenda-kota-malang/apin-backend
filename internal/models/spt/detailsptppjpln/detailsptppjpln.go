package detailsptppjpln

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptPpjPln struct {
	Id              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id          uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	JenisPPJ_Id     uint      `json:"jenisPpj_Id"`
	JumlahPelanggan uint      `json:"jumlahPelanggan"`
	JumlahRekening  uint      `json:"jumlahRekening"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id          uuid.UUID
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
	JenisPPJ        *jenisppj.JenisPPJ
}

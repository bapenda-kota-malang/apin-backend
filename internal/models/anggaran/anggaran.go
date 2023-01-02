package anggaran

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	skpd "github.com/bapenda-kota-malang/apin-backend/internal/models/satuankerja"
)

type Anggaran struct {
	Id          uint64             `json:"id" gorm:"primaryKey"`
	SatuanKerja_Id     *uint64            `json:"skpd_id"`
	SatuanKerja        *skpd.SatuanKerja  `json:"skpd,omitempty" gorm:"foreignKey:Skpd_Id"`
	Tahun       *uint64            `json:"tahun" gorm:"type:integer"`
	Rekening_Id *uint64            `json:"rekening_id"`
	Rekening    *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	PaguMurni   *float64           `json:"paguMurni" gorm:"type:decimal"`
	PaguPak     *float64           `json:"paguPak" gorm:"type:decimal"`
}

type CreateDto struct {
	SatuanKerja_Id     uint64  `json:"skpd_id" validate:"required"`
	Tahun       uint64  `json:"tahun" validate:"required"`
	Rekening_Id uint64  `json:"rekening_id" validate:"required"`
	PaguMurni   float64 `json:"paguMurni" validate:"required"`
	PaguPak     float64 `json:"paguPak" validate:"required"`
}

type UpdateDto struct {
	SatuanKerja_Id     uint64  `json:"skpd_id"`
	Tahun       uint64  `json:"tahun"`
	Rekening_Id uint64  `json:"rekening_id"`
	PaguMurni   float64 `json:"paguMurni"`
	PaguPak     float64 `json:"paguPak"`
}

type FilterDto struct {
	SatuanKerja_Id     *uint64  `json:"skpd_id"`
	Tahun       *uint64  `json:"tahun"`
	Rekening_Id *uint64  `json:"rekening_id"`
	PaguMurni   *float64 `json:"paguMurni"`
	PaguPak     *float64 `json:"paguPak"`
	Page        int      `json:"page"`
	PageSize    int      `json:"page_size"`
}

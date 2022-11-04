package detailsptresto

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptResto struct {
	Id               uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id           uuid.UUID `json:"spt_id" gorm:"type:uuid"`
	JumlahMeja       *uint     `json:"jumlahMeja"`
	JumlahKursi      *uint     `json:"jumlahKursi"`
	TarifMinuman     *float32  `json:"tarifMinuman" gorm:"type:decimal"`
	TarifMakanan     *float32  `json:"tarifMakanan" gorm:"type:decimal"`
	JumlahPengunjung *uint     `json:"jumlahPengunjung"`
	IdBt             *uint     `json:"idBt"`
	gormhelper.DateModel
}

type CreateDto struct {
	Spt_Id           uuid.UUID
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

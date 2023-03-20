package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiResto struct {
	Id                  uint      `json:"id" gorm:"primaryKey"`
	Potensiop_Id        uuid.UUID `json:"potensiop_id" gorm:"type:uuid"`
	JumlahMeja          int64     `json:"jumlahMeja"`
	JumlahKursi         int64     `json:"jumlahKursi"`
	RentangHargaMakanan string    `json:"rentangHargaMakanan"`
	RentangHargaMinuman string    `json:"rentangHargaMinuman"`
	gormhelper.DateModel
}

type CreateDtoDPResto struct {
	Potensiop_Id        uuid.UUID `json:"-"`
	JumlahMeja          int64     `json:"jumlahMeja"`
	JumlahKursi         int64     `json:"jumlahKursi"`
	RentangHargaMakanan string    `json:"rentangHargaMakanan"`
	RentangHargaMinuman string    `json:"rentangHargaMinuman"`
}

type CreateDtoResto struct {
	CreateDto
	DetailPajakDtos CreateDtoDPResto `json:"detailPajaks"`
}

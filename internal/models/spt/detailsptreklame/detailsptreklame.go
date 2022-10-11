package detailsptreklame

import (
	mtr "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifreklame"
)

type DetailSptReklame struct {
	Id              uint64            `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id          uint64            `json:"spt_id"`
	TarifReklame_Id *uint64           `json:"tarifReklame_id"`
	TarifReklame    *mtr.TarifReklame `json:"tarifreklame,omitempty" gorm:"foreignKey:TarifReklame_Id"`
	Jumlah          *uint64           `json:"jumlah"`
	Sisi            *uint64           `json:"sisi"`
	Panjang         *float64          `json:"panjang" gorm:"type:decimal"`
	Lebar           *float64          `json:"lebar" gorm:"type:decimal"`
	Diameter        *float64          `json:"diameter" gorm:"type:decimal"`
	Diskon          *float64          `json:"diskon" gorm:"type:decimal"`
	Lokasi          *string           `json:"lokasi" gorm:"type:varchar(200)"`
	TarifHari       *float64          `json:"tarifHari" gorm:"type:decimal"`
	TarifMinggu     *float64          `json:"tarifMinggu" gorm:"type:decimal"`
	TarifBulan      *float64          `json:"tarifBulan" gorm:"type:decimal"`
	TarifTahun      *float64          `json:"tarifTahun" gorm:"type:decimal"`
	JumlahRp        *float64          `json:"jumlahRp"`
	NomorPersil     *string           `json:"nomorPersil" gorm:"type:varchar(100)"`
}

type CreateDto struct {
	// Spt_Id          uint     `json:"spt_id" validate:"required"`
	TarifReklame_Id uint64   `json:"tarifReklame_id" validate:"required"`
	Jumlah          uint64   `json:"jumlah" validate:"required"`
	Sisi            uint64   `json:"sisi" validate:"required"`
	Panjang         float64  `json:"panjang" validate:"required"`
	Lebar           float64  `json:"lebar" validate:"required"`
	Diameter        *float64 `json:"diameter" validate:"required"`
	Diskon          float64  `json:"diskon" validate:"required"`
	Lokasi          string   `json:"lokasi" validate:"required"`
	TarifHari       float64  `json:"tarifHari" validate:"required"`
	TarifMinggu     float64  `json:"tarifMinggu" validate:"required"`
	TarifBulan      float64  `json:"tarifBulan" validate:"required"`
	TarifTahun      float64  `json:"tarifTahun" validate:"required"`
	JumlahRp        float64  `json:"jumlahRp" validate:"required"`
	NomorPersil     string   `json:"nomorPersil" validate:"required"`
}

type UpdateDto struct {
	Spt_Id          uint    `json:"spt_id"`
	TarifReklame_Id uint64  `json:"tarifReklame_id"`
	Jumlah          uint64  `json:"jumlah"`
	Sisi            uint64  `json:"sisi"`
	Panjang         float64 `json:"panjang"`
	Lebar           float64 `json:"lebar"`
	Diameter        float64 `json:"diameter"`
	Diskon          float64 `json:"diskon"`
	Lokasi          string  `json:"lokasi"`
	TarifHari       float64 `json:"tarifHari"`
	TarifMinggu     float64 `json:"tarifMinggu"`
	TarifBulan      float64 `json:"tarifBulan"`
	TarifTahun      float64 `json:"tarifTahun"`
	JumlahRp        float64 `json:"jumlahRp"`
	NomorPersil     string  `json:"nomorPersil"`
}

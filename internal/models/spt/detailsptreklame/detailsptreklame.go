package detailsptreklame

import (
	mtr "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifreklame"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailSptReklame struct {
	Id              uint64       `json:"id" gorm:"primaryKey;autoIncrement"`
	Spt_Id          uuid.UUID    `json:"spt_id" gorm:"type:uuid"`
	TarifReklame_Id *uint64      `json:"tarifReklame_id"`
	Jumlah          uint64       `json:"jumlah"`
	Sisi            *uint64      `json:"sisi"`
	Panjang         *float64     `json:"panjang" gorm:"type:decimal"`
	Lebar           *float64     `json:"lebar" gorm:"type:decimal"`
	Diameter        *float64     `json:"diameter" gorm:"type:decimal"`
	Diskon          *float64     `json:"diskon" gorm:"type:decimal"`
	Lokasi          *string      `json:"lokasi" gorm:"type:varchar(200)"`
	TarifHari       *float64     `json:"tarifHari" gorm:"type:decimal"`
	TarifMinggu     *float64     `json:"tarifMinggu" gorm:"type:decimal"`
	TarifBulan      *float64     `json:"tarifBulan" gorm:"type:decimal"`
	TarifTahun      *float64     `json:"tarifTahun" gorm:"type:decimal"`
	JumlahRp        float64      `json:"jumlahRp"`
	NomorPersil     *string      `json:"nomorPersil" gorm:"type:varchar(100)"`
	JenisDimensi    JenisDimensi `json:"jenisDimensi"`
	gormhelper.DateModel
	TarifReklame *mtr.TarifReklame `json:"tarifreklame,omitempty" gorm:"foreignKey:TarifReklame_Id"`
}

type CreateDto struct {
	Spt_Id          uuid.UUID
	TarifReklame_Id uint64       `json:"tarifReklame_Id" validate:"required"`
	Jumlah          uint64       `json:"jumlah" validate:"required"`
	Sisi            uint64       `json:"sisi"`
	Panjang         float64      `json:"panjang"`
	Lebar           float64      `json:"lebar"`
	Diameter        *float64     `json:"diameter"`
	Diskon          float64      `json:"diskon"`
	Lokasi          string       `json:"lokasi" validate:"required"`
	TarifHari       float64      `json:"tarifHari"`
	TarifMinggu     float64      `json:"tarifMinggu"`
	TarifBulan      float64      `json:"tarifBulan"`
	TarifTahun      float64      `json:"tarifTahun"`
	JumlahRp        float64      `json:"-"`
	NomorPersil     string       `json:"nomorPersil"`
	JenisDimensi    JenisDimensi `json:"jenisDimensi"`
	TarifReklame    *mtr.TarifReklame
}

type UpdateDto struct {
	Spt_Id          *uint    `json:"spt_id"`
	TarifReklame_Id *uint64  `json:"tarifReklame_id"`
	Jumlah          *uint64  `json:"jumlah"`
	Sisi            *uint64  `json:"sisi"`
	Panjang         *float64 `json:"panjang"`
	Lebar           *float64 `json:"lebar"`
	Diameter        *float64 `json:"diameter"`
	Diskon          *float64 `json:"diskon"`
	Lokasi          *string  `json:"lokasi"`
	TarifHari       *float64 `json:"tarifHari"`
	TarifMinggu     *float64 `json:"tarifMinggu"`
	TarifBulan      *float64 `json:"tarifBulan"`
	TarifTahun      *float64 `json:"tarifTahun"`
	JumlahRp        *float64 `json:"jumlahRp"`
	NomorPersil     *string  `json:"nomorPersil"`
}

package detailobjek

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Detail struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Potensiop_Id uint   `json:"potensiop_id"`
	JenisOp      string `json:"jenisOp"`
	JumlahOp     string `json:"jumlahOp"`
	TarifOp      string `json:"tarifOp"`
	UnitOp       string `json:"unitOp"`
	Notes        string `json:"notes"`
	gormhelper.DateModel
}

type DetailPotensiAirTanah struct {
	Detail
}

type DetailPotensiHiburan struct {
	Detail
}

type DetailPotensiHotel struct {
	Detail
}

type DetailPotensiParkir struct {
	Detail
}

type DetailPotensiPPJ struct {
	Detail
}

type DetailPotensiReklame struct {
	Detail
}

type DetailPotensiResto struct {
	Detail
}

type DetailPajakDto struct {
	Potensiop_Id uint   `json:"-"`
	JenisOp      string `json:"jenisOp" validate:"oneof='Resto' 'Reklame' 'PPJ' 'Parkir' 'Hotel' 'Hiburan' 'Air Tanah'"`
	JumlahOp     string `json:"jumlahOp"`
	TarifOp      string `json:"tarifOp"`
	UnitOp       string `json:"unitOp"`
	Notes        string `json:"notes"`
}

type UpdateDto struct {
	Id           *uint   `json:"id"`
	Potensiop_Id *uint   `json:"-"`
	JenisOp      string  `json:"jenisOp"`
	JumlahOp     *string `json:"jumlahOp"`
	TarifOp      *string `json:"tarifOp"`
	UnitOp       *string `json:"unitOp"`
	Notes        *string `json:"notes"`
}

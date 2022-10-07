package npwpd

import (
	"reflect"
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func insertDetailOp(objek string, data *[]npwpd.DetailOp, registerForm *npwpd.Npwpd) error {
	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailOpHotel    npwpd.DetailOpHotel
		detailOpResto    npwpd.DetailOpResto
		detailOpHiburan  npwpd.DetailOpHiburan
		detailOpReklame  npwpd.DetailOpReklame
		detailOpPpj      npwpd.DetailOpPpj
		detailOpParkir   npwpd.DetailOpParkir
		detailOpAirTanah npwpd.DetailOpAirTanah
	)

	mActions[`detailOpHotel`] = reflect.TypeOf(detailOpHotel)
	mActions[`detailOpResto`] = reflect.TypeOf(detailOpResto)
	mActions[`detailOpHiburan`] = reflect.TypeOf(detailOpHiburan)
	mActions[`detailOpReklame`] = reflect.TypeOf(detailOpReklame)
	mActions[`detailOpPpj`] = reflect.TypeOf(detailOpPpj)
	mActions[`detailOpParkir`] = reflect.TypeOf(detailOpParkir)
	mActions[`detailOpAirTanah`] = reflect.TypeOf(detailOpAirTanah)

	switch objek {
	case "01":
		model = reflect.Zero(mActions["detailOpHotel"]).Interface()
	case "02":
		model = reflect.Zero(mActions["detailOpResto"]).Interface()
	case "03":
		model = reflect.Zero(mActions["detailOpHiburan"]).Interface()
	case "04":
		model = reflect.Zero(mActions["detailOpReklame"]).Interface()
	case "05":
		model = reflect.Zero(mActions["detailOpPpj"]).Interface()
	case "06":
		model = reflect.Zero(mActions["detailOpParkir"]).Interface()
	case "07":
		model = reflect.Zero(mActions["detailOpAirTanah"]).Interface()
	}

	for _, dop := range *data {
		dop.JenisOp = registerForm.Rekening.Nama
		dop.Npwpd_Id = registerForm.Id

		m := make(map[string]interface{})
		m["Npwpd_Id"] = dop.Npwpd_Id
		m["JenisOp"] = dop.JenisOp
		m["JumlahOp"] = dop.JumlahOp
		m["TarifOp"] = dop.TarifOp
		m["UnitOp"] = dop.UnitOp
		m["Notes"] = dop.Notes

		err = a.DB.Model(&model).Create(&m).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func parseTime(input string) *time.Time {
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		return nil
	}
	return &t
}

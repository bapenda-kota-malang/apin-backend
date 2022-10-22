package npwpd

import (
	"reflect"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func insertDetailOp(objek string, data *[]npwpd.DetailObjekPajak, registerForm *npwpd.Npwpd) error {
	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailObjekPajakHotel    npwpd.DetailObjekPajakHotel
		detailObjekPajakResto    npwpd.DetailObjekPajakResto
		detailObjekPajakHiburan  npwpd.DetailObjekPajakHiburan
		detailObjekPajakReklame  npwpd.DetailObjekPajakReklame
		detailObjekPajakPpj      npwpd.DetailObjekPajakPpj
		detailObjekPajakParkir   npwpd.DetailObjekPajakParkir
		detailObjekPajakAirTanah npwpd.DetailObjekPajakAirTanah
	)

	mActions[`detailObjekPajakHotel`] = reflect.TypeOf(detailObjekPajakHotel)
	mActions[`detailObjekPajakResto`] = reflect.TypeOf(detailObjekPajakResto)
	mActions[`detailObjekPajakHiburan`] = reflect.TypeOf(detailObjekPajakHiburan)
	mActions[`detailObjekPajakReklame`] = reflect.TypeOf(detailObjekPajakReklame)
	mActions[`detailObjekPajakPpj`] = reflect.TypeOf(detailObjekPajakPpj)
	mActions[`detailObjekPajakParkir`] = reflect.TypeOf(detailObjekPajakParkir)
	mActions[`detailObjekPajakAirTanah`] = reflect.TypeOf(detailObjekPajakAirTanah)

	switch objek {
	case "01":
		model = reflect.Zero(mActions["detailObjekPajakHotel"]).Interface()
	case "02":
		model = reflect.Zero(mActions["detailObjekPajakResto"]).Interface()
	case "03":
		model = reflect.Zero(mActions["detailObjekPajakHiburan"]).Interface()
	case "04":
		model = reflect.Zero(mActions["detailObjekPajakReklame"]).Interface()
	case "05":
		model = reflect.Zero(mActions["detailObjekPajakPpj"]).Interface()
	case "06":
		model = reflect.Zero(mActions["detailObjekPajakParkir"]).Interface()
	case "07":
		model = reflect.Zero(mActions["detailObjekPajakAirTanah"]).Interface()
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

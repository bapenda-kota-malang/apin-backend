package registrasinpwpd

import (
	"reflect"

	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func insertDetailOp(objek string, data *[]rn.DetailRegObjekPajak, registerForm *rn.RegistrasiNpwpd) error {
	if data == nil {
		return nil
	}

	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailRegObjekPajakHotel    rn.DetailRegObjekPajakHotel
		detailRegObjekPajakResto    rn.DetailRegObjekPajakResto
		detailRegObjekPajakHiburan  rn.DetailRegObjekPajakHiburan
		detailRegObjekPajakReklame  rn.DetailRegObjekPajakReklame
		detailRegObjekPajakPpj      rn.DetailRegObjekPajakPpj
		detailRegObjekPajakParkir   rn.DetailRegObjekPajakParkir
		detailRegObjekPajakAirTanah rn.DetailRegObjekPajakAirTanah
	)

	mActions[`detailRegObjekPajakHotel`] = reflect.TypeOf(detailRegObjekPajakHotel)
	mActions[`detailRegObjekPajakResto`] = reflect.TypeOf(detailRegObjekPajakResto)
	mActions[`detailRegObjekPajakHiburan`] = reflect.TypeOf(detailRegObjekPajakHiburan)
	mActions[`detailRegObjekPajakReklame`] = reflect.TypeOf(detailRegObjekPajakReklame)
	mActions[`detailRegObjekPajakPpj`] = reflect.TypeOf(detailRegObjekPajakPpj)
	mActions[`detailRegObjekPajakParkir`] = reflect.TypeOf(detailRegObjekPajakParkir)
	mActions[`detailRegObjekPajakAirTanah`] = reflect.TypeOf(detailRegObjekPajakAirTanah)

	switch objek {
	case "01":
		model = reflect.Zero(mActions["detailRegObjekPajakHotel"]).Interface()
	case "02":
		model = reflect.Zero(mActions["detailRegObjekPajakResto"]).Interface()
	case "03":
		model = reflect.Zero(mActions["detailRegObjekPajakHiburan"]).Interface()
	case "04":
		model = reflect.Zero(mActions["detailRegObjekPajakReklame"]).Interface()
	case "05":
		model = reflect.Zero(mActions["detailRegObjekPajakPpj"]).Interface()
	case "06":
		model = reflect.Zero(mActions["detailRegObjekPajakParkir"]).Interface()
	case "07":
		model = reflect.Zero(mActions["detailRegObjekPajakAirTanah"]).Interface()
	}

	for _, dop := range *data {
		dop.JenisOp = registerForm.Rekening.Nama
		dop.RegistrasiNpwpd_Id = registerForm.Id

		m := make(map[string]interface{})
		m["RegistrasiNpwpd_Id"] = dop.RegistrasiNpwpd_Id
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

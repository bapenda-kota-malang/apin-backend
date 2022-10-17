package registrasinpwpd

import (
	"reflect"

	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
)

func insertDetailOp(objek string, data *[]rn.DetailRegOp, registerForm *rn.RegistrasiNpwpd) error {
	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailRegOpHotel    rn.DetailRegOpHotel
		detailRegOpResto    rn.DetailRegOpResto
		detailRegOpHiburan  rn.DetailRegOpHiburan
		detailRegOpReklame  rn.DetailRegOpReklame
		detailRegOpPpj      rn.DetailRegOpPpj
		detailRegOpParkir   rn.DetailRegOpParkir
		detailRegOpAirTanah rn.DetailRegOpAirTanah
	)

	mActions[`detailRegOpHotel`] = reflect.TypeOf(detailRegOpHotel)
	mActions[`detailRegOpResto`] = reflect.TypeOf(detailRegOpResto)
	mActions[`detailRegOpHiburan`] = reflect.TypeOf(detailRegOpHiburan)
	mActions[`detailRegOpReklame`] = reflect.TypeOf(detailRegOpReklame)
	mActions[`detailRegOpPpj`] = reflect.TypeOf(detailRegOpPpj)
	mActions[`detailRegOpParkir`] = reflect.TypeOf(detailRegOpParkir)
	mActions[`detailRegOpAirTanah`] = reflect.TypeOf(detailRegOpAirTanah)

	switch objek {
	case "01":
		model = reflect.Zero(mActions["detailRegOpHotel"]).Interface()
	case "02":
		model = reflect.Zero(mActions["detailRegOpResto"]).Interface()
	case "03":
		model = reflect.Zero(mActions["detailRegOpHiburan"]).Interface()
	case "04":
		model = reflect.Zero(mActions["detailRegOpReklame"]).Interface()
	case "05":
		model = reflect.Zero(mActions["detailRegOpPpj"]).Interface()
	case "06":
		model = reflect.Zero(mActions["detailRegOpParkir"]).Interface()
	case "07":
		model = reflect.Zero(mActions["detailRegOpAirTanah"]).Interface()
	}

	for _, dop := range *data {
		dop.JenisOp = registerForm.Rekening.Nama
		dop.RegistrasiNpwpd_Id = registerForm.Id
		// dop.Pendaftaran_Npwpd = registerForm.Npwpd

		m := make(map[string]interface{})
		m["RegistrasiNpwpd_Id"] = dop.RegistrasiNpwpd_Id
		// m["Pendaftaran_Npwpd"] = dop.Pendaftaran_Npwpd
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

package registrasinpwpd

import (
	"errors"
	"reflect"

	nm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

func filePreProcess(b64String string, userId uint, docsName string) (fileName, path, extFile string, err error) {
	extFile, err = base64helper.GetExtensionBase64(b64String)
	if err != nil {
		return
	}
	path = sh.GetResourcesPath()
	switch extFile {
	case "png", "jpeg":
		path = sh.GetImgPath()
	default:
		err = errors.New("file bukan gambar")
		return
	}
	id, err := sh.GetUuidv4()
	if err != nil {
		err = errors.New("gagal generate uuid")
		return
	}
	fileName = sh.GenerateFilename(docsName, id, userId, extFile)
	return
}

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

func generateNomor() int {
	var tmp int
	var tmpNpwpd nm.Npwpd
	nomor := a.DB.Last(&tmpNpwpd)
	if nomor.Error != nil {
		return 1
	} else {
		tmp = tmpNpwpd.Nomor
		tmp++
	}
	return tmp
}

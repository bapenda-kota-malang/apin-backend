package npwpd

import (
	"reflect"
	"strconv"

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
	case "07":
		model = reflect.Zero(mActions["detailObjekPajakParkir"]).Interface()
	case "08":
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

func generateNomor(isNomorRegistrasiAuto bool, nomor int) int {
	if isNomorRegistrasiAuto {
		var tmp int
		var tmpNpwpd npwpd.Npwpd
		nomor := a.DB.Last(&tmpNpwpd)
		if nomor.Error != nil {
			return 1
		} else {
			tmp = tmpNpwpd.Nomor
			tmp++
		}
		return tmp
	}
	return nomor
}

func GenerateNpwpd(nomor int, kecamatan_id uint64, kodeJenisUsaha string) string {
	kecamatanIdString := strconv.Itoa(int(kecamatan_id))
	if kodeJenisUsaha == "" {
		kodeJenisUsaha = "xxx"
	}
	tmpNomorString := strconv.Itoa(nomor)
	if len(tmpNomorString) == 1 {
		tmpNomorString = "000" + tmpNomorString
	} else if len(tmpNomorString) == 2 {
		tmpNomorString = "00" + tmpNomorString
	} else if len(tmpNomorString) == 3 {
		tmpNomorString = "0" + tmpNomorString
	}
	npwpdString := tmpNomorString + "." + kecamatanIdString + "." + kodeJenisUsaha
	return npwpdString
}

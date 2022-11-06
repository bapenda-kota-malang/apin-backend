package regnpwpd

import (
	"errors"
	"reflect"

	nm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
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

func insertDetailOp(objek string, data *[]rn.DetailRegObjekPajakCreateDto, registerForm *rn.RegNpwpd, tx *gorm.DB) error {
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
	case "07":
		model = reflect.Zero(mActions["detailRegObjekPajakParkir"]).Interface()
	case "08":
		model = reflect.Zero(mActions["detailRegObjekPajakAirTanah"]).Interface()
	}

	for _, dop := range *data {
		dop.JenisOp = registerForm.Rekening.Nama
		dop.RegNpwpd_Id = registerForm.Id

		m := make(map[string]interface{})
		m["RegNpwpd_Id"] = dop.RegNpwpd_Id
		m["JenisOp"] = dop.JenisOp
		m["JumlahOp"] = dop.JumlahOp
		m["TarifOp"] = dop.TarifOp
		m["UnitOp"] = dop.UnitOp
		m["Notes"] = dop.Notes

		err = tx.Model(&model).Create(&m).Error
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

func updateDetailObjekPajak(input []rn.DetailRegObjekPajak, regNpwpd_Id uint64, rekeningObjek, rekeningNama string, tx *gorm.DB) error {
	switch rekeningObjek {
	case "01":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakHotel
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail reg objek pajak hotel tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak hotel")
			}
		}
	case "02":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakResto
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data reg objek pajak resto tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak resto")
			}
		}
	case "03":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakHiburan
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data reg objek pajak hiburan tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak hiburan")
			}
		}
	case "04":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakReklame
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data reg objek pajak reklame tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak reklame")
			}
		}
	case "05":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakPpj
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data reg objek pajak ppj tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak ppj")
			}
		}
	case "07":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakParkir
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data reg objek pajak parkir tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak parkir")
			}
		}
	case "08":
		for _, v := range input {
			var dataDetail *rn.DetailRegObjekPajakAirTanah
			result := tx.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data reg objek pajak air tanah tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return errors.New("gagal mengambil data payload detail reg objek pajak")
			}
			dataDetail.RegNpwpd_Id = regNpwpd_Id
			dataDetail.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetail); result.Error != nil {
				return errors.New("gagal menyimpan data detail reg objek pajak air tanah")
			}
		}
	}
	return nil
}

func deleteDetailObjekPajak(regNpwpd_Id uint64, rekeningObjek string, tx *gorm.DB) error {
	switch rekeningObjek {
	case "01":
		var DataOp []*rn.DetailRegObjekPajakHotel
		result := tx.Where(rn.DetailRegObjekPajakHotel{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak hotel tidak dapat ditemukan")
		}

		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}

	case "02":
		var DataOp []*rn.DetailRegObjekPajakResto
		result := tx.Where(rn.DetailRegObjekPajakResto{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak resto tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}
	case "03":
		var DataOp []*rn.DetailRegObjekPajakHiburan
		result := tx.Where(rn.DetailRegObjekPajakHiburan{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak hiburan tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}
	case "04":
		var DataOp []*rn.DetailRegObjekPajakReklame
		result := tx.Where(rn.DetailRegObjekPajakReklame{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak reklame tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}
	case "05":
		var DataOp []*rn.DetailRegObjekPajakPpj
		result := tx.Where(rn.DetailRegObjekPajakPpj{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak ppj tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}
	case "07":
		var DataOp []*rn.DetailRegObjekPajakParkir
		result := tx.Where(rn.DetailRegObjekPajakParkir{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak parkir tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}
	case "08":
		var DataOp []*rn.DetailRegObjekPajakAirTanah
		result := tx.Where(rn.DetailRegObjekPajakAirTanah{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak air tanah tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus detail reg objek pajak hotel")
			}
		}
	}
	return nil
}

func verifyDetailRegObjekPajak(regNpwpd_Id, npwpd_Id uint64, rekeningObjek string, tx *gorm.DB) error {

	switch rekeningObjek {
	case "01":
		var dataReg []*rn.DetailRegObjekPajakHotel
		result := tx.Where(rn.DetailRegObjekPajakHotel{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak hotel tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakHotel
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak hotel")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak hotel")
			}
		}

	case "02":
		var dataReg []*rn.DetailRegObjekPajakResto
		result := tx.Where(rn.DetailRegObjekPajakResto{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak resto tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakResto
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak resto")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak resto")
			}
		}
	case "03":
		var dataReg []*rn.DetailRegObjekPajakHiburan
		result := tx.Where(rn.DetailRegObjekPajakHiburan{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak hiburan tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakHiburan
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak hiburan")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak hiburan")
			}
		}
	case "04":
		var dataReg []*rn.DetailRegObjekPajakReklame
		result := tx.Where(rn.DetailRegObjekPajakReklame{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak reklame tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakReklame
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak reklame")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak reklame")
			}
		}
	case "05":
		var dataReg []*rn.DetailRegObjekPajakPpj
		result := tx.Where(rn.DetailRegObjekPajakPpj{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak ppj tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakPpj
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak ppj")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak ppj")
			}
		}
	case "07":
		var dataReg []*rn.DetailRegObjekPajakParkir
		result := tx.Where(rn.DetailRegObjekPajakParkir{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak parkir tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakParkir
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak parkir")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak parkir")
			}
		}
	case "08":
		var dataReg []*rn.DetailRegObjekPajakAirTanah
		result := tx.Where(rn.DetailRegObjekPajakAirTanah{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak air tanah tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakAirTanah
			if err := sc.Copy(&dataOp, v); err != nil {
				return errors.New("tidak dapat menyalin data detail reg objek pajak air tanah")
			}

			dataOp.Npwpd_Id = npwpd_Id
			dataOp.Id = 0
			err := tx.Create(&dataOp).Error
			if err != nil {
				return errors.New("tidak dapat membuat data detail objek pajak air tanah")
			}
		}
	}
	return nil
}

func checkDataDetailObjekPajak(regNpwpd_Id uint64, rekeningObjek string, tx *gorm.DB) error {
	switch rekeningObjek {
	case "01":
		var dataReg []*rn.DetailRegObjekPajakHotel
		result := tx.Where(rn.DetailRegObjekPajakHotel{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak hotel tidak dapat ditemukan")
		}

	case "02":
		var dataReg []*rn.DetailRegObjekPajakResto
		result := tx.Where(rn.DetailRegObjekPajakResto{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak resto tidak dapat ditemukan")
		}

	case "03":
		var dataReg []*rn.DetailRegObjekPajakHiburan
		result := tx.Where(rn.DetailRegObjekPajakHiburan{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak hiburan tidak dapat ditemukan")
		}

	case "04":
		var dataReg []*rn.DetailRegObjekPajakReklame
		result := tx.Where(rn.DetailRegObjekPajakReklame{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak reklame tidak dapat ditemukan")
		}

	case "05":
		var dataReg []*rn.DetailRegObjekPajakPpj
		result := tx.Where(rn.DetailRegObjekPajakPpj{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak ppj tidak dapat ditemukan")
		}

	case "07":
		var dataReg []*rn.DetailRegObjekPajakParkir
		result := tx.Where(rn.DetailRegObjekPajakParkir{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak parkir tidak dapat ditemukan")
		}

	case "08":
		var dataReg []*rn.DetailRegObjekPajakAirTanah
		result := tx.Where(rn.DetailRegObjekPajakAirTanah{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: regNpwpd_Id}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return errors.New("data detail reg objek pajak air tanah tidak dapat ditemukan")
		}

	}
	return nil
}

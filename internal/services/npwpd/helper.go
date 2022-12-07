package npwpd

import (
	"errors"
	"reflect"
	"strconv"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
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

func insertDetailObjekPajak(objek string, rekeningNama string, data []m.DetailObjekPajakCreateDto, dataNpwpd *m.Npwpd, tx *gorm.DB) error {
	if data == nil {
		return nil
	}
	var (
		err      error
		mActions map[string]reflect.Type = make(map[string]reflect.Type)
		model    interface{}

		detailObjekPajakHotel    m.DetailObjekPajakHotel
		detailObjekPajakResto    m.DetailObjekPajakResto
		detailObjekPajakHiburan  m.DetailObjekPajakHiburan
		detailObjekPajakReklame  m.DetailObjekPajakReklame
		detailObjekPajakPpj      m.DetailObjekPajakPpj
		detailObjekPajakParkir   m.DetailObjekPajakParkir
		detailObjekPajakAirTanah m.DetailObjekPajakAirTanah
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

	for _, dop := range data {
		dop.JenisOp = &rekeningNama
		dop.Npwpd_Id = dataNpwpd.Id

		m := make(map[string]interface{})
		m["Npwpd_Id"] = dop.Npwpd_Id
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

func generateNomor(isNomorRegistrasiAuto bool, nomor int) int {
	if isNomorRegistrasiAuto {
		var tmp int
		var tmpNpwpd m.Npwpd
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

func updateDetailObjekPajak(input []m.DetailObjekPajakUpdateDto, npwpd_id uint64, rekeningObjek, rekeningNama string, tx *gorm.DB) error {
	switch rekeningObjek {
	case "01":
		var dataUpdate []m.DetailObjekPajakHotel
		var dataDelete []m.DetailObjekPajakHotel
		var tmpDataUpdate m.DetailObjekPajakHotel
		var tmpDataDelete m.DetailObjekPajakHotel
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak hotel: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak hotel: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakHotel
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak hotel tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak hotel: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak hotel: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakHotel
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak hotel: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak hotel: " + result.Error.Error())
			}
		}
	case "02":
		var dataUpdate []m.DetailObjekPajakResto
		var dataDelete []m.DetailObjekPajakResto
		var tmpDataUpdate m.DetailObjekPajakResto
		var tmpDataDelete m.DetailObjekPajakResto
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak resto: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak resto: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakResto
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak resto tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak resto: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak resto: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakResto
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak resto: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak resto: " + result.Error.Error())
			}
		}
	case "03":
		var dataUpdate []m.DetailObjekPajakHiburan
		var dataDelete []m.DetailObjekPajakHiburan
		var tmpDataUpdate m.DetailObjekPajakHiburan
		var tmpDataDelete m.DetailObjekPajakHiburan
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak hiburan: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak hiburan: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakHiburan
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak hiburan tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak hiburan: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak hiburan: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakHiburan
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak hiburan: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak hiburan: " + result.Error.Error())
			}
		}
	case "04":
		var dataUpdate []m.DetailObjekPajakReklame
		var dataDelete []m.DetailObjekPajakReklame
		var tmpDataUpdate m.DetailObjekPajakReklame
		var tmpDataDelete m.DetailObjekPajakReklame
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak reklame: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak reklame: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakReklame
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak reklame tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak reklame: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak reklame: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakReklame
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak reklame: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak reklame: " + result.Error.Error())
			}
		}
	case "05":
		var dataUpdate []m.DetailObjekPajakPpj
		var dataDelete []m.DetailObjekPajakPpj
		var tmpDataUpdate m.DetailObjekPajakPpj
		var tmpDataDelete m.DetailObjekPajakPpj
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak ppj: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak ppj: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakPpj
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak ppj tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak ppj: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak ppj: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakPpj
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak ppj: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak ppj: " + result.Error.Error())
			}
		}
	case "07":
		var dataUpdate []m.DetailObjekPajakParkir
		var dataDelete []m.DetailObjekPajakParkir
		var tmpDataUpdate m.DetailObjekPajakParkir
		var tmpDataDelete m.DetailObjekPajakParkir
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak parkir: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak parkir: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakParkir
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak parkir tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak parkir: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak parkir: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakParkir
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak parkir: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak parkir: " + result.Error.Error())
			}
		}
	case "08":
		var dataUpdate []m.DetailObjekPajakAirTanah
		var dataDelete []m.DetailObjekPajakAirTanah
		var tmpDataUpdate m.DetailObjekPajakAirTanah
		var tmpDataDelete m.DetailObjekPajakAirTanah
		for _, v := range input {
			if v.Id != 0 && v.IsDeleted {
				if err := sc.Copy(&tmpDataDelete, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak air tanah: " + err.Error())
				}
				dataDelete = append(dataDelete, tmpDataDelete)
			} else {
				if err := sc.Copy(&tmpDataUpdate, &v); err != nil {
					return errors.New("gagal mengambil data payload detail objek pajak air tanah: " + err.Error())
				}
				dataUpdate = append(dataUpdate, tmpDataUpdate)
			}

		}

		// update data
		for _, v := range dataUpdate {
			var dataDetailUpdate *m.DetailObjekPajakAirTanah
			result := tx.First(&dataDetailUpdate, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("data detail objek pajak air tanah tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetailUpdate, &v); err != nil {
				return errors.New("gagal mengambil data payload detail objek pajak air tanah: " + err.Error())
			}
			// set directur value to null if golongan orang pribadi
			dataDetailUpdate.Npwpd_Id = npwpd_id
			dataDetailUpdate.JenisOp = &rekeningNama
			if result := tx.Save(&dataDetailUpdate); result.Error != nil {
				return errors.New("gagal menyimpan data detail objek pajak air tanah: " + result.Error.Error())
			}
		}
		for k := range dataUpdate {
			dataUpdate[k].Npwpd_Id = npwpd_id
		}

		// delete data
		for _, v := range dataDelete {
			var dataDetailDelete *m.DetailObjekPajakAirTanah
			result := tx.First(&dataDetailDelete, v.Id)
			if result.RowsAffected == 0 {
				return errors.New("gagal mengambil data payload detail objek pajak air tanah: " + result.Error.Error())
			}

			result = tx.Delete(&dataDetailDelete)
			if result.RowsAffected == 0 {
				return errors.New("gagal menghapus data detail objek pajak air tanah: " + result.Error.Error())
			}
		}
	}
	return nil
}

func deleteDetailObjekPajak(npwpd_Id uint64, rekeningObjek string, tx *gorm.DB) error {
	switch rekeningObjek {
	case "01":
		var DataOp []*m.DetailObjekPajakHotel
		result := tx.Where(m.DetailObjekPajakHotel{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak hotel tidak dapat ditemukan")
		}

		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}

	case "02":
		var DataOp []*m.DetailObjekPajakResto
		result := tx.Where(m.DetailObjekPajakResto{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak resto tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}
	case "03":
		var DataOp []*m.DetailObjekPajakHiburan
		result := tx.Where(m.DetailObjekPajakHiburan{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak hiburan tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}
	case "04":
		var DataOp []*m.DetailObjekPajakReklame
		result := tx.Where(m.DetailObjekPajakReklame{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak reklame tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}
	case "05":
		var DataOp []*m.DetailObjekPajakPpj
		result := tx.Where(m.DetailObjekPajakPpj{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak ppj tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}
	case "07":
		var DataOp []*m.DetailObjekPajakParkir
		result := tx.Where(m.DetailObjekPajakParkir{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak parkir tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}
	case "08":
		var DataOp []*m.DetailObjekPajakAirTanah
		result := tx.Where(m.DetailObjekPajakAirTanah{DetailObjekPajak: m.DetailObjekPajak{Npwpd_Id: npwpd_Id}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return errors.New("data detail objek pajak air tanah tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = tx.Delete(&v)
			if result.RowsAffected == 0 {
				return errors.New("tidak dapat menghapus data objek pajak air tanah")
			}
		}
	}
	return nil
}

package detailobjek

import (
	"errors"
	"reflect"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailobjek"
)

const source = "detailobjek"

func Create(input []m.DetailPajakDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var dPAirTanahs []m.DetailPotensiAirTanah
	var dPHiburans []m.DetailPotensiHiburan
	var dPHotels []m.DetailPotensiHotel
	var dPParkirs []m.DetailPotensiParkir
	var dPPpjs []m.DetailPotensiPPJ
	var dPReklames []m.DetailPotensiReklame
	var dPRestos []m.DetailPotensiResto
	var listTransactionData []interface{}

	for _, v := range input {
		switch v.JenisOp {
		case "airTanah":
			var tmp m.DetailPotensiAirTanah
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPAirTanahs = append(dPAirTanahs, tmp)
		case "hiburan":
			var tmp m.DetailPotensiHiburan
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPHiburans = append(dPHiburans, tmp)
		case "hotel":
			var tmp m.DetailPotensiHotel
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPHotels = append(dPHotels, tmp)
		case "parkir":
			var tmp m.DetailPotensiParkir
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPParkirs = append(dPParkirs, tmp)
		case "ppj":
			var tmp m.DetailPotensiPPJ
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPPpjs = append(dPPpjs, tmp)
		case "reklame":
			var tmp m.DetailPotensiReklame
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPReklames = append(dPReklames, tmp)
		case "resto":
			var tmp m.DetailPotensiResto
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", input)
			}
			dPRestos = append(dPRestos, tmp)
		default:
			return sh.SetError("request", "create-data", source, "failed", "jenis tidak diketahui", input)
		}
	}

	// filter data to listTransactionData before execute db
	tmp := []interface{}{&dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
	for _, x := range tmp {
		xVal := reflect.ValueOf(x)
		if xVal.Kind() == reflect.Pointer {
			xVal = xVal.Elem()
		}
		if xVal.Kind() == reflect.Slice && xVal.IsNil() {
			continue
		} else if xVal.Kind() == reflect.Struct {
			newEmpty := reflect.New(reflect.TypeOf(x).Elem()).Elem()
			if xVal.Interface() == newEmpty.Interface() {
				continue
			}
		}
		listTransactionData = append(listTransactionData, x)
	}

	// save data to db after create listTransactionData
	err := tx.Transaction(func(tx2 *gorm.DB) error {
		for _, x := range listTransactionData {
			xVal := reflect.ValueOf(x)
			if xVal.Kind() == reflect.Pointer {
				xVal = xVal.Elem()
			}
			if result := tx2.Save(xVal.Interface()); result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: listTransactionData}, nil
}

func Update(potensiOp_Id int, input []m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var dPAirTanahs []m.DetailPotensiAirTanah
	var dPHiburans []m.DetailPotensiHiburan
	var dPHotels []m.DetailPotensiHotel
	var dPParkirs []m.DetailPotensiParkir
	var dPPpjs []m.DetailPotensiPPJ
	var dPReklames []m.DetailPotensiReklame
	var dPRestos []m.DetailPotensiResto
	var listTransactionData []interface{}
	var listDeleteData []interface{}

	for _, v := range input {
		switch v.JenisOp {
		case "airTanah":
			var tmp m.DetailPotensiAirTanah
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPAirTanahs = append(dPAirTanahs, tmp)
			}
		case "hiburan":
			var tmp m.DetailPotensiHiburan
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPHiburans = append(dPHiburans, tmp)
			}
		case "hotel":
			var tmp m.DetailPotensiHotel
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPHotels = append(dPHotels, tmp)
			}
		case "parkir":
			var tmp m.DetailPotensiParkir
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPParkirs = append(dPParkirs, tmp)
			}
		case "ppj":
			var tmp m.DetailPotensiPPJ
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPPpjs = append(dPPpjs, tmp)
			}
		case "reklame":
			var tmp m.DetailPotensiReklame
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPReklames = append(dPReklames, tmp)
			}
		case "resto":
			var tmp m.DetailPotensiResto
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
				if tmp.Potensiop_Id != uint(potensiOp_Id) {
					return sh.SetError("request", "update-data", source, "failed", "tidak bisa mengubah data ini", v)
				}
			} else {
				tmp.Potensiop_Id = uint(potensiOp_Id)
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
			}
			if v.Delete != nil && *v.Delete {
				listDeleteData = append(listDeleteData, tmp)
			} else {
				dPRestos = append(dPRestos, tmp)
			}
		default:
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", v)
		}
	}

	// filter data to listTransactionData before execute db
	tmp := []interface{}{&dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
	for _, x := range tmp {
		xVal := reflect.ValueOf(x)
		if xVal.Kind() == reflect.Pointer {
			xVal = xVal.Elem()
		}
		if xVal.Kind() == reflect.Slice && xVal.IsNil() {
			continue
		}
		listTransactionData = append(listTransactionData, x)
	}

	// save data to db after create listTransactionData
	err := tx.Transaction(func(tx2 *gorm.DB) error {
		for _, x := range listTransactionData {
			if result := tx2.Save(x); result.Error != nil {
				return result.Error
			}
		}

		for _, x := range listDeleteData {
			if result := tx2.Delete(x); result.RowsAffected == 0 {
				return errors.New("data tidak terhapus")
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return rp.OKSimple{Data: listTransactionData}, nil
}

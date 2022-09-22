package potensiop

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	registration "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "potensiop"

func GetList(input m.FilterDto) (any, error) {
	var data []m.PotensiOp
	var count int64
	a.DB.Model(&m.PotensiOp{}).Count(&count)

	var pagination gh.Pagination
	result := a.DB.Scopes(gh.Filter(input, &pagination)).Preload("DetailPotensiOp").Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount": strconv.Itoa(int(count)),
			// "currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":     strconv.Itoa(pagination.Page),
			"pageSize": strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (any, error) {
	var data *m.PotensiOp

	result := a.DB.Preload("DetailPotensiOp").Preload("PotensiPemilikWp").Preload("PotensiNarahubung").First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Create(input m.CreateDto) (any, error) {
	var dataPotensiOp m.PotensiOp
	var dataDetailPotensiOp m.DetailPotensiOp
	var dataPemilikWp m.PotensiPemilikWp
	var dataPemilikNarahub m.PotensiNarahubung
	var dPAirTanahs []m.DetailPotensiAirTanah
	var dPHiburans []m.DetailPotensiHiburan
	var dPHotels []m.DetailPotensiHotel
	var dPParkirs []m.DetailPotensiParkir
	var dPPpjs []m.DetailPotensiPPJ
	var dPReklames []m.DetailPotensiReklame
	var dPRestos []m.DetailPotensiResto
	// listTransactionData := []interface{}{&dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&dataPotensiOp, &input.CreatePotensiOp); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
	}
	if err := sc.Copy(&dataDetailPotensiOp, &input.CreateDetailPotensiOp); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataDetailPotensiOp)
	}
	if err := sc.Copy(&dataPemilikWp, &input.CreatePotensiPemilikWp); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	}
	if input.CreatePotensiNarahubung != nil {
		if err := sc.Copy(&dataPemilikNarahub, &input.CreatePotensiNarahubung); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikNarahub)
		}
	}
	for _, v := range input.DetailPajakDtos {
		switch v.JenisOp {
		case "airTanah":
			var tmp m.DetailPotensiAirTanah
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPAirTanahs = append(dPAirTanahs, tmp)
		case "hiburan":
			var tmp m.DetailPotensiHiburan
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPHiburans = append(dPHiburans, tmp)
		case "hotel":
			var tmp m.DetailPotensiHotel
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPHotels = append(dPHotels, tmp)
		case "parkir":
			var tmp m.DetailPotensiParkir
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPParkirs = append(dPParkirs, tmp)
		case "ppj":
			var tmp m.DetailPotensiPPJ
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPPpjs = append(dPPpjs, tmp)
		case "reklame":
			var tmp m.DetailPotensiReklame
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPReklames = append(dPReklames, tmp)
		case "resto":
			var tmp m.DetailPotensiResto
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPRestos = append(dPRestos, tmp)
		default:
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
		}
	}

	// var dataMRek mr.Rekening
	// if result := a.DB.First(&dataMRek, data.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// var dataMUser mu.User
	// if result := a.DB.First(&dataMUser, data.User_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	dataPotensiOp.Status = registration.StatusAktif

	// Transaction save to db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&dataPotensiOp); result.Error != nil {
			return result.Error
		}

		dataDetailPotensiOp.Potensiop_Id = dataPotensiOp.Id
		if result := tx.Create(&dataDetailPotensiOp); result.Error != nil {
			return result.Error
		}
		dataPemilikWp.Potensiop_Id = dataPotensiOp.Id
		if result := tx.Create(&dataPemilikWp); result.Error != nil {
			return result.Error
		}
		if input.CreatePotensiNarahubung != nil {
			dataPemilikNarahub.Potensiop_Id = dataPotensiOp.Id
			if result := tx.Create(&dataPemilikNarahub); result.Error != nil {
				return result.Error
			}
		}
		// for _, v := range listTransactionData {
		// 	vType := reflect.ValueOf(v)
		// 	if vType.Kind() == reflect.Pointer {
		// 		vType = vType.Elem()
		// 	}

		// 	if vType.Kind() != reflect.Slice {
		// 		return errors.New("failed parsing not slice")
		// 	}

		// 	if vType.Len() != 0 {
		// 		if result := tx.Create(&v); result.Error != nil {
		// 			return result.Error
		// 		}
		// 	}
		// }
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
	}

	return rp.OKSimple{Data: dataPotensiOp}, nil
}

func Update(id int, input m.CreateDto) (any, error) {
	var dataPotensiOp *m.PotensiOp
	var dataDetailPotensiOp *m.DetailPotensiOp
	var dataPemilikWp *m.PotensiPemilikWp
	var dataPemilikNarahub *m.PotensiNarahubung
	// var dPAirTanahs []m.DetailPotensiAirTanah
	// var dPHiburans []m.DetailPotensiHiburan
	// var dPHotels []m.DetailPotensiHotel
	// var dPParkirs []m.DetailPotensiParkir
	// var dPPpjs []m.DetailPotensiPPJ
	// var dPReklames []m.DetailPotensiReklame
	// var dPRestos []m.DetailPotensiResto
	// listTransactionData := []interface{}{&dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}

	// validate data exist
	dataResPotensiOp := a.DB.First(&dataPotensiOp, id).RowsAffected
	if dataResPotensiOp == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if result := a.DB.Where(&m.DetailPotensiOp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataDetailPotensiOp); result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if result := a.DB.Where(&m.PotensiPemilikWp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataDetailPotensiOp); result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	// copy to model struct
	if err := sc.Copy(&dataPotensiOp, &input.CreatePotensiOp); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
	}
	if err := sc.Copy(&dataDetailPotensiOp, &input.CreateDetailPotensiOp); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataDetailPotensiOp)
	}
	if err := sc.Copy(&dataPemilikWp, &input.CreatePotensiPemilikWp); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	}
	if input.CreatePotensiNarahubung != nil {
		if err := sc.Copy(&dataPemilikNarahub, &input.CreatePotensiNarahubung); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikNarahub)
		}
	}

	// var dataMRek mr.Rekening
	// if result := a.DB.First(&dataMRek, data.Rekening_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// var dataMUser mu.User
	// if result := a.DB.First(&dataMUser, data.User_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	// transaction update db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Save(&dataPotensiOp); result.Error != nil {
			return result.Error
		}

		dataDetailPotensiOp.Potensiop_Id = dataPotensiOp.Id
		if result := tx.Save(&dataDetailPotensiOp); result.Error != nil {
			return result.Error
		}
		dataPemilikWp.Potensiop_Id = dataPotensiOp.Id
		if result := tx.Save(&dataPemilikWp); result.Error != nil {
			return result.Error
		}
		if input.CreatePotensiNarahubung != nil {
			dataPemilikNarahub.Potensiop_Id = dataPotensiOp.Id
			if result := tx.Save(&dataPemilikNarahub); result.Error != nil {
				return result.Error
			}
		}
		// for _, v := range listTransactionData {
		// 	vType := reflect.ValueOf(v)
		// 	if vType.Kind() == reflect.Pointer {
		// 		vType = vType.Elem()
		// 	}

		// 	if vType.Kind() != reflect.Slice {
		// 		return errors.New("failed parsing not slice")
		// 	}

		// 	if vType.Len() != 0 {
		// 		if result := tx.Create(&v); result.Error != nil {
		// 			return result.Error
		// 		}
		// 	}
		// }
		return nil
	})

	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataResPotensiOp)),
		},
		Data: dataPotensiOp,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.PotensiOp
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

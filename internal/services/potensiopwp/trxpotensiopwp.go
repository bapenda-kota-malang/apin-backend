package potensiopwp

import (
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"

	sdetailobjek "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailobjek"
	sdpotensiop "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/detailpotensiop"
	spnarahubung "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/potensinarahubung"
	sppemilikiwp "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiopwp/potensipemilikwp"
	// mvetax "github.com/bapenda-kota-malang/apin-backend/internal/models/vendoretax"
)

func CreateTrx(input m.CreateDto, userId uint) (any, error) {
	var dataPotensiOp m.PotensiOp

	input.PotensiOp.User_Id = userId

	// Transaction save to db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		// save potensi op
		respPotensiOp, err := Create(input.PotensiOp, tx)
		if err != nil {
			return err
		}
		dataPotensiOp = respPotensiOp.(rp.OKSimple).Data.(m.PotensiOp)

		// replace potensi op id
		input.DetailPotensiOp.Potensiop_Id = dataPotensiOp.Id
		for v := range input.PotensiPemilikWps {
			input.PotensiPemilikWps[v].Potensiop_Id = dataPotensiOp.Id
		}
		for v := range input.PotensiNarahubungs {
			input.PotensiNarahubungs[v].Potensiop_Id = dataPotensiOp.Id
		}
		for v := range input.DetailPajakDtos {
			input.DetailPajakDtos[v].Potensiop_Id = dataPotensiOp.Id
		}

		// save detail potensi op
		_, err = sdpotensiop.Create(input.DetailPotensiOp, tx)
		if err != nil {
			return err
		}

		// save pemilik wps
		_, err = sppemilikiwp.Create(input.PotensiPemilikWps, tx)
		if err != nil {
			return err
		}

		// save narahubungs if len narahubung > 0
		if len(input.PotensiNarahubungs) > 0 {
			_, err := spnarahubung.Create(input.PotensiNarahubungs, tx)
			if err != nil {
				return err
			}
		}

		_, err = sdetailobjek.Create(input.DetailPajakDtos, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data potensi: "+err.Error(), dataPotensiOp)
	}

	return rp.OKSimple{Data: dataPotensiOp}, nil
}

func UpdateTrx(id int, input m.UpdateDto) (any, error) {
	// var dataPotensiOp m.PotensiOp
	// var dataDetailPotensiOp m.DetailPotensiOp
	// var dataPemilikWp m.PotensiPemilikWp
	// var dataPemilikNarahub m.PotensiNarahubung
	// var dPAirTanahs []m.DetailPotensiAirTanah
	// var dPHiburans []m.DetailPotensiHiburan
	// var dPHotels []m.DetailPotensiHotel
	// var dPParkirs []m.DetailPotensiParkir
	// var dPPpjs []m.DetailPotensiPPJ
	// var dPReklames []m.DetailPotensiReklame
	// var dPRestos []m.DetailPotensiResto
	// // multiListData := [][]interface{}{
	// // 	// {&dataPotensiOp, &input.CreatePotensiOp},
	// // 	{&dataDetailPotensiOp, &input.CreateDetailPotensiOp},
	// // 	{&dataPemilikWp, &input.CreatePotensiPemilikWp},
	// // 	{&dataPemilikNarahub, input.CreatePotensiNarahubung},
	// // }
	// listTransactionData := []interface{}{&dataDetailPotensiOp, &dataPemilikWp}

	// if result := a.DB.Where(&m.DetailPotensiOp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataDetailPotensiOp); result.RowsAffected == 0 {
	// 	return nil, errors.New("data tidak dapat ditemukan")
	// }
	// if err := sc.Copy(&dataDetailPotensiOp, &input.CreatePotensiOp); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// }

	// // if result := a.DB.Where(&m.PotensiPemilikWp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataPemilikWp); result.RowsAffected == 0 {
	// // 	return nil, errors.New("data tidak dapat ditemukan")
	// // }
	// if err := sc.Copy(&dataPemilikWp, &input.CreatePotensiPemilikWp); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// }

	// if input.CreatePotensiNarahubung != nil {
	// 	if dataPemilikNarahub != (m.PotensiNarahubung{}) {
	// 		if result := a.DB.Where(&m.PotensiNarahubung{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataPemilikNarahub); result.RowsAffected == 0 {
	// 			return nil, errors.New("data tidak dapat ditemukan")
	// 		}
	// 	}
	// 	if err := sc.Copy(&dataPemilikNarahub, &input.CreatePotensiNarahubung); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikNarahub)
	// 	}
	// }
	// // for _, v := range multiListData {
	// // 	mVal := reflect.ValueOf(v[0])
	// // 	iVal := reflect.ValueOf(v[1])
	// // 	if iVal.IsNil() {
	// // 		continue
	// // 	}
	// // 	if iVal.Kind() == reflect.Pointer {
	// // 		iVal = iVal.Elem()
	// // 	}
	// // 	if result := a.DB.Where(&m.DetailPotensiOp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataDetailPotensiOp); result.RowsAffected == 0 {
	// // 		return nil, errors.New("data tidak dapat ditemukan")
	// // 	}
	// // 	if err := sc.Copy(mVal.Interface(), iVal.Interface()); err != nil {
	// // 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetailPotensiOp)
	// // 	}
	// // }

	// for _, v := range input.DetailPajakDtos {
	// 	switch v.JenisOp {
	// 	case "airTanah":
	// 		var tmp m.DetailPotensiAirTanah
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPAirTanahs = append(dPAirTanahs, tmp)
	// 	case "hiburan":
	// 		var tmp m.DetailPotensiHiburan
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPHiburans = append(dPHiburans, tmp)
	// 	case "hotel":
	// 		var tmp m.DetailPotensiHotel
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPHotels = append(dPHotels, tmp)
	// 	case "parkir":
	// 		var tmp m.DetailPotensiParkir
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPParkirs = append(dPParkirs, tmp)
	// 	case "ppj":
	// 		var tmp m.DetailPotensiPPJ
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPPpjs = append(dPPpjs, tmp)
	// 	case "reklame":
	// 		var tmp m.DetailPotensiReklame
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPReklames = append(dPReklames, tmp)
	// 	case "resto":
	// 		var tmp m.DetailPotensiResto
	// 		if v.Id != nil {
	// 			if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
	// 				return nil, errors.New("data tidak dapat ditemukan")
	// 			}
	// 		}
	// 		if err := sc.Copy(&tmp, &v); err != nil {
	// 			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	// 		}
	// 		dPRestos = append(dPRestos, tmp)
	// 	default:
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
	// 	}
	// }

	// // filter data to listTransactionData before execute db
	// tmp := []interface{}{&dataPemilikNarahub, &dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
	// for _, x := range tmp {
	// 	xVal := reflect.ValueOf(x)
	// 	if xVal.Kind() == reflect.Pointer {
	// 		xVal = xVal.Elem()
	// 	}
	// 	if xVal.Kind() == reflect.Slice {
	// 		if xVal.IsNil() {
	// 			continue
	// 		}
	// 	} else if xVal.Kind() == reflect.Struct {
	// 		newEmpty := reflect.New(reflect.TypeOf(x).Elem()).Elem()
	// 		if xVal.Interface() == newEmpty.Interface() {
	// 			continue
	// 		}
	// 	}
	// 	listTransactionData = append(listTransactionData, x)
	// }

	// // transaction update db
	// err := a.DB.Transaction(func(tx *gorm.DB) error {
	// 	// simpan data ke db satu if karena result dipakai sekali, +error

	// 	_, err := Update(id, input.PotensiOp, tx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// save detail potensi op
	// 	_, err = sdpotensiop.Update(id, input.DetailPotensiOp, tx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// save pemilik wps
	// 	_, err = sppemilikiwp.Update(input.PotensiPemilikWps, tx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	// save narahubungs if len narahubung > 0
	// 	if len(input.PotensiNarahubungs) > 0 {
	// 		_, err := spnarahubung.Update(input.PotensiNarahubungs, tx)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}

	// 	_, err = sdetailobjek.Update(input.DetailPajakDtos, tx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	for _, x := range listTransactionData {
	// 		xVal := reflect.ValueOf(x)
	// 		if xVal.Kind() == reflect.Pointer {
	// 			xVal = xVal.Elem()
	// 		}
	// 		if xVal.Kind() == reflect.Struct {
	// 			if field := xVal.FieldByName("Potensiop_Id"); field.IsValid() {
	// 				field.SetUint(uint64(dataPotensiOp.Id))
	// 			}
	// 			xVal = xVal.Addr()
	// 		} else if xVal.Kind() == reflect.Slice {
	// 			for i := 0; i < xVal.Len(); i++ {
	// 				iVal := xVal.Index(i)
	// 				if field := iVal.FieldByName("Potensiop_Id"); field.IsValid() {
	// 					field.SetUint(uint64(dataPotensiOp.Id))
	// 				}
	// 			}
	// 		}
	// 		if result := tx.Save(xVal.Interface()); result.Error != nil {
	// 			return result.Error
	// 		}
	// 	}
	// 	return nil
	// })
	// if err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
	// }

	// return rp.OK{
	// 	Meta: t.IS{
	// 		// "affected": strconv.Itoa(int(dataRowPotensiOp)),
	// 	},
	// 	Data: dataPotensiOp,
	// }, nil
	return nil, nil
}

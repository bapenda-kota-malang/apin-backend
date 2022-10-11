package esptd

import (
	"errors"
	"strconv"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/esptd"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "esptd"

func GetList(input m.FilterDto) (any, error) {
	var data []m.Espt
	var count int64

	var pagination gh.Pagination
	result := a.DB.Model(&m.Espt{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (any, error) {
	var data *m.Espt

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

// func Create(input m.CreateDto) (any, error) {
// 	var dataPotensiOp m.Espt
// 	var dataDetailPotensiOp m.DetailPotensiOp
// 	var dataPemilikWp m.PotensiPemilikWp
// 	var dataPemilikNarahub m.PotensiNarahubung
// 	var dPAirTanahs []m.DetailPotensiAirTanah
// 	var dPHiburans []m.DetailPotensiHiburan
// 	var dPHotels []m.DetailPotensiHotel
// 	var dPParkirs []m.DetailPotensiParkir
// 	var dPPpjs []m.DetailPotensiPPJ
// 	var dPReklames []m.DetailPotensiReklame
// 	var dPRestos []m.DetailPotensiResto
// 	multiListData := [][]interface{}{
// 		{&dataPotensiOp, &input.CreatePotensiOp},
// 		{&dataDetailPotensiOp, &input.CreateDetailPotensiOp},
// 		{&dataPemilikWp, &input.CreatePotensiPemilikWp},
// 		{&dataPemilikNarahub, input.CreatePotensiNarahubung},
// 	}
// 	listTransactionData := []interface{}{&dataDetailPotensiOp, &dataPemilikWp}

// 	//  copy input (payload) ke struct data jika tidak ada akan error
// 	for _, v := range multiListData {
// 		mVal := reflect.ValueOf(v[0])
// 		iVal := reflect.ValueOf(v[1])
// 		if iVal.IsNil() {
// 			continue
// 		}
// 		if iVal.Kind() == reflect.Pointer {
// 			iVal = iVal.Elem()
// 		}
// 		if err := sc.Copy(mVal.Interface(), iVal.Interface()); err != nil {
// 			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataDetailPotensiOp)
// 		}
// 	}
// 	for _, v := range input.DetailPajakDtos {
// 		switch v.JenisOp {
// 		case "airTanah":
// 			var tmp m.DetailPotensiAirTanah
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPAirTanahs = append(dPAirTanahs, tmp)
// 		case "hiburan":
// 			var tmp m.DetailPotensiHiburan
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPHiburans = append(dPHiburans, tmp)
// 		case "hotel":
// 			var tmp m.DetailPotensiHotel
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPHotels = append(dPHotels, tmp)
// 		case "parkir":
// 			var tmp m.DetailPotensiParkir
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPParkirs = append(dPParkirs, tmp)
// 		case "ppj":
// 			var tmp m.DetailPotensiPPJ
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPPpjs = append(dPPpjs, tmp)
// 		case "reklame":
// 			var tmp m.DetailPotensiReklame
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPReklames = append(dPReklames, tmp)
// 		case "resto":
// 			var tmp m.DetailPotensiResto
// 			if err := sc.Copy(&tmp, &v); err != nil {
// 				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
// 			}
// 			dPRestos = append(dPRestos, tmp)
// 		default:
// 			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
// 		}
// 	}

// 	// check relasi id tabel ke tabel relasi
// 	// espt table
// 	// if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
// 	// 	return nil, nil
// 	// }

// 	// filter data to listTransactionData before execute db
// 	tmp := []interface{}{&dataPemilikNarahub, &dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
// 	for _, x := range tmp {
// 		xVal := reflect.ValueOf(x)
// 		if xVal.Kind() == reflect.Pointer {
// 			xVal = xVal.Elem()
// 		}
// 		if xVal.Kind() == reflect.Slice && xVal.IsNil() {
// 			continue
// 		} else if xVal.Kind() == reflect.Struct {
// 			newEmpty := reflect.New(reflect.TypeOf(x).Elem()).Elem()
// 			if xVal.Interface() == newEmpty.Interface() {
// 				continue
// 			}
// 		}
// 		listTransactionData = append(listTransactionData, x)
// 	}

// 	// static add value to field
// 	dataPotensiOp.Status = registration.StatusAktif

// 	// Transaction save to db
// 	err := a.DB.Transaction(func(tx *gorm.DB) error {
// 		// simpan data ke db satu if karena result dipakai sekali, +error
// 		if result := tx.Create(&dataPotensiOp); result.Error != nil {
// 			return result.Error
// 		}

// 		for _, x := range listTransactionData {
// 			xVal := reflect.ValueOf(x)
// 			if xVal.Kind() == reflect.Pointer {
// 				xVal = xVal.Elem()
// 			}
// 			if xVal.Kind() == reflect.Struct {
// 				if field := xVal.FieldByName("Potensiop_Id"); field.IsValid() {
// 					field.SetUint(uint64(dataPotensiOp.Id))
// 				}
// 				xVal = xVal.Addr()
// 			} else if xVal.Kind() == reflect.Slice {
// 				for i := 0; i < xVal.Len(); i++ {
// 					iVal := xVal.Index(i)
// 					if field := iVal.FieldByName("Potensiop_Id"); field.IsValid() {
// 						field.SetUint(uint64(dataPotensiOp.Id))
// 					}
// 				}
// 			}
// 			if result := tx.Create(xVal.Interface()); result.Error != nil {
// 				return result.Error
// 			}
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
// 	}

// 	return rp.OKSimple{Data: dataPotensiOp}, nil
// }

// func Update(id int, input m.CreateDto) (any, error) {
// 	var data m.Espt
// 	var dataDetailPotensiOp m.DetailPotensiOp
// 	var dataPemilikWp m.PotensiPemilikWp
// 	var dataPemilikNarahub m.PotensiNarahubung
// 	var dPAirTanahs []m.DetailPotensiAirTanah
// 	var dPHiburans []m.DetailPotensiHiburan
// 	var dPHotels []m.DetailPotensiHotel
// 	var dPParkirs []m.DetailPotensiParkir
// 	var dPPpjs []m.DetailPotensiPPJ
// 	var dPReklames []m.DetailPotensiReklame
// 	var dPRestos []m.DetailPotensiResto
// 	// multiListData := [][]interface{}{
// 	// 	// {&dataPotensiOp, &input.CreatePotensiOp},
// 	// 	{&dataDetailPotensiOp, &input.CreateDetailPotensiOp},
// 	// 	{&dataPemilikWp, &input.CreatePotensiPemilikWp},
// 	// 	{&dataPemilikNarahub, input.CreatePotensiNarahubung},
// 	// }
// 	listTransactionData := []interface{}{&dataDetailPotensiOp, &dataPemilikWp}

// 	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
// 	dataRow := a.DB.First(&data, id).RowsAffected
// 	if dataRow == 0 {
// 		return nil, errors.New("data tidak dapat ditemukan")
// 	}

// 	// check relasi id tabel ke tabel relasi
// 	// potensiop table
// 	// if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
// 	// 	return nil, nil
// 	// }

// 	// filter data to listTransactionData before execute db
// 	tmp := []interface{}{&dataPemilikNarahub, &dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
// 	for _, x := range tmp {
// 		xVal := reflect.ValueOf(x)
// 		if xVal.Kind() == reflect.Pointer {
// 			xVal = xVal.Elem()
// 		}
// 		if xVal.Kind() == reflect.Slice {
// 			if xVal.IsNil() {
// 				continue
// 			}
// 		} else if xVal.Kind() == reflect.Struct {
// 			newEmpty := reflect.New(reflect.TypeOf(x).Elem()).Elem()
// 			if xVal.Interface() == newEmpty.Interface() {
// 				continue
// 			}
// 		}
// 		listTransactionData = append(listTransactionData, x)
// 	}

// 	// transaction update db
// 	err := a.DB.Transaction(func(tx *gorm.DB) error {
// 		// simpan data ke db satu if karena result dipakai sekali, +error
// 		if result := tx.Save(&dataPotensiOp); result.Error != nil {
// 			return result.Error
// 		}

// 		for _, x := range listTransactionData {
// 			xVal := reflect.ValueOf(x)
// 			if xVal.Kind() == reflect.Pointer {
// 				xVal = xVal.Elem()
// 			}
// 			if xVal.Kind() == reflect.Struct {
// 				if field := xVal.FieldByName("Potensiop_Id"); field.IsValid() {
// 					field.SetUint(uint64(dataPotensiOp.Id))
// 				}
// 				xVal = xVal.Addr()
// 			} else if xVal.Kind() == reflect.Slice {
// 				for i := 0; i < xVal.Len(); i++ {
// 					iVal := xVal.Index(i)
// 					if field := iVal.FieldByName("Potensiop_Id"); field.IsValid() {
// 						field.SetUint(uint64(dataPotensiOp.Id))
// 					}
// 				}
// 			}
// 			if result := tx.Save(xVal.Interface()); result.Error != nil {
// 				return result.Error
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
// 	}

// 	return rp.OK{
// 		Meta: t.IS{
// 			"affected": strconv.Itoa(int(dataRow)),
// 		},
// 		Data: data,
// 	}, nil
// }

func Delete(id int) (any, error) {
	var data *m.Espt
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

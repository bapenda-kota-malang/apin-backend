package potensiop

import (
	"errors"
	"reflect"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"

	mr "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"

	// mvetax "github.com/bapenda-kota-malang/apin-backend/internal/models/vendoretax"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "potensiop"

func GetList(input m.FilterDto) (any, error) {
	var data []m.PotensiOp
	var count int64

	var pagination gh.Pagination
	result := a.DB.Model(&m.PotensiOp{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload("DetailPotensiOp").
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
	var data *m.PotensiOp

	result := a.DB.Preload(clause.Associations).First(&data, id)
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
	multiListData := [][]interface{}{
		{&dataPotensiOp, &input.CreatePotensiOp},
		{&dataDetailPotensiOp, &input.CreateDetailPotensiOp},
		{&dataPemilikWp, &input.CreatePotensiPemilikWp},
		{&dataPemilikNarahub, input.CreatePotensiNarahubung},
	}
	listTransactionData := []interface{}{&dataDetailPotensiOp, &dataPemilikWp}

	//  copy input (payload) ke struct data jika tidak ada akan error
	for _, v := range multiListData {
		mVal := reflect.ValueOf(v[0])
		iVal := reflect.ValueOf(v[1])
		if iVal.IsNil() {
			continue
		}
		if iVal.Kind() == reflect.Pointer {
			iVal = iVal.Elem()
		}
		if err := sc.Copy(mVal.Interface(), iVal.Interface()); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataDetailPotensiOp)
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

	// check relasi id tabel ke tabel relasi
	// potensiop table
	if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	// if result := a.DB.First(&mvetax.VendorEtax{}, dataPotensiOp.VendorEtax_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }
	if result := a.DB.First(&mad.Kecamatan{}, dataDetailPotensiOp.Kecamatan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kelurahan{}, dataDetailPotensiOp.Kelurahan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kecamatan{}, dataPemilikWp.Kecamatan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kelurahan{}, dataPemilikWp.Kelurahan_Id); result.RowsAffected == 0 {
		return nil, nil
	}

	// filter data to listTransactionData before execute db
	tmp := []interface{}{&dataPemilikNarahub, &dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
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

	// static add value to field
	dataPotensiOp.Status = nt.StatusAktif

	// Transaction save to db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&dataPotensiOp); result.Error != nil {
			return result.Error
		}

		for _, x := range listTransactionData {
			xVal := reflect.ValueOf(x)
			if xVal.Kind() == reflect.Pointer {
				xVal = xVal.Elem()
			}
			if xVal.Kind() == reflect.Struct {
				if field := xVal.FieldByName("Potensiop_Id"); field.IsValid() {
					field.SetUint(uint64(dataPotensiOp.Id))
				}
				xVal = xVal.Addr()
			} else if xVal.Kind() == reflect.Slice {
				for i := 0; i < xVal.Len(); i++ {
					iVal := xVal.Index(i)
					if field := iVal.FieldByName("Potensiop_Id"); field.IsValid() {
						field.SetUint(uint64(dataPotensiOp.Id))
					}
				}
			}
			if result := tx.Create(xVal.Interface()); result.Error != nil {
				return result.Error
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
	}

	return rp.OKSimple{Data: dataPotensiOp}, nil
}

func Update(id int, input m.CreateDto) (any, error) {
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
	// multiListData := [][]interface{}{
	// 	// {&dataPotensiOp, &input.CreatePotensiOp},
	// 	{&dataDetailPotensiOp, &input.CreateDetailPotensiOp},
	// 	{&dataPemilikWp, &input.CreatePotensiPemilikWp},
	// 	{&dataPemilikNarahub, input.CreatePotensiNarahubung},
	// }
	listTransactionData := []interface{}{&dataDetailPotensiOp, &dataPemilikWp}

	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRowPotensiOp := a.DB.First(&dataPotensiOp, id).RowsAffected
	if dataRowPotensiOp == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	// copy to model struct
	if err := sc.Copy(&dataPotensiOp, &input.CreatePotensiOp); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
	}

	if result := a.DB.Where(&m.DetailPotensiOp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataDetailPotensiOp); result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&dataDetailPotensiOp, &input.CreatePotensiOp); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	}

	if result := a.DB.Where(&m.PotensiPemilikWp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataPemilikWp); result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&dataPemilikWp, &input.CreatePotensiPemilikWp); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
	}

	if input.CreatePotensiNarahubung != nil {
		if dataPemilikNarahub != (m.PotensiNarahubung{}) {
			if result := a.DB.Where(&m.PotensiNarahubung{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataPemilikNarahub); result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
		}
		if err := sc.Copy(&dataPemilikNarahub, &input.CreatePotensiNarahubung); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikNarahub)
		}
	}
	// for _, v := range multiListData {
	// 	mVal := reflect.ValueOf(v[0])
	// 	iVal := reflect.ValueOf(v[1])
	// 	if iVal.IsNil() {
	// 		continue
	// 	}
	// 	if iVal.Kind() == reflect.Pointer {
	// 		iVal = iVal.Elem()
	// 	}
	// 	if result := a.DB.Where(&m.DetailPotensiOp{Potensi: m.Potensi{Potensiop_Id: uint(id)}}).First(&dataDetailPotensiOp); result.RowsAffected == 0 {
	// 		return nil, errors.New("data tidak dapat ditemukan")
	// 	}
	// 	if err := sc.Copy(mVal.Interface(), iVal.Interface()); err != nil {
	// 		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetailPotensiOp)
	// 	}
	// }

	for _, v := range input.DetailPajakDtos {
		switch v.JenisOp {
		case "airTanah":
			var tmp m.DetailPotensiAirTanah
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPAirTanahs = append(dPAirTanahs, tmp)
		case "hiburan":
			var tmp m.DetailPotensiHiburan
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPHiburans = append(dPHiburans, tmp)
		case "hotel":
			var tmp m.DetailPotensiHotel
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPHotels = append(dPHotels, tmp)
		case "parkir":
			var tmp m.DetailPotensiParkir
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPParkirs = append(dPParkirs, tmp)
		case "ppj":
			var tmp m.DetailPotensiPPJ
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPPpjs = append(dPPpjs, tmp)
		case "reklame":
			var tmp m.DetailPotensiReklame
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPReklames = append(dPReklames, tmp)
		case "resto":
			var tmp m.DetailPotensiResto
			if v.Id != nil {
				if result := a.DB.First(&tmp, v.Id); result.RowsAffected == 0 {
					return nil, errors.New("data tidak dapat ditemukan")
				}
			}
			if err := sc.Copy(&tmp, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilikWp)
			}
			dPRestos = append(dPRestos, tmp)
		default:
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPotensiOp)
		}
	}

	// check relasi id tabel ke tabel relasi
	// potensiop table
	if result := a.DB.First(&mr.Rekening{}, dataPotensiOp.Rekening_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	// if result := a.DB.First(&mvetax.VendorEtax{}, dataPotensiOp.VendorEtax_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }
	if result := a.DB.First(&mad.Kecamatan{}, dataDetailPotensiOp.Kecamatan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kelurahan{}, dataDetailPotensiOp.Kelurahan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kecamatan{}, dataPemilikWp.Kecamatan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kelurahan{}, dataPemilikWp.Kelurahan_Id); result.RowsAffected == 0 {
		return nil, nil
	}

	// filter data to listTransactionData before execute db
	tmp := []interface{}{&dataPemilikNarahub, &dPAirTanahs, &dPHiburans, &dPHotels, &dPParkirs, &dPPpjs, &dPReklames, &dPRestos}
	for _, x := range tmp {
		xVal := reflect.ValueOf(x)
		if xVal.Kind() == reflect.Pointer {
			xVal = xVal.Elem()
		}
		if xVal.Kind() == reflect.Slice {
			if xVal.IsNil() {
				continue
			}
		} else if xVal.Kind() == reflect.Struct {
			newEmpty := reflect.New(reflect.TypeOf(x).Elem()).Elem()
			if xVal.Interface() == newEmpty.Interface() {
				continue
			}
		}
		listTransactionData = append(listTransactionData, x)
	}

	// transaction update db
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Save(&dataPotensiOp); result.Error != nil {
			return result.Error
		}

		for _, x := range listTransactionData {
			xVal := reflect.ValueOf(x)
			if xVal.Kind() == reflect.Pointer {
				xVal = xVal.Elem()
			}
			if xVal.Kind() == reflect.Struct {
				if field := xVal.FieldByName("Potensiop_Id"); field.IsValid() {
					field.SetUint(uint64(dataPotensiOp.Id))
				}
				xVal = xVal.Addr()
			} else if xVal.Kind() == reflect.Slice {
				for i := 0; i < xVal.Len(); i++ {
					iVal := xVal.Index(i)
					if field := iVal.FieldByName("Potensiop_Id"); field.IsValid() {
						field.SetUint(uint64(dataPotensiOp.Id))
					}
				}
			}
			if result := tx.Save(xVal.Interface()); result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataPotensiOp)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRowPotensiOp)),
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

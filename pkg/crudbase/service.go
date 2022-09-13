package crudbase

import (
	"fmt"
	"reflect"
	"strings"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	sc "github.com/jinzhu/copier"

	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

func Create(data any, input any) (any, error) {
	dataV := reflect.ValueOf(data)
	for dataV.Kind() == reflect.Ptr {
		dataV = dataV.Elem()
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, input); err != nil {
		dataT := dataV.Type().String()
		dataTs := strings.Split(dataT, ".")
		return sh.SetError("request", "create-data", dataTs[1], "failed", "gagal mengambil payload", data)
	}
	fmt.Println(data)

	// save lewat map, karena direct param ga bisa
	// transfered := []string{}
	// dataV = reflect.ValueOf(data)
	// for dataV.Kind() == reflect.Ptr {
	// 	dataV = dataV.Elem()
	// }
	dataT := dataV.Type()
	for i := 0; i < dataV.NumField(); i++ {
		// transfered = append(transfered, dataT.Field(i).Name)
		fmt.Println(dataT.Field(i).Tag.Get("gorm"))
	}
	// fmt.Println(transfered)

	result := a.DB.Create(&data)
	if result.Error != nil {
		// 	dataVT := dataV.Type().String()
		// 	dataVTs := strings.Split(dataVT, ".")
		// 	return sh.SetError("request", "create-data", dataVTs[1], "failed", "gagal menyimpan data", data)
	}
	// result.

	// datax := dataV.Interface()
	// datay := &datax
	// fmt.Println(datax)
	// if result := a.DB.Select(transfered[0], transfered[1:]).Create(&data); result.Error != nil {
	// 	dataVT := dataV.Type().String()
	// 	dataVTs := strings.Split(dataVT, ".") // alwways part of package
	// 	return sh.SetError("request", "create-data", dataVTs[1], "failed", "gagal mengambil menyimpan data", data)
	// }

	return rp.OKSimple{Data: data}, nil
}

// func GetList(r *http.Request) (any, error) {
// 	var data []m.Group
// 	var count int64
// 	var pagination gh.Pagination

// 	filtered := a.DB.Table("Group").Scopes(gh.Filter(r.URL, m.Filter{}))
// 	filtered.Count(&count)

// 	result := filtered.Scopes(gh.Paginate(r.URL, &pagination)).Find(&data)
// 	if result.Error != nil {
// 		return sh.SetError("request", "get-data-list", "", "failed", "gagal mengambil data", data)
// 	}

// 	return rp.OK{
// 		Meta: t.IS{
// 			"totalCount":   strconv.Itoa(int(count)),
// 			"currentCount": strconv.Itoa(int(result.RowsAffected)),
// 			"page":         strconv.Itoa(pagination.Page),
// 			"pageSize":     strconv.Itoa(pagination.PageSize),
// 		},
// 		Data: data,
// 	}, nil
// }

// func GetDetail(id int) (any, error) {
// 	var data *m.Group

// 	result := a.DB.First(&data, id)
// 	if result.RowsAffected == 0 {
// 		return nil, nil
// 	} else if result.Error != nil {
// 		return sh.SetError("request", "get-data-detail", "", "failed", "gagal mengambil data", data)
// 	}

// 	return rp.OKSimple{
// 		Data: data,
// 	}, nil
// }

// func Update(id int, input m.Update) (any, error) {
// 	var data *m.Group
// 	result := a.DB.First(&data, id)
// 	if result.RowsAffected == 0 {
// 		return nil, errors.New("data tidak dapat ditemukan")
// 	}

// 	if err := sc.Copy(&data, &input); err != nil {
// 		return sh.SetError("request", "update-data", "", "failed", "gagal mengambil payload", data)
// 	}

// 	if result := a.DB.Save(&data); result.Error != nil {
// 		return sh.SetError("request", "update-data", "", "failed", "gagal menyimpan data", data)
// 	}

// 	return rp.OK{
// 		Meta: t.IS{
// 			"affected": strconv.Itoa(int(result.RowsAffected)),
// 		},
// 		Data: data,
// 	}, nil
// }

// func Delete(id int) (any, error) {
// 	var data *m.Group
// 	result := a.DB.First(&data, id)
// 	if result.RowsAffected == 0 {
// 		return nil, errors.New("data tidak dapat ditemukan")
// 	}

// 	result = a.DB.Delete(&data, id)
// 	status := "deleted"
// 	if result.RowsAffected == 0 {
// 		data = nil
// 		status = "no deletion"
// 	}

// 	return rp.OK{
// 		Meta: t.IS{
// 			"count":  strconv.Itoa(int(result.RowsAffected)),
// 			"status": status,
// 		},
// 		Data: data,
// 	}, nil
// }

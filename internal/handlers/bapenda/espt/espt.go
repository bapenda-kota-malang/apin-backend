package espt

import (
	"errors"
	"io"
	"net/http"
	"reflect"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/espt"
)

func validateDetail(w http.ResponseWriter, body io.ReadCloser, data interface{}) (err error) {
	if !hh.ValidateStructByIOR(w, body, &data) {
		err = errors.New("failed")
		return
	}
	// using reflect check len dataDetails array
	dVal := reflect.ValueOf(data)
	if dVal.Kind() == reflect.Pointer {
		dVal = dVal.Elem()
	}
	if field := dVal.FieldByName("DataDetails"); field.IsValid() {
		if field.Kind() == reflect.Slice && field.Len() == 0 {
			err = errors.New("data detail kosong")
			return
		}
	}
	return
}

func checkCategoryCreate(w http.ResponseWriter, body io.ReadCloser, category string) (result interface{}, err error) {
	switch category {
	case "air":
		var input m.CreateDetailAirDto
		if err = validateDetail(w, body, &input); err != nil {
			break
		}
		result, err = s.CreateAir(input)
	case "hotel":
		var input m.CreateDetailHotelDto
		if err = validateDetail(w, body, &input); err != nil {
			break
		}
		result, err = s.CreateHotel(input)
	case "parkir":
		var input m.CreateDetailParkirDto
		if err = validateDetail(w, body, &input); err != nil {
			break
		}
		result, err = s.CreateParkir(input)
	case "reklame":
		var input m.CreateDetailReklameDto
		if err = validateDetail(w, body, &input); err != nil {
			break
		}
		result, err = s.CreateReklame(input)
	case "resto":
		var input m.CreateDetailRestoDto
		if err = validateDetail(w, body, &input); err != nil {
			break
		}
		result, err = s.CreateResto(input)
	default:
		err = errors.New("category tidak diketahui")
	}
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	category := query.Get("category")
	result, err := checkCategoryCreate(w, r.Body, category)
	if err.Error() == "failed" {
		return
	}
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

// func Update(w http.ResponseWriter, r *http.Request) {
// 	id := hh.ValidateAutoInc(w, r, "id")
// 	if id < 1 {
// 		return
// 	}

// 	var data m.DaerahUpdateDto
// 	if !hh.ValidateStructByIOR(w, r.Body, &data) {
// 		return
// 	}

// 	result, err := s.Update(id, data)
// 	hh.DataResponse(w, result, err)
// }

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

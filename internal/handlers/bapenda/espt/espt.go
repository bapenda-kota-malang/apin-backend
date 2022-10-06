package espt

import (
	"errors"
	"io"
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/espt"
)

func checkCategory(w http.ResponseWriter, body io.ReadCloser, category, typeProcess string) (result interface{}, err error) {
	switch category {
	case "air":
		var input m.CreateDetailAirDto
		if !hh.ValidateStructByIOR(w, body, &input) {
			err = errors.New("failed")
			break
		}
		if len(input.DataDetails) == 0 {
			err = errors.New("data detail kosong")
			break
		}
		result, err = s.CreateAir(input)
	case "hotel":
		var input m.CreateDetailHotelDto
		if !hh.ValidateStructByIOR(w, body, &input) {
			err = errors.New("failed")
			break
		}
	default:
		err = errors.New("category tidak diketahui")
	}
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	category := query.Get("category")
	result, err := checkCategory(w, r.Body, category, "create")
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

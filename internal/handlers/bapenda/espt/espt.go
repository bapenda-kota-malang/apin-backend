package espt

import (
	"errors"
	"io"
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailesptair"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/detailespthotel"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/espt"
)

func checkCategory(w http.ResponseWriter, body io.ReadCloser, category string) (dataInput []interface{}, err error) {
	switch category {
	case "air":
		var inputDetail []detailesptair.CreateDto
		if !hh.ValidateStructByIOR(w, body, &inputDetail) {
			err = errors.New("failed")
			break
		}
		dataInput = append(dataInput, inputDetail)
	case "hotel":
		var inputDetail []detailespthotel.CreateDto
		if !hh.ValidateStructByIOR(w, body, &inputDetail) {
			err = errors.New("failed")
			break
		}
		dataInput = append(dataInput, inputDetail)
	default:
		err = errors.New("category tidak diketahui")
		hh.DataResponse(w, nil, err)
	}
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDto
	if !hh.ValidateStructByIOR(w, r.Body, &input) {
		return
	}
	query := r.URL.Query()
	category := query.Get("category")
	// dataInput, err := checkCategory(w, r.Body, category)
	// if err != nil {
	// 	return
	// }
	dataInput := input.DataDetails

	result, err := s.Create(input, dataInput, category)
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

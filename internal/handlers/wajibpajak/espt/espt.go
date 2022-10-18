package espt

import (
	"errors"
	"io"
	"net/http"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/espt"
)

func validateDetail(w http.ResponseWriter, body io.ReadCloser, data interface{}) (err error) {
	// validation body payload with struct data, if check failed will create new error and return
	if !hh.ValidateStructByIOR(w, body, &data) {
		err = errors.New("failed")
		return
	}
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateInput
	var err error
	var result any
	query := r.URL.Query()
	category := query.Get("category")
	switch category {
	case "air":
		var tmp m.CreateDetailAirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "hotel":
		var tmp m.CreateDetailHotelDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "hiburan":
		var tmp m.CreateDetailHiburanDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "parkir":
		var tmp m.CreateDetailParkirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "resto":
		var tmp m.CreateDetailRestoDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "ppjnonpln":
		var tmp m.CreateDetailPpjNonPlnDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	default:
		err = errors.New("category tidak diketahui")
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	if err != nil {
		return
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err = s.CreateDetail(input, uint(authInfo.User_Id))
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

func Update(w http.ResponseWriter, r *http.Request) {
	var input m.UpdateInput
	var err error
	var result any
	query := r.URL.Query()
	category := query.Get("category")
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	switch category {
	case "air":
		var tmp m.UpdateDetailAirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "hotel":
		var tmp m.UpdateDetailHotelDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "hiburan":
		var tmp m.UpdateDetailHiburanDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "parkir":
		var tmp m.UpdateDetailParkirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "resto":
		var tmp m.UpdateDetailRestoDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	case "ppjnonpln":
		var tmp m.UpdateDetailPpjNonPlnDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = tmp
	default:
		err = errors.New("category tidak diketahui")
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	if err != nil {
		return
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err = s.UpdateDetail(id, input, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

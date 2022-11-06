package spt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"

	// gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
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
	var input m.Input
	var err error
	var result any
	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	var baseDto m.CreateDetailBaseDto
	err = validateDetail(w, r.Body, &baseDto)
	if err != nil {
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	category := r.URL.Query().Get("category")
	switch category {
	case "air":
		var tmp m.CreateDetailAirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		tmp.CreateDetailBaseDto = baseDto
		input = &tmp
	case "hiburan":
		var tmp m.CreateDetailHiburanDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "hotel":
		var tmp m.CreateDetailHotelDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "parkir":
		var tmp m.CreateDetailParkirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "ppjnonpln":
		var tmp m.CreateDetailPpjNonPlnDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "ppjpln":
		var tmp m.CreateDetailPpjPlnDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "reklame":
		var tmp m.CreateDetailReklameDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "resto":
		var tmp m.CreateDetailRestoDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	default:
		err = errors.New("category tidak diketahui")
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	if err != nil {
		return
	}
	re := regexp.MustCompile(`^\/\w*`)
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	opts := make(map[string]interface{})
	opts["userId"] = uint(authInfo.User_Id)
	opts["newFile"] = true
	opts["baseUri"] = re.FindString(r.RequestURI)[1:]

	result, err = s.CreateDetail(input, opts, nil)
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

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	category := chi.URLParam(r, "category")
	fmt.Println("category: ", category)
	switch category {
	case "air":
		// var input m.UpdateAirDto
		// if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 	return
		// }

		// result, err := sdsa.Update(id, input)
		// hh.DataResponse(w, result, err)
		// case "hotel":
		// 	var input m.UpdateHotelDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsh.Create(input)
		// 	hh.DataResponse(w, result, err)
		// case "parkir":
		// 	var input m.UpdateParkirDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsp.Update(input)
		// 	hh.DataResponse(w, result, err)
		// case "reklame":
		// 	var input m.UpdateReklameDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsrek.Update(input)
		// 	hh.DataResponse(w, result, err)
		// case "resto":
		// 	var input m.UpdateRestoDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sdsres.Update(input)
		// 	hh.DataResponse(w, result, err)
		// case "jaminan":
		// 	var input m.UpdateJaminanDto
		// 	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		// 		return
		// 	}

		// 	result, err := sjbr.Update(input)
		// 	hh.DataResponse(w, result, err)
	}

}

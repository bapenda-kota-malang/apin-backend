package spt

import (
	"errors"
	"io"
	"net/http"
	"regexp"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
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
	re := regexp.MustCompile(`^\/\w*`)
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	opts := make(map[string]interface{})
	opts["baseUri"] = re.FindString(r.RequestURI)[1:]
	opts["userId"] = uint(authInfo.User_Id)
	opts["newFile"] = true
	category := r.URL.Query().Get("category")
	switch category {
	case "air":
		var tmp m.CreateDetailAirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
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
		if opts["baseUri"].(string) == "sptpd" {
			err = errors.New("category tidak didukung pada tautan ini")
			hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
			return
		}
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

	result, err = s.CreateDetail(r.Context(), input, opts, nil)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}
	re := regexp.MustCompile(`^\/\w*`)
	switch re.FindString(r.RequestURI)[1:] {
	case "sptpd":
		input.Type = mtypes.JenisPajakSA
	case "skpd":
		input.Type = mtypes.JenisPajakOA
	}

	jkOpt := "IS NULL"
	input.JenisKetetapan_Opt = &jkOpt

	result, err := s.GetList(input, 0, "bapenda", nil)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	jenisPajak := string(mtypes.JenisPajakSA)
	re := regexp.MustCompile(`^\/\w*`)
	switch re.FindString(r.RequestURI)[1:] {
	case "skpd":
		jenisPajak = string(mtypes.JenisPajakOA)
	}

	result, err := s.GetDetail(id, jenisPajak, 0)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	var err error
	var result any
	var input m.Input
	re := regexp.MustCompile(`^\/\w*`)
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	opts := make(map[string]interface{})
	opts["baseUri"] = re.FindString(r.RequestURI)[1:]
	opts["userId"] = uint(authInfo.User_Id)
	category := r.URL.Query().Get("category")
	switch category {
	case "air":
		var tmp m.UpdateDetailAirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "hiburan":
		var tmp m.UpdateDetailHiburanDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "hotel":
		var tmp m.UpdateDetailHotelDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "parkir":
		var tmp m.UpdateDetailParkirDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "ppjnonpln":
		var tmp m.UpdateDetailPpjNonPlnDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "ppjpln":
		var tmp m.UpdateDetailPpjPlnDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "reklame":
		if opts["baseUri"].(string) == "sptpd" {
			err = errors.New("category tidak didukung pada tautan ini")
			hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
			return
		}
		var tmp m.UpdateDetailReklameDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	case "resto":
		var tmp m.UpdateDetailRestoDto
		err = validateDetail(w, r.Body, &tmp)
		if err != nil {
			return
		}
		input = &tmp
	default:
		err = errors.New("category tidak diketahui")
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	result, err = s.UpdateDetail(id, input, opts)
	hh.DataResponse(w, result, err)
}

func UpdateVa(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	var input m.UpdateVaDto
	if !hh.ValidateStructByIOR(w, r.Body, &input) {
		return
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.UpdateVa(r.Context(), id, input, uint64(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

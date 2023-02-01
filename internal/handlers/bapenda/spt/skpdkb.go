package spt

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func SkpdkbGetList(w http.ResponseWriter, r *http.Request) {
	jkOptNotNull := "IS NOT NULL"
	jkOptEq := "="
	var input m.SkpdkbListDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	var jk *m.JenisKetetapan

	if input.JenisKetetapan == nil {
		input.JenisKetetapan_Opt = &jkOptNotNull
	} else if string(m.JenisKetetapanSkpdkb) == *input.JenisKetetapan {
		tmp := m.JenisKetetapanSkpdkb
		jk = &tmp
		input.JenisKetetapan_Opt = &jkOptEq
	} else if string(m.JenisKetetapanSkpdkbt) == *input.JenisKetetapan {
		tmp := m.JenisKetetapanSkpdkbt
		jk = &tmp
		input.JenisKetetapan_Opt = &jkOptEq
	} else {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{
			Message: "jenis ketetapan tidak valid",
		}, nil)
		return
	}

	inputFilter := m.FilterDto{
		Rekening_Id:        input.Rekening_Id,
		Type:               mtypes.JenisPajak(strings.ToUpper(input.Type)),
		JenisKetetapan:     jk,
		JenisKetetapan_Opt: input.JenisKetetapan_Opt,
		Page:               input.Page,
		PageSize:           input.PageSize,
		NoPagination:       input.NoPagination,
	}

	result, err := s.GetList(inputFilter, 0, "bapenda", nil)
	hh.DataResponse(w, result, err)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var input m.VerifyDto
	if !hh.ValidateStructByIOR(w, r.Body, &input) {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Verify(id, input, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func SkpdkbExisting(w http.ResponseWriter, r *http.Request) {
	typeSpt := strings.ToUpper(chi.URLParam(r, "type"))
	var input m.SkpdkbExisting
	err := validateDetail(w, r.Body, &input)
	if err != nil {
		return
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	opts := make(map[string]interface{})
	opts["userId"] = uint(authInfo.User_Id)
	opts["newFile"] = false
	opts["baseUri"] = typeSpt
	result, err := s.CreateSkpdkbExisting(input, opts)
	hh.DataResponse(w, result, err)
}

func SkpdNew(w http.ResponseWriter, r *http.Request) {
	var input m.Input
	var err error
	typeSpt := strings.ToUpper(chi.URLParam(r, "type"))
	if typeSpt != "SA" {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{
			Message: "type tidak valid",
		}, nil)
		return
	}
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
	result, err := s.CreateSkpdkbNew(input, opts)
	hh.DataResponse(w, result, err)
}

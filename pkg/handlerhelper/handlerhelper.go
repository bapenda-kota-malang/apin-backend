package handlerhelper

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
)

// write error response if validation fails, return boool true if success
func ValidateAutoInc(w http.ResponseWriter, r *http.Request, input string) int {
	id, err := strconv.Atoi(chi.URLParam(r, input))
	if err != nil || id < 1 {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "format harus berupa integer positive"}, nil)
		return 0
	}
	return id
}

// write error response if validation fails, return boool true on success
func ValidateStruct(w http.ResponseWriter, data any) bool {
	err := sv.Validate(data)
	if err != nil {
		httpStatus := http.StatusUnprocessableEntity
		if _, ok := err["struct"]; ok {
			httpStatus = http.StatusBadRequest
		}
		hj.WriteJSON(w, httpStatus, rp.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return boool true on success
func ValidateStructByIOR(w http.ResponseWriter, body io.Reader, data any) bool {
	err := sv.ValidateIoReader(&data, body)
	if err != nil {
		httpStatus := http.StatusUnprocessableEntity
		if _, ok := err["struct"]; ok {
			httpStatus = http.StatusBadRequest
		}
		hj.WriteJSON(w, httpStatus, rp.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
		return false
	}

	return true
}

// by io reader version of ValidateStruct, to cover request.body, return boool true on success
func ValidateStructByURL(w http.ResponseWriter, url url.URL, data any) bool {
	err := sv.ValidateURL(&data, url)
	if err != nil {
		httpStatus := http.StatusUnprocessableEntity
		if _, ok := err["struct"]; ok {
			httpStatus = http.StatusBadRequest
		}
		hj.WriteJSON(w, httpStatus, rp.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
		return false
	}

	return true
}

// respond at the service level that related to data
func DataResponse(w http.ResponseWriter, result any, err error) {
	if result == nil && err == nil {
		hj.WriteJSON(w, http.StatusNotFound, rp.ErrSimple{Message: "data tidak dapat ditemukan"}, nil)
	} else if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rp.ErrSimple{Message: err.Error()}, nil)
	} else {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	}

}

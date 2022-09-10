package handlerhelper

import (
	"io"
	"net/http"
	"strconv"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
)

// send error response if validation fails, return boool with success as true
func StructValidation(w http.ResponseWriter, body io.Reader, data any) bool {
	if err := sv.ValidateIoReader(&data, body); err != nil {
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

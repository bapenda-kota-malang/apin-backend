package pendaftaran

import (
	"net/http"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/registration"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	result, err := registration.GetAll(r)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	result, err := registration.GetDetail(r, id)
	hh.DataResponse(w, result, err)
}

func RegisterByOperator(w http.ResponseWriter, r *http.Request) {
	var register rm.RegisterByOperator
	if hh.ValidateStructByIOR(w, r.Body, &register) == false {
		return
	}

	result, err := registration.RegisterByOperator(r, register)
	hh.DataResponse(w, result, err)
}

package pendaftaran

import (
	"net/http"
	"strconv"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/registration"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	rq "github.com/bapenda-kota-malang/apin-backend/pkg/requester"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	if result, err := registration.GetAll(r); err == nil {
		hj.WriteJSON(w, http.StatusOK, responses.OKSimple{
			Data: result,
		}, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, responses.ErrCustom{
			Meta:     nil,
			Messages: err,
		}, nil)
	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := rq.GetParam("id", r)
	if result, err := registration.GetDetail(r, id); err == nil {
		hj.WriteJSON(w, http.StatusOK, responses.OKSimple{
			Data: result,
		}, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, responses.ErrCustom{
			Meta:     nil,
			Messages: err,
		}, nil)
	}
}

func RegisterByOperator(w http.ResponseWriter, r *http.Request) {
	var register rm.RegisterByOperator
	if err := sv.ValidateIoReader(&register, r.Body); err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, responses.ErrCustom{
			Meta:     types.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
	} else {
		if result, err := registration.RegisterByOperator(r, register); err == nil {
			hj.WriteJSON(w, http.StatusOK, result, nil)
		} else {
			hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
		}
	}
}

package pendaftaran

import (
	"net/http"
	"strconv"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/registration"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	rq "github.com/bapenda-kota-malang/apin-backend/pkg/requester"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	pagination, err := gormhelper.ParseQueryPagination(r.URL.Query())
	if err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, responses.ErrSimple{Message: err.Error()}, nil)
	}

	if result, err := registration.GetAll(pagination); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := rq.GetParam("id", r)
	if result, err := registration.GetDetail(r, id); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
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

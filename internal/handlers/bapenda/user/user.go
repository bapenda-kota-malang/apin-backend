package user

import (
	"net/http"
	"strconv"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rs "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	rq "github.com/bapenda-kota-malang/apin-backend/pkg/requester"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	us "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user um.UserCreate
	if err := sv.ValidateIoReader(&user, r.Body); err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rs.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
	}

	if result, err := us.Create(user); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if result, err := us.GetList(r); err == nil {
		hj.WriteJSON(w, http.StatusOK, rs.OKSimple{
			Data: result,
		}, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rs.ErrCustom{
			Meta:     nil,
			Messages: err,
		}, nil)
	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := rq.GetParam("id", r)
	data := t.II{
		"message": "You are visiting user detail for id " + id + " of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are updating user of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

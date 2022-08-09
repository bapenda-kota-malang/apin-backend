package user

import (
	"net/http"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	rq "github.com/bapenda-kota-malang/apin-backend/pkg/requester"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	us "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data um.UserCreate
	if err := sv.ValidateIoReader(&data, r.Body); err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}

	result := us.Create(data)
	hj.WriteJSON(w, http.StatusOK, result, nil)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are visiting user list of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
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

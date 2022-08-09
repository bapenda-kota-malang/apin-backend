package user

import (
	"net/http"

	"github.com/bapenda-kota-malang/apin-backend/internal/core"
	hj "github.com/bapenda-kota-malang/apin-backend/internal/core/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/internal/core/types"
	sv "github.com/bapenda-kota-malang/apin-backend/internal/libraries/structvalidator"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/usermodel"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var user um.UserCreate

	err := sv.ValidateIoReader(&user, r.Body)
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}

	data := t.II{
		"data": user,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are visiting user list of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := core.GetParam("id", r)
	data := t.II{
		"message": "You are visiting user detail for id " + id + " of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are updating user of app: " + core.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

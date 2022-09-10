package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	rq "github.com/bapenda-kota-malang/apin-backend/pkg/requester"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// complicated here, since there are 2 schemes
	position := r.FormValue("position")
	if position == "" || position == "pegawai" {
		var data pegawai.CreateByUser
		if err := sv.ValidateIoReader(&data, r.Body); err != nil {
			hj.WriteJSON(w, http.StatusUnprocessableEntity, rp.ErrCustom{
				Meta:     t.IS{"count": strconv.Itoa(len(err))},
				Messages: err,
			}, nil)
			return
		}

		if result, err := s.Create(data); err == nil {
			hj.WriteJSON(w, http.StatusOK, result, nil)
		} else {
			hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
		}
	} else {
		var data ppat.CreateByUser
		if err := sv.ValidateIoReader(&data, r.Body); err != nil {
			hj.WriteJSON(w, http.StatusUnprocessableEntity, rp.ErrCustom{
				Meta:     t.IS{"count": strconv.Itoa(len(err))},
				Messages: err,
			}, nil)
		}

		if result, err := s.Create(data); err == nil {
			hj.WriteJSON(w, http.StatusOK, result, nil)
		} else {
			hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
		}
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if result, err := s.GetList(r); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
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

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{
			"message": "format id tidak dapat di kenali",
		}, nil)
		return
	}

	if result, err := s.Delete(id); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{"message": err.Error()}, nil)
	}
}

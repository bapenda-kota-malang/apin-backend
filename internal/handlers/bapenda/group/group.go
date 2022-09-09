package group

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/group"   // model
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/group" // service
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.Create
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
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if result, err := s.GetList(r); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{
			"Messages": "format id tidak dapat di kenali",
		}, nil)
		return
	}

	if result, err := s.GetDetail(id); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{
			"Messages": "format id tidak dapat di kenali",
		}, nil)
		return
	}

	var data m.Update
	if err := sv.ValidateIoReader(&data, r.Body); err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rp.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
	}

	if result, err := s.Update(id, data); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{
			"Messages": "format id tidak dapat di kenali",
		}, nil)
		return
	}

	if result, err := s.Delete(id); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

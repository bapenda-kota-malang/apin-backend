package ppat

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/ppat"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.Create
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Create(data)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	result, err := s.GetList(r)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.Update
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

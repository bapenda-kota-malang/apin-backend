package sppt

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
)

func CreateSimulasi(w http.ResponseWriter, r *http.Request) {
	var data m.RequestSimulasiDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.CreateSimulasi(data)
	hh.DataResponse(w, result, err)
}

func GetListSimulasi(w http.ResponseWriter, r *http.Request) {
	var input m.FilterSimulasiDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListSimulasi(input)
	hh.DataResponse(w, result, err)
}

func GetSimulasiByNop(w http.ResponseWriter, r *http.Request) {
	flag := hh.ValidateString(w, r, "flag")
	if flag == "" {
		return
	}

	var data m.SimulasiNOP
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.GetSimulasiByNop(data, flag)
	hh.DataResponse(w, result, err)
}

func GetDetailSimulasi(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailSimulasi(id)
	hh.DataResponse(w, result, err)
}

func UpdateSimulasi(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.RequestSimulasiDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.UpdateSimulasi(id, data)
	hh.DataResponse(w, result, err)
}

func DeleteSimulasi(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.DeleteSimulasi(id)
	hh.DataResponse(w, result, err)
}

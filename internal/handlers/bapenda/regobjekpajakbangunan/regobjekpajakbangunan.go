package regobjekpajakbangunan

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/regobjekpajakbangunan"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
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

func VerifyLspop(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input m.VerifyDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.VerifyLspop(id, input)
	hh.DataResponse(w, result, err)
}

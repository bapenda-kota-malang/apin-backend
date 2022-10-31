package regnpwpd

import (
	"net/http"

	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func VerifyRegistrasiNpwpd(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input rn.VerifikasiDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.VerifyNpwpd(id, input)
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

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input rn.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.Update(id, input)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input rn.FilterDto
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
	result, err := s.GetDetail(r, id)
	hh.DataResponse(w, result, err)
}

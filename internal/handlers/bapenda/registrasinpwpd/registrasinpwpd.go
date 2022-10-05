package registrasinpwpd

import (
	"net/http"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrasinpwpd"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/registrasinpwpd"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var register rm.RegisterNpwpdCreate
	if hh.ValidateStructByIOR(w, r.Body, &register) == false {
		return
	}

	result, err := s.Create(r, register)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	// parameter dan getAll service harus diganti filterDTO
	var pagination gh.Pagination
	result, err := s.GetAll(pagination)
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

func VerifyRegistrasiNpwpd(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input rm.VerifikasiDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.VerifyNpwpd(id, input)
	hh.DataResponse(w, result, err)
}

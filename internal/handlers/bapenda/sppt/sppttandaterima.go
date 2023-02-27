package sppt

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/tandaterima"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt/tandaterima"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

type CrudSpptTandaTerima struct{}

func (c CrudSpptTandaTerima) Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.Create(data, authInfo.User_Id)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptTandaTerima) GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptTandaTerima) GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptTandaTerima) Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.Update(id, data, authInfo.User_Id)
	hh.DataResponse(w, result, err)
}

func (c CrudSpptTandaTerima) Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

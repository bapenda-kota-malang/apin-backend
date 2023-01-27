package objekpajakpbb

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreateDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.Create(data)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetList(input, authInfo.Ref_Id)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id, authInfo.Ref_Id)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id, nil)
	hh.DataResponse(w, result, err)
}

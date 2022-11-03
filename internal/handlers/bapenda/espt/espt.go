package espt

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/espt"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/espt"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}

	result, err := s.GetList(input, 0)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.GetDetail(id, 0)
	hh.DataResponse(w, result, err)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	var data m.VerifyDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Update(id, data, uint(authInfo.User_Id), nil)
	hh.DataResponse(w, result, err)
}

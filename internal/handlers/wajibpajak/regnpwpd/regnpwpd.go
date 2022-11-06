package regnpwpd

import (
	"net/http"

	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var register rn.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &register) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Create(register, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func UpdateForWp(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var input rn.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.UpdateForWp(id, input, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func GetListForWp(w http.ResponseWriter, r *http.Request) {
	var input rn.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	user_id := uint64(authInfo.User_Id)
	input.User_Id = &user_id

	result, err := s.GetListForWp(input)
	hh.DataResponse(w, result, err)
}

func GetDetailForWp(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.GetDetailForWp(id, uint64(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

func DeleteForWp(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.DeleteForWp(id, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

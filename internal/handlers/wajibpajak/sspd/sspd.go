package sspd

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sspd"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterWpDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	npwpd := hh.ValidateString(w, r, "npwpd")
	if npwpd == "" {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.GetListWp(authInfo.User_Id, npwpd, input)
	hh.DataResponse(w, result, err)
}

package spt

import (
	"net/http"
	"regexp"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if !hh.ValidateStructByURL(w, *r.URL, &input) {
		return
	}
	re := regexp.MustCompile(`^\/\w*`)
	switch re.FindString(r.RequestURI)[1:] {
	case "sptpd":
		input.Type = mtypes.JenisPajakSA
	case "skpd":
		input.Type = mtypes.JenisPajakOA
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.GetList(input, uint(authInfo.User_Id), "wp")
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}
	jenisPajak := string(mtypes.JenisPajakSA)
	re := regexp.MustCompile(`^\/\w*`)
	switch re.FindString(r.RequestURI)[1:] {
	case "skpd":
		jenisPajak = string(mtypes.JenisPajakOA)
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.GetDetail(id, jenisPajak, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

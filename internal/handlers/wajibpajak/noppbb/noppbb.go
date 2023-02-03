package noppbb

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakpbb"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.GetList(m.FilterDto{NoPagination: true}, authInfo.Ref_Id, nil)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	result, err := s.GetDetail(id, authInfo.Ref_Id)
	hh.DataResponse(w, result, err)
}

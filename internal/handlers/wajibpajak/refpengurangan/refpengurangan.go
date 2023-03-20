package pengurangan

import (
	"fmt"
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan/refpengurangan"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDtoRefPengurangan
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.Create(input, uint64(authInfo.User_Id))
	if err != nil {
		err = fmt.Errorf("gagal membuat pengajuan pengurangan baru: %w", err)
	}
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoRefPengurangan
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.GetList(input, authInfo.User_Id)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)
	result, err := s.GetDetail(id, uint(authInfo.User_Id))
	hh.DataResponse(w, result, err)
}

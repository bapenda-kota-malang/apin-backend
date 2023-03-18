package pengurangan

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan/refpengurangan"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"github.com/go-chi/chi/v5"
)

func GetListRefPengurangan(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoRefPengurangan
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input, 0)
	hh.DataResponse(w, result, err)
}

func GetDetailRefPengurangan(w http.ResponseWriter, r *http.Request) {
	id, pass := hh.ValidateIdUuid(w, chi.URLParam(r, "id"))
	if !pass {
		return
	}

	result, err := s.GetDetail(id, 0)
	hh.DataResponse(w, result, err)
}

package sppt

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func ListCatataSejarahWp(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListCatatanSejarahWP(input)
	hh.DataResponse(w, result, err)
}

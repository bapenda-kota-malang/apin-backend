package sppt

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func ListCatataSejarahOp(w http.ResponseWriter, r *http.Request) {
	var input m.CatatanSejarahOPDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListCatatanSejarahOp(input)
	hh.DataResponse(w, result, err)
}

package sppt

import (
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt/sejarah"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt/sejarah"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	"net/http"
)

func GetListSejarahSppt(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.GetSejarahSppt(input)
	hh.DataResponse(w, result, err)
}

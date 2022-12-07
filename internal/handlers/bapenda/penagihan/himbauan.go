package penagihan

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/penagihan"
)

func GetReportHimbauan(w http.ResponseWriter, r *http.Request) {
	var input m.RequestDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetReportHimbauan(input)
	hh.DataResponse(w, result, err)
}

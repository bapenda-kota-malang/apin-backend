package copydbkb

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/copydbkb"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/copydbkb"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func CopyDbkb(w http.ResponseWriter, r *http.Request) {
	var input m.CopyDBKBDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.CopyDBKB(input)
	hh.DataResponse(w, result, err)
}

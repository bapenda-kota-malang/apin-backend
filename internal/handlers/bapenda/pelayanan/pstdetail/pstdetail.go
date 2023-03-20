package pstdetail

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pelayanan/pstdetail"
)

func GetByNoPelayanan(w http.ResponseWriter, r *http.Request) {
	var input m.GetByNoPelayananPstDetailDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.GetByNoPelayanan(input)
	hh.DataResponse(w, result, err)
}

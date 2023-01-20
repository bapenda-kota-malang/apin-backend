package penilaian

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/sppt"
)

func Massal(w http.ResponseWriter, r *http.Request) {
	var data m.PenilaianDto
	if !hh.ValidateStructByIOR(w, r.Body, &data) {
		return
	}

	result, err := s.PenilaianMassal(data)
	hh.DataResponse(w, result, err)
}

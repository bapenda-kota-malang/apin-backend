package permohonan

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	// m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pelayanan"
)

func GetDetailbyNop(w http.ResponseWriter, r *http.Request) {
	nop := hh.ValidateString(w, r, "nop")
	if nop == "" {
		return
	}

	result, err := s.GetDetailbyNop(nop)
	hh.DataResponse(w, result, err)
}

// nyusul
// func UpdateByNop(w http.ResponseWriter, r *http.Request) {
// 	id := hh.ValidateAutoInc(w, r, "nop")
// 	if id < 1 {
// 		return
// 	}

// 	var data m.PermohonanRequestDto
// 	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
// 		return
// 	}

// 	result, err := s.UpdateByNop(id, data)
// 	hh.DataResponse(w, result, err)
// }

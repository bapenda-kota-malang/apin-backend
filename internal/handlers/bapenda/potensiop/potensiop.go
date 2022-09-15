package potensiop

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiop"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

// GetList Data Potensi Objek Pajak with pagination
func GetList(w http.ResponseWriter, r *http.Request) {
	result, err := s.GetList(r)
	hh.DataResponse(w, result, err)
}

// Return data detail Objek Pajak
func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var data m.CreatePotensiOp
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	switch data.Golongan {
	case rm.GolonganBadan:
		break
	case rm.GolonganOrangPribadi:
		break
	default:
		hj.WriteJSON(w, http.StatusBadRequest, t.IS{"message": "golongan tidak diketahui"}, nil)
		return
	}

	result, err := s.Create(data)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.CreatePotensiOp
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data)
	hh.DataResponse(w, result, err)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.Delete(id)
	hh.DataResponse(w, result, err)
}

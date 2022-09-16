package potensiop

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/potensiop"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

// GetList Data Potensi Objek Pajak with pagination
func GetList(w http.ResponseWriter, r *http.Request) {
	pagination, err := gormhelper.ParseQueryPagination(r.URL.Query())
	if err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	result, err := s.GetList(r.URL.Query(), pagination)
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
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: "golongan tidak diketahui"}, nil)
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

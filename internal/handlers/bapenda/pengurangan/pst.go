package pengurangan

import (
	"net/http"
	"strings"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func CreatePST(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDtoPST
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*input.JenisSk)
	input.JenisSk = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.CreatePST(input)
	hh.DataResponse(w, result, err)
}

func GetListPST(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoPST
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListPST(input)
	hh.DataResponse(w, result, err)
}

func GetDetailPST(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailPST(id)
	hh.DataResponse(w, result, err)
}

func UpdatePST(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDtoPST
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*data.JenisSk)
	data.JenisSk = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.UpdatePST(id, data)
	hh.DataResponse(w, result, err)
}

func DeletePST(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.DeletePST(id)
	hh.DataResponse(w, result, err)
}

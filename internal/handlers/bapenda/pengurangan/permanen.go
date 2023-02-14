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

func CreatePermanen(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDtoPermanen
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*input.JenisSk)
	input.JenisSk = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.CreatePermanen(input)
	hh.DataResponse(w, result, err)
}

func GetListPermanen(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoPermanen
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListPermanen(input)
	hh.DataResponse(w, result, err)
}

func GetDetailPermanen(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailPermanen(id)
	hh.DataResponse(w, result, err)
}

func UpdatePermanen(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDtoPermanen
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*data.JenisSk)
	data.JenisSk = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.UpdatePermanen(id, data)
	hh.DataResponse(w, result, err)
}

func DeletePermanen(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.DeletePermanen(id)
	hh.DataResponse(w, result, err)
}

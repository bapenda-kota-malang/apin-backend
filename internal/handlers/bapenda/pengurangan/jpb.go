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

func CreateJPB(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDtoJPB
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	input.SkSk.JnsSK = strings.ToUpper(input.SkSk.JnsSK)
	if err := checkJnsSk(input.SkSk.JnsSK); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.CreateJPB(input)
	hh.DataResponse(w, result, err)
}

func GetListJPB(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoJPB
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListJPB(input)
	hh.DataResponse(w, result, err)
}

func GetDetailJPB(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailJPB(id)
	hh.DataResponse(w, result, err)
}

func UpdateJPB(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDtoJPB
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	// input.SkSk.JnsSK = strings.ToUpper(input.SkSk.JnsSK)
	// if err := checkJnsSk(input.SkSk.JnsSK); err != nil {
	// 	hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	// 	return
	// }

	result, err := s.UpdateJPB(id, data)
	hh.DataResponse(w, result, err)
}

func DeleteJPB(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.DeleteJPB(id)
	hh.DataResponse(w, result, err)
}

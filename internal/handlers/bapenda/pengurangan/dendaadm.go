package pengurangan

import (
	"fmt"
	"net/http"
	"strings"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pengurangan"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
)

func checkJnsSk(jnsSk string) error {
	if jnsSk != "K" {
		return fmt.Errorf("jenis sk tidak valid")
	}
	return nil
}

func CreateDendaADM(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDtoDendaADM
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	input.SkSk.JnsSK = strings.ToUpper(input.SkSk.JnsSK)
	if err := checkJnsSk(input.SkSk.JnsSK); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.CreateDendaADM(input)
	hh.DataResponse(w, result, err)
}

func GetListDendaADM(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoDendaADM
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetListDendaADM(input)
	hh.DataResponse(w, result, err)
}

func GetDetailDendaADM(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetailDendaADM(id)
	hh.DataResponse(w, result, err)
}

func UpdateDendaADM(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDtoDendaADM
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*data.JenisSk)
	data.JenisSk = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.UpdateDendaADM(id, data)
	hh.DataResponse(w, result, err)
}

func DeleteDendaADM(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.DeleteDendaADM(id)
	hh.DataResponse(w, result, err)
}

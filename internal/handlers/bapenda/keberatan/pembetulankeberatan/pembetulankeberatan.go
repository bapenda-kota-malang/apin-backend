package pembetulankeberatan

import (
	"fmt"
	"net/http"
	"strings"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/keberatan/pembetulankeberatan"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
)

func checkJnsSk(jnsSk string) error {
	if jnsSk != "A" {
		return fmt.Errorf("jenis sk tidak valid")
	}
	return nil
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.InputPembetulanSkKeberatan
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	jnsSkUpLama := strings.ToUpper(*input.JenisSKLama)
	input.JenisSKLama = &jnsSkUpLama
	if err := checkJnsSk(jnsSkUpLama); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	jnsSkUpBaru := strings.ToUpper(*input.JenisSKBaru)
	input.JenisSKBaru = &jnsSkUpBaru
	if err := checkJnsSk(jnsSkUpBaru); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.Create(input)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoPemKeb
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDtoPemKeb
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*data.JenisSK)
	data.JenisSK = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.Update(id, data, nil)
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

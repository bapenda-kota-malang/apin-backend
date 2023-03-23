package keputusankeberatanpbb

import (
	"fmt"
	"net/http"
	"strings"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/keberatan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/keberatan/keputusankeberatanpbb"
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
	var input m.CreateDtoKepKebPbb
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*input.JnsSK)
	input.JnsSK = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.Create(input, nil)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDtoKepKebPbb
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

func GetByNopDanSk(w http.ResponseWriter, r *http.Request) {
	var input m.FilterNopSkDtoKepKebPbb
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.GetByNopDanSk(input)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDtoKepKebPbb
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	jnsSkUp := strings.ToUpper(*data.JnsSK)
	data.JnsSK = &jnsSkUp
	if err := checkJnsSk(jnsSkUp); err != nil {
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
		return
	}

	result, err := s.Update(id, data, nil)
	hh.DataResponse(w, result, err)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var verifyDto m.VerifyDtoKepKebPbb
	if hh.ValidateStructByIOR(w, r.Body, &verifyDto) == false {
		return
	}

	result, err := s.Update(id, m.UpdateDtoKepKebPbb{JnsKeputusan: verifyDto.JnsKeputusan}, nil)
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

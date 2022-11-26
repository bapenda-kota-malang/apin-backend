package permohonan

import (
	"fmt"
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/pelayanan"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.PermohonanRequestDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}
	fmt.Println("reached 1")

	result, err := s.Create(input)
	hh.DataResponse(w, result, err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var input m.FilterDto
	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
		return
	}

	result, err := s.GetList(input)
	hh.DataResponse(w, result, err)
}

func GetNoPelayanan(w http.ResponseWriter, r *http.Request) {
	jt := hh.ValidateString(w, r, "jt")
	if jt != "" {
		return
	}

	result, err := s.GetNoPelayanan(&jt)
	hh.DataResponse(w, result, err)
}

func GetStatusNOP(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateString(w, r, "nop")
	if id != "" {
		return
	}
	nop := fmt.Sprint(id)
	jp := r.URL.Query().Get("jp")

	result, err := s.GetStatusNOP(&nop, &jp)
	hh.DataResponse(w, result, err)
}

func GetDetailNOP(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "nop")
	if id < 1 {
		return
	}
	nop := fmt.Sprint(id)
	jp := r.URL.Query().Get("jp")

	result, err := s.GetDetailNOP(&nop, &jp)
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

	var data m.PermohonanRequestDto
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

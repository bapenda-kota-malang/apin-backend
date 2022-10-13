package npwpd

import (
	"fmt"
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	// parameter dan getAll service harus diganti filterDTO
	var pagination gh.Pagination
	result, err := s.GetAll(pagination)
	hh.DataResponse(w, result, err)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}
	result, err := s.GetDetail(r, id)
	hh.DataResponse(w, result, err)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var input m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	result, err := s.Create(r, input)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}
	fmt.Println("datahandler: ", *data.DetailOp.JumlahOp)
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

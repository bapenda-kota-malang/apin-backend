package pendaftaran

import (
	"fmt"
	"net/http"

	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/registrationmodel"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/registration"
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

func RegisterByOperator(w http.ResponseWriter, r *http.Request) {
	var register rm.RegisterByOperator
	if hh.ValidateStructByIOR(w, r.Body, &register) == false {
		return
	}

	result, err := s.RegisterByOperator(r, register)
	hh.DataResponse(w, result, err)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	var data rm.RegisterUpdate
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}
	fmt.Println("datahandler: ", data)
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

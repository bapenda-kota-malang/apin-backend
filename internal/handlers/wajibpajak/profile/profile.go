package profile

import (
	"net/http"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/wajibpajak"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

// func GetList(w http.ResponseWriter, r *http.Request) {
// 	var input m.CreateDto
// 	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
// 		return
// 	}
// 	result, err := s.GetList(input)
// 	hh.DataResponse(w, result, err)
// }

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
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
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

	var data m.UpdateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Update(id, data)
	hh.DataResponse(w, result, err)
}

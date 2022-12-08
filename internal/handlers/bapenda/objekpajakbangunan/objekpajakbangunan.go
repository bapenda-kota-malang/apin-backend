package objekpajakbangunan

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakbangunan"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// var payload t.II
	var result any
	var err error

	// decodedInput := json.NewDecoder(r.Body)
	// err = decodedInput.Decode(&payload)
	// if err != nil {
	// 	hj.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	// 	return
	// }
	// fmt.Println("test: ", payload["jpb_kode"])
	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err = s.Create(data, nil)
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

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

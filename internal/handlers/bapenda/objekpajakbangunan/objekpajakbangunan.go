package objekpajakbangunan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajakbangunan"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// var payload t.II
	var result any
	var data m.CreateDto
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	// decodedInput := json.NewDecoder(r.Body)
	// err = decodedInput.Decode(&payload)
	// if err != nil {
	// 	hj.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	// 	return
	// }
	if err := json.Unmarshal(body, &data); err != nil {
		hj.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	fmt.Println("test: ", data.Jpb_Kode)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err = s.Create(data)
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

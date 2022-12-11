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
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// var payload t.II
	var result any
	var data m.CreateDto
	var dataJpb2 m.OpbJpb2CreateDto
	var dataJpb3 m.CreateDto
	var dataJpb4 m.CreateDto
	var dataJpb5 m.CreateDto
	var dataJpb6 m.CreateDto
	var dataJpb7 m.CreateDto
	var dataJpb8 m.CreateDto
	var dataJpb9 m.CreateDto
	var dataJpb12 m.CreateDto
	var dataJpb13 m.CreateDto
	var dataJpb14 m.CreateDto
	var dataJpb15 m.CreateDto
	var dataJpb16 m.CreateDto
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

	switch data.Jpb_Kode {
	case "02":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb2) == false {
			return
		}
		fmt.Println("nop: ", *dataJpb2.Nop)
		fmt.Println("kode: ", dataJpb2.Jpb_Kode)
		fmt.Println("Jpb2: ", dataJpb2.Jpbs.KelasBangunan2)
		result, err = s.Create(dataJpb2)

	case "03":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb3) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "04":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb4) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "05":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb5) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "06":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb6) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "07":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb7) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "08":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb8) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "09":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb9) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "12":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb12) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "13":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb13) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "14":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb14) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "15":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb15) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	case "16":
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb16) == false {
			return
		}
		result, err = s.Create(dataJpb2)
	default:
		err = fmt.Errorf("jpb kode %s tidak diketahui", data.Jpb_Kode)
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	if err != nil {
		return
	}
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

package regobjekpajakbangunan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/regobjekpajakbangunan"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// var payload t.II
	var result any
	var data m.CreateDto
	var input m.RegInput

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &data); err != nil {
		hj.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	switch data.Jpb_Kode {
	case "02":
		var dataJpb2 m.RegOpbJpb2CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb2) == false {
			return
		}
		input = &dataJpb2
	case "03":
		var dataJpb3 m.RegOpbJpb3CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb3) == false {
			return
		}
		input = &dataJpb3
	case "04":
		var dataJpb4 m.RegOpbJpb4CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb4) == false {
			return
		}
		input = &dataJpb4
	case "05":
		var dataJpb5 m.RegOpbJpb5CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb5) == false {
			return
		}
		input = &dataJpb5
	case "06":
		var dataJpb6 m.RegOpbJpb6CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb6) == false {
			return
		}
		input = &dataJpb6
	case "07":
		var dataJpb7 m.RegOpbJpb7CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb7) == false {
			return
		}
		input = &dataJpb7
	case "08":
		var dataJpb8 m.RegOpbJpb8CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb8) == false {
			return
		}
		input = &dataJpb8
	case "09":
		var dataJpb9 m.RegOpbJpb9CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb9) == false {
			return
		}
		input = &dataJpb9
	case "12":
		var dataJpb12 m.RegOpbJpb12CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb12) == false {
			return
		}
		input = &dataJpb12
	case "13":
		var dataJpb13 m.RegOpbJpb13CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb13) == false {
			return
		}
		input = &dataJpb13
	case "14":
		var dataJpb14 m.RegOpbJpb14CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb14) == false {
			return
		}
		input = &dataJpb14
	case "15":
		var dataJpb15 m.RegOpbJpb15CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb15) == false {
			return
		}
		input = &dataJpb15
	case "16":
		var dataJpb16 m.RegOpbJpb16CreateDto
		if hh.ValidateStructByIOR(w, r.Body, &dataJpb16) == false {
			return
		}
		input = &dataJpb16
	default:
		err = fmt.Errorf("jpb kode %s tidak diketahui", data.Jpb_Kode)
		hj.WriteJSON(w, http.StatusBadRequest, rp.ErrSimple{Message: err.Error()}, nil)
	}
	if err != nil {
		return
	}
	result, err = s.Create(input)
	hh.DataResponse(w, result, err)
}

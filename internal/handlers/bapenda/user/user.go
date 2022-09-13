package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	sc "github.com/jinzhu/copier"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

// complicated process where data depends on position
func Create(w http.ResponseWriter, r *http.Request) {
	var payload t.II
	var result any
	var err error

	decodedInput := json.NewDecoder(r.Body)
	err = decodedInput.Decode(&payload)
	if err != nil {
		hj.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if _, ok := payload["position"]; !ok {
		payload["position"] = "1"
	}
	fmt.Println(payload)
	if payload["position"] == "" || payload["position"] == "1" {
		var data pegawai.Create
		sc.Copy(&data, payload)
		if hh.ValidateStruct(w, &data) == false {
			return
		}
		result, err = s.Create(data)
	} else {
		var data ppat.CreateByUser
		sc.Copy(&data, payload)
		if hh.ValidateStruct(w, &data) == false {
			return
		}
	}

	if err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if result, err := s.GetList(r); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, err, nil)
	}
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{
			"message": "format id tidak dapat di kenali",
		}, nil)
		return
	}

	data := t.II{
		"message": "You are visiting user detail for id " + strconv.Itoa(id) + " of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	data := t.II{
		"message": "You are updating user of app: " + ac.Self.Name,
	}
	hj.WriteJSON(w, http.StatusOK, data, nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{
			"message": "format id tidak dapat di kenali",
		}, nil)
		return
	}

	if result, err := s.Delete(id); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, t.IS{"message": err.Error()}, nil)
	}
}

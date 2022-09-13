package crudbase

import (
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
)

func HandleCreate(m CrudBase) func(http.ResponseWriter, *http.Request) {
	data := m.GetCreateStruct()
	return func(w http.ResponseWriter, r *http.Request) {
		if hh.ValidateStructByIOR(w, r.Body, &data) == false {
			return
		}

		result, err := Create(m, data)
		hh.DataResponse(w, result, err)
	}
}

// func HandleCreate(m CrudBase) http.Handler {
// 	data := m.GetCreateStruct()
// 	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
// 		if hh.ValidateStructByIOR(w, r.Body, &data) == false {
// 			return
// 		}

// 		result, err := s.Create(data)
// 		hh.DataResponse(w, result, err)
// 	})
// }

func HandleGetList(w http.ResponseWriter, r *http.Request) {
	// result, err := s.GetList(r)
	// hh.DataResponse(w, result, err)
}

func HandleGetDetail(w http.ResponseWriter, r *http.Request) {
	// id := hh.ValidateAutoInc(w, r, "id")
	// if id < 1 {
	// 	return
	// }

	// result, err := s.GetDetail(id)
	// hh.DataResponse(w, result, err)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	// id := hh.ValidateAutoInc(w, r, "id")
	// if id < 1 {
	// 	return
	// }

	// var data m.Update
	// if hh.ValidateStructByIOR(w, r.Body, &data) == false {
	// 	return
	// }

	// result, err := s.Update(id, data)
	// hh.DataResponse(w, result, err)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	// id := hh.ValidateAutoInc(w, r, "id")
	// if id < 1 {
	// 	return
	// }

	// result, err := s.Delete(id)
	// hh.DataResponse(w, result, err)
}

package regobjekpajakpbb

import (
	"context"
	"net/http"

	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
	s "github.com/bapenda-kota-malang/apin-backend/internal/services/regobjekpajakpbb"
)

type ctxKey string

const ctxKeyFrom ctxKey = "from"

// middleware specific for create handler, this middleware will add from value to context request with key from
func CreateMw(next http.Handler, from string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ctxKeyFrom, from)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Create(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	fromCtx := reqCtx.Value(ctxKeyFrom)
	from := ""
	if fromCtx != nil {
		from = fromCtx.(string)
	}
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	var data m.CreateDto
	if hh.ValidateStructByIOR(w, r.Body, &data) == false {
		return
	}

	result, err := s.Create(data, from, authInfo)
	hh.DataResponse(w, result, err)
}

// func GetList(w http.ResponseWriter, r *http.Request) {
// 	var input m.FilterDto
// 	if hh.ValidateStructByURL(w, *r.URL, &input) == false {
// 		return
// 	}

// 	result, err := s.GetList(input)
// 	hh.DataResponse(w, result, err)
// }

func GetDetail(w http.ResponseWriter, r *http.Request) {
	id := hh.ValidateAutoInc(w, r, "id")
	if id < 1 {
		return
	}

	result, err := s.GetDetail(id)
	hh.DataResponse(w, result, err)
}

func GetListNop(w http.ResponseWriter, r *http.Request) {
	authInfo := r.Context().Value("authInfo").(*auth.AuthInfo)

	id := authInfo.User_Id

	result, err := s.GetListNop(id)
	hh.DataResponse(w, result, err)
}

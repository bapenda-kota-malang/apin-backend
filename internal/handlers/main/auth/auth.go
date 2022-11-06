package auth

import (
	"context"
	"net/http"
	"strconv"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rs "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	hh "github.com/bapenda-kota-malang/apin-backend/pkg/handlerhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/slicehelper"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	as "github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
)

var SkipAuhPaths []string
var Position int16 = 0

func Login(w http.ResponseWriter, r *http.Request) {
	var input um.LoginDto
	if hh.ValidateStructByIOR(w, r.Body, &input) == false {
		return
	}

	input.Position = Position
	result, err := as.GenToken(input)
	if err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnauthorized, rs.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(1)},
			Messages: err.Error(),
		}, nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func GuardMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sh.StringInSlice(r.URL.Path, SkipAuhPaths) {
			next.ServeHTTP(w, r)
			return
		}
		accessDetail, err := as.ExtractToken(r, as.AccessToken)
		if err != nil {
			http.Error(w, "user tidak memiliki akses untuk data terkait.", 403)
			return
		}
		ctx := context.WithValue(r.Context(), "authInfo", accessDetail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

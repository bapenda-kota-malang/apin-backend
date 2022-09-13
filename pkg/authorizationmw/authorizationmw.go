package authorizationmw

import (
	"net/http"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func CheckAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			hj.WriteJSON(w, http.StatusUnauthorized,
				t.II{"message": "unathorized user"},
				nil,
			)
		} else {
			// if strings.ToLower(authHeader[0:7]) == "bearer " {
			// 	authHeader = authHeader[7:]
			// }
			h.ServeHTTP(w, r)
		}
	}
}

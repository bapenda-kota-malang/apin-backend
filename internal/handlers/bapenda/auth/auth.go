package auth

import (
	"net/http"
	"strconv"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rs "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	as "github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
)

func login(w http.ResponseWriter, r *http.Request) {
	var user um.UserLogin
	if err := sv.ValidateIoReader(&user, r.Body); err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rs.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
	}

	if result, err := as.GenToken(user); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		hj.WriteJSON(w, http.StatusUnauthorized, err, nil)
	}
}

func logout() {

}

package auth

import (
	"fmt"
	"net/http"
	"strconv"

	hj "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/httpjson"
	rs "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	as "github.com/bapenda-kota-malang/apin-backend/internal/services/auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user um.UserLogin
	if err := sv.ValidateIoReader(&user, r.Body); err != nil {
		hj.WriteJSON(w, http.StatusUnprocessableEntity, rs.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(len(err))},
			Messages: err,
		}, nil)
		return
	}

	if result, err := as.GenToken(user); err == nil {
		hj.WriteJSON(w, http.StatusOK, result, nil)
	} else {
		fmt.Println(err)
		hj.WriteJSON(w, http.StatusUnauthorized, rs.ErrCustom{
			Meta:     t.IS{"count": strconv.Itoa(1)},
			Messages: err.Error(),
		}, nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {

}

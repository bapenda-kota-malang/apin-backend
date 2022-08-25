package userservice

import (
	"fmt"
	"net/http"
	"strconv"

	sc "github.com/jinzhu/copier"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	rs "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

func init() {
	ac.AutoMigrate(&um.User{})
}

func Create(data um.UserCreate) (interface{}, error) {
	var user um.User
	sc.Copy(&user, &data) // might results error
	ac.DB.Create(&user)   // might results error

	// if error happens, the return should be something like this
	// return nil, rs.ErrCustom{
	// 	Meta:     nil,
	// 	Messages: err,
	// }

	return rs.OKSimple{Data: user}, nil
}

func GetList(r *http.Request) (interface{}, error) {
	result := ac.DB.Scopes(gh.Paginate(r)).Find(&um.User{})
	resultx := ac.DB.Find(&um.User{})
	fmt.Println(resultx)
	return rp.OK{
		Meta: t.IS{
			"Count": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: result,
	}, nil
}

func GetDetail(data um.UserCreate) (interface{}, error) {
	return "data here", nil
}

package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	r "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

func GenToken(userLogin um.UserLogin) (interface{}, error) {
	var user um.User
	result := ac.DB.Where("name = ?", userLogin.Name).Find(&user)
	if result.Error != nil {
		return nil, errors.New("Internal server error")
	} else if result.RowsAffected == 0 {
		return nil, errors.New("Data can not be found")
	}

	// Prep
	screetKey := viper.GetString("authConf.secretKey")
	atClaims := jwt.MapClaims{}

	// Creating Access Token
	atClaims["user_id"] = user.Id
	atClaims["user_name"] = user.Name
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(screetKey))

	if err != nil {
		return "", err
	}

	return r.OKSimple{
		Data: t.II{
			"user": t.IS{
				"id":    strconv.Itoa(user.Id),
				"name":  user.Name,
				"email": user.Email,
			},
			"accessToken": token,
		},
	}, nil
}

func RevokeToken() {

}

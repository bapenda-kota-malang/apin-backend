package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

func GenToken(userLogin um.UserLogin) (string, error) {
	var user um.User
	result := ac.DB.Where("name = ?", userLogin.Name).Find(&user)
	if result.Error != nil {
		return "", errors.New("Failed to fetch data")
	}
	if result.RowsAffected == 0 {
		return "", errors.New("Invalid password")
	}

	screetKey := viper.GetString("authConf.secretKey")
	fmt.Println(screetKey)

	// Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") // TODO: move to config file
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id
	atClaims["user_name"] = user.Name
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	if token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET"))); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func RevokeToken() {

}

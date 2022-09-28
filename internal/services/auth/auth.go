package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	r "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	p "github.com/bapenda-kota-malang/apin-backend/pkg/password"

	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

type AccessDetails struct {
	Uuid    string
	User_Id uint64
}

var rdClient *redis.Client

// type TokenDetails struct {
// 	AccessToken  string
// 	RefreshToken string
// 	AccessUuid   string
// 	RefreshUuid  string
// 	AtExpires    int64
// 	RtExpires    int64
// }

func init() {
	ac.RegisterExtrCall(RegisterRedis)
}

// Generates token and store in redis at one place
func GenToken(input um.LoginDto) (interface{}, error) {
	var user um.User
	result := ac.DB.Where(um.User{Name: input.Name, Position: input.Position}).Find(&user)
	if result.Error != nil {
		return nil, errors.New("gagal mengambil data")
	} else if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if p.Check(input.Password, *user.Password) == false {
		return nil, errors.New("username atau password tidak sesuai")
	}

	// Access token uuid prep
	id, err := uuid.NewV4()
	if err != nil {
		panic(fmt.Sprintf("gagal mebuat UUID access token: %v", err))
	}
	aUuid := id.String()
	atExpires := time.Now().Add(time.Minute * 15).Unix()
	atSecretKey := viper.GetString("authConf.atSecretKey")

	// Refresh token uuid prep
	id, err = uuid.NewV4()
	if err != nil {
		panic(fmt.Sprintf("gagal mebuat UUID untuk refresh token: %v", err))
	}
	rUuid := id.String()
	rtExpires := time.Now().Add(time.Hour * 24 * 7).Unix()
	rtSecretKey := viper.GetString("authConf.rtSecretKey")

	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id
	atClaims["user_name"] = user.Name
	atClaims["exp"] = atExpires
	atClaims["uuid"] = aUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	ats, err := at.SignedString([]byte(atSecretKey))
	if err != nil {
		return "", err
	}

	// Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = user.Id
	rtClaims["ref_id"] = user.Ref_Id
	rtClaims["exp"] = rtExpires
	rtClaims["uuid"] = rUuid
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rts, err := rt.SignedString([]byte(rtSecretKey))
	if err != nil {
		return nil, err
	}

	// Save to redis
	atx := time.Unix(atExpires, 0) //converting Unix to UTC(to Time object)
	rtx := time.Unix(rtExpires, 0)
	now := time.Now()
	errAccess := rdClient.Set(aUuid, strconv.Itoa(user.Id), atx.Sub(now)).Err()
	if errAccess != nil {
		return nil, errAccess
	}
	errRefresh := rdClient.Set(rUuid, strconv.Itoa(user.Id), rtx.Sub(now)).Err()
	if errRefresh != nil {
		return nil, errRefresh
	}

	return r.OKSimple{
		Data: t.II{
			"user": t.IS{
				"id":    strconv.Itoa(user.Id),
				"name":  user.Name,
				"email": user.Email,
			},
			"accessToken":  ats,
			"refreshToken": rts,
		},
	}, nil
}

func RevokeToken() {
}

func VerifyToken(r *http.Request, tokenType string) (*jwt.Token, error) {
	auth := r.Header.Get("Authorization")
	authArr := strings.Split(auth, " ")
	if len(authArr) == 2 {
		auth = authArr[1]
	}

	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if tokenType == "access" {
			return []byte(viper.GetString("authConf.atSecretKey")), nil
		} else {
			return []byte(viper.GetString("authConf.rtSecretKey")), nil
		}
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request, tokenType string) (interface{}, error) {
	token, err := VerifyToken(r, tokenType)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			Uuid:    accessUuid,
			User_Id: userId,
		}, nil
	}
	return nil, err
}

func RegisterRedis() {
	//Initializing redis
	dsn := viper.GetString("redisConf.dsn")
	if len(dsn) == 0 {
		dsn = "127.0.0.1:6379"
	}

	rdClient = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := rdClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	r "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	p "github.com/bapenda-kota-malang/apin-backend/pkg/password"

	pm "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	am "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	um "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	wm "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
)

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
	// Get User
	var user um.User
	var result *gorm.DB
	if err := getAndCheck(result, &user, um.User{Name: input.Name, Position: input.Position}); err != nil {
		return nil, err
	}
	if user.Status == 3 {
		return nil, errors.New("user tidak dapat login karena diblokir")
	} else if user.Status == 4 {
		return nil, errors.New("user tidak dapat login karena pengajuan pendaftaran ditolak")
	} else if p.Check(input.Password, *user.Password) == false {
		return nil, errors.New("username atau password tidak sesuai")
	}

	// Get Ref
	ref := t.II{}
	ref_type := ""
	jabatan_id := 0
	if user.Position == 1 {
		var refData pm.Pegawai
		if err := getAndCheck(result, &refData, pm.Pegawai{Id: user.Ref_Id}); err != nil {
			if user.SysAdmin == false {
				return nil, err
			}
		}
		ref_type = "pegawai"
		ref["nama"] = refData.Nama
		ref["nip"] = refData.Nip
		ref["jabatan_id"] = refData.Jabatan_Id
		jabatan_id = refData.Jabatan_Id
	} else if user.Position == 2 {
		var refData am.Ppat
		if err := getAndCheck(result, &refData, am.Ppat{Id: user.Ref_Id}); err != nil {
			return nil, err
		}
		ref_type = "ppat"
		ref["nama"] = refData.Nama
		ref["nik"] = refData.Nik
	} else if user.Position == 3 {
		var refData wm.WajibPajak
		if err := getAndCheck(result, &refData, wm.WajibPajak{Id: user.Ref_Id}); err != nil {
			return nil, err
		}
		ref_type = "wajibPajak"
		ref["nama"] = refData.Nama
		ref["nik"] = refData.Nik
	}

	// Access token prep
	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf("gagal mebuat UUID access token: %v", err))
	}
	duration := time.Hour * 24
	if input.LongTerm {
		duration = time.Hour * 24 * 30
	}
	aUuid := id.String()
	atExpires := time.Now().Add(duration).Unix()
	atSecretKey := viper.GetString("authConf.atSecretKey")
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = user.Id
	atClaims["ref_id"] = user.Ref_Id
	atClaims["jabatan_id"] = jabatan_id
	atClaims["exp"] = atExpires
	atClaims["uuid"] = aUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	ats, err := at.SignedString([]byte(atSecretKey))
	if err != nil {
		return "", err
	}
	// Save to redis
	now := time.Now()
	atx := time.Unix(atExpires, 0) //converting Unix to UTC(to Time object)
	errAccess := rdClient.Set(aUuid, strconv.Itoa(user.Id), atx.Sub(now)).Err()
	if errAccess != nil {
		return nil, errAccess
	}

	// // Refresh token uuid prep // WE DON'T NEED THIS FOR NOW
	// id, err = uuid.NewV4()
	// if err != nil {
	// 	panic(fmt.Sprintf("gagal mebuat UUID untuk refresh token: %v", err))
	// }
	// rUuid := id.String()
	// rtExpires := time.Now().Add(time.Hour * 24 * 7).Unix()
	// rtSecretKey := viper.GetString("authConf.rtSecretKey")
	// // Creating Refresh Token // WE DON'T NEED THIS FOR NOW
	// rtClaims := jwt.MapClaims{}
	// rtClaims["user_id"] = user.Id
	// rtClaims["ref_id"] = user.Ref_Id
	// rtClaims["exp"] = rtExpires
	// rtClaims["uuid"] = rUuid
	// rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	// rts, err := rt.SignedString([]byte(rtSecretKey))
	// if err != nil {
	// 	return nil, err
	// }
	// Save to redis
	// rtx := time.Unix(rtExpires, 0)
	// errRefresh := rdClient.Set(rUuid, strconv.Itoa(user.Id), rtx.Sub(now)).Err()
	// if errRefresh != nil {
	// 	return nil, errRefresh
	// }

	// Current data
	data := t.II{
		"accessToken": ats,
		// "refreshToken": rts,
		"user": t.IS{
			"id":     strconv.Itoa(user.Id),
			"name":   user.Name,
			"email":  user.Email,
			"status": strconv.Itoa(int(user.Status)),
		},
		ref_type: ref,
	}

	return r.OKSimple{Data: data}, nil
}

func RevokeToken() {
}

func VerifyToken(r *http.Request, tokenType TokenType) (*jwt.Token, error) {
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
		if tokenType == AccessToken {
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

func ExtractToken(r *http.Request, tokenType TokenType) (*AuthInfo, error) {
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
		user_id, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		ref_id, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["ref_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		jabatan_id, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["jabatan_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AuthInfo{
			Uuid:       accessUuid,
			User_Id:    int(user_id),
			Ref_Id:     int(ref_id),
			Jabatan_Id: int(jabatan_id),
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

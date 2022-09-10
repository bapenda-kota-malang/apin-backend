package userservice

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	sc "github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

func init() {
	a.AutoMigrate(&pegawai.Pegawai{})
	a.AutoMigrate(&ppat.Ppat{})
	a.AutoMigrate(&m.User{})
}

func myErrLogger(xtype, source, status, message string, data any) {
	dataString, _ := json.Marshal(data)
	a.Logger.Error("request",
		zap.String("type", xtype),
		zap.String("source", source),
		zap.String("status", status),
		zap.String("message", message),
		zap.String("data", string(dataString)))
}

func Create(input any) (any, error) {
	switch input.(type) {
	case pegawai.CreateByUser:
		var pegawai pegawai.Pegawai
		var user m.User

		// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
		if err := sc.Copy(&pegawai, &input); err != nil {
			myErrLogger("create-data", "pegawai", "failed", "failed to copy payload", pegawai)
			return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
		}
		if err := sc.Copy(&user, &input); err != nil {
			myErrLogger("create-data", "user", "failed", "failed to copy payload", user)
			return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
		}

		a.DB.Transaction(func(tx *gorm.DB) error {
			// simpan data ke db satu if karena result dipakai sekali, +error
			if result := a.DB.Create(&pegawai); result.Error != nil {
				myErrLogger("create-data", "pegawai", "failed", "failed to create", pegawai)
				return errors.New("penyimpanan data gagal")
			}

			user.Position = 1
			user.Ref_Id = pegawai.Id
			if result := a.DB.Create(&user); result.Error != nil {
				myErrLogger("create-data", "user", "failed", "failed to create", user)
				return errors.New("penyimpanan data gagal")
			}

			// return nil will commit the whole transaction
			return nil
		})

		return rp.OKSimple{Data: input}, nil
	case ppat.CreateByUser:
		var ppat ppat.Ppat
		var user m.User

		// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
		if err := sc.Copy(&ppat, &input); err != nil {
			myErrLogger("create-data", "pegawai", "failed", "failed to copy payload", ppat)
			return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
		}
		if err := sc.Copy(&user, &input); err != nil {
			myErrLogger("create-data", "user", "failed", "failed to copy payload", user)
			return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
		}

		a.DB.Transaction(func(tx *gorm.DB) error {
			// simpan data ke db satu if karena result dipakai sekali, +error
			if result := a.DB.Create(&ppat); result.Error != nil {
				myErrLogger("create-data", "pegawai", "failed", "failed to create", ppat)
				return errors.New("penyimpanan data gagal")
			}

			user.Position = 1
			user.Ref_Id = ppat.Id
			if result := a.DB.Create(&user); result.Error != nil {
				myErrLogger("create-data", "user", "failed", "failed to create", user)
				return errors.New("penyimpanan data gagal")
			}

			// return nil will commit the whole transaction
			return nil
		})

		return rp.OKSimple{Data: input}, nil
	}

	return nil, errors.New("unknown data type")
}

func GetList(r *http.Request) (interface{}, error) {
	var data []m.User
	var count int64

	filtered := a.DB.Table("Group").Scopes(gh.Filter(r, m.Filter{}))
	filtered.Count(&count)

	result := filtered.Scopes(gh.Paginate(r)).Find(&data)
	if result.Error != nil {
		myErrLogger("get-list", "user", "failed", result.Error.Error(), "")
		return nil, errors.New("proses pengambilan data gagal")
	}

	return rp.OK{
		Meta: t.IS{
			"count":        strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.User

	result := a.DB.First(&data, id)
	// if result.Error != nil {
	// 	myErrLogger("get-detail", "failed", "failed fetch data", result.Error.Error())
	// 	return nil, errors.New("Proses pengambilan data gagal")
	// } else
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.Update) (interface{}, error) {
	var data *m.User
	result := a.DB.First(&data, id)
	// if result.Error != nil {
	// 	myErrLogger("get-detail", "failed", "failed fetch data", "")
	// 	return nil, errors.New("Proses pengambilan data gagal")
	// } else
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		dataString, _ := json.Marshal(data)
		myErrLogger("update-data", "user", "failed", "failed to copy payload", string(dataString))
		return nil, errors.New("proses pengambilan data dari \"body-payload\" gagal")
	}

	if result := a.DB.Save(&data); result.Error != nil {
		myErrLogger("update-data", "user", "failed", "failed to update", data)
		return nil, errors.New("proses penyimpanan data gagal")
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.User
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
	status := "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil
}

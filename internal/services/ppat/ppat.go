// service
package ppat

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"

	sc "github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
	sl "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/stringval"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

///// Private funcs start here

// migrate and register validator
func init() {
	a.AutoMigrate(&m.Ppat{})
	sv.RegisterValidator("validemail", sl.ValEmailValidator)
}

// local logger
func myErrLogger(xtype, source, status, message string, data any) {
	dataString, _ := json.Marshal(data)
	a.Logger.Error("request",
		zap.String("type", xtype),
		zap.String("source", source),
		zap.String("status", status),
		zap.String("message", message),
		zap.String("data", string(dataString)))
}

///// Exported funcs start here

func Create(input m.Create) (any, error) {
	var data m.Ppat
	var dataU mu.User

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, input); err != nil {
		myErrLogger("create-data", "ppat", "failed", "failed to copy payload", data)
		return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
	}
	if err := sc.Copy(&dataU, input); err != nil {
		myErrLogger("create-data", "user", "failed", "failed to copy payload", dataU)
		return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
	}

	if err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ppat ke db satu if karena result dipakai sekali, +error
		if result := a.DB.Create(&data); result.Error != nil {
			myErrLogger("create-data", "ppat", "failed", "failed to create", data)
			return errors.New("penyimpanan data ppat gagal")
		}

		// simpan data user ke db satu if karena result dipakai sekali, +error
		dataU.Position = 2
		dataU.Ref_Id = data.Id
		if result := a.DB.Create(&dataU); result.Error != nil {
			myErrLogger("create-data", "user", "failed", "failed to create", dataU)
			return errors.New("penyimpanan data user gagal")
		}

		// return nil will commit the whole transaction
		return nil
	}); err != nil {
		return nil, err
	}

	dataU.Password = "********" // cara cepat
	return rp.OKSimple{Data: t.II{
		"ppat": data,
		"user": dataU,
	}}, nil
}

func GetList(query url.Values, pagination gh.Pagination) (interface{}, error) {
	var data []m.Ppat
	var count int64

	// filtered := a.DB.Table("Ppat").Scopes(gh.Filter(query, m.Filter{}))
	// filtered.Count(&count)

	// result := filtered.Scopes(gh.Paginate(&pagination)).Find(&data)
	// if result.Error != nil {
	// 	myErrLogger("get-list", "user", "failed", result.Error.Error(), "")
	// 	return nil, errors.New("proses pengambilan data gagal")
	// }

	return rp.OK{
		Meta: t.IS{
			"totalCount": strconv.Itoa(int(count)),
			// "currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":     strconv.Itoa(pagination.Page),
			"pageSize": strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.Ppat
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		myErrLogger("get-data", "user", "failed", "failed to get data", data)
		return nil, errors.New("proses pengambilan data gagal")
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.Update) (interface{}, error) {
	var data m.Ppat
	var dataU mu.User

	if err := a.DB.Transaction(func(tx *gorm.DB) error {
		result := a.DB.First(&data, id)
		if result.RowsAffected == 0 {
			return errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&data, &input); err != nil {
			myErrLogger("update-data", "ppat", "failed", "failed to copy payload", data)
			return errors.New("proses pengambilan data dari \"body-payload\" gagal")
		}
		if result := a.DB.Save(&data); result.Error != nil {
			myErrLogger("update-data", "ppat", "failed", "failed to update data", data)
			return errors.New("proses penyimpanan data gagal")
		}

		result = a.DB.Where("Ref_Id = ?", data.Id).First(&dataU, id)
		if result.RowsAffected == 0 {
			return errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataU, &input); err != nil {
			myErrLogger("update-data", "user", "failed", "failed to copy payload", dataU)
			return errors.New("proses pengambilan data dari \"body-payload\" gagal")
		}
		if result := a.DB.Save(&dataU); result.Error != nil {
			myErrLogger("update-data", "user", "failed", "failed to update data", dataU)
			return errors.New("proses penyimpanan data gagal")
		}

		return nil
	}); err != nil {
		return nil, err
	}

	dataU.Password = "********" // cara cepat

	return rp.OKSimple{Data: t.II{
		"ppat": data,
		"user": dataU,
	}}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.Ppat
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
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

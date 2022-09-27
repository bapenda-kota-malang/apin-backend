// service
package pegawai

import (
	"encoding/json"
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"go.uber.org/zap"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
	sl "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/stringval"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

///// Private funcs start here
const source = "pegawai"

// migrate and register validator
func init() {
	a.AutoMigrate(&m.Pegawai{})
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
	var data m.Pegawai
	var dataU mu.User

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	if err := sc.Copy(&dataU, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataU)
	}

	// simpan data pegawai ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	// simpan data user ke db satu if karena result dipakai sekali, +error
	dataU.Position = 1
	dataU.Ref_Id = data.Id
	if result := a.DB.Create(&dataU); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", dataU)
	}

	// return nil will commit the whole transaction

	dataU.Password = "********" // cara cepat
	return rp.OKSimple{Data: t.II{
		"pegawai": data,
		"user":    dataU,
	}}, nil
}

func GetList(input m.Filter) (any, error) {
	var data []m.Pegawai
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Pegawai{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.Pegawai
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
	var data *m.Pegawai
	var dataU *mu.User

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	result = a.DB.Where(mu.User{Ref_Id: data.Id}).First(&dataU, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if err := sc.Copy(&dataU, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&dataU); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataU)
	}

	dataU.Password = "********" // cara cepat

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.Pegawai
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

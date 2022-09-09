package groupservice

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	sc "github.com/jinzhu/copier"
	"go.uber.org/zap"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/group"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

func init() {
	a.AutoMigrate(&m.Group{})
}

func myLogger(xtype, status, message string, data any) {
	dataString, _ := json.Marshal(data)
	a.Logger.Error("request",
		zap.String("type", xtype),
		zap.String("source", "group"),
		zap.String("status", status),
		zap.String("message", message),
		zap.String("data", string(dataString)))
}

func Create(input m.Create) (interface{}, error) {
	var data m.Group

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		myLogger("create-data", "failed", "failed to copy payload", data)
		return nil, errors.New("Proses pengambilan data dari \"body-payload\" gagal")
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		myLogger("create-data", "failed", "failed to create", data)
		return nil, errors.New("Proses penyimpanan data gagal")
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(r *http.Request) (interface{}, error) {
	var data []m.Group
	var count int64

	filtered := a.DB.Table("Group").Scopes(gh.Filter(r, m.Filter{}))
	filtered.Count(&count)

	result := filtered.Scopes(gh.Paginate(r)).Find(&data)
	if result.Error != nil {
		myLogger("get-list", "failed", result.Error.Error(), "")
		return nil, errors.New("Proses pengambilan data gagal")
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
	var data *m.Group

	result := a.DB.First(&data, id)
	if result.Error != nil {
		myLogger("get-detail", "failed", "failed fetch data", "")
		return nil, errors.New("Proses pengambilan data gagal")
	} else if result.RowsAffected == 0 {
		data = nil // set to nil since we dont want empty struct
	}

	return rp.OK{
		Meta: t.IS{"count": strconv.Itoa(int(result.RowsAffected))},
		Data: data,
	}, nil
}

func Update(id int, input m.Update) (interface{}, error) {
	var data *m.Group
	result := a.DB.First(&data, id)
	if result.Error != nil {
		myLogger("get-detail", "failed", "failed fetch data", "")
		return nil, errors.New("Proses pengambilan data gagal")
	} else if result.RowsAffected == 0 {
		return nil, errors.New("Data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		dataString, _ := json.Marshal(data)
		myLogger("update-data", "failed", "failed to copy payload", string(dataString))
		return nil, errors.New("Proses pengambilan data dari \"body-payload\" gagal")
	}

	if result := a.DB.Save(&data); result.Error != nil {
		myLogger("update-data", "failed", "failed to update", data)
		return nil, errors.New("Proses penyimpanan data gagal")
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.Group
	result := a.DB.Delete(&data, id)
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

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

func myErrLogger(xtype, status, message string, data any) {
	dataString, _ := json.Marshal(data)
	a.Logger.Error("request",
		zap.String("type", xtype),
		zap.String("source", "group"),
		zap.String("status", status),
		zap.String("message", message),
		zap.String("data", string(dataString)))
}

func Create(input m.Create) (any, error) {
	var data m.Group

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		myErrLogger("create-data", "failed", "failed to copy payload", data)
		return nil, errors.New("pengambilan data dari \"body-payload\" gagal")
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := a.DB.Create(&data); result.Error != nil {
		myErrLogger("create-data", "failed", "failed to create", data)
		return nil, errors.New("penyimpanan data gagal")
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(r *http.Request) (any, error) {
	var data []m.Group
	var count int64
	var pagination gh.Pagination

	filtered := a.DB.Table("Group").Scopes(gh.Filter(r.URL, m.Filter{}))
	filtered.Count(&count)

	result := filtered.Scopes(gh.Paginate(r.URL, &pagination)).Find(&data)
	if result.Error != nil {
		myErrLogger("get-list", "failed", result.Error.Error(), "")
		return nil, errors.New("proses pengambilan data gagal")
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

func GetDetail(id int) (any, error) {
	var data *m.Group

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		myErrLogger("get-data", "failed", "failed to get data", data)
		return nil, errors.New("proses pengambilan data gagal")
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.Update) (any, error) {
	var data *m.Group
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		dataString, _ := json.Marshal(data)
		myErrLogger("update-data", "failed", "failed to copy payload", string(dataString))
		return nil, errors.New("proses pengambilan data dari \"body-payload\" gagal")
	}

	if result := a.DB.Save(&data); result.Error != nil {
		myErrLogger("update-data", "failed", "failed to update", data)
		return nil, errors.New("proses penyimpanan data gagal")
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.Group
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

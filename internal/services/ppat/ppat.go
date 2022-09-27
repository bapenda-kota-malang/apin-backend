// service
package ppat

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sv "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator"
	sl "github.com/bapenda-kota-malang/apin-backend/pkg/structvalidator/stringval"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/ppat"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	su "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

// /// Private funcs start here
const source = "pegawai"

// migrate and register validator
func init() {
	a.AutoMigrate(&m.Ppat{})
	sv.RegisterValidator("validemail", sl.ValEmailValidator)
}

///// Exported funcs start here

func Create(input m.Create) (any, error) {
	var data m.Ppat
	var dataU mu.CreateDto
	var dataUX any

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	if err := sc.Copy(&dataU, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataU)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data pegawai ke db satu if karena result dipakai sekali, +error
		if result := a.DB.Create(&data); result.Error != nil {
			return errors.New("penyimpanan data pegawai gagal")
		}

		// simpan data user melalui user service
		dataU.Position = 2
		dataU.Ref_Id = data.Id
		dataU.RegMode = 1
		dataUXTemp, err := su.Create(dataU)
		if err != nil {
			return err
		}
		dataUX = dataUXTemp

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	dataUP := dataUX.(rp.OKSimple)

	return rp.OKSimple{Data: t.II{
		"ppat": data,
		"user": dataUP.Data,
	}}, nil
}

func GetList(input m.Filter) (any, error) {
	var data []m.Ppat
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.Ppat{}).
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
	var data *m.Ppat
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(id int, input m.Update) (interface{}, error) {
	var data *m.Ppat
	var dataU *mu.User

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data ppat tidak dapat ditemukan", data)
	}
	result = a.DB.Where(mu.User{Ref_Id: data.Id, Position: 2}).First(&dataU, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data user tidak dapat ditemukan", data)
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload ppat", data)
	}
	if err := sc.Copy(&dataU, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload user", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := a.DB.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data ppat")
		}
		if result := a.DB.Save(&dataU); result.Error != nil {
			return errors.New("gagal menyimpan data user")
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", err.Error(), input)
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
	var data *m.Ppat
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

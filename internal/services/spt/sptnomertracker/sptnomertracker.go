package sptnomertracker

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/sptnomertracker"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "sptnomertracker"

func getSpecificLastNumber(input m.Dto, tx *gorm.DB) (data m.SptNomerTracker, err error) {
	if tx == nil {
		tx = a.DB
	}

	result := tx.
		Where(&m.SptNomerTracker{Tahun: *input.Tahun, KodeJenisPajak: *input.KodeJenisPajak}).
		Attrs(&m.SptNomerTracker{LastNumber: 0}).
		FirstOrCreate(&data)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func GetDetail(id int) (any, error) {
	var data *m.SptNomerTracker

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.Dto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.SptNomerTracker
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id uint) (any, error) {
	var data *m.SptNomerTracker
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

func TrxNewNumber(input m.Dto, tx *gorm.DB) (number string, err error) {
	if tx == nil {
		tx = a.DB
	}
	for i := 0; i < 5; i++ {
		err = tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Begin(&sql.TxOptions{Isolation: sql.LevelSerializable})

			sptNumberTracker, err := getSpecificLastNumber(input, tx2)
			if err != nil {
				tx2.Rollback()
				return err
			}

			newNumber := sptNumberTracker.LastNumber + 1

			_, err = Update(int(sptNumberTracker.Id), m.Dto{LastNumber: &newNumber}, tx2)
			if err != nil {
				tx2.Rollback()
				return err
			}

			number = fmt.Sprintf("%05d", newNumber)

			return err
		})
		if err == nil {
			break
		}
		time.Sleep(2 * time.Nanosecond)
	}

	return
}

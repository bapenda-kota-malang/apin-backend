package auth

import (
	"errors"

	ac "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	"gorm.io/gorm"
)

func getAndCheck(r *gorm.DB, input, condition any) error {
	result := ac.DB.Where(condition).Find(input)
	if result.Error != nil {
		return errors.New("gagal mengambil data")
	} else if result.RowsAffected == 0 {
		return errors.New("data tidak dapat ditemukan")
	}

	return nil
}

package sptpd

import (
	"strconv"
	"time"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/bphtb/sptpd"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"gorm.io/gorm"
)

// this function change status to either 01 or 02 based from input, update ppat_id and verifikasiPpatAt
func VerifyPpat(id int, input m.VerifyPpatDto, userId int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.BphtbSptpd

	result := a.DB.Where("\"Status\" IN ?", []string{"00", "01", "02"}).First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data tidak ditemukan", nil)
	} else if result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mencari data existing: "+result.Error.Error(), nil)
	}

	// prepare data
	statusString := ""
	switch input.Status {
	case 1:
		statusString = "01"
	case 2:
		statusString = "02"
	}
	now := time.Now()

	// change field
	data.Status = &statusString
	data.VerifikasiPpatAt = &now
	data.AlasanReject = input.AlasanReject

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah data status oleh ppat", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

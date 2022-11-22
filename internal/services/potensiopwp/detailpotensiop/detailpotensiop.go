package detailpotensiop

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"

	mareadivision "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"

	skecamatan "github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"
	skelurahan "github.com/bapenda-kota-malang/apin-backend/internal/services/kelurahan"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "detailpotensiop"

func checkLastUpdate(condition m.DetailPotensiOp, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DetailPotensiOp
	err = tx.Where(condition).Find(&data).Error
	if err != nil {
		return
	}
	if time.Now().Unix() <= sh.Midnight(data.UpdatedAt).AddDate(0, 0, 1).Unix() {
		err = errors.New("telah melakukan entry data pada hari ini")
	}
	return
}

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DetailPotensiOp

	// TODO: CHECK IF NEW DATA EXISTING WITHIN 1 DAY

	// TODO: CROSS CHECK DATA EXISTING AT OBJEK PAJAK NPWPD

	// check linked area
	respKecamatan, err := skecamatan.GetDetail(int(input.Kecamatan_Id))
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data kecamatan", data)
	}
	respKelurahan, err := skelurahan.GetDetail(int(input.Kelurahan_Id))
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data kecamatan", data)
	}
	kecamatan := respKecamatan.(rp.OKSimple).Data.(*mareadivision.Kecamatan)
	kelurahan := respKelurahan.(rp.OKSimple).Data.(*mareadivision.Kelurahan)
	if kelurahan.Kecamatan_Kode != kecamatan.Kode {
		return sh.SetError("request", "create-data", source, "failed", "kelurahan tidak berada pada kecamatan terkait", data)
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Update(potensiOp_Id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.DetailPotensiOp
	result := tx.Where("\"Potensiop_Id\" = ?", potensiOp_Id).First(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
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

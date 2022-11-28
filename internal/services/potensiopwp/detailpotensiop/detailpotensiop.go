package detailpotensiop

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"

	mareadivision "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"

	skecamatan "github.com/bapenda-kota-malang/apin-backend/internal/services/kecamatan"
	skelurahan "github.com/bapenda-kota-malang/apin-backend/internal/services/kelurahan"
	sobjekpajak "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajak"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "detailpotensiop"

func checkLastUpdate(data m.DetailPotensiOp) (err error) {
	if time.Now().Unix() <= sh.Midnight(data.UpdatedAt).AddDate(0, 0, 1).Unix() {
		err = errors.New("telah melakukan entry data pada hari ini")
	}
	return
}

func Create(input m.CreateDto, rekeningId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DetailPotensiOp

	data, err := GetExisting(input.Nama,
		input.Alamat,
		input.RtRw,
		input.Kecamatan_Id,
		input.Kelurahan_Id,
		rekeningId,
		tx)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	if err := checkLastUpdate(data); err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	kecamatanId := uint64(input.Kecamatan_Id)
	kelurahanId := uint64(input.Kelurahan_Id)
	input.IsNpwpd = sobjekpajak.CheckIsNpwpd(&input.Nama, &input.Alamat, &input.RtRw, &kecamatanId, &kelurahanId, rekeningId, tx)

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

func Update(potensiOp_Id uuid.UUID, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.DetailPotensiOp
	result := tx.Where("\"Potensiop_Id\" = ?", potensiOp_Id.String()).First(&data)
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

func GetExisting(nama, alamat, rtRw string, kecamatanId, kelurahanId uint, rekeningId uint, tx *gorm.DB) (data m.DetailPotensiOp, err error) {
	if tx == nil {
		tx = a.DB
	}
	if result := tx.
		Joins("JOIN \"PotensiOp\" ON \"PotensiOp\".\"Id\" = \"DetailPotensiOp\".\"Potensiop_Id\" AND \"PotensiOp\".\"Rekening_Id\" = ?", rekeningId).
		Where(m.DetailPotensiOp{Nama: nama,
			Alamat:       alamat,
			RtRw:         rtRw,
			Kecamatan_Id: kecamatanId,
			Kelurahan_Id: kelurahanId}).
		First(&data); result.Error != nil {
		err = result.Error
	}
	return
}

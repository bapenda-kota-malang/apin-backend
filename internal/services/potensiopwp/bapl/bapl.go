package bapl

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	sbk "github.com/bapenda-kota-malang/apin-backend/internal/services/bidangkerja"
	"github.com/bapenda-kota-malang/apin-backend/internal/services/pegawai"

	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "bapl"

func Create(input m.CreateDto, userId uint, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.PotensiBapl

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload bapl", data)
	}

	data.CreateBy_Pegawai_Id = userId
	data.Status = mtypes.StatusVerifikasiBaru

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data bapl", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func Verify(potensiOp_Id uuid.UUID, input m.VerifyDto, userId uint, bidangKerjaKode string) (any, error) {
	var data m.PotensiBapl

	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := a.DB.Where("\"PotensiOp_Id\" = ?", potensiOp_Id.String()).First(&data).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if data.Status == mtypes.StatusVerifikasiDisetujuiKabid {
		return sh.SetError("request", "update-data", source, "failed", "data telah disetujui", data)
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	pegawai, err := pegawai.GetFromUserid(uint64(userId))
	if err != nil {
		return nil, err
	}

	bidangKerjaData, err := sbk.GetByKode(bidangKerjaKode)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal data bidang kerja: "+err.Error(), data)
	}

	switch bidangKerjaData.Level {
	case mtypes.LevelJabatanKasubidKasubag:
		data.Kasubid_Pegawai_Id = &pegawai.Id
		if input.Status == "disetujui" {
			data.Status = mtypes.StatusVerifikasiDisetujuiKasubid
			break
		}
		data.Status = mtypes.StatusVerifikasiDitolakKasubid
	case mtypes.LevelJabatanKabid:
		data.Kabid_Pegawai_Id = &pegawai.Id
		if input.Status == "disetujui" {
			data.Status = mtypes.StatusVerifikasiDisetujuiKabid
			break
		}
		data.Status = mtypes.StatusVerifikasiDitolakKabid
	default:
		return sh.SetError("request", "update-data", source, "failed", "level jabatan tidak diperbolehkan verifikasi", data)
	}

	switch data.Status {
	case mtypes.StatusVerifikasiBaru,
		mtypes.StatusVerifikasiDisetujuiKasubid,
		mtypes.StatusVerifikasiDisetujuiKabid,
		mtypes.StatusVerifikasiDitolakKasubid,
		mtypes.StatusVerifikasiDitolakKabid:
		// do nothing
	default:
		return sh.SetError("request", "update-data", source, "failed", "status penetapan tidak diketahui", data)
	}
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

func Update(potensiOp_Id uuid.UUID, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.PotensiBapl
	result := tx.Where("\"Potensiop_Id\" = ?", potensiOp_Id.String()).First(&data)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
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

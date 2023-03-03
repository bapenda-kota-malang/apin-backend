package indukobjekpajak

import (
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/indukobjekpajak"
	m_objekpajakpbb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajakpbb"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "induk objek pajak"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.IndukObjekPajak

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data ", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var (
		data     []m.IndukObjekPajak
		response m.ListDto
	)

	// get objek pajak pbb
	var objekPajaPbb m_objekpajakpbb.ObjekPajakPbb
	if err := a.DB.
		Where("Provinsi_Kode = ?", input.Provinsi_Kode).
		Where("Daerah_Kode = ?", input.Dati2_Kode).
		Where("Kecamatan_Kode = ?", input.Kecamatan_Kode).
		Where("Kelurahan_Kode = ?", input.Kelurahan_Kode).
		Where("Blok_Kode = ?", input.Blok_Kode).
		Where("NoUrut = ?", input.NoUrut).
		Where("JenisOP = ?", input.JenisOP_Kode).
		Preload("Kelurahan").
		First(&objekPajaPbb).Error; err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data ", data)
	}
	response.AlamatObjekBersama = objekPajaPbb.Jalan

	if objekPajaPbb.Kelurahan != nil {
		response.Kelurahan = objekPajaPbb.Kelurahan.Nama
	}

	response.LuasBangunan = objekPajaPbb.TotalLuasBangunan
	response.LuasBumi = objekPajaPbb.TotalLuasBumi

	// relation wajib pajak pbb
	var wajibPajakPbb wajibpajakpbb.WajibPajakPbb
	if err := a.DB.First(&wajibPajakPbb, objekPajaPbb.WajibPajakPbb_Id).Error; err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data ", data)
	}

	// get anggota objek pajak
	var anggotaObjekPajak []anggotaobjekpajak.AnggotaObjekPajak
	if err := a.DB.
		Where("Induk_Provinsi_Kode = ?", input.Provinsi_Kode).
		Where("Induk_Daerah_Kode = ?", input.Dati2_Kode).
		Where("Induk_Kecamatan_Kode = ?", input.Kecamatan_Kode).
		Where("Induk_Kelurahan_Kode = ?", input.Kelurahan_Kode).
		Where("Induk_Blok_Kode = ?", input.Blok_Kode).
		Where("Induk_NoUrut = ?", input.NoUrut).
		Where("Induk_JenisOp = ?", input.JenisOP_Kode).
		Find(&anggotaObjekPajak).Error; err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data ", data)
	}

	var listDetailDto []m.ListDetailDto
	for _, item := range anggotaObjekPajak {
		var detailDto m.ListDetailDto
		// str provinsi_kode + dati2_kode + kecamatan_kode + kelurahan_kode + blok_kode + no_urut + jenis_op
		nopAnggota := fmt.Sprintf("%s%s%s%s%s%s%s", item.Provinsi_Kode, item.Daerah_Kode, item.Kecamatan_Kode, item.Kelurahan_Kode, item.Blok_Kode, item.NoUrut, item.JenisOp)
		detailDto.NOPAnggota = nopAnggota
		detailDto.NamaWajibPajak = wajibPajakPbb.Nama
		detailDto.LuasBumiBeban = item.LuasBumiBeban
		detailDto.LuasBangunanBeban = item.LuasBangunanBeban

		var sppt sppt.Sppt
		if err := a.DB.Where("NJOP_sppt = ?", nopAnggota).First(&sppt).Error; err != nil {
			return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data ", data)
		}
		detailDto.NJOP = sppt.NJOP_sppt

		listDetailDto = append(listDetailDto, detailDto)
	}

	response.Details = listDetailDto

	return rp.OKSimple{
		Data: response,
	}, nil
}

func GetDetail(id int) (any, error) {
	var data *m.IndukObjekPajak

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data ", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.IndukObjekPajak
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data ", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.IndukObjekPajak
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = tx.Delete(&data, id)
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

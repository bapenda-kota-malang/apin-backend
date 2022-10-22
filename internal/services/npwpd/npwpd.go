package npwpd

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	sop "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajak"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "npwpd"

func GetAll(pagination gh.Pagination) (interface{}, error) {
	var (
		register []*npwpd.Npwpd
		count    int64
	)

	result := a.DB.Model(&npwpd.Npwpd{}).
		Preload(clause.Associations).
		//nested preload
		//Preload("PemilikWps.Kelurahan").
		Count(&count).
		// Scopes(gormhelper.Paginate(&pagination)).
		Find(&register)

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: register,
	}, nil
}

func GetDetail(r *http.Request, regID int) (interface{}, error) {
	var register *npwpd.Npwpd
	err := a.DB.Model(&npwpd.Npwpd{}).
		Preload(clause.Associations).First(&register, regID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: register,
	}, err
}

func Create(r *http.Request, reg npwpd.CreateDto) (interface{}, error) {
	// objekpajak
	var dataObjekPajak op.ObjekPajakCreate
	if err := sc.Copy(&dataObjekPajak, reg.ObjekPajak); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", reg.ObjekPajak)
	}
	resultObjekPajak, err := sop.Create(dataObjekPajak)
	if err != nil {
		return nil, err
	}
	resultCastObjekPajak := resultObjekPajak.(rp.OKSimple).Data.(op.ObjekPajak)

	// data rekening
	var rekening *rm.Rekening
	err = a.DB.Model(&rm.Rekening{}).First(&rekening, reg.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	// var tmpverify = npwpd.VerifiyPendaftaranDisetujui
	var tmpNomor = func() int {

		if reg.IsNomorRegistrasiAuto {
			var tmp int
			var tmpNpwpd npwpd.Npwpd
			nomor := a.DB.Last(&tmpNpwpd)
			if nomor.Error != nil {
				return 1
			} else {
				tmp = tmpNpwpd.Nomor
				tmp++
			}
			return tmp
		}
		return reg.Nomor
	}()

	kecamatanIdString := strconv.Itoa(int(*reg.ObjekPajak.Kecamatan_Id))
	kodeJenisUsahaString := *rekening.KodeJenisUsaha
	if kodeJenisUsahaString == "" {
		kodeJenisUsahaString = "xxx"
	}
	tmpNomorString := strconv.Itoa(tmpNomor)
	if len(tmpNomorString) == 1 {
		tmpNomorString = "000" + tmpNomorString
	} else if len(tmpNomorString) == 2 {
		tmpNomorString = "00" + tmpNomorString
	} else if len(tmpNomorString) == 3 {
		tmpNomorString = "0" + tmpNomorString
	}
	npwpdString := tmpNomorString + "." + kecamatanIdString + "." + kodeJenisUsahaString
	register := npwpd.Npwpd{
		JalurRegistrasi:   nt.JalurRegistrasiOperator,
		Status:            nt.StatusAktif,
		JenisPajak:        reg.JenisPajak,
		ObjekPajak_Id:     resultCastObjekPajak.Id,
		Golongan:          reg.Golongan,
		Npwp:              reg.Npwp,
		VerifiedAt:        th.TimeNow(),
		Nomor:             tmpNomor,
		Npwpd:             &npwpdString,
		TanggalPengukuhan: th.ParseTime(*reg.TanggalPengukuhan),
		TanggalNpwpd:      th.ParseTime(*reg.TanggalNpwpd),
		Rekening_Id:       &reg.Rekening_Id,
		Rekening:          rekening,
		TanggalMulaiUsaha: th.ParseTime(*reg.TanggalMulaiUsaha),
		LuasBangunan:      reg.LuasBangunan,
		JamBukaUsaha:      reg.JamBukaUsaha,
		JamTutupUsaha:     reg.JamTutupUsaha,
		Pengunjung:        reg.Pengunjung,
		OmsetOp:           reg.OmsetOp,
		Genset:            &reg.Genset,
		AirTanah:          &reg.AirTanah,
	}
	err = a.DB.Create(&register).Error
	if err != nil {
		return nil, err
	}
	err = insertDetailOp(*rekening.Objek, reg.DetailOp, &register)
	if err != nil {
		return nil, err
	}

	for _, p := range *reg.Pemilik {
		p.Npwpd_Id = register.Id
		// p.Status = npwpd.StatusPemilikBaru
		err := a.DB.Create(&p).Error
		if err != nil {
			return nil, err
		}
	}
	for _, n := range *reg.Narahubung {
		n.Npwpd_Id = register.Id
		// n.Status = npwpd.StatusNarahubungBaru
		err := a.DB.Create(&n).Error
		if err != nil {
			return nil, err
		}
	}
	return rp.OKSimple{
		Data: register,
	}, nil
}

func Update(id int, input npwpd.UpdateDto, user_Id uint) (any, error) {
	//data npwpd
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if *data.User_Id != userIdConv {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	var rekening *rm.Rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	switch *rekening.Objek {
	case "01":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakHotel
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "02":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakResto
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "03":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakHiburan
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "04":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakReklame
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "05":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakPpj
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "06":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakParkir
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	case "07":
		for _, v := range input.DetailObjekPajak {
			var dataDetail *npwpd.DetailObjekPajakAirTanah
			result := a.DB.First(&dataDetail, v.Id)
			if result.RowsAffected == 0 {
				return nil, errors.New("data tidak dapat ditemukan")
			}
			if err := sc.Copy(&dataDetail, &v); err != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetail)
			}
			dataDetail.Npwpd_Id = data.Id
			dataDetail.JenisOp = rekening.Nama
			if result := a.DB.Save(&dataDetail); result.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataDetail)
			}
		}
	}

	var dataObjekPajak op.ObjekPajakUpdate
	if err := sc.Copy(&dataObjekPajak, &input.ObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataObjekPajak)
	}

	_, err = sop.Update(int(input.ObjekPajak.Id), dataObjekPajak)
	if err != nil {
		return nil, err
	}

	for _, v := range input.Narahubung {
		var dataN *npwpd.Narahubung
		result := a.DB.First(&dataN, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataN, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataN)
		}
		dataN.Npwpd_Id = data.Id
		if result := a.DB.Save(&dataN); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataN)
		}
	}

	for _, v := range input.Pemilik {
		var dataP *npwpd.PemilikWp
		result := a.DB.First(&dataP, v.Id)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		if err := sc.Copy(&dataP, &v); err != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataP)
		}
		dataP.Npwpd_Id = data.Id
		if result := a.DB.Save(&dataP); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", dataP)
		}
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	status := "no deletion"
	// data pemilikwp
	var dataPemilik []*npwpd.PemilikWp
	result := a.DB.Where(npwpd.PemilikWp{Npwpd_Id: uint64(id)}).Find(&dataPemilik)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	for _, v := range dataPemilik {
		result = a.DB.Where(npwpd.PemilikWp{Npwpd_Id: uint64(id)}).Delete(&v)
		if result.RowsAffected == 0 {
			dataPemilik = nil
		}
	}

	// data narahubung
	var dataNarahubung []*npwpd.Narahubung
	result = a.DB.Where(npwpd.Narahubung{Npwpd_Id: uint64(id)}).Find(&dataNarahubung)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	for _, v := range dataNarahubung {
		result = a.DB.Where(npwpd.Narahubung{Npwpd_Id: uint64(id)}).Delete(&v)
		if result.RowsAffected == 0 {
			dataPemilik = nil
		}
	}

	// data regis
	var data *npwpd.Npwpd
	result = a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	objekPajakId := int(data.ObjekPajak_Id)
	result = a.DB.Delete(&data, id)
	status = "deleted"
	if result.RowsAffected == 0 {
		data = nil
		status = "no deletion"
	}

	// data objekpajak

	_, err := sop.Delete(objekPajakId)
	if err != nil {
		status = "no deletion"
	}

	// data rekening
	var rekening *rm.Rekening
	err = a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	// delete detail OP
	switch *rekening.Objek {
	case "01":
		// model = reflect.Zero(mActions["detailOpHotel"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakHotel
		result := a.DB.Where(npwpd.DetailObjekPajakHotel{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}

		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakHotel{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}

	case "02":
		// model = reflect.Zero(mActions["detailOpResto"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakResto
		result := a.DB.Where(npwpd.DetailObjekPajakResto{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakResto{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "03":
		// model = reflect.Zero(mActions["detailOpHiburan"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakHiburan
		result := a.DB.Where(npwpd.DetailObjekPajakHiburan{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakHiburan{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "04":
		// model = reflect.Zero(mActions["detailOpReklame"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakReklame
		result := a.DB.Where(npwpd.DetailObjekPajakReklame{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakReklame{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "05":
		// model = reflect.Zero(mActions["detailOpPpj"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakPpj
		result := a.DB.Where(npwpd.DetailObjekPajakPpj{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakPpj{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "06":
		// model = reflect.Zero(mActions["detailOpParkir"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakParkir
		result := a.DB.Where(npwpd.DetailObjekPajakParkir{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakParkir{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	case "07":
		// model = reflect.Zero(mActions["detailOpAirTanah"]).Interface()
		var DataOp []*npwpd.DetailObjekPajakAirTanah
		result := a.DB.Where(npwpd.DetailObjekPajakAirTanah{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Find(&DataOp)
		if result.RowsAffected == 0 {
			return nil, errors.New("data tidak dapat ditemukan")
		}
		for _, v := range DataOp {
			result = a.DB.Where(npwpd.DetailObjekPajakAirTanah{npwpd.DetailObjekPajak{Npwpd_Id: uint64(id)}}).Delete(&v)
			status = "deleted"
			if result.RowsAffected == 0 {
				dataPemilik = nil
				status = "no deletion"
			}
		}
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil

}

func VerifyUser(id int, input mu.VerifikasiDto) (any, error) {
	var data *mu.User
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func GetListForWp(input npwpd.FilterDto) (any, error) {
	var data []npwpd.Npwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&npwpd.Npwpd{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload("Rekening").
		Preload("User").
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

func GetDetailByUser(regID int, user_id uint) (interface{}, error) {
	user_IdConv := uint64(user_id)
	var register *npwpd.Npwpd
	err := a.DB.Model(&npwpd.Npwpd{}).
		Where(npwpd.Npwpd{User_Id: &user_IdConv, Id: uint64(regID)}).
		Preload(clause.Associations).First(&register, regID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return rp.OKSimple{
		Data: register,
	}, err
}

func AddPhotoLainLain(id int, input npwpd.PhotoUpdate, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	data.LainLain = sh.AddMorePhotos(input.LainLain, data.LainLain)
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}
func AddPhotoSuratIzin(id int, input npwpd.PhotoUpdate, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	data.SuratIzinUsaha = sh.AddMorePhotos(input.SuratIzinUsaha, data.SuratIzinUsaha)
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}
func AddPhotoObject(id int, input npwpd.PhotoUpdate, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	data.FotoObjek = sh.AddMorePhotos(input.FotoObjek, data.FotoObjek)
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func AddPhotoKtp(id int, input npwpd.PhotoUpdate, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	var imgNameChan = make(chan string)
	var errChan = make(chan error)
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	go sh.ReplaceImage(data.FotoKtp, input.FotoKtp, imgNameChan, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", data)
	}
	data.FotoKtp = <-imgNameChan
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func DeletePhotoLainLain(id int, input string, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultDelete, errDelete := sh.DeletePhoto(input, data.LainLain)
	if errDelete != nil {
		return sh.SetError("request", "delete-data", source, "failed", errDelete.Error(), input)
	}
	data.LainLain = resultDelete
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func DeletePhotoObject(id int, input string, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultDelete, errDelete := sh.DeletePhoto(input, data.LainLain)
	if errDelete != nil {
		return sh.SetError("request", "delete-data", source, "failed", errDelete.Error(), input)
	}
	data.FotoObjek = resultDelete
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func DeletePhotoSuratIzin(id int, input string, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultDelete, errDelete := sh.DeletePhoto(input, data.LainLain)
	if errDelete != nil {
		return sh.SetError("request", "delete-data", source, "failed", errDelete.Error(), input)
	}
	data.SuratIzinUsaha = resultDelete
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "deletedata", source, "failed", "gagal menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

package npwpd

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	sn "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd/narahubung"
	sp "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd/pemilik"
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

func GetList(input npwpd.FilterDto) (interface{}, error) {
	var data []npwpd.Npwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&npwpd.Npwpd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak").
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
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

func GetDetail(r *http.Request, regID int) (interface{}, error) {
	var register *npwpd.Npwpd
	err := a.DB.Model(&npwpd.Npwpd{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload("PemilikWps.Daerah").
		Preload("PemilikWps.Kelurahan").
		Preload("Narahubungs.Daerah").
		Preload("Narahubungs.Kelurahan").
		First(&register, regID).Error
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

func Create(r *http.Request, input npwpd.CreateDto) (interface{}, error) {
	// objekpajak
	var dataObjekPajak op.ObjekPajakCreateDto
	var rekening *rm.Rekening
	var dataNpwpd npwpd.Npwpd
	var dataPemilik []npwpd.PemilikWpCreateDto
	var dataNarahubung []npwpd.NarahubungCreateDto
	var respDataObjekPajak interface{}
	var respDataPemilik interface{}
	var respDataNarahubung interface{}
	var resp t.II

	// data rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, input.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	// copy data objek pajak
	if err := sc.Copy(&dataObjekPajak, input.ObjekPajak); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload objek pajak", input.ObjekPajak)
	}

	// copy data pemilik
	if err := sc.Copy(&dataPemilik, input.Pemilik); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload pemilik", input.Pemilik)
	}

	// copy data pemilik
	if err := sc.Copy(&dataNarahubung, input.Narahubung); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload narahubung", input.Narahubung)
	}

	// copy data npwpd
	if err := sc.Copy(&dataNpwpd, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload npwpd", input)
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		resultObjekPajak, err := sop.Create(dataObjekPajak, tx)
		if err != nil {
			return err
		}
		resultCastObjekPajak := resultObjekPajak.(rp.OKSimple).Data.(op.ObjekPajak)
		respDataObjekPajak = resultObjekPajak

		// add static field
		// var tmpverify = npwpd.VerifiyPendaftaranDisetujui
		var tmpNomor = generateNomor(input.IsNomorRegistrasiAuto, input.Nomor)
		var tmpNpwpd = GenerateNpwpd(tmpNomor, *input.ObjekPajak.Kecamatan_Id, *rekening.KodeJenisUsaha)
		dataNpwpd.JalurRegistrasi = nt.JalurRegistrasiOperator
		dataNpwpd.Status = mt.StatusAktif
		dataNpwpd.ObjekPajak_Id = resultCastObjekPajak.Id
		dataNpwpd.VerifiedAt = th.TimeNow()
		dataNpwpd.Nomor = tmpNomor
		dataNpwpd.Npwpd = &tmpNpwpd
		dataNpwpd.TanggalPengukuhan = th.ParseTime(input.TanggalPengukuhan)
		dataNpwpd.TanggalNpwpd = th.ParseTime(input.TanggalNpwpd)
		dataNpwpd.TanggalMulaiUsaha = th.ParseTime(input.TanggalMulaiUsaha)

		err = tx.Create(&dataNpwpd).Error
		if err != nil {
			return err
		}

		err = insertDetailObjekPajak(*rekening.Objek, *rekening.Nama, input.DetailObjekPajak, &dataNpwpd, tx)
		if err != nil {
			return err
		}

		// data pemilik
		resultPemilik, err := sp.Create(dataPemilik, dataNpwpd.Id, tx)
		if err != nil {
			return err
		}
		respDataPemilik = resultPemilik

		// data narahubung
		resultNarahubung, err := sn.Create(dataNarahubung, dataNpwpd.Id, tx)
		if err != nil {
			return err
		}
		respDataNarahubung = resultNarahubung

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataNpwpd)
	}
	resp = t.II{
		"objekPajak": respDataObjekPajak.(rp.OKSimple).Data,
		"npwpd":      dataNpwpd,
		"pemilikWp":  respDataPemilik.(rp.OKSimple).Data,
		"narahubung": respDataNarahubung.(rp.OKSimple).Data,
	}
	return rp.OKSimple{
		Data: resp,
	}, nil
}

func Update(id int, input npwpd.UpdateDto, user_Id uint) (any, error) {
	var data *npwpd.Npwpd
	var dataObjekPajak op.ObjekPajakUpdateDto
	var dataNarahubung []npwpd.NarahubungUpdateDto
	var dataPemilik []npwpd.PemilikWpUpdateDto
	var dataDetailObjekPajak []npwpd.DetailObjekPajak
	var respDataObjekPajak interface{}
	var respDataPemilik interface{}
	var respDataNarahubung interface{}
	var resp t.II
	var rekening *rm.Rekening

	//data npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if *data.User_Id != userIdConv {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	// copy data objek pajak
	if err := sc.Copy(&dataObjekPajak, &input.ObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataObjekPajak)
	}

	// copy data npwpd
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	// copy data pemilik
	if err := sc.Copy(&dataPemilik, &input.Pemilik); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataPemilik)
	}

	// copy data narahubung
	if err := sc.Copy(&dataNarahubung, &input.Narahubung); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataNarahubung)
	}

	// copy data detail objek pajak
	if err := sc.Copy(&dataDetailObjekPajak, &input.DetailObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", dataDetailObjekPajak)
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		if result := a.DB.Save(&data); result.Error != nil {
			return errors.New("tidak dapat menyimpan data npwpd")
		}

		// update detail objek pajak
		err = updateDetailObjekPajak(dataDetailObjekPajak, data.Id, *rekening.Objek, *rekening.Nama, tx)
		if err != nil {
			return nil
		}

		// update objek pajak
		resultObjekPajak, err := sop.Update(int(input.ObjekPajak.Id), dataObjekPajak, tx)
		if err != nil {
			return err
		}
		respDataObjekPajak = resultObjekPajak

		// update narahubung
		resultNarahubung, err := sn.Update(dataNarahubung, data.Id, tx)
		if err != nil {
			return err
		}
		respDataNarahubung = resultNarahubung

		// update pemilik
		resultPemilik, err := sp.Update(dataPemilik, data.Id, data.Golongan, tx)
		if err != nil {
			return err
		}
		respDataPemilik = resultPemilik

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data npwpd transaction", data)
	}
	resp = t.II{
		"affected":   strconv.Itoa(int(result.RowsAffected)),
		"objekPajak": respDataObjekPajak.(rp.OK).Data,
		"Npwpd":      data,
		"pemilikWp":  respDataPemilik.(rp.OKSimple).Data,
		"narahubung": respDataNarahubung.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func Delete(id int) (any, error) {
	var status string = "deleted"
	var data *npwpd.Npwpd
	var rekening *rm.Rekening

	// cek data npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data npwpd tidak dapat ditemukan")
	}

	// untuk mengambil data objekpajak
	tmpObjekPajakId := data.ObjekPajak_Id

	// cek data rekening untuk mengambil data objek untuk menghapus detail objek pajak
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// cek dan hapus data pemilik
		err := sp.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// cek dan hapus data narahubung
		err = sn.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete reg detail objek pajak
		err = deleteDetailObjekPajak(data.Id, *rekening.Objek, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// delete data npwpd
		result = tx.Delete(&data, id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data npwpd")
		}

		// delete reg objek pajak
		_, err = sop.Delete(int(tmpObjekPajakId), tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menghapus data: "+err.Error(), data)
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
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
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
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		First(&register, regID).Error
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

	slcLainLain, err := sh.AddMoreFile(input.LainLain, data.LainLain, "npwpdLainLain", uint(user_id))
	if err != nil {
		return nil, err
	}
	data.LainLain = slcLainLain
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

	slcSuratIzin, err := sh.AddMorePdf(input.SuratIzinUsaha, data.SuratIzinUsaha, "npwpdIzinUsaha", uint(user_id))
	if err != nil {
		return nil, err
	}
	data.SuratIzinUsaha = slcSuratIzin
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

	slcFotoObjek, err := sh.AddMorePhotos(input.FotoObjek, data.FotoObjek, "npwpdFotoObjek", uint(user_id))
	if err != nil {
		return nil, err
	}
	data.FotoObjek = slcFotoObjek
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
	var errChan = make(chan error)
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	fileName, path, extFile, err := filePreProcess(input.FotoKtp, uint(user_id), "npwpdFotoKtp")
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}

	go sh.ReplaceFile(data.FotoKtp, input.FotoKtp, fileName, path, extFile, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", data)
	}
	data.FotoKtp = fileName
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

func DeleteFileLainLain(id int, input string, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultDelete, errDelete := sh.DeleteFile(input, data.LainLain)
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

	resultDelete, errDelete := sh.DeleteFile(input, data.FotoObjek)
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

func DeletePdfSuratIzin(id int, input string, user_id uint64) (any, error) {
	var data *npwpd.Npwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if *data.User_Id != user_id {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultDelete, errDelete := sh.DeleteFile(input, data.SuratIzinUsaha)
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

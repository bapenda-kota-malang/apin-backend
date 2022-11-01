package regnpwpd

import (
	"errors"
	"net/http"
	"strconv"

	nm "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	nt "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd/types"
	op "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	rn "github.com/bapenda-kota-malang/apin-backend/internal/models/regnpwpd"
	rop "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajak"
	rm "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	sn "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd"
	rsn "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd/regnarahubung"
	rsop "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd/regobjekpajak"
	rsp "github.com/bapenda-kota-malang/apin-backend/internal/services/regnpwpd/regpemilik"
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

const source = "registrasi npwpd"

func Create(input rn.CreateDto, user_Id uint) (interface{}, error) {
	user_IdConv := uint64(user_Id)
	var rekening *rm.Rekening
	var dataOp rop.RegObjekPajakCreateDto
	var register rn.RegNpwpd
	var dataPemilik []rn.RegPemilikWpCreateDto
	var dataNarahubung []rn.RegNarahubungCreateDto
	var imgNameChan = make(chan string)
	var errChan = make(chan error)
	var respDataOp interface{}
	var respDataPemilik interface{}
	var respDataNarahubung interface{}
	var resp t.II

	go sh.SaveImage(input.FotoKtp, imgNameChan, errChan)
	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", input)
	}

	// get data rekening
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, input.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error

	// data registrasi objek pajak
	if err := sc.Copy(&dataOp, &input.RegObjekPajak); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload regObjekPajak", dataOp)
	}

	// data pemilik wajib pajak
	if err := sc.Copy(&dataPemilik, &input.RegPemilik); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload regPemilik", dataPemilik)
	}

	// data narahubung
	if err := sc.Copy(&dataNarahubung, &input.RegNarahubung); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload regNarahubung", dataNarahubung)
	}

	// data registrasi npwpd
	if err := sc.Copy(&register, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload registrasi npwpd", register)
	}
	register.FotoKtp = <-imgNameChan

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// objekpajak
		resultRegObjekPajak, err := rsop.Create(dataOp, tx)
		if err != nil {
			return err
		}
		resultCastRegObjekPajak := resultRegObjekPajak.(rp.OKSimple).Data.(rop.RegObjekPajak)
		respDataOp = resultRegObjekPajak

		// add static field
		register.JenisPajak = mt.JenisPajakSA
		register.Status = nt.StatusAktif
		register.User_Id = user_IdConv
		register.RegObjekPajak_Id = resultCastRegObjekPajak.Id
		register.VerifyStatus = rn.VerifyStatusBaru
		register.Rekening = rekening
		register.TanggalMulaiUsaha = th.ParseTime(*input.TanggalMulaiUsaha)
		register.LainLain = sh.GetArrayPhoto(input.LainLain)
		register.SuratIzinUsaha = sh.GetArrayPhoto(input.SuratIzinUsaha)
		register.FotoObjek = sh.GetArrayPhoto(input.FotoObjek)

		err = tx.Create(&register).Error
		if err != nil {
			return err
		}
		err = insertDetailOp(*rekening.Objek, input.DetailRegOp, &register, tx)
		if err != nil {
			return err
		}

		// set directur value to null if golongan orang pribadi
		if input.Golongan == 1 {
			for k := range dataPemilik {
				dataPemilik[k].Direktur_Nama = nil
				dataPemilik[k].Direktur_Alamat = nil
				dataPemilik[k].Direktur_Daerah_Id = nil
				dataPemilik[k].Direktur_Kelurahan_Id = nil
				dataPemilik[k].Direktur_Nik = nil
				dataPemilik[k].Direktur_Telp = nil
			}
		}

		resultRegPemilik, err := rsp.Create(dataPemilik, register.Id, tx)
		if err != nil {
			return err
		}
		respDataPemilik = resultRegPemilik

		resultRegNarahubung, err := rsn.Create(dataNarahubung, register.Id, tx)
		if err != nil {
			return err
		}
		respDataNarahubung = resultRegNarahubung

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), register)
	}

	resp = t.II{
		"objekPajak": respDataOp.(rp.OKSimple).Data,
		"regNpwpd":   register,
		"pemilikWp":  respDataPemilik.(rp.OKSimple).Data,
		"narahubung": respDataNarahubung.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil
}

func GetList(input rn.FilterDto) (interface{}, error) {
	var data []rn.RegNpwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&rn.RegNpwpd{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Rekening").
		Preload("RegObjekPajak").
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
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
	var register *rn.RegNpwpd
	err := a.DB.Model(&rn.RegNpwpd{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Rekening").
		Preload("RegObjekPajak").
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
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

func VerifyNpwpd(id int, input rn.VerifikasiDto) (any, error) {
	if input.VerifyStatus > 2 {
		return nil, errors.New("status tidak diketahui")
	}

	var data *rn.RegNpwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}
	if data.VerifyStatus != rn.VerifyStatusBaru {
		if data.VerifyStatus == rn.VerifyStatusDisetujui {
			return nil, errors.New("data sudah disetujui")
		} else if data.VerifyStatus == rn.VerifyStatusDitolak {
			return nil, errors.New("data sudah ditolak")
		}
	}

	if input.VerifyStatus == 2 {
		data.VerifyStatus = rn.VerifyStatusDitolak
		if result := a.DB.Save(&data); result.Error != nil {
			return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data registrasi npwpd", data)
		}
		return rp.OK{
			Meta: t.IS{
				"affected": strconv.Itoa(int(result.RowsAffected)),
			},
			Data: data,
		}, nil
	}
	var dataROP *rop.RegObjekPajak
	result = a.DB.Where(rop.RegObjekPajak{Id: data.RegObjekPajak_Id}).First(&dataROP)
	if result.RowsAffected == 0 {
		return nil, errors.New("data reg objek pajak tidak dapat ditemukan")
	}

	var dataObjekPajak op.ObjekPajak
	if err := sc.Copy(&dataObjekPajak, dataROP); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload objek pajak", dataROP)
	}

	dataObjekPajak.Id = 0
	err := a.DB.Create(&dataObjekPajak).Error
	if err != nil {
		return nil, err
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload registrasi npwpd", data)
	}

	data.VerifiedAt = th.TimeNow()
	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data registrasi npwpd", data)
	}

	var dataNpwpd nm.Npwpd
	if err := sc.Copy(&dataNpwpd, &data); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload npwpd", data)
	}

	//rekening
	var rekening *rm.Rekening
	err = a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	// kecamatan_id from regobjekpajak
	var dataRegOp *rop.RegObjekPajak
	err = a.DB.Where(rop.RegObjekPajak{Id: data.RegObjekPajak_Id}).First(&dataRegOp).Error
	if err != nil {
		return nil, err
	}

	//creating npwpd
	var tmpNomor = generateNomor()
	tmpNpwpd := sn.GenerateNpwpd(tmpNomor, *dataRegOp.Kecamatan_Id, *rekening.KodeJenisUsaha)

	dataNpwpd.Nomor = tmpNomor
	dataNpwpd.Npwpd = &tmpNpwpd

	//tanggal
	dataNpwpd.TanggalPengukuhan = th.TimeNow()
	dataNpwpd.TanggalNpwpd = th.TimeNow()
	dataNpwpd.Id = 0
	dataNpwpd.ObjekPajak_Id = dataObjekPajak.Id
	err = a.DB.Create(&dataNpwpd).Error
	if err != nil {
		return nil, err
	}

	//transfer detail from reg

	switch *rekening.Objek {
	case "01":
		var dataReg []*rn.DetailRegObjekPajakHotel
		result := a.DB.Where(rn.DetailRegObjekPajakHotel{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak hotel tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakHotel
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail objek pajak hotel", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}

	case "02":
		var dataReg []*rn.DetailRegObjekPajakResto
		result := a.DB.Where(rn.DetailRegObjekPajakResto{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak resto tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakResto
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail reg objek pajak resto", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "03":
		var dataReg []*rn.DetailRegObjekPajakHiburan
		result := a.DB.Where(rn.DetailRegObjekPajakHiburan{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak hiburan tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakHiburan
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail reg objek pajak hiburan", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "04":
		var dataReg []*rn.DetailRegObjekPajakReklame
		result := a.DB.Where(rn.DetailRegObjekPajakReklame{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak reklame tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakReklame
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail reg objek pajak reklame", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "05":
		var dataReg []*rn.DetailRegObjekPajakPpj
		result := a.DB.Where(rn.DetailRegObjekPajakPpj{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak ppj tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakPpj
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail reg objek pajak ppj", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "07":
		var dataReg []*rn.DetailRegObjekPajakParkir
		result := a.DB.Where(rn.DetailRegObjekPajakParkir{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak parkir tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakParkir
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail reg objek pajak parkir", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	case "08":
		var dataReg []*rn.DetailRegObjekPajakAirTanah
		result := a.DB.Where(rn.DetailRegObjekPajakAirTanah{DetailRegObjekPajak: rn.DetailRegObjekPajak{RegNpwpd_Id: uint64(id)}}).Find(&dataReg)
		if result.RowsAffected == 0 {
			return nil, errors.New("data detail reg objek pajak air tanah tidak dapat ditemukan")
		}

		for _, v := range dataReg {
			var dataOp nm.DetailObjekPajakAirTanah
			if err := sc.Copy(&dataOp, v); err != nil {
				return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload detail reg objek pajak air tanah", v)
			}

			dataOp.Npwpd_Id = dataNpwpd.Id
			dataOp.Id = 0
			err = a.DB.Create(&dataOp).Error
			if err != nil {
				return nil, err
			}
		}
	}

	var dataRP []*rn.RegPemilikWp
	result = a.DB.Where(rn.RegPemilikWp{RegNpwpd_Id: uint64(id)}).Find(&dataRP)
	if result.RowsAffected == 0 {
		return nil, errors.New("data reg pemilik tidak dapat ditemukan")
	}

	for _, v := range dataRP {
		var dataP nm.PemilikWp
		if err := sc.Copy(&dataP, v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg pemilik", v)
		}

		dataP.Npwpd_Id = dataNpwpd.Id
		dataP.Id = 0
		err = a.DB.Create(&dataP).Error
		if err != nil {
			return nil, err
		}
	}

	var dataRN []*rn.RegNarahubung
	result = a.DB.Where(rn.RegNarahubung{RegNpwpd_Id: uint64(id)}).Find(&dataRN)
	if result.RowsAffected == 0 {
		return nil, errors.New("data reg narahubung tidak dapat ditemukan")
	}

	for _, v := range dataRN {
		var dataN nm.Narahubung
		if err := sc.Copy(&dataN, v); err != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload reg narahubung", v)
		}

		dataN.Npwpd_Id = dataNpwpd.Id
		dataN.Id = 0
		err = a.DB.Create(&dataN).Error
		if err != nil {
			return nil, err
		}
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: dataNpwpd,
	}, nil
}

func Delete(id int) (any, error) {
	var status string = "deleted"
	var data *rn.RegNpwpd
	var rekening *rm.Rekening

	// cek data regNpwpd
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	// untuk mengambil data regobjekpajak
	tmpObjekPajakId := data.RegObjekPajak_Id

	// cek data rekening untuk mengambil data objek untuk menghapus detail objek pajak
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}
	err = a.DB.Transaction(func(tx *gorm.DB) error {

		// cek dan hapus data pemilik
		err := rsp.Delete(data.Id, tx)
		if err != nil {
			status = "no deletion"
			return err
		}

		// cek dan hapus data narahubung
		err = rsn.Delete(data.Id, tx)
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

		// delete data regNpwpd
		result = tx.Delete(&data, id)
		if result.RowsAffected == 0 {
			return errors.New("tidak dapat menghapus data registrasi npwpd")
		}

		// delete reg objek pajak
		err = rsop.Delete(tmpObjekPajakId, tx)
		if err != nil {
			status = "no deletion"
			return err
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "delete-data", source, "failed", "gagal menghapus data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"count":  strconv.Itoa(int(result.RowsAffected)),
			"status": status,
		},
		Data: data,
	}, nil

}

func Update(id int, input rn.UpdateDto) (any, error) {
	var data *rn.RegNpwpd
	var dataNarahubung []rn.RegNarahubungUpdateDto
	var dataPemilik []rn.RegPemilikWpUpdateDto
	var dataObjekPajak rop.RegObjekPajakUpdateDto
	var dataDetailRegObjekPajak []rn.DetailRegObjekPajak
	var respDataObjekPajak interface{}
	var respDataPemilik interface{}
	var respDataNarahubung interface{}
	var resp t.II
	var rekening *rm.Rekening

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	//cek rekening objek, dgn mengambil data rekening menggunakan rekening id
	err := a.DB.Model(&rm.Rekening{}).First(&rekening, data.Rekening_Id).Error
	if err != nil {
		return nil, err
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload registrasi npwpd", data)
	}

	if err := sc.Copy(&dataNarahubung, &input.RegNarahubung); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg narahubung", data)
	}

	if err := sc.Copy(&dataPemilik, &input.RegPemilik); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg pemilik", data)
	}

	if err := sc.Copy(&dataObjekPajak, &input.RegObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload reg objek pajak", data)
	}

	if err := sc.Copy(&dataDetailRegObjekPajak, &input.DetailRegObjekPajak); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload detail reg objek pajak", data)
	}

	err = a.DB.Transaction(func(tx *gorm.DB) error {

		if result := tx.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data registrasi npwpd")
		}

		resultRegPemilik, err := rsp.Update(dataPemilik, data.Id, data.Golongan, tx)
		if err != nil {
			return err
		}
		respDataPemilik = resultRegPemilik

		resultRegNarahubung, err := rsn.Update(dataNarahubung, data.Id, tx)
		if err != nil {
			return err
		}
		respDataNarahubung = resultRegNarahubung

		resultRegObjekPajak, err := rsop.Update(dataObjekPajak, data.RegObjekPajak_Id, tx)
		if err != nil {
			return err
		}
		respDataObjekPajak = resultRegObjekPajak

		err = updateDetailObjekPajak(dataDetailRegObjekPajak, data.Id, *rekening.Objek, *rekening.Nama, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp = t.II{
		"affected":   strconv.Itoa(int(result.RowsAffected)),
		"objekPajak": respDataObjekPajak.(rp.OKSimple).Data,
		"regNpwpd":   data,
		"pemilikWp":  respDataPemilik.(rp.OKSimple).Data,
		"narahubung": respDataNarahubung.(rp.OKSimple).Data,
	}

	return rp.OKSimple{
		Data: resp,
	}, nil

}

func GetListForWp(input rn.FilterDto) (any, error) {
	var data []rn.RegNpwpd
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&rn.RegNpwpd{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Preload(clause.Associations).
		Preload("User").
		Preload("Rekening").
		Preload("RegObjekPajak").
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
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

func GetDetailForWp(id int, user_Id uint64) (interface{}, error) {
	var data *rn.RegNpwpd
	err := a.DB.Model(&rn.RegNpwpd{}).
		Preload(clause.Associations).
		Preload("User").
		Preload("Rekening").
		Preload("RegObjekPajak").
		Preload("RegObjekPajak.Kecamatan").
		Preload("RegObjekPajak.Kelurahan").
		First(&data, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if data.User_Id != user_Id {
		return nil, errors.New("tidak dapat melihat data yang bukan milik anda")
	}
	return rp.OKSimple{
		Data: data,
	}, err
}

func UpdateForWp(id int, input rn.UpdateDto, user_Id uint) (any, error) {
	var dataCheck *rn.RegNpwpd
	result := a.DB.First(&dataCheck, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if dataCheck.User_Id != userIdConv {
		return nil, errors.New("tidak dapat merubah data yang bukan milik anda")
	}

	resultUpdate, err := Update(id, input)
	return resultUpdate, err
}

func DeleteForWp(id int, user_Id uint) (any, error) {
	var dataCheck *rn.RegNpwpd
	result := a.DB.First(&dataCheck, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data registrasi npwpd tidak dapat ditemukan")
	}

	userIdConv := uint64(user_Id)
	if dataCheck.User_Id != userIdConv {
		return nil, errors.New("tidak dapat menghapus data yang bukan milik anda")
	}

	resultDelete, err := Delete(id)
	return resultDelete, err
}

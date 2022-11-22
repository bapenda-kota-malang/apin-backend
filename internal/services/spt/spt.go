package spt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	mrek "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mnomertracker "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/sptnomertracker"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/google/uuid"

	srek "github.com/bapenda-kota-malang/apin-backend/internal/services/configuration/rekening"
	snomertracker "github.com/bapenda-kota-malang/apin-backend/internal/services/spt/sptnomertracker"
	suser "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "spt"

func filePreProcess(b64String, docsname string, userId uint, oldId uuid.UUID) (fileName, path, extFile string, id uuid.UUID, err error) {
	extFile, err = base64helper.GetExtensionBase64(b64String)
	if err != nil {
		return
	}
	path = sh.GetResourcesPath()
	switch extFile {
	case "pdf":
		path = sh.GetPdfPath()
	case "png", "jpeg":
		path = sh.GetImgPath()
	case "xlsx", "xls":
		path = sh.GetExcelPath()
	default:
		err = errors.New("file tidak diketahui")
		return
	}
	if oldId == uuid.Nil {
		id, err = sh.GetUuidv4()
		if err != nil {
			err = errors.New("gagal generate uuid")
			return
		}
	} else {
		id = oldId
	}
	fileName = sh.GenerateFilename(docsname, id, userId, extFile)
	return
}

func Create(input m.CreateDto, opts map[string]interface{}, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	// get data rekening
	respRek, err := srek.GetDetail(int(*input.Rekening_Id))
	if err != nil {
		return nil, err
	}
	dataRekening := respRek.(rp.OKSimple).Data.(*mrek.Rekening)

	// if kodeJenisPajak nil then search data rekening from parent id
	kodeJenisPajak := dataRekening.KodeJenisPajak
	if kodeJenisPajak == nil {
		parentId, err := strconv.Atoi(*dataRekening.Parent_Id)
		if err != nil {
			return nil, err
		}
		respRek, err := srek.GetDetail(parentId)
		if err != nil {
			return nil, err
		}
		dataRekening = respRek.(rp.OKSimple).Data.(*mrek.Rekening)
		kodeJenisPajak = dataRekening.KodeJenisPajak
	}

	// if mark as new file not clone from esptpd then save new file
	if opts["newFile"].(bool) {
		var errChan = make(chan error)
		fileName, path, extFile, id, err := filePreProcess(input.Lampiran, "Lampiran"+opts["baseUri"].(string), opts["userId"].(uint), uuid.Nil)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(input.Lampiran, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save pdf", data)
		}
		input.Lampiran = fileName
		data.Id = id
	}

	// if gambar reklame not nil save image
	if input.Gambar != nil {
		errCh := make(chan error)
		fileName, path, extFile, _, err := filePreProcess(*input.Gambar, "GambarReklame"+opts["baseUri"].(string), opts["userId"].(uint), data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.SaveFile(*input.Gambar, fileName, path, extFile, errCh)
		if err := <-errCh; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "failed save pdf", data)
		}
		input.Gambar = &fileName
	}

	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	if data.TanggalSpt.IsZero() {
		data.TanggalSpt = time.Now()
	}
	yearNow := uint(data.TanggalSpt.Year())

	nomerUrut, err := snomertracker.TrxNewNumber(mnomertracker.Dto{Tahun: &yearNow, KodeJenisPajak: kodeJenisPajak}, tx)
	if err != nil {
		return nil, err
	}
	kode := *kodeJenisPajak
	yearTwoDigit := yearNow % 1e2
	nomerSpt := fmt.Sprintf("%d%s", yearTwoDigit, nomerUrut)
	data.NomorSpt = fmt.Sprintf("%c-%s", kode[0], nomerSpt)
	data.KodeBilling = fmt.Sprintf("%s%s", *dataRekening.KodeBilling, nomerSpt)
	switch data.Type {
	case mtypes.JenisPajakOA, mtypes.JenisPajakSA:
	default:
		data.Type = mtypes.JenisPajakSA
	}
	data.StatusPembayaran = m.StatusBelumLunas
	if data.CreateBy_User_Id == 0 {
		data.CreateBy_User_Id = opts["userId"].(uint)
	}

	periodeVal, _ := data.PeriodeAwal.Value()
	periodeTime := periodeVal.(time.Time)
	if periodeTime.IsZero() {
		periodeTime = sh.BeginningOfPreviosMonth()
	} else {
		periodeTime = sh.BeginningOfMonth(periodeTime)
	}

	jatuhTempoVal, _ := data.JatuhTempo.Value()
	jatuhTempoTime := jatuhTempoVal.(time.Time)
	if jatuhTempoTime.IsZero() {
		jatuhTempoTime = time.Now()
	}

	data.PeriodeAwal = datatypes.Date(periodeTime)
	data.PeriodeAkhir = datatypes.Date(sh.EndOfMonth(periodeTime))
	data.JatuhTempo = datatypes.Date(sh.EndOfMonth(jatuhTempoTime))

	err = tx.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto, userId uint, cmdBase string) (any, error) {
	var data []m.ListDataDto
	var count int64
	var pagination gh.Pagination
	baseQuery := a.DB.Model(&m.Spt{})
	if userId != 0 {
		baseQuery.Joins("JOIN \"Npwpd\" ON \"Spt\".\"Npwpd_Id\" = \"Npwpd\".\"Id\" AND \"Npwpd\".\"User_Id\" = ?", userId)
	}
	opts := "gte"
	dataDateNow := datatypes.Date(time.Now())
	if input.JatuhTempo != nil {
		val, _ := input.JatuhTempo.Value()
		endOfMonthDate := datatypes.Date(sh.EndOfMonth(val.(time.Time)))
		input.JatuhTempo = &endOfMonthDate
		opts = "="
	} else {
		input.JatuhTempo = &dataDateNow
	}
	stringJoin := "FULL JOIN \"DetailTbp\" ON \"DetailTbp\".\"Spt_Id\" = \"Spt\".\"Id\""
	statusAll := "Baru"
	if input.StatusData != nil {
		stringJoin = "LEFT JOIN \"DetailTbp\" ON \"DetailTbp\".\"Spt_Id\" = \"Spt\".\"Id\""
		statusData := *input.StatusData
		if cmdBase == "wp" {
			switch statusData {
			case m.TbpStatusFilterBaru, m.TbpStatusFilterPembayaran, m.TbpStatusFilterLunas, m.TbpStatusFilterJatuhTempo:
			default:
				return sh.SetError("request", "get-data-list", source, "failed", "status data invalid", data)
			}
		}
		switch statusData {
		case m.TbpStatusFilterBaru:
			baseQuery = baseQuery.Where("\"StatusPembayaran\" = ?", m.StatusBelumLunas)
		case m.TbpStatusFilterPembayaran:
			stringJoin = "JOIN \"DetailTbp\" ON \"DetailTbp\".\"Spt_Id\" = \"Spt\".\"Id\" AND \"DetailTbp\".\"NominalBayar\" <> '0' OR \"DetailTbp\".\"NominalBayar\" IS NULL"
			statusAll = "Pembayaran"
		case m.TbpStatusFilterPenyetoran:
			// TODO: Wait for STS
			// stringJoin = "JOIN \"DetailTbp\" ON \"DetailTbp\".\"Spt_Id\" = \"Spt\".\"Id\" AND \"DetailTbp\".\"NominalBayar\""
			// statusAll = "Penyetoran"
		case m.TbpStatusFilterLunas:
			stringJoin = "JOIN \"DetailTbp\" ON \"DetailTbp\".\"Spt_Id\" = \"Spt\".\"Id\" AND \"DetailTbp\".\"NominalBayar\" = '0'"
			statusAll = "Lunas"
		case m.TbpStatusFilterJatuhTempo:
			opts = "lt"
			stringJoin = "JOIN \"DetailTbp\" ON \"DetailTbp\".\"Spt_Id\" = \"Spt\".\"Id\" AND \"DetailTbp\".\"NominalBayar\" <> '0' OR \"DetailTbp\".\"NominalBayar\" IS NULL"
			statusAll = "Jatuh Tempo"
		case m.TbpStatusFilterPenetapan:
			baseQuery = baseQuery.Where("\"StatusPenetapan\" = ?", mtypes.StatusVerifikasiDisetujuiKabid)
			statusAll = "Penetapan"
		}
		input.StatusData = nil
	}
	input.JatuhTempo_Opt = &opts
	result := baseQuery.
		Select("\"Spt\".*, \"DetailTbp\".\"NominalBayar\"").
		Preload("Rekening").
		Preload("ObjekPajak").
		Preload("Npwpd").
		Joins(stringJoin).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	for v := range data {
		status := statusAll
		if statusAll == "" {
			checkBaru := data[v].StatusPembayaran == m.StatusBelumLunas
			checkLunas := data[v].NominalBayar != nil && *data[v].NominalBayar == 0
			checkPenetapan := data[v].StatusPenetapan == mtypes.StatusVerifikasiDisetujuiKabid
			sqlTime, _ := data[v].JatuhTempo.Value()
			checkDueDate := sqlTime.(time.Time).Unix() < time.Now().Unix()

			if checkBaru {
				status = "Baru"
			} else if checkPenetapan && cmdBase == "bapenda" {
				status = "Penetapan"
			} else if checkLunas {
				status = "Lunas"
			} else {
				switch cmdBase {
				case "bapenda":
					status = "Pembayaran"
				case "wp":
					status = "Belum Lunas"
				}
			}
			if checkDueDate && !checkLunas {
				status = "Jatuh Tempo"
			}
		}
		data[v].StatusFinal = &status
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

func GetDetail(id uuid.UUID, typeSpt string, userId uint) (any, error) {
	var data *m.Spt

	baseQuery := a.DB.Model(&m.Spt{})
	if userId != 0 {
		baseQuery.Joins("JOIN \"Npwpd\" ON \"Spt\".\"Npwpd_Id\" = \"Npwpd\".\"Id\" AND \"Npwpd\".\"User_Id\" = ?", userId)
	}
	if typeSpt == string(mtypes.JenisPajakSA) || typeSpt == string(mtypes.JenisPajakOA) {
		baseQuery.Where("\"Spt\".\"Type\" = ?", typeSpt)
	}
	result := baseQuery.
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		First(&data, "\"Spt\".\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id uuid.UUID, input m.UpdateDto, opts map[string]interface{}, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.Spt
	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := tx.First(&data, "\"Id\" = ?", id.String()).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	// copy to model struct
	if input.Lampiran != nil {
		var errChan = make(chan error)
		fileName, path, extFile, _, err := filePreProcess(*input.Lampiran, "Lampiran"+opts["baseUri"].(string), opts["userId"].(uint), data.Id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
		}
		go sh.ReplaceFile(data.Lampiran, *input.Lampiran, fileName, path, extFile, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "update-data", source, "failed", fmt.Sprintf("failed save file: %s", err), data)
		}
		input.Lampiran = &fileName
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(dataRow)),
		},
		Data: data,
	}, nil
}

func Verify(id uuid.UUID, input m.VerifyDto, userId uint) (any, error) {
	var data *m.Spt
	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	dataRow := a.DB.First(&data, "\"Id\" = ? AND \"Type\" = ?", id.String(), mtypes.JenisPajakOA).RowsAffected
	if dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if data.StatusPenetapan == mtypes.StatusVerifikasiDisetujuiKabid {
		return sh.SetError("request", "update-data", source, "failed", "data telah disetujui", data)
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	resp, err := suser.GetJabatanPegawai(userId)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal data pegawai: "+err.Error(), data)
	}
	jabatan := strings.ToUpper(resp.(string))
	userRole := ""
	if kasubid := strings.Contains(jabatan, "KEPALA SUB BIDANG"); kasubid {
		data.Kasubid_User_Id = &userId
		userRole = "kasubid"
	} else if kabid := strings.Contains(jabatan, "KEPALA BIDANG"); kabid {
		data.Kabid_User_Id = &userId
		userRole = "kabid"
	}
	if userRole == "" {
		return sh.SetError("request", "update-data", source, "failed", "pegawai bukan kabid atau kasubid", data)
	}
	switch input.StatusPenetapan {
	case "disetujui":
		if userRole == "kasubid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDisetujuiKasubid
		} else if userRole == "kabid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDisetujuiKabid
		}
	case "ditolak":
		if userRole == "kasubid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDitolakKasubid
		} else if userRole == "kabid" {
			data.StatusPenetapan = mtypes.StatusVerifikasiDitolakKabid
		}
	default:
		return sh.SetError("request", "update-data", source, "failed", "status tidak diketahui", data)
	}
	switch data.StatusPenetapan {
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

func Delete(id uuid.UUID) (any, error) {
	//data spt
	var data *m.Spt
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	status := "deleted"

	result = a.DB.Delete(&data, id)
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

func SkpdkbExisting(input m.SkpdkbExisting, opts map[string]interface{}) (any, error) {
	var newDataInput m.Input
	var existingData m.Spt
	var createdNewData m.Spt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// get old data from db
		dataRow := tx.Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).First(&existingData, "\"Id\" = ?", input.Spt_Id.String()).RowsAffected
		if dataRow == 0 {
			return errors.New("data tidak dapat ditemukan")
		}

		// casting to dto
		if existingData.DetailSptAir != nil {
			newDataInput = &m.CreateDetailAirDto{}
		} else if existingData.DetailSptHiburan != nil {
			newDataInput = &m.CreateDetailHiburanDto{}
		} else if existingData.DetailSptHotel != nil {
			newDataInput = &m.CreateDetailHotelDto{}
		} else if existingData.DetailSptParkir != nil {
			newDataInput = &m.CreateDetailParkirDto{}
		} else if existingData.DetailSptNonPln != nil {
			newDataInput = &m.CreateDetailPpjNonPlnDto{}
		} else if existingData.DetailSptPln != nil {
			newDataInput = &m.CreateDetailPpjPlnDto{}
		} else if existingData.DetailSptReklame != nil {
			newDataInput = &m.CreateDetailReklameDto{}
		} else if existingData.DetailSptResto != nil {
			newDataInput = &m.CreateDetailRestoDto{}
		} else {
			newDataInput = &m.CreateDetailBaseDto{}
		}
		// copy data from existing
		if err := newDataInput.SkpdkbDuplicate(&existingData, &input); err != nil {
			return err
		}

		// calculate skpdkb process
		newDataInput.CalculateSkpdkb()
		// create new data
		respNewData, err := CreateDetail(newDataInput, opts, tx)
		if err != nil {
			return err
		}
		createdNewData = respNewData.(rp.OKSimple).Data.(m.Spt)

		// change existing data value
		// existingData.Ref_Spt_Id = &createdNewData.Id
		// TODO: FLAG SKPDKB
		// update existing data
		if result := tx.Save(&existingData); result.Error != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-skpdkb-existing", source, "failed", "transaction skpd existing: "+err.Error(), createdNewData)
	}
	return rp.OKSimple{Data: createdNewData}, nil
}

func SkpdkbNew(input m.Input, opts map[string]interface{}) (any, error) {
	var createdData m.Spt
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// save new data
		respNewData, err := CreateDetail(input, opts, tx)
		if err != nil {
			return err
		}
		createdData = respNewData.(rp.OKSimple).Data.(m.Spt)
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-skpdkb-existing", source, "failed", "transaction skpd existing: "+err.Error(), createdData)
	}
	return rp.OKSimple{Data: createdData}, nil
}

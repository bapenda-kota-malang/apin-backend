package spt

import (
	"context"
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
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	ibj "github.com/bapenda-kota-malang/apin-backend/pkg/integration/bankjatim"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sc "github.com/jinzhu/copier"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "spt"

func Create(ctx context.Context, input m.CreateDto, opts map[string]interface{}, tx *gorm.DB) (any, error) {
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
	for kodeJenisPajak == nil {
		if dataRekening.Parent_Id == nil {
			return nil, errors.New("unknown kode jenis pajak")
		}
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

	data.Total = data.JumlahPajak

	if data.TanggalSpt.IsZero() {
		data.TanggalSpt = time.Now()
	}
	yearNow := uint(data.TanggalSpt.Year())

	nomerUrut, err := snomertracker.TrxNewNumber(mnomertracker.Dto{Tahun: &yearNow, KodeJenisPajak: kodeJenisPajak}, tx)
	if err != nil {
		return nil, err
	}
	kode := strings.ToUpper(*kodeJenisPajak)
	yearTwoDigit := yearNow % 1e2
	nomerSpt := fmt.Sprintf("%d%s", yearTwoDigit, nomerUrut)
	data.NomorSpt = fmt.Sprintf("%c-%s", kode[0], nomerSpt)

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

	// set jenis ketetapan
	if data.DasarPengenaan == nil {
		data.JenisKetetapan = m.JenisKetetapanSptpd
		if data.Type == mtypes.JenisPajakOA {
			data.JenisKetetapan = m.JenisKetetapanSkpd
		}
	} else {
		// normalize dasar pengenaan to uppercase
		dasarPengenaan := strings.ToUpper(*data.DasarPengenaan)
		data.DasarPengenaan = &dasarPengenaan
		switch dasarPengenaan {
		case m.DasarPengenaanSptpd, m.DasarPengenaanSkpd, m.DasarPengenaanTeguran:
			data.JenisKetetapan = m.JenisKetetapanSkpdkb
		case m.DasarPengenaanSkpdkb, m.DasarPengenaanSkpdkbt:
			data.JenisKetetapan = m.JenisKetetapanSkpdkbt
		}
	}

	switch data.JenisKetetapan {
	case m.JenisKetetapanSptpd:
		data.KodeBilling = generateKodeBilling(dataRekening.KodeBilling, nomerSpt)
		va, err := vaManager(ctx, data, ibj.ProsesCreate)
		if err != nil {
			return nil, err
		}
		data.VaJatim = &va
	case m.JenisKetetapanSkpdkb, m.JenisKetetapanSkpdkbt:
		data.KodeBilling = generateKodeBilling(dataRekening.KodeBilling, nomerSpt)
		_, err := vaManager(ctx, data, ibj.ProsesUpdate)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto, userId uint, cmdBase string, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data []m.ListDataDto
	var count int64
	var pagination gh.Pagination
	baseQuery := tx.Model(&m.Spt{})
	if userId != 0 {
		baseQuery.Joins("JOIN \"Npwpd\" ON \"Spt\".\"Npwpd_Id\" = \"Npwpd\".\"Id\" AND \"Npwpd\".\"User_Id\" = ?", userId)
	}
	// jatuh tempo query
	opts := "gte"
	if input.JatuhTempo != nil {
		val, _ := input.JatuhTempo.Value()
		endOfMonthDate := datatypes.Date(sh.EndOfMonth(val.(time.Time)))
		input.JatuhTempo = &endOfMonthDate
		opts = "="
	}

	if input.JenisKetetapan_Opt != nil && *input.JenisKetetapan_Opt == "unfilter skpdkb" {
		baseQuery.Where("\"JenisKetetapan\" = ? AND \"JenisKetetapan\" = ?", m.JenisKetetapanSkpdkb, m.JenisKetetapanSkpdkbt)
		input.JenisKetetapan = nil
		input.JenisKetetapan_Opt = nil
	}

	statusAll := ""
	if input.StatusData != nil {
		statusData := *input.StatusData
		// if from wp and status data not in baru, pembayaran, lunas, jatuh tempo then error
		if cmdBase == "wp" {
			switch statusData {
			case m.TbpStatusFilterBaru, m.TbpStatusFilterPembayaran, m.TbpStatusFilterLunas, m.TbpStatusFilterJatuhTempo:
				// do nothing
			default:
				return sh.SetError("request", "get-data-list", source, "failed", "status data invalid", data)
			}
		}
		switch statusData {
		case m.TbpStatusFilterBaru:
			baseQuery = baseQuery.Where("\"StatusPembayaran\" = ?", m.StatusBelumLunas)
			statusAll = "Baru"
		case m.TbpStatusFilterPembayaran:
			baseQuery = baseQuery.Where("\"StatusPembayaran\" LIKE ?", "1%")
			statusAll = "Pembayaran"
		case m.TbpStatusFilterPenyetoran:
			baseQuery = baseQuery.Where("\"StatusPembayaran\" = ?", m.StatusPenyetoran)
			statusAll = "Penyetoran"
		case m.TbpStatusFilterLunas:
			baseQuery = baseQuery.Where("\"StatusPembayaran\" = ?", m.StatusLunas)
			statusAll = "Lunas"
		case m.TbpStatusFilterJatuhTempo:
			opts = "lt"
			listData := []string{
				string(m.StatusBelumLunas),
				string(m.StatusKurangBayar),
				string(m.StatusKurangBayarAngsuran),
			}
			baseQuery = baseQuery.Where("\"StatusPembayaran\" IN ?", listData)
			statusAll = "Jatuh Tempo"
		case m.TbpStatusFilterPenetapan:
			baseQuery = baseQuery.Where("\"StatusPenetapan\" = ?", mtypes.StatusVerifikasiDisetujuiKabid)
			statusAll = "Penetapan"
		}
		input.StatusData = nil
	}
	input.JatuhTempo_Opt = &opts
	result := baseQuery.Debug().
		Select("\"Spt\".*").
		Joins("Rekening").
		Preload("Rekening").
		Preload("ObjekPajak").
		Preload("Npwpd.PemilikWps").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Order("\"JatuhTempo\" DESC").
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}

	for v := range data {
		status := statusAll
		if statusAll == "" {
			switch data[v].StatusPembayaran {
			case m.StatusBelumLunas:
				status = "Baru"
			case m.StatusLunas:
				status = "Lunas"
			case m.StatusKurangBayar, m.StatusKurangBayarAngsuran:
				switch cmdBase {
				case "bapenda":
					status = "Pembayaran"
				case "wp":
					status = "Belum Lunas"
				}
			case m.StatusLebihBayar:
				switch cmdBase {
				case "bapenda":
					status = "Pembayaran"
				case "wp":
					status = "Lebih Bayar"
				}
			case m.StatusPenyetoran:
				status = "Penyetoran"
			}

			if data[v].StatusPenetapan == mtypes.StatusVerifikasiDisetujuiKabid && cmdBase == "bapenda" {
				status = "Penetapan"
			}

			// jika sudah lewat masa jatuh tempo dan belum lunas maka jatuh tempo
			sqlTime, _ := data[v].JatuhTempo.Value()
			checkDueDate := sqlTime.(time.Time).Unix() < time.Now().Unix()
			if checkDueDate && !(data[v].StatusPembayaran == m.StatusLunas) {
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
		Preload("Npwpd").
		Preload("Npwpd.PemilikWps").
		Preload("ObjekPajak.Kecamatan").
		Preload("ObjekPajak.Kelurahan").
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("DetailSptReklame.TarifReklame.KlasifikasiJalan").
		Preload("DetailSptPln.JenisPPJ").
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

func GetByVa(va *string) (m.Spt, error) {
	var data m.Spt
	if err := a.DB.Preload("Npwpd.ObjekPajak").Where("VaJatim", &va).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
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

// Update status pembayaran berdasarkan jumlah pajak - nominal bayar yang diterima
func UpdateStatus(id uuid.UUID, nominalBayar float64, tx *gorm.DB) error {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	if dataRow := tx.First(&data, "\"Id\" = ?", id.String()).RowsAffected; dataRow == 0 {
		return errors.New("data tidak dapat ditemukan")
	}

	denda := float64(0)
	if data.Denda != nil {
		denda = *data.Denda
	}

	sisa := (data.JumlahPajak + denda) - nominalBayar

	if sisa == 0 {
		data.StatusPembayaran = m.StatusLunas
	} else if sisa > 0 {
		data.StatusPembayaran = m.StatusKurangBayar
	} else {
		data.StatusPembayaran = m.StatusLebihBayar
	}

	if result := tx.Save(&data); result.Error != nil {
		return result.Error
	}

	return nil
}

// update status pembayaran menjadi penyetoran (40)
func UpdatePenyetoran(id uuid.UUID, tx *gorm.DB) error {
	if tx == nil {
		tx = a.DB
	}
	var data m.Spt

	if dataRow := tx.First(&data, "\"Id\" = ?", id.String()).RowsAffected; dataRow == 0 {
		return errors.New("data tidak dapat ditemukan")
	}
	data.StatusPembayaran = m.StatusPenyetoran
	if result := tx.Save(&data); result.Error != nil {
		return result.Error
	}
	return nil
}

// update virtual account bank jatim
func UpdateVa(ctx context.Context, id uuid.UUID, input m.UpdateVaDto, userId uint64) (any, error) {
	var data m.Spt
	// validate data exist and copy input (payload) ke struct data jika tidak ada akan error
	if dataRow := a.DB.Preload("Npwpd").Preload("ObjekPajak").First(&data, "\"Id\" = ?", id.String()).RowsAffected; dataRow == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}
	if data.VaJatim == nil {
		return nil, errors.New("data spt ini masih belum ada virtual account")
	}
	data.JumlahPajak = *input.JumlahPajak

	data.Total = data.JumlahPajak
	data.Denda = input.Denda
	if data.Denda != nil {
		data.Total = data.JumlahPajak + *data.Denda
	}

	_, err := vaManager(ctx, data, ibj.ProsesUpdate)
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengubah va: "+err.Error(), data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data", data)
	}
	return rp.OK{
		Meta: t.IS{
			"affected": "1",
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

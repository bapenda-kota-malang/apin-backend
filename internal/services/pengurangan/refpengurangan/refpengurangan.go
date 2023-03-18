package refpengurangan

import (
	"errors"
	"strconv"
	"sync"
	"time"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pengurangan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const source = "refpengurangan"

func Create(input m.CreateDtoRefPengurangan, user_Id uint64) (any, error) {
	var (
		dataRefPengurangan m.RefPengurangan
		errChan            = make(chan error, 4)
		wg                 sync.WaitGroup
	)

	// TODO: FILTER USER ID TO EXECUTE THIS FUNCTION
	// resp, err := suser.GetJabatanPegawai(uint(user_Id))
	// if err != nil {
	// 	return sh.SetError("request", "create-data", source, "failed", "gagal data pegawai: "+err.Error(), nil)
	// }
	// if ok := strings.Contains(strings.ToUpper(resp.(string)), "PETUGAS"); !ok {
	// 	return sh.SetError("request", "create-data", source, "failed", "pegawai bukan petugas, sehingga tidak bisa membuat data baru", nil)
	// }

	var dataSpt spt.Spt
	resCheckSpt := a.DB.First(&dataSpt, "\"Id\" = ?", input.Spt_Id.String())
	if resCheckSpt.RowsAffected == 0 {
		return nil, errors.New("tidak ditemukan data spt")
	}

	var dataExisting *m.RefPengurangan
	if res := a.DB.Model(m.RefPengurangan{}).Where("\"Spt_Id\" = ?", input.Spt_Id.String()).First(&dataExisting); res.RowsAffected == 1 {
		if dataExisting.Pengurangan_Id != nil {
			return nil, errors.New("sudah pernah mengajukan pengurangan yang telah diproses staff")
		}
		if dataExisting.Status != m.StatusDitolak {
			return nil, errors.New("ditemukan data spt sama yang masih dalam proses")
		}
	}

	id, err := sh.GetUuidv4()
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat uuid: "+err.Error(), nil)
	}

	// foto ktp
	fileName, path, extFile, _, err := sh.FilePreProcess(input.FotoKtp, source+"FotoKtp", uint(user_Id), id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}
	wg.Add(1)
	go sh.BulkSaveFile(&wg, input.FotoKtp, fileName, path, extFile, errChan)
	input.FotoKtp = fileName

	// laporan keuangan
	fileName, path, extFile, _, err = sh.FilePreProcess(input.LaporanKeuangan, source+"LaporanKeuangan", uint(user_Id), id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}
	wg.Add(1)
	go sh.BulkSaveFile(&wg, input.LaporanKeuangan, fileName, path, extFile, errChan)
	input.LaporanKeuangan = fileName

	// Surat Permohonan
	fileName, path, extFile, _, err = sh.FilePreProcess(input.SuratPermohonan, source+"SuratPermohonan", uint(user_Id), id)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
	}
	wg.Add(1)
	go sh.BulkSaveFile(&wg, input.SuratPermohonan, fileName, path, extFile, errChan)
	input.SuratPermohonan = fileName

	// dokumen lainnya
	if input.DokumenLainnya != nil {
		fileName, path, extFile, _, err = sh.FilePreProcess(*input.DokumenLainnya, source+"DokumenLainnya", uint(user_Id), id)
		if err != nil {
			return sh.SetError("request", "create-data", source, "failed", err.Error(), nil)
		}
		wg.Add(1)
		go sh.BulkSaveFile(&wg, *input.DokumenLainnya, fileName, path, extFile, errChan)
		input.DokumenLainnya = &fileName
	}

	// data pengurangan
	if err := sc.Copy(&dataRefPengurangan, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload ref pengurangan", dataRefPengurangan)
	}

	wg.Wait()
	close(errChan)
	for v := range errChan {
		if v != nil {
			err = v
			return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan file: "+err.Error(), dataRefPengurangan)
		}
	}

	// static value
	dataRefPengurangan.Id = id
	dataRefPengurangan.Pemohon_User_Id = user_Id
	dataRefPengurangan.Status = m.StatusDiproses
	now := time.Now()
	dataRefPengurangan.TanggalPengajuan = now
	dataRefPengurangan.NominalKetetapan_Spt = dataSpt.JumlahPajak
	// create data pengurangan
	err = a.DB.Create(&dataRefPengurangan).Error
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), dataRefPengurangan)
	}

	return rp.OKSimple{
		Data: dataRefPengurangan,
	}, nil
}

func GetList(input m.FilterDtoRefPengurangan, userid int) (any, error) {
	var data []m.RefPengurangan
	var count int64
	var pagination gh.Pagination

	baseQuery := a.DB.
		Model(&m.RefPengurangan{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination))

	if userid != 0 {
		baseQuery.Where("Pemohon_User_Id", userid)
	}

	result := baseQuery.Find(&data)
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

func GetDetail(id uuid.UUID, userid uint) (any, error) {
	var data *m.RefPengurangan

	baseQuery := a.DB.Model(&m.RefPengurangan{}).
		Preload(clause.Associations, func(tx *gorm.DB) *gorm.DB {
			return tx.Omit("Password")
		}).
		Preload("Spt.Npwpd.ObjekPajak")

	if userid != 0 {
		baseQuery.Where("Pemohon_User_Id", userid)
	}

	result := baseQuery.First(&data, "\"Id\" = ?", id.String())
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data RefPengurangan", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id uuid.UUID, input m.UpdateDto, userId uint64) (any, error) {
	//ambil data based on id
	var data *m.RefPengurangan

	if result := a.DB.Preload("Pemohon").First(&data, "\"Id\" = ?", id.String()); result.RowsAffected == 0 {
		return nil, nil
	}
	//ambil nama jabatan based on userId
	// resp, err := suser.GetJabatanPegawai(uint(userId))
	// if err != nil {
	// 	return sh.SetError("request", "verify-data", source, "failed", "gagal data pegawai: "+err.Error(), data)
	// }
	// jabatan := strings.ToUpper(resp.(string))
	if err := sc.CopyWithOption(&data, &input, sc.Option{IgnoreEmpty: true}); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload ref pengurangan", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "verify-data", source, "failed", "gagal menyimpan data RefPengurangan: "+result.Error.Error(), data)
	}
	return rp.OKSimple{
		Data: data,
	}, nil
}

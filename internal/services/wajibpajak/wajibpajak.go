package wajibpajak

import (
	"errors"
	"strconv"
	"time"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	mad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/wajibpajak"
	su "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

const source = "wajibpajak"

func Create(input m.RegistrasiWajibPajak) (any, error) {
	var data m.WajibPajak
	var dataU mu.CreateDto
	var respDataU interface{}
	var imgNameChan = make(chan string)
	var errChan = make(chan error)

	go sh.SaveImage(input.WajibPajak.FotoKtp, imgNameChan, errChan)

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input.WajibPajak); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	if err := sc.Copy(&dataU, &input.User); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// if result := a.DB.First(&mad.Kecamatan{}, data.Kecamatan_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }
	// if result := a.DB.First(&mad.Kelurahan{}, data.Kelurahan_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }
	// if result := a.DB.First(&mad.Daerah{}, data.Kota_id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }
	// if result := a.DB.First(&mad.Provinsi{}, data.Provinsi_Id); result.RowsAffected == 0 {
	// 	return nil, nil
	// }

	if err := <-errChan; err != nil {
		return sh.SetError("request", "create-data", source, "failed", "image unsupported", data)
	}
	data.FotoKtp = <-imgNameChan
	data.Email = dataU.Email

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return result.Error
		}

		// add static field
		dataU.Ref_Id = int(data.Id)
		dataU.Position = 3
		dataU.Status = 0
		dataU.RegMode = 2
		dataU.ValidPeriod = time.Now().AddDate(20, 0, 0)

		tmpResp, err := su.Create(dataU, tx)
		if err != nil {
			return err
		}
		respDataU = tmpResp

		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	resp := t.II{
		"wajibPajak": data,
		"user":       respDataU.(rp.OKSimple).Data,
	}

	return rp.OKSimple{Data: resp}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.WajibPajak
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.WajibPajak{}).
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

func GetDetail(id int) (any, error) {
	var data *m.WajibPajak

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto) (any, error) {
	var data *m.WajibPajak
	var imgNameChan = make(chan string)
	var errChan = make(chan error)
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if input.FotoKtp != nil {
		go sh.ReplaceImage(data.FotoKtp, *input.FotoKtp, imgNameChan, errChan)
		if err := <-errChan; err != nil {
			return sh.SetError("request", "create-data", source, "failed", "image unsupported", data)
		}
		data.FotoKtp = <-imgNameChan
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.First(&mad.Kecamatan{}, data.Kecamatan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Kelurahan{}, data.Kelurahan_Id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Daerah{}, data.Kota_id); result.RowsAffected == 0 {
		return nil, nil
	}
	if result := a.DB.First(&mad.Provinsi{}, data.Provinsi_Id); result.RowsAffected == 0 {
		return nil, nil
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

func CheckerPOne(input m.CheckerPOneDto) (interface{}, error) {
	var data m.WajibPajak
	if result := a.DB.Where(&m.WajibPajak{Nik: input.Nik}).First(&data); result.RowsAffected != 0 {
		return nil, errors.New("nik telah terdaftar")
	}
	return rp.OKSimple{
		Data: input,
	}, nil
}

func CheckerPTwo(input m.CheckerPTwoDto) (interface{}, error) {
	if result := a.DB.First(&mad.Kecamatan{}, input.Kecamatan_Id); result.RowsAffected == 0 {
		return nil, errors.New("kecamatan tidak ditemukan")
	}
	if result := a.DB.First(&mad.Kelurahan{}, input.Kelurahan_Id); result.RowsAffected == 0 {
		return nil, errors.New("kelurahan tidak ditemukan")
	}
	if result := a.DB.First(&mad.Daerah{}, input.Kota_id); result.RowsAffected == 0 {
		return nil, errors.New("daerah tidak ditemukan")
	}
	if result := a.DB.First(&mad.Provinsi{}, input.Provinsi_Id); result.RowsAffected == 0 {
		return nil, errors.New("provinsi tidak ditemukan")
	}
	return rp.OKSimple{
		Data: input,
	}, nil
}

func CheckerPFour(input m.CheckerPFourDto) (interface{}, error) {
	if err := sh.CheckImage(input.FotoKtp); err != nil {
		return nil, err
	}
	return rp.OKSimple{
		Data: nil,
	}, nil
}

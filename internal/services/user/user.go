package userservice

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	sc "github.com/jinzhu/copier"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/excelhelper"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	p "github.com/bapenda-kota-malang/apin-backend/pkg/password"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/usertoken"
)

const source = "user"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.User

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}
	password, err := p.Hash(*data.Password)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat password", data)
	} else {
		data.Password = &password
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data user: "+result.Error.Error(), data)
	}

	data.Password = nil

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.User
	var count int64
	var pagination gh.Pagination

	result := a.DB.
		Model(&m.User{}).
		Omit("Password").
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	}
	for i := range data {
		data[i].Password = nil
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

func DownloadExcelList(input m.FilterDto) (*excelize.File, error) {
	var data []m.User

	result := a.DB.
		Model(&m.User{}).
		Omit("Password").
		Scopes(gh.Filter(input)).
		Find(&data)
	if result.Error != nil {
		_, err := sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
		return nil, err
	}

	var excelData = func() []interface{} {
		var tmp []interface{}
		tmp = append(tmp, map[string]interface{}{
			"A": "No",
			"B": "Username",
			"C": "Email",
			"D": "Nama",
			"E": "NIK",
			"F": "Tgl Daftar",
			"G": "Status",
		})
		for i, v := range data {
			n := i + 1
			tmp = append(tmp, map[string]interface{}{
				"A": n,
				"B": v.Name,
				"C": v.Email,
				"D": v.Name,
				"E": "",
				"F": func() string {
					return v.CreatedAt.Format("2006-01-02")
				}(),
				"G": func() string {
					if v.Status == 0 {
						return "Baru"
					} else if v.Status == 1 {
						return "Aktif"
					} else {
						return "Tidak Diketahui"
					}
				}(),
			})
		}
		return tmp
	}()
	return excelhelper.ExportList(excelData, "List")
}

func GetDetail(id int) (any, error) {
	var data *m.User

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
	var data *m.User
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+result.Error.Error(), data)
	}
	data.Password = nil

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int) (any, error) {
	var data *m.User
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("data tidak dapat ditemukan")
	}

	result = a.DB.Delete(&data, id)
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

func CheckerPThree(input m.CheckerPThreeDto) (interface{}, error) {
	var data m.User
	if result := a.DB.Unscoped().Where(&m.User{Name: input.Name}).Or(&m.User{Email: input.Email}).First(&data); result.RowsAffected != 0 {
		return nil, errors.New("username atau email telah terdaftar")
	}
	return rp.OKSimple{
		Data: input,
	}, nil
}

func Verifikasi(id int, input m.VerifikasiDto) (any, error) {
	var data *m.User
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

func GetJabatanPegawai(userId uint) (any, error) {
	var data string
	res := a.DB.
		Model(m.User{}).
		Select("\"Jabatan\".\"Nama\"").
		Joins("JOIN \"Pegawai\" ON \"User\".\"Ref_Id\" = \"Pegawai\".\"Id\"").
		Joins("JOIN \"Jabatan\" ON \"Pegawai\".\"Jabatan_Id\" = \"Jabatan\".\"Id\"").
		First(&data, userId)
	if res.Error != nil {
		return sh.SetError("request", "get-data", source, "failed", res.Error.Error(), data)
	}
	return data, nil
}

func ChangePass(id int, input m.ChangePassDto) (any, error) {
	// TODO: PINDAH KE VALIDATOR
	if *input.NewPassword != *input.RePassword {
		return nil, errors.New("password baru dan konfirmasi tidak sama")
	}

	var data *m.User
	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data)
	} else if !p.Check(*input.OldPassword, *data.Password) {
		return nil, errors.New("password lama tidak sesuai")
	}

	password, err := p.Hash(*input.NewPassword)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat password", data)
	} else {
		data.Password = &password
	}

	if result := a.DB.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+result.Error.Error(), data)
	}
	data.Password = nil

	return data, nil
}

func RequestResetPass(input m.RequestResetPassDto) (any, error) {
	var dataUser *m.User
	var dataUserToken *mt.UserToken

	result := a.DB.Where(&m.User{Email: input.Email}).First(&dataUser)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "create-data", source, "failed", "user dengan email terkait tidak dapat ditemukan", dataUser)
	}

	token, err := uuid.NewRandom()
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat token", dataUserToken)
	}

	expiredAt := time.Now().Local().Add(time.Hour * time.Duration(1))
	result = a.DB.Where(&mt.UserToken{User_Email: input.Email, Jenis: mt.JenisResetPass}).First(&dataUserToken)
	if result.RowsAffected == 0 {
		dataUserToken.User_Email = input.Email
		dataUserToken.Jenis = mt.JenisResetPass
		dataUserToken.Token = token
		dataUserToken.ExpiredAt = &expiredAt
		if result := a.DB.Create(&dataUserToken); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data token user: "+result.Error.Error(), dataUserToken)
		}
		// TODO : REMOVE TOKEN
		return rp.OKSimple{Data: dataUserToken}, nil
	} else if result.Error == nil {
		if err := checkExpiration(*dataUserToken.ExpiredAt); err == nil {
			return nil, errors.New("permintaan reset pasword sidah pernah dilakukan dan masih berlaku")
		}
		dataUserToken.Token = token
		dataUserToken.ExpiredAt = &expiredAt
		if result := a.DB.Save(&dataUserToken); result.Error != nil {
			return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data token user: "+result.Error.Error(), dataUserToken)
		}
		// TODO : REMOVE TOKEN
		return rp.OKSimple{Data: dataUserToken}, nil
	} else {
		return sh.SetError("request", "create-data", source, "failed", "gagal mendapatkan data token user", dataUserToken)
	}
}

func CheckResetPass(input m.CheckResetPassDto) (any, error) {
	var data *mt.UserToken
	result := a.DB.Where(&mt.UserToken{User_Email: input.Email, Jenis: mt.JenisResetPass}).First(&data)
	if result.RowsAffected == 0 {
		return nil, errors.New("data request resest password tidak dapat ditemukan")
	} else if data.Token.String() != input.Token {
		return nil, errors.New("token tidak valid")
	} else if err := checkExpiration(*data.ExpiredAt); err != nil {
		return nil, err
	}

	return data, nil
}

func ResetPass(input1 m.CheckResetPassDto, input2 m.ResetPassDto) (any, error) {
	var data1 *mt.UserToken
	result := a.DB.Where(&mt.UserToken{User_Email: input1.Email, Jenis: mt.JenisResetPass}).First(&data1)
	if result.RowsAffected == 0 {
		return nil, errors.New("data request resest password tidak dapat ditemukan")
	} else if err := checkExpiration(*data1.ExpiredAt); err != nil {
		return nil, err
	} else if data1.Token.String() != input1.Token {
		return nil, errors.New("token tidak valid")
	}

	// TODO: PINDAH KE VALIDATOR
	// fmt.Println(input2.NewPassword)
	// fmt.Println(input2.RePassword)
	if *input2.NewPassword != *input2.RePassword {
		return nil, errors.New("password baru dan konfirmasi tidak sama")
	}

	var data2 *m.User
	result = a.DB.Where(&m.User{Email: input1.Email}).First(&data2)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data", data2)
	}

	password, err := p.Hash(*input2.NewPassword)
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal membuat password", data2)
	} else {
		data2.Password = &password
	}

	if result := a.DB.Save(&data2); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data: "+result.Error.Error(), data2)
	}

	a.DB.Delete(data1)
	data2.Password = nil
	return data2, nil

}

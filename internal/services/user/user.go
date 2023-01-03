package userservice

import (
	"errors"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	p "github.com/bapenda-kota-malang/apin-backend/pkg/password"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
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
	if *input.OldPassword != *input.RePassword {
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

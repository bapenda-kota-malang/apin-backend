// service
package pegawai

import (
	"errors"
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	su "github.com/bapenda-kota-malang/apin-backend/internal/services/user"
)

// /// Private funcs start here
const source = "pegawai"

///// Exported funcs start here

func Create(input m.CreateDto) (any, error) {
	var data m.Pegawai
	var dataU mu.CreateDto
	var dataUX any

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		// simpan data pegawai ke db satu if karena result dipakai sekali, +error
		if result := tx.Create(&data); result.Error != nil {
			return errors.New("penyimpanan data pegawai gagal")
		}

		// simpan data user melalui user service
		dataU.Name = input.User_Name
		dataU.Password = &input.User_Password
		dataU.Email = input.User_Email
		dataU.Notes = input.User_Notes
		dataU.Group_Id = input.User_Group_Id
		dataU.SysAdmin = input.User_SysAdmin
		dataU.Position = 1
		dataU.Ref_Id = data.Id
		dataU.RegMode = 1
		dataU.Status = 1
		dataUXTemp, err := su.Create(dataU, tx)
		if err != nil {
			return err
		}
		dataUX = dataUXTemp

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", err.Error(), data)
	}

	dataUP := dataUX.(rp.OKSimple)

	dataU.Password = nil

	return rp.OKSimple{Data: t.II{
		"pegawai": data,
		"user":    dataUP.Data,
	}}, nil
}

// NOTE: this is special since pegawai is not related to user by structure
func GetList(input m.FilterDto) (any, error) {
	var data []m.OutputDto
	var count int64

	var pagination gh.Pagination

	rows, err := a.DB.Table("\"Pegawai\" \"P\"").
		Scopes(gh.Filter(input)).
		Joins("LEFT JOIN \"User\" \"U\" ON \"P\".\"Id\"=\"U\".\"Ref_Id\" AND \"U\".\"Position\"=1").
		Joins("LEFT JOIN \"Group\" \"G\" ON \"U\".\"Group_Id\"=\"G\".\"Id\"").
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Select("\"P\".\"Id\", \"P\".\"Nama\", \"P\".\"Nip\", " +
			"\"P\".\"Jabatan_Id\", \"P\".\"BidangKerja_Kode\", \"P\".\"StartDate\", \"P\".\"EndDate\", \"P\".\"Aktif\"," +
			"\"U\".\"Name\", \"U\".\"Email\", \"U\".\"SysAdmin\", \"U\".\"Group_Id\", \"G\".\"Name\"").
		Rows()
	if err != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data: "+err.Error(), data)
	}
	defer rows.Close()

	data = make([]m.OutputDto, 0)
	row := m.OutputDto{}
	for rows.Next() {
		rows.Scan(&row.Id, &row.Nama, &row.Nip,
			&row.Jabatan_Id, &row.BidangKerja_Kode, &row.StartDate, &row.EndDate, &row.Aktif,
			&row.User_Name, &row.User_Email, &row.User_SysAdmin, &row.User_Group_Id, &row.User_Group_Name)
		data = append(data, row)
	}

	// sql := a.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Table("\"Pegawai\" \"P\"", "P").
	// 		Scopes(gh.Filter(input)).
	// 		Joins("LEFT JOIN \"User\" \"U\" ON \"P\".\"Id\"=\"U\".\"Ref_Id\" AND \"U\".\"Position\"=1").
	// 		// Count(&count).
	// 		Scopes(gh.Paginate(input, &pagination)).
	// 		Select("\"P\".\"Id\", \"P\".\"Nama\", \"P\".\"Nip\", " +
	// 			"\"P\".\"Jabatan_Id\", \"P\".\"Pangkat_Id\", \"P\".\"SatuanKerja_Id\", \"P\".\"StartDate\", \"P\".\"EndDate\", \"P\".\"Aktif\"," +
	// 			"\"U\".\"Name\", \"U\".\"Email\"").
	// 		// Rows()
	// 		Find(&m.Pegawai{})
	// })
	// fmt.Println(sql)

	// result := a.DB.
	// 	Find(&data)

	// if result.Error != nil {
	// 	return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data", data)
	// }

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(len(data))),
			"page":         strconv.Itoa(pagination.Page),
			"pageSize":     strconv.Itoa(pagination.PageSize),
		},
		Data: data,
	}, nil
}

func GetDetail(id int) (interface{}, error) {
	var data *m.OutputDto = &m.OutputDto{}
	row := a.DB.Table("\"Pegawai\" \"P\"").
		Joins("LEFT JOIN \"User\" \"U\" ON \"P\".\"Id\"=\"U\".\"Ref_Id\" AND \"U\".\"Position\"=1").
		Joins("LEFT JOIN \"Group\" \"G\" ON \"U\".\"Group_Id\"=\"G\".\"Id\"").
		Where("\"P\".\"Id\"", id).
		Select("\"P\".\"Id\", \"P\".\"Nama\", \"P\".\"Nip\", " +
			"\"P\".\"Jabatan_Id\", \"P\".\"BidangKerja_Kode\", \"P\".\"StartDate\", \"P\".\"EndDate\", \"P\".\"Aktif\"," +
			"\"U\".\"Name\", \"U\".\"Email\", \"U\".\"SysAdmin\", \"U\".\"Notes\", \"U\".\"Group_Id\", \"G\".\"Name\"").
		Row()
	row.Scan(&data.Id, &data.Nama, &data.Nip,
		&data.Jabatan_Id, &data.BidangKerja_Kode, &data.StartDate, &data.EndDate, &data.Aktif,
		&data.User_Name, &data.User_Email, &data.User_SysAdmin, &data.User_Notes, &data.User_Group_Id, &data.User_Group_Name)
	// if result.RowsAffected == 0 {
	// 	return nil, nil
	// } else if result.Error != nil {
	// 	return sh.SetError("request", "get-data", source, "failed", "gagal mengambil data", data)
	// }

	// sql := a.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Table("\"Pegawai\" \"P\"", "P").
	// 		Table("\"Pegawai\" \"P\"").
	// 		Joins("LEFT JOIN \"User\" \"U\" ON \"P\".\"Id\"=\"U\".\"Ref_Id\" AND \"U\".\"Position\"=1").
	// 		Where("\"P\".\"Id\"", id).
	// 		Select("\"P\".\"Id\", \"P\".\"Nama\", \"P\".\"Nip\", " +
	// 			"\"P\".\"Jabatan_Id\", \"P\".\"BidangKerja_Kode\", \"P\".\"StartDate\", \"P\".\"EndDate\", \"P\".\"Aktif\"," +
	// 			"\"U\".\"Name\" \"User_Name\", \"U\".\"Email\" \"User_Email\"").
	// 		Find(&m.Pegawai{})
	// })
	// fmt.Println(sql)

	return rp.OKSimple{Data: data}, nil
}

func GetFromUserid(userId uint64) (m.Pegawai, error) {
	var data m.Pegawai
	if err := a.DB.Select("\"Pegawai\".*").Joins("JOIN \"User\" ON \"User\".\"Position\" = 1 AND \"User\".\"Ref_Id\" = \"Pegawai\".\"Id\" AND \"User\".\"Id\" = ?", userId).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func Update(id int, input m.UpdateDto) (interface{}, error) {
	var data *m.Pegawai
	var dataU *mu.User

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data pegawai tidak dapat ditemukan", input)
	}
	result = a.DB.Where(mu.User{Ref_Id: data.Id, Position: 1}).First(&dataU)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data user tidak dapat ditemukan", input)
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk pegawai", input)
	}
	// if err := sc.Copy(&dataU, &input); err != nil {
	// 	return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk user", input)
	// }

	rowsAffected := 0
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := a.DB.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data pegawai")
		}
		if result.RowsAffected > 0 {
			rowsAffected++
		}

		dataU.Group_Id = input.User_Group_Id
		dataU.SysAdmin = input.User_SysAdmin
		dataU.Email = input.User_Email
		dataU.Notes = input.User_Notes
		if result := a.DB.Save(&dataU); result.Error != nil {
			return errors.New("gagal menyimpan data user")
		}
		if result.RowsAffected > 0 {
			rowsAffected++
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", err.Error(), input)
	}

	dataU.Password = nil

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowsAffected),
		},
		Data: t.II{
			"pegawai": data,
			"user":    dataU,
		},
	}, nil
}

func UpdateNew(id int, input m.UpdateDto) (interface{}, error) {
	var data *m.Pegawai = &m.Pegawai{}
	var dataU *mu.User

	result := a.DB.Where(mu.User{Id: id, Position: 1}).First(&dataU)
	if result.RowsAffected == 0 {
		return sh.SetError("request", "update-data", source, "failed", "data user tidak dapat ditemukan", input)
	}
	if err := sc.Copy(&data, &input); err != nil {
		fmt.Println(err)
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk pegawai", input)
	}
	if err := sc.Copy(&dataU, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload untuk user", input)
	}

	rowsAffected := 0
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		if result := a.DB.Save(&data); result.Error != nil {
			return errors.New("gagal menyimpan data pegawai")
		}
		if result.RowsAffected > 0 {
			rowsAffected++
		}
		dataU.Ref_Id = data.Id
		if result := a.DB.Save(&dataU); result.Error != nil {
			return errors.New("gagal menyimpan data user")
		}
		if result.RowsAffected > 0 {
			rowsAffected++
		}
		return nil
	})
	if err != nil {
		return sh.SetError("request", "update-data", source, "failed", err.Error(), input)
	}

	dataU.Password = nil

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(rowsAffected),
		},
		Data: t.II{
			"pegawai": data,
			"user":    dataU,
		},
	}, nil
}

func Delete(id int) (interface{}, error) {
	var data *m.Pegawai
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

package jaminanbongkarreklame

// import (
// 	ms "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
// 	msjbr "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/jaminanbongkarreklame"
// 	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
// 	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
// 	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
// 	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
// 	th "github.com/bapenda-kota-malang/apin-backend/pkg/timehelper"
// 	sc "github.com/jinzhu/copier"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// const source = "jaminanBongkarReklame"

// func init() {
// 	a.AutoMigrate(&ms.Spt{})
// 	a.AutoMigrate(&msjbr.JaminanBongkarReklame{})
// }

// func Create(input ms.CreateJaminanDto) (any, error) {
// 	var dataS ms.Spt
// 	var dataD msjbr.JaminanBongkarReklame

// 	if err := sc.Copy(&dataS, input); err != nil {
// 		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataS)
// 	}
// 	if err := sc.Copy(&dataD, input); err != nil {
// 		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataD)
// 	}
// 	dataD.Date = th.ParseTime(input.Date)
// 	dataD.DueDate = th.ParseTime(input.DueDate)

// var dataJ msjbr.JaminanBongkarReklame

// 	if err := sc.Copy(&dataJ, input); err != nil {
// 		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", dataJ)
// 	}
// 	dataJ.Date = th.ParseTime(input.JaminanBongkarReklame.Date)
// 	dataJ.DueDate = th.ParseTime(input.JaminanBongkarReklame.DueDate)

// 	dataJ.Spt_Id = dataS.Id

// 	err = a.DB.Create(&dataJ).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	var dataResult msjbr.JaminanBongkarReklame
// 	err = a.DB.Model(&msjbr.JaminanBongkarReklame{}).
// 		Preload(clause.Associations).First(&dataResult, dataD.Id).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	err := a.DB.Create(&dataS).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	dataD.Spt_Id = dataS.Id

// 	err = a.DB.Create(&dataD).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	var dataResult msjbr.JaminanBongkarReklame
// 	err = a.DB.Model(&msjbr.JaminanBongkarReklame{}).
// 		Preload(clause.Associations).First(&dataResult, dataD.Id).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return rp.OKSimple{Data: t.II{
// 		"spt":                   dataS,
// 		"jaminanBongkarReklame": dataResult,
// 	}}, nil
// }

package objekpajakbumi

import (
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	d "github.com/bapenda-kota-malang/apin-backend/internal/models/dokument"
	fasbng "github.com/bapenda-kota-malang/apin-backend/internal/models/fasilitasbangunan"
	bng "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi"
	msppt "github.com/bapenda-kota-malang/apin-backend/internal/models/sppt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "objekpajakbumi"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data m.ObjekPajakBumi

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	if result := tx.Create(&data); result.Error != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil menyimpan data objekpajakbumi", data)
	}

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.ObjekPajakBumi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.ObjekPajakBumi{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data objekpajakbumi", data)
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
	var data *m.ObjekPajakBumi

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data objekpajakbumi", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	var dataFasBangunan fasbng.FasilitasBangunan
	var dataBng bng.CreateDto
	var dataBangunan bng.ObjekPajakBangunan
	if tx == nil {
		tx = a.DB
	}
	var data m.ObjekPajakBumi
	// var dataNop pel.PermohonanNOP
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	err := a.DB.Transaction(func(tx *gorm.DB) error {

		if result := tx.Save(&data); result.Error != nil {
			return result.Error
		}

		nobangunan := 0
		tempKodeFas := ""
		if input.ObjekPajakBangunans != nil {
			if input.Nop != nil {
				if *input.Nop != "" {
					dataNop := *DecodeNOP(input.Nop)

					resBang := tx.
						Where("Provinsi_Kode", &dataNop.PermohonanProvinsiID).
						Where("Daerah_Kode", &dataNop.PermohonanKotaID).
						Where("Kecamatan_Kode", &dataNop.PermohonanKecamatanID).
						Where("Kelurahan_Kode", &dataNop.PermohonanKelurahanID).
						Where("Blok_Kode", &dataNop.PermohonanBlokID).
						Where("NoUrut", &dataNop.NoUrutPemohon).
						Where("JenisOp", &dataNop.PemohonJenisOPID).
						First(&dataBangunan)

					var errTemp error
					tempRegOpBng := *input.ObjekPajakBangunans
					for idx, opb := range tempRegOpBng {

						nobangunan = 000 //ngarang

						dataBng.FasilitasBangunans = nil
						dataBng.Provinsi_Kode = &dataNop.PermohonanProvinsiID
						dataBng.Daerah_Kode = &dataNop.PermohonanKotaID
						dataBng.Kecamatan_Kode = &dataNop.PermohonanKecamatanID
						dataBng.Kelurahan_Kode = &dataNop.PermohonanKelurahanID
						dataBng.Blok_Kode = &dataNop.PermohonanBlokID
						dataBng.NoUrut = &dataNop.NoUrutPemohon
						dataBng.JenisOp = &dataNop.PemohonJenisOPID

						dataBng.Area_Kode = input.Area_Kode
						dataBng.NoBangunan = &nobangunan

						dataBng.JenisAtap = opb.JenisAtap
						dataBng.JenisKonstruksi = opb.JenisKonstruksi
						dataBng.JenisTransaksi = opb.JenisTransaksi
						// dataBng.JmlLantaiBng = opb.JmlLantaiBng
						dataBng.Jpb_Kode = opb.Jpb_Kode
						dataBng.KodeDinding = opb.KodeDinding
						dataBng.KodeLangitLangit = opb.KodeLangitLangit
						dataBng.KodeLantai = opb.KodeLantai
						dataBng.Kondisi = opb.Kondisi
						dataBng.LuasBangunan = opb.LuasBangunan
						dataBng.NilaiSistem = opb.NilaiSistem
						dataBng.NoFormulirSpop = opb.NoFormulirSpop
						dataBng.Nop = opb.Nop
						dataBng.TahunDibangun = opb.TahunDibangun
						dataBng.TahunRenovasi = opb.TahunRenovasi

						if err := sc.Copy(&dataBangunan, &dataBng); err != nil {
							return err
						}

						if resBang.RowsAffected == 0 {
							//create data bangunan
							if result := tx.Save(&dataBangunan); result.Error != nil {
								return result.Error
							}
							dataBangunan.Id = dataBangunan.Id + 1
						}
						if result := tx.Create(&dataBangunan); result.Error != nil {
							return result.Error
						}

						if opb.RegFasBangunan != nil {
							val := opb.RegFasBangunan
							var fas fasbng.FasilitasBangunan

							nobangunan := dataBangunan.NoBangunan

							resFas := tx.
								Where("Provinsi_Kode", &dataNop.PermohonanProvinsiID).
								Where("Daerah_Kode", &dataNop.PermohonanKotaID).
								Where("Kecamatan_Kode", &dataNop.PermohonanKecamatanID).
								Where("Kelurahan_Kode", &dataNop.PermohonanKelurahanID).
								Where("Blok_Kode", &dataNop.PermohonanBlokID).
								Where("NoUrut", &dataNop.NoUrutPemohon).
								Where("JenisOp", &dataNop.PemohonJenisOPID).
								First(&dataFasBangunan)

							if resFas.RowsAffected == 0 {
								_ = a.DB.Order("\"Id\" DESC").First(&fas)
								if fas.Id != 0 && idx == 0 {
									dataFasBangunan.Id = fas.Id + 1
								}
							}

							dataFasBangunan.Provinsi_Kode = input.Provinsi_Kode
							dataFasBangunan.Daerah_Kode = input.Daerah_Kode
							dataFasBangunan.Kecamatan_Kode = input.Kecamatan_Kode
							dataFasBangunan.Kelurahan_Kode = input.Kelurahan_Kode
							dataFasBangunan.NoUrut = input.NoUrut
							dataFasBangunan.JenisOp = input.JenisOp
							dataFasBangunan.NoBangunan = nobangunan

							//create data fasilitas/jpb berdasar fasilitas/jpb di isi
							if val.FBJumlahACSplit != nil {
								tempKodeFas = "01"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBJumlahACSplit
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBJumlahACSplit
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Create(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBJumlahACWindow != nil {
								tempKodeFas = "02"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBJumlahACWindow
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBJumlahACWindow
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBIsACCentral != nil {
								tempKodeFas = "11"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBIsACCentral
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBIsACCentral
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTipeLapisanKolam != nil {
								tempKodeFas = strconv.Itoa(*val.FBTipeLapisanKolam)
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									if val.FBLuasKolamRenang != nil {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.FBLuasKolamRenang
									}
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								if val.FBLuasKolamRenang != nil {
									dataFasBang.JumlahSatuan = val.FBLuasKolamRenang
								}
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBHalamanBerat != nil {
								tempKodeFas = "16"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBHalamanBerat
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBHalamanBerat
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBHalamanSendang != nil {
								tempKodeFas = "15"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBHalamanSendang
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBHalamanSendang
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBHalamanRingan != nil {
								tempKodeFas = "14"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBHalamanRingan
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBHalamanRingan
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBHalamanLantai != nil {
								tempKodeFas = "17"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBHalamanLantai
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBHalamanLantai
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTenisLampuBeton != nil {
								tempKodeFas = "18"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTenisLampuBeton
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTenisLampuBeton
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTenisTanpaLampuBeton != nil {
								tempKodeFas = "21"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTenisTanpaLampuBeton
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTenisTanpaLampuBeton
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTenisAspal1 != nil {
								tempKodeFas = "19"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTenisAspal1
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTenisAspal1
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTenisAspal2 != nil {
								tempKodeFas = "22"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTenisAspal2
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTenisAspal2
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTenisLiatRumput1 != nil {
								tempKodeFas = "20"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTenisLiatRumput1
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTenisLiatRumput1
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTenisLiatRumput2 != nil {
								tempKodeFas = "23"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTenisLiatRumput2
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTenisLiatRumput2
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBLiftPenumpang != nil {
								tempKodeFas = "30"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBLiftPenumpang
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBLiftPenumpang
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBLiftKapsul != nil {
								tempKodeFas = "31"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBLiftKapsul
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBLiftKapsul
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBLiftBarang != nil {
								tempKodeFas = "32"
								dataFasBangunan.KodeFasilitas = &tempKodeFas
								dataFasBangunan.JumlahSatuan = val.FBLiftBarang
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBangunan).Error
								if errTemp != nil {
									return errTemp
								}
								dataFasBangunan.Id = dataFasBangunan.Id + 1
							}
							if val.FBTangga80 != nil {
								tempKodeFas = "33"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTangga80
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTangga80
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBTangga81 != nil {
								tempKodeFas = "34"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBTangga81
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBTangga81
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBPagarBahan != nil {
								tempKodeFas = strconv.Itoa(*val.FBPagarBahan)
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									if val.FBPagarPanjang != nil {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.FBPagarPanjang
									}
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								if val.FBPagarPanjang != nil {
									dataFasBang.JumlahSatuan = val.FBPagarPanjang
								}
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBPKHydrant != nil {
								tempKodeFas = "37"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBPKHydrant
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBPKHydrant
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBPKSplinkler != nil {
								tempKodeFas = "39"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBPKSplinkler
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBPKSplinkler
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBPKFireAI != nil {
								tempKodeFas = "38"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBPKFireAI
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBPKFireAI
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBPABX != nil {
								tempKodeFas = "41"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBPABX
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBPABX
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}
							if val.FBSumur != nil {
								tempKodeFas = "42"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.FBSumur
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.FBSumur
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}
							}

							if val.JpbKlinikACCentralKamar != nil {
								var dataRJpb bng.Jpb5

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								tempKodeFas = "07"
								var dataFasBang fasbng.FasilitasBangunan
								resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
								if resFasTemp.RowsAffected == 0 {
									dataFasBangunan.KodeFasilitas = &tempKodeFas
									dataFasBangunan.JumlahSatuan = val.JpbKlinikACCentralKamar
									errTemp = tx.Create(&dataFasBangunan).Error
									if errTemp != nil {
										return errTemp
									}
									dataFasBangunan.Id = dataFasBangunan.Id + 1
								}
								dataFasBang.JumlahSatuan = val.JpbKlinikACCentralKamar
								errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
								if errTemp != nil {
									return errTemp
								}

								dataRJpb.LuasKamarAcCentral = val.JpbKlinikACCentralKamar

								if val.JpbKlinikACCentralRuang != nil {
									tempKodeFas = "08"
									var dataFasBang fasbng.FasilitasBangunan
									resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
									if resFasTemp.RowsAffected == 0 {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.JpbKlinikACCentralRuang
										errTemp = tx.Create(&dataFasBangunan).Error
										if errTemp != nil {
											return errTemp
										}
										dataFasBangunan.Id = dataFasBangunan.Id + 1
									}
									dataFasBang.JumlahSatuan = val.JpbKlinikACCentralRuang
									errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
									if errTemp != nil {
										return errTemp
									}

									dataRJpb.LuasRuangLainAcCentral = val.JpbKlinikACCentralRuang
								}

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if val.JpbHotelJenis != nil {
								var dataRJpb bng.Jpb7

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								dataRJpb.JenisHotel = (bng.JenisHotel)(strconv.Itoa(*val.JpbHotelJenis))
								if val.JpbHotelBintang != nil {
									dataRJpb.JumlahBintang = (bng.JumlahBintang)(strconv.Itoa(*val.JpbHotelJenis))
								}
								if val.JpbHotelJmlKamar != nil {
									dataRJpb.JumlahKamar = val.JpbHotelJmlKamar
								}
								if val.JpbHotelACCentralKamar != nil {
									tempKodeFas = "04"
									var dataFasBang fasbng.FasilitasBangunan
									resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
									if resFasTemp.RowsAffected == 0 {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.JpbHotelACCentralKamar
										errTemp = tx.Create(&dataFasBangunan).Error
										if errTemp != nil {
											return errTemp
										}
										dataFasBangunan.Id = dataFasBangunan.Id + 1
									}
									dataFasBang.JumlahSatuan = val.JpbHotelACCentralKamar
									errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
									if errTemp != nil {
										return errTemp
									}

									dataRJpb.LuasKamarAcCentral = val.JpbHotelACCentralKamar
								}
								if val.JpbHotelACCentralRuang != nil {
									tempKodeFas = "05"
									var dataFasBang fasbng.FasilitasBangunan
									resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
									if resFasTemp.RowsAffected == 0 {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.JpbHotelACCentralRuang
										errTemp = tx.Create(&dataFasBangunan).Error
										if errTemp != nil {
											return errTemp
										}
										dataFasBangunan.Id = dataFasBangunan.Id + 1
									}
									dataFasBang.JumlahSatuan = val.JpbHotelACCentralRuang
									errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
									if errTemp != nil {
										return errTemp
									}

									dataRJpb.LuasRuangLainAcCentral = val.JpbHotelACCentralRuang
								}

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if val.JpbApartemenJumlah != nil {
								var dataRJpb bng.Jpb13

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								dataRJpb.JumlahApartment = val.JpbApartemenJumlah

								if val.JpbApartemenACCentralKamar != nil {
									tempKodeFas = "09"
									var dataFasBang fasbng.FasilitasBangunan
									resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
									if resFasTemp.RowsAffected == 0 {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.JpbApartemenACCentralKamar
										errTemp = tx.Create(&dataFasBangunan).Error
										if errTemp != nil {
											return errTemp
										}
										dataFasBangunan.Id = dataFasBangunan.Id + 1
									}
									dataFasBang.JumlahSatuan = val.JpbApartemenACCentralKamar
									errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
									if errTemp != nil {
										return errTemp
									}

									dataRJpb.LuasApartAcCentral = val.JpbHotelACCentralRuang
								}
								if val.JpbApartemenACCentralLain != nil {
									tempKodeFas = "10"
									var dataFasBang fasbng.FasilitasBangunan
									resFasTemp := tx.Where("NoBangunan", &nobangunan).Where("KodeFasilitas", &tempKodeFas).First(&dataFasBang)
									if resFasTemp.RowsAffected == 0 {
										dataFasBangunan.KodeFasilitas = &tempKodeFas
										dataFasBangunan.JumlahSatuan = val.JpbApartemenACCentralLain
										errTemp = tx.Create(&dataFasBangunan).Error
										if errTemp != nil {
											return errTemp
										}
										dataFasBangunan.Id = dataFasBangunan.Id + 1
									}
									dataFasBang.JumlahSatuan = val.JpbApartemenACCentralLain
									errTemp = tx.Where("NoBangunan", nobangunan).Where("KodeFasilitas", tempKodeFas).Save(&dataFasBang).Error
									if errTemp != nil {
										return errTemp
									}

									dataRJpb.LuasRuangLainAcCentral = val.JpbHotelACCentralRuang
								}

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if val.JpbTankiKapasitas != nil {
								var dataRJpb bng.Jpb15

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								dataRJpb.KapasitasTanki = val.JpbTankiKapasitas
								if val.JpbTankiLetak != nil {
									dataRJpb.LetakTanki = (bng.LetakTanki)(*val.JpbTankiLetak)
								}

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if val.JpbProdTinggi != nil {
								if opb.Jpb_Kode == "03" {
									var dataRJpb bng.Jpb3

									resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

									dataRJpb.TinggiKolom3 = val.JpbProdTinggi
									if val.JpbProdLebar != nil {
										dataRJpb.LebarBentang3 = val.JpbProdLebar
									}
									if val.JpbProdDaya != nil {
										dataRJpb.DayaDukungLantai = val.JpbProdDaya
									}
									if val.JpbProdKeliling != nil {
										dataRJpb.KelilingDinding3 = val.JpbProdKeliling
									}
									if val.JpbProdLuas != nil {
										dataRJpb.LuasMezzanine3 = val.JpbProdLuas
									}

									if resJpbTemp.RowsAffected == 0 {
										dataRJpb.NoBangunan = nobangunan
										errTemp = tx.Create(&dataRJpb).Error
										if errTemp != nil {
											return errTemp
										}
									} else {
										errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
										if errTemp != nil {
											return errTemp
										}
									}
								} else if opb.Jpb_Kode == "08" {
									var dataRJpb bng.Jpb8

									resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

									dataRJpb.TinggiKolom8 = val.JpbProdTinggi
									if val.JpbProdLebar != nil {
										dataRJpb.LebarBentang8 = val.JpbProdLebar
									}
									if val.JpbProdDaya != nil {
										dataRJpb.DayaDukungLantai = val.JpbProdDaya
									}
									if val.JpbProdKeliling != nil {
										dataRJpb.KelilingDinding8 = val.JpbProdKeliling
									}
									if val.JpbProdLuas != nil {
										dataRJpb.LuasMezzanine8 = val.JpbProdLuas
									}

									if resJpbTemp.RowsAffected == 0 {
										dataRJpb.NoBangunan = nobangunan
										errTemp = tx.Create(&dataRJpb).Error
										if errTemp != nil {
											return errTemp
										}
									} else {
										errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
										if errTemp != nil {
											return errTemp
										}
									}
								}
							}

							if opb.Jpb_Kode == "00" || opb.Jpb_Kode == "01" || opb.Jpb_Kode == "10" || opb.Jpb_Kode == "11" {
							}

							if opb.Jpb_Kode == "02" {
								var dataRJpb bng.Jpb2

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if opb.Jpb_Kode == "04" {
								var dataRJpb bng.Jpb4

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if opb.Jpb_Kode == "06" {
								var dataRJpb bng.Jpb6

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if opb.Jpb_Kode == "09" {
								var dataRJpb bng.Jpb9

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if opb.Jpb_Kode == "12" {
								var dataRJpb bng.Jpb12

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)
								dataRJpb.TipeBangunan = "1" // 1-4 gak jelas dapat dari mana

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = *nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if opb.Jpb_Kode == "14" {
								var dataRJpb bng.Jpb14

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)
								dataRJpb.LuasKanopi = nil // gak jelas dapat dari mana

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
							if opb.Jpb_Kode == "16" {
								var dataRJpb bng.Jpb16

								resJpbTemp := tx.Where("NoBangunan", &nobangunan).First(&dataRJpb)

								if resJpbTemp.RowsAffected == 0 {
									dataRJpb.NoBangunan = nobangunan
									errTemp = tx.Create(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								} else {
									errTemp = tx.Where("NoBangunan", &nobangunan).Save(&dataRJpb).Error
									if errTemp != nil {
										return errTemp
									}
								}
							}
						}
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal menyimpan data: "+err.Error(), data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func UpdateBulk(input m.UpdateBulkDto) (any, error) {
	var datas []m.ObjekPajakBumi
	var dok d.Dokument

	err := a.DB.Transaction(func(tx *gorm.DB) error {
		resDok := tx.Where("NoDokumen", &input.NomerDokumen).First(&dok)
		if resDok.RowsAffected > 0 {
			_, errTemp := sh.SetError("request", "get-data-dok", source, "failed", "no dokumen telah digunakan", input)
			return errTemp
		}

		errDok := tx.Create(&d.Dokument{
			Kanwil_Kd:      input.Kanwil_Id,
			Kppbb_Kd:       input.Kpbb_Id,
			JenisDokumen:   input.JenisDokumen,
			NoDokumen:      &input.NomerDokumen,
			TglPendataan:   input.TglPendataan,
			NipPendataan:   input.NipPendataan,
			TglPemeriksaan: input.TglPemeriksaan,
			NipPemeriksaan: input.NipPemeriksaan,
		}).Error
		if errDok != nil {
			return errDok
		}

		for _, v := range input.Datas {
			var data = new(m.ObjekPajakBumi)
			checkExist := tx.
				Where("Provinsi_Kode", input.Provinsi_Kode).
				Where("Daerah_Kode", input.Daerah_Kode).
				Where("Kecamatan_Kode", input.Kecamatan_Kode).
				Where("Kelurahan_Kode", input.Kelurahan_Kode).
				Where("Blok_Kode", input.Blok_Kode).
				Where("NoUrut", v.NoUrut).
				Where("JnsOp", v.JenisOp).
				Where("KodeZNT", input.KodeZNT).
				First(&data)
			if checkExist.Error != nil || checkExist.RowsAffected == 0 {
				return checkExist.Error
			} else {
				var cdto m.UpdateDto
				if err := sc.Copy(&cdto, &data); err != nil {
					return err
				}
				cdto.KodeZNT = input.KodeZNTBaru
				respData, err := Update(v.Id, cdto, tx)
				if err != nil {
					return err
				}
				datas = append(datas, respData.(rp.OK).Data.(m.ObjekPajakBumi))
			}
		}
		return nil
	})

	if err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal create bulk: "+err.Error(), datas)
	}

	return rp.OKSimple{Data: datas}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.ObjekPajakBumi
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}

	result = tx.Delete(&data, id)
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

// func PerubahanZnt(nop , kodeZnt string) (any, error) {

// }

func PenilaianSppt(input []msppt.Sppt, tx *gorm.DB) (any, error) {
	var data *m.ObjekPajakBumi
	for _, v := range input {
		condition := nopSearcher(v)
		result := tx.Where(condition).First(&data)
		// if result.Error != nil {
		// 	return sh.SetError("request", "penilaian-data", source, "failed", "gagal mengambil data", data)
		// }
		if result.RowsAffected != 0 {
			// TODO: nilai sistem
			if resultSave := tx.Save(&data); resultSave.Error != nil {
				return sh.SetError("request", "update-data", source, "failed", "gagal menyimpan data objek pajak pbb", data)
			}
		}

	}
	return rp.OKSimple{Data: data}, nil
}

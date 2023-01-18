package regobjekpajakbumi

import (
	"fmt"
	"strconv"

	sc "github.com/jinzhu/copier"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"

	fasbng "github.com/bapenda-kota-malang/apin-backend/internal/models/regfasilitasbangunan"
	bng "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbangunan"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/regobjekpajakbumi"
	sropbng "github.com/bapenda-kota-malang/apin-backend/internal/services/regobjekpajakbangunan"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "regobjekpajakbumi"

func Create(input m.CreateDto, tx *gorm.DB) (any, error) {
	var dataFasBangunan fasbng.RegFasilitasBangunan
	var dataBng bng.CreateDto
	var dataBangunan bng.RegObjekPajakBangunan
	if tx == nil {
		tx = a.DB
	}
	var data m.RegObjekPajakBumi

	// copy input (payload) ke struct data satu if karene error dipakai sekali, +error
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "create-data", source, "failed", "gagal mengambil data payload", data)
	}

	// simpan data ke db satu if karena result dipakai sekali, +error
	err := a.DB.Transaction(func(tx *gorm.DB) error {
		//create data bumi
		if result := tx.Create(&data); result.Error != nil {
			return result.Error
		}

		nobangunan := 0
		tempKodeFas := ""
		if input.RegObjekPajakBangunans != nil {
			var errTemp error
			tempRegOpBng := *input.RegObjekPajakBangunans
			for _, opb := range tempRegOpBng {
				nobangunan = sropbng.GetLast() + 1

				dataBng.PstPermohonan_id = opb.PstPermohonan_id
				dataBng.RegFasilitasBangunans = nil
				dataBng.Provinsi_Kode = input.Provinsi_Kode
				dataBng.Daerah_Kode = input.Daerah_Kode
				dataBng.Kecamatan_Kode = input.Kecamatan_Kode
				dataBng.Kelurahan_Kode = input.Kelurahan_Kode
				dataBng.Blok_Kode = input.Blok_Kode
				dataBng.NoUrut = input.NoUrut
				dataBng.JenisOp = input.JenisOp

				dataBng.Area_Kode = input.Area_Kode
				dataBng.NoBangunan = &nobangunan

				dataBng.JenisAtap = opb.JenisAtap
				dataBng.JenisKonstruksi = opb.JenisKonstruksi
				dataBng.JenisTransaksi = opb.JenisTransaksi
				dataBng.JmlLantaiBng = opb.JmlLantaiBng
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

				//create data bangunan
				if result := tx.Create(&dataBangunan); result.Error != nil {
					return result.Error
				}

				if opb.RegFasBangunan != nil {
					val := opb.RegFasBangunan
					var fas fasbng.RegFasilitasBangunan
					_ = a.DB.Order("\"Id\" DESC").First(&fas)
					if fas.Id != 0 {
						dataFasBangunan.Id = fas.Id + 1
					}
					dataFasBangunan.Provinsi_Kode = input.Provinsi_Kode
					dataFasBangunan.Daerah_Kode = input.Daerah_Kode
					dataFasBangunan.Kecamatan_Kode = input.Kecamatan_Kode
					dataFasBangunan.Kelurahan_Kode = input.Kelurahan_Kode
					dataFasBangunan.NoUrut = input.NoUrut
					dataFasBangunan.JenisOp = input.JenisOp
					dataFasBangunan.NoBangunan = &nobangunan
					//create data fasilitas/jpb berdasar fasilitas/jpb di isi
					if val.FBJumlahACSplit != nil {
						tempKodeFas = "01"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBJumlahACSplit
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBJumlahACWindow != nil {
						tempKodeFas = "02"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBJumlahACWindow
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBIsACCentral != nil {
						tempKodeFas = "11"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBIsACCentral
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTipeLapisanKolam != nil {
						tempKodeFas = fmt.Sprint(val.FBTipeLapisanKolam)
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
					if val.FBHalamanBerat != nil {
						tempKodeFas = "16"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBHalamanBerat
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBHalamanSendang != nil {
						tempKodeFas = "15"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBHalamanSendang
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBHalamanRingan != nil {
						tempKodeFas = "14"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBJumlahACSplit
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBHalamanLantai != nil {
						tempKodeFas = "17"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBHalamanLantai
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTenisLampuBeton != nil {
						tempKodeFas = "18"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTenisLampuBeton
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTenisTanpaLampuBeton != nil {
						tempKodeFas = "21"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTenisTanpaLampuBeton
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTenisAspal1 != nil {
						tempKodeFas = "19"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTenisAspal1
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTenisAspal2 != nil {
						tempKodeFas = "22"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTenisAspal2
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTenisLiatRumput1 != nil {
						tempKodeFas = "20"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTenisLiatRumput1
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTenisLiatRumput2 != nil {
						tempKodeFas = "23"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTenisLiatRumput2
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBLiftPenumpang != nil {
						tempKodeFas = "30"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBLiftPenumpang
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBLiftKapsul != nil {
						tempKodeFas = "31"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBLiftKapsul
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBLiftBarang != nil {
						tempKodeFas = "32"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBLiftBarang
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTangga80 != nil {
						tempKodeFas = "33"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTangga80
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBTangga81 != nil {
						tempKodeFas = "34"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBTangga81
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBPagarBahan != nil {
						tempKodeFas = fmt.Sprint(val.FBPagarBahan)
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
					if val.FBPKHydrant != nil {
						tempKodeFas = "37"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBPKHydrant
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBPKSplinkler != nil {
						tempKodeFas = "39"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBPKSplinkler
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBPKFireAI != nil {
						tempKodeFas = "38"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBPKFireAI
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBPABX != nil {
						tempKodeFas = "41"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBPABX
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}
					if val.FBSumur != nil {
						tempKodeFas = "42"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.FBSumur
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}
						dataFasBangunan.Id = dataFasBangunan.Id + 1
					}

					if val.JpbKlinikACCentralKamar != nil {
						var dataRJpb bng.RegJpb5

						dataRJpb.NopDetail.Provinsi_Kode = input.Provinsi_Kode
						dataRJpb.NopDetail.Daerah_Kode = input.Daerah_Kode
						dataRJpb.NopDetail.Kecamatan_Kode = input.Kecamatan_Kode
						dataRJpb.NopDetail.Kelurahan_Kode = input.Kelurahan_Kode
						dataRJpb.NopDetail.Blok_Kode = input.Blok_Kode
						dataRJpb.NopDetail.NoUrut = input.NoUrut
						dataRJpb.NopDetail.JenisOp = input.JenisOp
						dataRJpb.NopDetail.Area_Kode = input.Area_Kode
						dataRJpb.NoBangunan = &nobangunan

						tempKodeFas = "07"
						dataFasBangunan.KodeFasilitas = &tempKodeFas
						dataFasBangunan.JumlahSatuan = val.JpbKlinikACCentralKamar
						errTemp = tx.Create(&dataFasBangunan).Error
						if errTemp != nil {
							return errTemp
						}

						dataRJpb.LuasKamarAcCentral = val.JpbKlinikACCentralKamar

						if val.JpbKlinikACCentralRuang != nil {
							tempKodeFas = "08"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.JpbKlinikACCentralRuang
							errTemp = tx.Create(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1

							dataRJpb.LuasRuangLainAcCentral = val.JpbKlinikACCentralRuang
						}

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if val.JpbHotelJenis != nil {
						var dataRJpb bng.RegJpb7

						dataRJpb.NopDetail.Provinsi_Kode = input.Provinsi_Kode
						dataRJpb.NopDetail.Daerah_Kode = input.Daerah_Kode
						dataRJpb.NopDetail.Kecamatan_Kode = input.Kecamatan_Kode
						dataRJpb.NopDetail.Kelurahan_Kode = input.Kelurahan_Kode
						dataRJpb.NopDetail.Blok_Kode = input.Blok_Kode
						dataRJpb.NopDetail.NoUrut = input.NoUrut
						dataRJpb.NopDetail.JenisOp = input.JenisOp
						dataRJpb.NopDetail.Area_Kode = input.Area_Kode
						dataRJpb.NoBangunan = &nobangunan

						dataRJpb.JenisHotel = (bng.JenisHotel)(fmt.Sprint(val.JpbHotelJenis))
						if val.JpbHotelBintang != nil {
							dataRJpb.JumlahBintang = (bng.JumlahBintang)(fmt.Sprint(val.JpbHotelJenis))
						}
						if val.JpbHotelJmlKamar != nil {
							dataRJpb.JumlahKamar = val.JpbHotelJmlKamar
						}
						if val.JpbHotelACCentralKamar != nil {
							tempKodeFas = "04"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.JpbHotelACCentralKamar
							errTemp = tx.Create(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1

							dataRJpb.LuasKamarAcCentral = val.JpbHotelACCentralKamar
						}
						if val.JpbHotelACCentralRuang != nil {
							tempKodeFas = "05"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.JpbHotelACCentralRuang
							errTemp = tx.Create(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1

							dataRJpb.LuasRuangLainAcCentral = val.JpbHotelACCentralRuang
						}

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if val.JpbApartemenJumlah != nil {
						var dataRJpb bng.RegJpb13

						dataRJpb.NopDetail.Provinsi_Kode = input.Provinsi_Kode
						dataRJpb.NopDetail.Daerah_Kode = input.Daerah_Kode
						dataRJpb.NopDetail.Kecamatan_Kode = input.Kecamatan_Kode
						dataRJpb.NopDetail.Kelurahan_Kode = input.Kelurahan_Kode
						dataRJpb.NopDetail.Blok_Kode = input.Blok_Kode
						dataRJpb.NopDetail.NoUrut = input.NoUrut
						dataRJpb.NopDetail.JenisOp = input.JenisOp
						dataRJpb.NopDetail.Area_Kode = input.Area_Kode
						dataRJpb.NoBangunan = &nobangunan

						dataRJpb.JumlahApartment = val.JpbApartemenJumlah
						// dataRJpb.KelasBangunan13 = (bng.KelasBangunan)(fmt.Sprint(val.))?? kelas bangunan tidak ada di inputan

						if val.JpbApartemenACCentralKamar != nil {
							tempKodeFas = "09"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.JpbApartemenACCentralKamar
							errTemp = tx.Create(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1

							dataRJpb.LuasApartAcCentral = val.JpbHotelACCentralRuang
						}
						if val.JpbApartemenACCentralLain != nil {
							tempKodeFas = "10"
							dataFasBangunan.KodeFasilitas = &tempKodeFas
							dataFasBangunan.JumlahSatuan = val.JpbApartemenACCentralLain
							errTemp = tx.Create(&dataFasBangunan).Error
							if errTemp != nil {
								return errTemp
							}
							dataFasBangunan.Id = dataFasBangunan.Id + 1

							dataRJpb.LuasRuangLainAcCentral = val.JpbHotelACCentralRuang
						}

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if val.JpbTankiKapasitas != nil {
						var dataRJpb bng.RegJpb15

						dataRJpb.NopDetail.Provinsi_Kode = input.Provinsi_Kode
						dataRJpb.NopDetail.Daerah_Kode = input.Daerah_Kode
						dataRJpb.NopDetail.Kecamatan_Kode = input.Kecamatan_Kode
						dataRJpb.NopDetail.Kelurahan_Kode = input.Kelurahan_Kode
						dataRJpb.NopDetail.Blok_Kode = input.Blok_Kode
						dataRJpb.NopDetail.NoUrut = input.NoUrut
						dataRJpb.NopDetail.JenisOp = input.JenisOp
						dataRJpb.NopDetail.Area_Kode = input.Area_Kode
						dataRJpb.NoBangunan = &nobangunan

						dataRJpb.KapasitasTanki = val.JpbTankiKapasitas
						if val.JpbTankiLetak != nil {
							dataRJpb.LetakTanki = (bng.LetakTanki)(fmt.Sprint(val.JpbTankiLetak))
						}

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if val.JpbProdTinggi != nil {
						if opb.Jpb_Kode == "03" {
							var dataRJpb bng.RegJpb3

							dataRJpb.NopDetail.Provinsi_Kode = input.Provinsi_Kode
							dataRJpb.NopDetail.Daerah_Kode = input.Daerah_Kode
							dataRJpb.NopDetail.Kecamatan_Kode = input.Kecamatan_Kode
							dataRJpb.NopDetail.Kelurahan_Kode = input.Kelurahan_Kode
							dataRJpb.NopDetail.Blok_Kode = input.Blok_Kode
							dataRJpb.NopDetail.NoUrut = input.NoUrut
							dataRJpb.NopDetail.JenisOp = input.JenisOp
							dataRJpb.NopDetail.Area_Kode = input.Area_Kode
							dataRJpb.NoBangunan = &nobangunan

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

							errTemp = tx.Create(&dataRJpb).Error
							if errTemp != nil {
								return errTemp
							}
						} else if opb.Jpb_Kode == "08" {
							var dataRJpb bng.RegJpb8

							dataRJpb.NopDetail.Provinsi_Kode = input.Provinsi_Kode
							dataRJpb.NopDetail.Daerah_Kode = input.Daerah_Kode
							dataRJpb.NopDetail.Kecamatan_Kode = input.Kecamatan_Kode
							dataRJpb.NopDetail.Kelurahan_Kode = input.Kelurahan_Kode
							dataRJpb.NopDetail.Blok_Kode = input.Blok_Kode
							dataRJpb.NopDetail.NoUrut = input.NoUrut
							dataRJpb.NopDetail.JenisOp = input.JenisOp
							dataRJpb.NopDetail.Area_Kode = input.Area_Kode
							dataRJpb.NoBangunan = &nobangunan

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

							errTemp = tx.Create(&dataRJpb).Error
							if errTemp != nil {
								return errTemp
							}
						}
					}

					if opb.Jpb_Kode == "00" || opb.Jpb_Kode == "01" || opb.Jpb_Kode == "10" || opb.Jpb_Kode == "11" {
					}

					if opb.Jpb_Kode == "02" {
						var dataRJpb bng.RegJpb2

						dataRJpb.NoBangunan = &nobangunan
						dataRJpb.KelasBangunan2 = "2" // gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if opb.Jpb_Kode == "04" {
						var dataRJpb bng.RegJpb4

						dataRJpb.NoBangunan = &nobangunan
						dataRJpb.KelasBangunan4 = "4" // gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if opb.Jpb_Kode == "06" {
						var dataRJpb bng.RegJpb6

						dataRJpb.NoBangunan = &nobangunan
						dataRJpb.KelasBanguna = "1" // gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if opb.Jpb_Kode == "09" {
						var dataRJpb bng.RegJpb9

						dataRJpb.NoBangunan = &nobangunan
						dataRJpb.KelasBangunan9 = "1" // gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if opb.Jpb_Kode == "12" {
						var dataRJpb bng.RegJpb12

						dataRJpb.NoBangunan = nobangunan
						dataRJpb.TipeBangunan = "1" // 1-4 gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if opb.Jpb_Kode == "14" {
						var dataRJpb bng.RegJpb14

						dataRJpb.NoBangunan = &nobangunan
						dataRJpb.LuasKanopi = nil // gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
						}
					}
					if opb.Jpb_Kode == "16" {
						var dataRJpb bng.RegJpb16

						dataRJpb.NoBangunan = &nobangunan
						dataRJpb.KelasBangunan16 = "1" // gak jelas dapat dari mana

						errTemp = tx.Create(&dataRJpb).Error
						if errTemp != nil {
							return errTemp
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

	return rp.OKSimple{Data: data}, nil
}

func GetList(input m.FilterDto) (any, error) {
	var data []m.RegObjekPajakBumi
	var count int64

	var pagination gh.Pagination
	result := a.DB.
		Model(&m.RegObjekPajakBumi{}).
		Scopes(gh.Filter(input)).
		Count(&count).
		Scopes(gh.Paginate(input, &pagination)).
		Find(&data)
	if result.Error != nil {
		return sh.SetError("request", "get-data-list", source, "failed", "gagal mengambil data reg objekpajakbumi", data)
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
	var data *m.RegObjekPajakBumi

	result := a.DB.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	} else if result.Error != nil {
		return sh.SetError("request", "get-data-detail", source, "failed", "gagal mengambil data reg objekpajakbumi", data)
	}

	return rp.OKSimple{
		Data: data,
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.RegObjekPajakBumi
	result := tx.First(&data, id)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if err := sc.Copy(&data, &input); err != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil data payload", data)
	}

	if result := tx.Save(&data); result.Error != nil {
		return sh.SetError("request", "update-data", source, "failed", "gagal mengambil menyimpan data reg objekpajakbumi", data)
	}

	return rp.OK{
		Meta: t.IS{
			"affected": strconv.Itoa(int(result.RowsAffected)),
		},
		Data: data,
	}, nil
}

func Delete(id int, tx *gorm.DB) (any, error) {
	if tx == nil {
		tx = a.DB
	}
	var data *m.RegObjekPajakBumi
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

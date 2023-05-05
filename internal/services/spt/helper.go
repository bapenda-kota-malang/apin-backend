package spt

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	mhdair "github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	mjppj "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisppj"
	mnpwpd "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mobjekpajak "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	mrek "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mdair "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptair"
	mdpln "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptppjpln"
	mdreklame "github.com/bapenda-kota-malang/apin-backend/internal/models/spt/detailsptreklame"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	mtarifreklame "github.com/bapenda-kota-malang/apin-backend/internal/models/tarifreklame"
	mtypes "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	sbjc "github.com/bapenda-kota-malang/apin-backend/internal/services/bankjatimclient"
	srek "github.com/bapenda-kota-malang/apin-backend/internal/services/configuration/rekening"
	shda "github.com/bapenda-kota-malang/apin-backend/internal/services/hargadasarair"
	sjppj "github.com/bapenda-kota-malang/apin-backend/internal/services/jenisppj"
	snpwpd "github.com/bapenda-kota-malang/apin-backend/internal/services/npwpd"
	sobjekpajak "github.com/bapenda-kota-malang/apin-backend/internal/services/objekpajak"
	stp "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifpajak"
	starifreklame "github.com/bapenda-kota-malang/apin-backend/internal/services/tarifreklame"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/bapenda-kota-malang/apin-backend/pkg/base64helper"
	ibj "github.com/bapenda-kota-malang/apin-backend/pkg/integration/bankjatim"
	"github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"github.com/google/uuid"
)

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

func generateKodeBilling(kodeBilling *string, nomerSpt string) *string {
	kode := fmt.Sprintf("%s%s", *kodeBilling, nomerSpt)
	return &kode
}

func vaManager(ctx context.Context, data m.Spt, flagProses ibj.FlagProses) (string, error) {
	if data.Rekening == nil {
		// get data rekening
		respRek, err := srek.GetDetail(int(data.Rekening_Id))
		if err != nil {
			return "", err
		}
		data.Rekening = respRek.(rp.OKSimple).Data.(*mrek.Rekening)
	}

	if data.Npwpd == nil {
		respNpwpd, err := snpwpd.GetDetail(nil, int(data.Npwpd_Id))
		if err != nil {
			return "", err
		}
		data.Npwpd = respNpwpd.(rp.OKSimple).Data.(*mnpwpd.Npwpd)
	}

	if data.ObjekPajak == nil {
		respObjePajak, err := sobjekpajak.GetDetail(int(data.ObjekPajak_Id))
		if err != nil {
			return "", err
		}
		data.ObjekPajak = respObjePajak.(rp.OKSimple).Data.(*mobjekpajak.ObjekPajak)
	}
	var va string
	tglExp := time.Time(data.JatuhTempo)
	switch flagProses {
	case ibj.ProsesCreate:
		// PDL => 5 digit no VA, 3 digit jenis pajak, 2 digit tahun, 6 digit nomor SPT
		va = fmt.Sprintf("%s0%s%s0%s", *data.Rekening.KodeVaJatim, *data.Rekening.Objek, data.NomorSpt[2:4], data.NomorSpt[4:])
	case ibj.ProsesUpdate, ibj.ProsesDelete:
		if data.VaJatim == nil {
			return "", errors.New("tidak dapat update pada nomor va kosong")
		}
		va = *data.VaJatim
		if servicehelper.IsJatuhTempo(&data.JatuhTempo) {
			tglExp = servicehelper.EndOfMonth(time.Now())
		}
	default:
		return "", errors.New("flag proses tidak diperbolehkan")
	}
	denda := float64(0)
	if data.Denda != nil {
		denda = *data.Denda
	}
	kenaikan := float64(0)
	if data.Kenaikan != nil {
		kenaikan = *data.Kenaikan
	}

	// payload registration bank jatim
	payload := ibj.ReqRegistration{
		VirtualAccount: va,
		Nama:           *data.ObjekPajak.Nama,
		TotalTagihan:   uint64(math.RoundToEven(data.Total)),
		TanggalExp:     tglExp.Format("20060102"),
		Berita1:        fmt.Sprintf("%s %s", *data.Npwpd.Npwpd, time.Now().Format("01-2006")),
		Berita2:        data.NomorSpt,
		Berita3:        fmt.Sprintf("DENDA %.2f", denda),
		Berita4:        fmt.Sprintf("KENAIKAN %.2f", kenaikan),
		Berita5:        fmt.Sprintf("UPDATE %s", time.Now().Format("02-01-2006")),
		FlagProses:     flagProses,
	}
	ctxTo, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := sbjc.Registration(ctxTo, payload, uint64(data.CreateBy_User_Id)); err != nil {
		return "", err
	}

	// deactivate va
	if data.JenisKetetapan == m.JenisKetetapanSkpdkb || data.JenisKetetapan == m.JenisKetetapanSkpdkbt {
		if data.Ref_Spt_Id == nil {
			return "", errors.New("gagal menghapus va lama: data ref spt id kosong")
		}
		respDetail, err := GetDetail(*data.Ref_Spt_Id, "", 0)
		if err != nil {
			return "", err
		}
		dataSptOld := respDetail.(rp.OKSimple).Data.(*m.Spt)
		if dataSptOld == nil {
			return "", fmt.Errorf("data spt ref tidak ditemukan")
		}
		if dataSptOld.Denda != nil {
			denda = *dataSptOld.Denda
		}
		if dataSptOld.Kenaikan != nil {
			kenaikan = *dataSptOld.Kenaikan
		}
		payload := ibj.ReqRegistration{
			VirtualAccount: va,
			Nama:           *dataSptOld.ObjekPajak.Nama,
			TotalTagihan:   uint64(math.RoundToEven(dataSptOld.Total)),
			TanggalExp:     time.Time(dataSptOld.JatuhTempo).Format("20060102"),
			Berita1:        fmt.Sprintf("%s %s", *dataSptOld.Npwpd.Npwpd, time.Now().Format("01-2006")),
			Berita2:        dataSptOld.NomorSpt,
			Berita3:        fmt.Sprintf("DENDA %f", denda),
			Berita4:        fmt.Sprintf("KENAIKAN %f", kenaikan),
			Berita5:        fmt.Sprintf("UPDATE %s", time.Now().Format("02-01-2006")),
			FlagProses:     ibj.ProsesDelete,
		}
		ctxTo, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		if err := sbjc.Registration(ctxTo, payload, uint64(data.CreateBy_User_Id)); err != nil {
			return "", err
		}
	}
	return va, nil
}

// Process calculate tax
func taxProcess(rekeningId *uint64, omset *float64, input m.Input) error {
	omsetOpt := "lte"

	// change detail for detail spt air & ppj pln before calculate tax
	if detail, ok := input.GetDetails().(mdair.CreateDto); ok {
		switch detail.Peruntukan {
		case mtypes.PeruntukanIndustriAir, mtypes.PeruntukanNiaga, mtypes.PeruntukanNonNiaga, mtypes.PeruntukanPdam:
			strPeruntukan := string(detail.Peruntukan)
			resphdair, err := shda.GetList(mhdair.FilterDto{
				Peruntukan:     &strPeruntukan,
				BatasBawah:     omset,
				BatasBawah_Opt: &omsetOpt,
			})
			if err != nil {
				return err
			}
			hdairs := resphdair.(rp.OK).Data.([]mhdair.HargaDasarAir)
			if len(hdairs) == 0 {
				return fmt.Errorf("harga dasar air not found")
			}
			hdair := hdairs[len(hdairs)-1]
			if detail.JenisAbt == mtypes.JenisABTNonMataAir {
				detail.Pengenaan = *hdair.TarifBukanMataAir
			} else {
				detail.JenisAbt = mtypes.JenisABTMataAir
				detail.Pengenaan = *hdair.TarifMataAir
			}
			input.ChangeDetails(detail)
		default:
			return fmt.Errorf("unknown peruntukan air")
		}
	} else if detail, ok := input.GetDetails().([]mdpln.CreateDto); ok {
		for v := range detail {
			resp, err := sjppj.GetDetail(int(detail[v].JenisPPJ_Id))
			if err != nil {
				return err
			}
			detail[v].JenisPPJ = resp.(rp.OKSimple).Data.(*mjppj.JenisPPJ)
		}
		input.ChangeDetails(detail)
	} else if detail, ok := input.GetDetails().([]mdreklame.CreateDto); ok {
		for v := range detail {
			resp, err := starifreklame.GetDetail(int(detail[v].TarifReklame_Id))
			if err != nil {
				return err
			}
			dataReklame := resp.(rp.OKSimple).Data.(*mtarifreklame.TarifReklame)
			if dataReklame.JenisMasa == mtypes.MasaPajakHari {
				resp, err := starifreklame.GetList(mtarifreklame.FilterDto{
					JenisMasa:    int16(mtypes.MasaPajakHari),
					JenisReklame: dataReklame.JenisReklame,
				})
				for _, vInner := range resp.(rp.OK).Data.([]mtarifreklame.TarifReklame) {
					if vInner.MasaPajak == nil {
						continue
					}
					if strings.ToLower(*vInner.MasaPajak) == "bulan" {
						detail[v].TarifBulan = *vInner.Tarif
					} else if strings.ToLower(*vInner.MasaPajak) == "minggu" {
						detail[v].TarifMinggu = *vInner.Tarif
					} else if strings.ToLower(*vInner.MasaPajak) == "hari" {
						detail[v].TarifHari = *vInner.Tarif
					}
				}
				if err != nil {
					return err
				}
			}
			detail[v].TarifReklame = dataReklame
		}
		input.ChangeDetails(detail)
	}

	// get tarif pajak data for calculate tax
	yearOrder := "order desc"
	rspTp, err := stp.GetList(mtp.FilterDto{
		Rekening_Id:   rekeningId,
		Tahun_Opt:     &yearOrder,
		OmsetAwal:     omset,
		OmsetAwal_Opt: &omsetOpt,
	})
	if err != nil {
		return err
	}

	tarifPajaks := rspTp.(rp.OK).Data.([]mtp.TarifPajak)
	if len(tarifPajaks) == 0 {
		return fmt.Errorf("tarif pajak not found")
	}
	tarifPajak := tarifPajaks[0]

	input.ReplaceTarifPajakId(uint(tarifPajak.Id))
	input.CalculateTax(tarifPajak.TarifPersen)
	return nil
}

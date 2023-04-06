package suratpemberitahuan

import (
	"path/filepath"
	"strconv"
	"time"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"

	mspt "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	msspd "github.com/bapenda-kota-malang/apin-backend/internal/models/sspd"
	m "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan"
	msuratdetail "github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan/detail"
	sspt "github.com/bapenda-kota-malang/apin-backend/internal/services/spt"
	servsspddetail "github.com/bapenda-kota-malang/apin-backend/internal/services/sspd/sspddetail"
	ssuratdetail "github.com/bapenda-kota-malang/apin-backend/internal/services/suratpemberitahuan/detail"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TrxSchedule(db *gorm.DB) (resp rp.OKSimple, err error) {
	if db == nil {
		db = a.DB
	}
	var data []m.SuratPemberitahuan
	err = db.Transaction(func(tx *gorm.DB) error {
		// get list data spt jatuh tempo
		dateNow := datatypes.Date(time.Now())
		filterJatuhTempo := uint8(5)
		respSptList, err := sspt.GetList(mspt.FilterDto{JatuhTempo: &dateNow, StatusData: &filterJatuhTempo, NoPagination: true}, 0, "bapenda", tx)
		if err != nil {
			return err
		}
		sptList := respSptList.(rp.OK).Data.([]mspt.ListDataDto)

		// calculate nominal, sum every spt calculate spt tax jumlah pajak + bunga - nominal bayar sspd detail
		// get spt id for surat pemberitahuan detail list array
		mapSuratDetail := make(map[uint64][]msuratdetail.CreateDto)
		mapSurat := make(map[uint64]m.CreateDto)
		mapPdf := make(map[uint64]m.Pdf)
		mapTotalPdf := make(map[uint64]map[string]float64)
		// loop business flow and prepare data
		for _, v := range sptList {
			jatuhTempo, _ := v.JatuhTempo.Value()
			tableContent := m.DataTableSurat{
				MasaPajak:  jatuhTempo.(time.Time).Format("January 2006"),
				JatuhTempo: jatuhTempo.(time.Time).Format("02-01-2006"),
				Skpd:       v.NomorSpt,
				Ketetapan:  sh.FormatCurrency(v.JumlahPajak),
			}
			diffMonth := sh.DendaMonth(time.Time(v.JatuhTempo))
			// denda = (jumlah pajak * 2%) * diff month
			denda := (v.JumlahPajak * 0.02) * float64(diffMonth)
			// sptTax = jumlah pajak (ketetapan) + denda
			sptTax := v.JumlahPajak + denda
			respSspdDetail, err := servsspddetail.GetListSspdDetail(msspd.SspdDetailFilterDto{Spt_Id: &v.Id})
			if err != nil {
				return err
			}
			telahDibayar := float64(0)
			sspdDetails := respSspdDetail.(rp.OK).Data.([]msspd.SspdDetail)
			// sum telah terbayar
			for i := range sspdDetails {
				telahDibayar += sspdDetails[i].NominalBayar
			}
			sisaPajak := sptTax - telahDibayar
			tableContent.Denda = sh.FormatCurrency(denda)
			tableContent.TelahDibayar = sh.FormatCurrency(telahDibayar)
			tableContent.SisaPajak = sh.FormatCurrency(sisaPajak)
			suratDetail := msuratdetail.CreateDto{Spt_Id: v.Id, Denda: denda, TelahDibayar: telahDibayar, SisaPajak: sisaPajak}
			mapSuratDetail[v.Npwpd_Id] = append(mapSuratDetail[v.Npwpd_Id], suratDetail)
			if suratDto, exist := mapSurat[v.Npwpd_Id]; exist {
				suratDto.Nominal += sisaPajak
				mapSurat[v.Npwpd_Id] = suratDto
				mapPdf[v.Npwpd_Id].AppendContent(tableContent)
				mapTotalPdf[v.Npwpd_Id]["ketetapan"] += v.JumlahPajak
				mapTotalPdf[v.Npwpd_Id]["denda"] += denda
				mapTotalPdf[v.Npwpd_Id]["telahBayar"] += telahDibayar
				mapTotalPdf[v.Npwpd_Id]["sisaPajak"] += sisaPajak
			} else {
				tableContent.No = 1
				nomorSurat := "973/. . . ./35.73.504/" + strconv.Itoa(time.Now().Year())
				mapSurat[v.Npwpd_Id] = m.CreateDto{
					Npwpd_Id:   v.Npwpd_Id,
					Tanggal:    dateNow,
					JatuhTempo: datatypes.Date(sh.EndOfMonth(sh.BeginningOfPreviosMonth())),
					Nominal:    sisaPajak,
					NoSurat:    nomorSurat,
				}

				jenisPajak := ""
				if v.Rekening.Nama != nil {
					jenisPajak = *v.Rekening.Nama
				}
				pengelola := ""
				if v.Npwpd.Nama != nil {
					pengelola = *v.Npwpd.Nama
				}
				npwpd := ""
				if v.Npwpd.Npwpd != nil {
					npwpd = *v.Npwpd.Npwpd
				}
				alamat := ""
				if v.ObjekPajak.Alamat != nil {
					alamat = *v.ObjekPajak.Alamat
				}

				perihal := "Surat Pemberitahuan"
				mapPdf[v.Npwpd_Id] = &Surat{
					Title:      perihal,
					Nomor:      nomorSurat,
					Sifat:      "Penting",
					Lampiran:   "-",
					Perihal:    perihal,
					JenisPajak: jenisPajak,
					Pengelola:  pengelola,
					Npwpd:      npwpd,
					Alamat:     alamat,
					DataTable:  []m.DataTableSurat{tableContent},
				}
				mapTotalPdf[v.Npwpd_Id] = make(map[string]float64)
				mapTotalPdf[v.Npwpd_Id]["ketetapan"] = v.JumlahPajak
				mapTotalPdf[v.Npwpd_Id]["denda"] = denda
				mapTotalPdf[v.Npwpd_Id]["telahBayar"] = telahDibayar
				mapTotalPdf[v.Npwpd_Id]["sisaPajak"] = sisaPajak
			}
		}
		// loop insert to database
		for k, v := range mapSurat {
			// check surat exist, if exist then continue loop
			respSuratExistingList, err := GetList(m.FilterDto{Npwpd_Id: k, JatuhTempo: v.JatuhTempo}, tx)
			if err != nil {
				return err
			}
			if len(respSuratExistingList.(rp.OK).Data.([]m.SuratPemberitahuan)) > 0 {
				continue
			}
			// set id
			id, err := sh.GetUuidv4()
			if err != nil {
				return err
			}
			v.Id = id

			// generate pdf
			mapPdf[k].SetTotal(mapTotalPdf[k])

			fileName := sh.GenerateFilename("suratPemberitahuan", v.Id, 0, "pdf")
			if err := mapPdf[k].GeneratePdf(filepath.Join(sh.GetPdfPath(), fileName)); err != nil {
				return err
			}
			v.Dokumen = fileName

			// create surat pemberitahuan
			respSurat, err := Create(v, tx)
			if err != nil {
				return err
			}
			suratData := respSurat.(rp.OKSimple).Data.(m.SuratPemberitahuan)
			// loop add surat pemberitahuan id to surat pemberitahuan detail
			if suratDetailDto, exist := mapSuratDetail[k]; exist {
				for i := range suratDetailDto {
					suratDetailDto[i].SuratPemberitahuan_Id = suratData.Id
				}
				respSuratDetail, err := ssuratdetail.Create(suratDetailDto, tx)
				if err != nil {
					return err
				}
				suratData.SuratPemberitahuanDetail = respSuratDetail.(rp.OKSimple).Data.([]msuratdetail.SuratPemberitahuanDetail)
			}
			data = append(data, suratData)
		}
		return nil
	})
	resp.Data = data
	return
}

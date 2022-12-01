package suratpemberitahuan

import (
	"math"
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
		for _, v := range sptList {
			// to get difference months basically:
			// 1. get now time change it to end of month and remove hour, min, second, nanosec
			nowMonth := sh.Midnight(sh.EndOfMonth(time.Now()))
			// 2. sub with jatuh tempo value to get time difference duration
			diffDuration := nowMonth.Sub(time.Time(v.JatuhTempo))
			// 3. get hours duration divided by 24 then 30 to get difference month
			diffMonth := diffDuration.Hours() / 24 / 30
			// 4. remove fractional from result step 3
			diffMonth, _ = math.Modf(diffMonth)
			// max diff 24 month
			if diffMonth > 24 {
				diffMonth = 24
			}
			// sptTax = (jumlah pajak * 2%) * diff month
			sptTax := (v.JumlahPajak * 0.02) * diffMonth
			respSspdDetail, err := servsspddetail.GetListSspdDetail(msspd.SspdDetailFilterDto{Spt_Id: &v.Id})
			if err != nil {
				return err
			}
			telahDibayar := float64(0)
			sspdDetails := respSspdDetail.(rp.OK).Data.([]msspd.SspdDetail)
			for i := range sspdDetails {
				telahDibayar += sspdDetails[i].NominalBayar
			}
			mapSuratDetail[v.Npwpd_Id] = append(mapSuratDetail[v.Npwpd_Id], msuratdetail.CreateDto{Spt_Id: v.Id})
			if suratDto, exist := mapSurat[v.Npwpd_Id]; exist {
				suratDto.Nominal += sptTax - telahDibayar
				mapSurat[v.Npwpd_Id] = suratDto
			} else {
				mapSurat[v.Npwpd_Id] = m.CreateDto{
					Npwpd_Id:   v.Npwpd_Id,
					Tanggal:    dateNow,
					JatuhTempo: datatypes.Date(sh.EndOfMonth(sh.BeginningOfPreviosMonth())),
					Nominal:    sptTax - telahDibayar,
				}
			}
		}
		for k, v := range mapSurat {
			// check surat exist, if exist then continue loop
			respSuratExistingList, err := GetList(m.FilterDto{Npwpd_Id: k, JatuhTempo: v.JatuhTempo}, tx)
			if err != nil {
				return err
			}
			if len(respSuratExistingList.(rp.OK).Data.([]m.SuratPemberitahuan)) > 0 {
				continue
			}
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

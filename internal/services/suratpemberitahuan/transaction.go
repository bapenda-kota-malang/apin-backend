package suratpemberitahuan

import (
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
			sptTax := v.JumlahPajak
			if v.Denda != nil {
				sptTax += *v.Denda
			}
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
				mapSurat[v.Npwpd_Id] = m.CreateDto{Npwpd_Id: v.Npwpd_Id, Tanggal: dateNow, JatuhTempo: v.JatuhTempo, Nominal: sptTax - telahDibayar}
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

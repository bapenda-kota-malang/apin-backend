package prosesjambong

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"
	"github.com/google/uuid"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
	sh "github.com/bapenda-kota-malang/apin-backend/pkg/servicehelper"
)

const source = "penetapanmassal"

func Create(input m.CreateDto, userId uint, tx *gorm.DB) (any, error) {

	return rp.OKSimple{Data: ""}, nil
}

func CopySpt(input m.CopySptDto) (any, error) {
	for _, sptId := range input.Spt_Id {
		// find spt by spt id
		var spt *m.Spt
		result := a.DB.First(&spt, sptId)
		if result.Error != nil {
			continue
		}

		var newSpt m.Spt

		if input.Types == "sa" {
			newSpt.PeriodeAwal = spt.PeriodeAwal
			newSpt.PeriodeAkhir = spt.PeriodeAkhir
			newSpt.JatuhTempo = spt.JatuhTempo
			newSpt.NomorSpt = spt.NomorSpt
			newSpt.KodeBilling = spt.KodeBilling
			newSpt.TanggalSpt = spt.TanggalSpt
			jenisKetetapan := m.JenisKetetapanSkpdkb
			newSpt.JenisKetetapan = &jenisKetetapan
			newSpt.Kenaikan = spt.Kenaikan
			newSpt.Denda = spt.Denda

			// hitung jumlah pajak + kenaikan 25% (dari jumlah pajak) + denda 2% (dari jumlah pajak)
			total := spt.JumlahPajak + (spt.JumlahPajak * 25 / 100) + (spt.JumlahPajak * 2 / 100)
			// int to float
			totalFloat := float64(total)
			newSpt.Total = &totalFloat
		} else if input.Types == "oa" {
			currentDate := sh.CurrentDate()
			beginningOfPreviosMonth := sh.BeginningOfPreviosMonth()
			lastOfPreviosMonth := sh.EndOfPreviousMonth()
			jatuhTempo := sh.EndOfMonth(sh.CurrentDate())

			newSpt.PeriodeAwal = datatypes.Date(beginningOfPreviosMonth)
			newSpt.PeriodeAkhir = datatypes.Date(lastOfPreviosMonth)
			newSpt.JatuhTempo = datatypes.Date(jatuhTempo)
			newSpt.TanggalSpt = currentDate
			jenisKetetapan := m.JenisKetetapanSkpd
			newSpt.JenisKetetapan = &jenisKetetapan
		}

		// create new spt
		newSpt.Npwpd_Id = spt.Npwpd_Id
		newSpt.ObjekPajak_Id = spt.ObjekPajak_Id
		newSpt.Rekening_Id = spt.Rekening_Id
		newSpt.LuasLokasi = spt.LuasLokasi
		newSpt.Description = spt.Description
		newSpt.TarifPajak_Id = spt.TarifPajak_Id
		newSpt.EtaxSubTotal = spt.EtaxSubTotal
		newSpt.EtaxTotal = spt.EtaxTotal
		newSpt.EtaxJumlahPajak = spt.EtaxJumlahPajak
		newSpt.Omset = spt.Omset
		newSpt.JumlahPajak = spt.JumlahPajak
		newSpt.Lampiran = spt.Lampiran
		newSpt.Sunset = spt.Sunset
		newSpt.CreateBy_User_Id = spt.CreateBy_User_Id
		newSpt.NomorSpt = spt.NomorSpt
		newSpt.KodeBilling = spt.KodeBilling
		newSpt.Type = spt.Type

		newSpt.TanggalLunas = spt.TanggalLunas
		newSpt.BatalPenetapan = spt.BatalPenetapan
		newSpt.IdBT = spt.IdBT
		newSpt.JumlahTahun = spt.JumlahTahun
		newSpt.JumlahBulan = spt.JumlahBulan
		newSpt.JumlahMinggu = spt.JumlahMinggu
		newSpt.JumlahHari = spt.JumlahHari
		newSpt.Gambar = spt.Gambar
		newSpt.KeteranganPajak = spt.KeteranganPajak
		newSpt.KoefisienPajak = spt.KoefisienPajak
		newSpt.NamaProduk = spt.NamaProduk
		newSpt.NomorRegister = spt.NomorRegister
		newSpt.JudulReklame = spt.JudulReklame
		newSpt.NamaPenyewa = spt.NamaPenyewa
		newSpt.AlamatPenyewa = spt.AlamatPenyewa
		newSpt.VaJatim = spt.VaJatim

		newSpt.DasarPengenaan = spt.DasarPengenaan
		newSpt.Kenaikan = spt.Kenaikan
		newSpt.Denda = spt.Denda
		newSpt.Pengurangan = spt.Pengurangan
		newSpt.Total = spt.Total
		newSpt.Ref_Spt_Id = spt.Ref_Spt_Id
		newSpt.BillingPenetapan = spt.BillingPenetapan
		newSpt.Teguran_Id = spt.Teguran_Id
		newSpt.IsTeguran = spt.IsTeguran
		newSpt.KeteranganPenetapan = spt.KeteranganPenetapan
		newSpt.StatusPenetapan = spt.StatusPenetapan
		newSpt.Kasubid_User_Id = spt.Kasubid_User_Id
		newSpt.Kabid_User_Id = spt.Kabid_User_Id
		newSpt.CancelledAt = spt.CancelledAt
		newSpt.DateModel = spt.DateModel

		// create new spt
		result = a.DB.Create(&newSpt)
		if result.Error != nil {
			return nil, result.Error
		}

	}

	return rp.OK{Meta: t.IS{
		"message": "success copy spt",
	}}, nil
}

func GetList(input m.FilterDto, types, kodeJenisPajak string) (any, error) {
	var (
		count int64
		query string
	)

	if input.PageSize == 0 {
		input.PageSize = 10
	}

	if input.Page == 0 {
		input.Page = 1
	}

	offset := input.PageSize * (input.Page - 1)

	type item struct {
		Id             uuid.UUID `json:"id" gorm:"column:Id"`
		NomorSPt       string    `json:"nomor_spt" gorm:"column:NomorSpt"`
		Npwpd          string    `json:"npwpd" gorm:"column:Npwpd"`
		JenisUsaha     string    `json:"jenis_usaha" gorm:"column:JenisUsaha"`
		PeriodeAwal    string    `json:"periode_awal" gorm:"column:PeriodeAwal"`
		PeriodeAkhir   string    `json:"periode_akhir" gorm:"column:PeriodeAkhir"`
		TanggalSpt     string    `json:"tanggal_spt" gorm:"column:TanggalSpt"`
		JatuhTempo     string    `json:"jatuh_tempo" gorm:"column:JatuhTempo"`
		ObjekPajakNama string    `json:"objek_pajak_nama" gorm:"column:Nama"`
	}

	// "211, 200, 300" to '211','200','300'
	kodeJenisPajak = strings.Replace(kodeJenisPajak, ",", "','", -1)
	kodeJenisPajak = fmt.Sprintf("'%s'", kodeJenisPajak)

	// default query
	query = fmt.Sprintf(`
			SELECT
				spt."Id",
				spt."NomorSpt",
				spt."PeriodeAwal",
				spt."PeriodeAkhir",
				spt."TanggalSpt",
				spt."JatuhTempo",
				rek."JenisUsaha",
				npwpd."Npwpd",
				op."Nama"
			FROM
				"Spt" spt
				INNER JOIN "Rekening" rek ON rek."Id" = spt."Rekening_Id"
				INNER JOIN "Npwpd" npwpd ON npwpd."Id" = spt."Npwpd_Id"
				INNER JOIN "ObjekPajak" op ON op."Id" = spt."ObjekPajak_Id"
			WHERE
				rek."KodeJenisUsaha" IN (%s) 
				AND spt."Type" = '%s' 
				AND spt."DeletedAt" IS NULL 
				AND npwpd."DeletedAt" IS NULL 
		`, kodeJenisPajak, types)

	// query
	if types == "OA" {
		// query += ` AND spt."PeriodeAwal" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '1 day' AND spt."PeriodeAkhir" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '1 day'`

	} else if types == "SA" {
		// query += ` AND spt."PeriodeAwal" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '10 day' AND spt."PeriodeAkhir" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '10 day'`
	}

	// query count
	queryCount := fmt.Sprintf(`SELECT COUNT(*) FROM (%s) AS count`, query)
	resultCount := a.DB.Raw(queryCount).Count(&count)
	if resultCount.Error != nil {
		return nil, resultCount.Error
	}

	// pagination
	var items []item
	query = fmt.Sprintf(`%s ORDER BY spt."NomorSpt" DESC LIMIT %d OFFSET %d`, query, input.PageSize, offset)
	result := a.DB.Debug().Raw(query).Scan(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return rp.OK{
		Meta: t.IS{
			"totalCount":   strconv.Itoa(int(count)),
			"currentCount": strconv.Itoa(int(result.RowsAffected)),
			"page":         strconv.Itoa(input.Page),
			"pageSize":     strconv.Itoa(input.PageSize),
		},
		Data: items,
	}, nil
}

func GetDetail(id int) (any, error) {

	return rp.OKSimple{
		Data: "",
	}, nil
}

func Update(id int, input m.UpdateDto, tx *gorm.DB) (any, error) {

	return rp.OK{
		Meta: t.IS{
			"affected": "",
		},
		Data: "",
	}, nil
}

func Delete(id int) (any, error) {

	return rp.OK{
		Meta: t.IS{
			"count":  "",
			"status": "",
		},
		Data: "",
	}, nil
}

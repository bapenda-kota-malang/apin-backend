package prosesjambong

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	a "github.com/bapenda-kota-malang/apin-backend/pkg/apicore"
	rp "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/responses"

	m "github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	t "github.com/bapenda-kota-malang/apin-backend/pkg/apicore/types"
)

const source = "penetapanmassal"

func Create(input m.CreateDto, userId uint, tx *gorm.DB) (any, error) {

	return rp.OKSimple{Data: ""}, nil
}

func GetList(input m.FilterDto, types, kodeJenisUsaha string) (any, error) {
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
		NomorSPt       string `json:"nomor_spt" gorm:"column:NomorSpt"`
		Npwpd          string `json:"npwpd" gorm:"column:Npwpd"`
		JenisUsaha     string `json:"jenis_usaha" gorm:"column:JenisUsaha"`
		PeriodeAwal    string `json:"periode_awal" gorm:"column:PeriodeAwal"`
		PeriodeAkhir   string `json:"periode_akhir" gorm:"column:PeriodeAkhir"`
		TanggalSpt     string `json:"tanggal_spt" gorm:"column:TanggalSpt"`
		JatuhTempo     string `json:"jatuh_tempo" gorm:"column:JatuhTempo"`
		ObjekPajakNama string `json:"objek_pajak_nama" gorm:"column:Nama"`
	}

	// default query
	query = fmt.Sprintf(`
			SELECT
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
				rek."KodeJenisUsaha" = '%s' 
				AND npwpd."JenisPajak" = '%s' 
				AND spt."DeletedAt" IS NULL 
				AND npwpd."DeletedAt" IS NULL 
		`, kodeJenisUsaha, types)

	// query
	if types == "OA" {
		// TODO check tanggal :
		// AND spt."PeriodeAwal" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '1 day' AND spt."PeriodeAkhir" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '1 day'

	} else if types == "SA" {
		// TODO check tanggal
		// AND spt."PeriodeAwal" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '10 day' AND spt."PeriodeAkhir" BETWEEN date_trunc('month', current_date - interval '1 month') AND date_trunc('month', current_date) - interval '10 day'
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

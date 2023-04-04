package pelaporanppat

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type PelaporanPpat struct {
	Id                 uint64          `json:"id" gorm:"primaryKey"`
	Sptpd_Id           *string         `json:"sptpd_id" gorm:"type:varchar(50)"`
	Ppat_Id            *string         `json:"ppat_id" gorm:"type:varchar(50)"`
	NoAkta             *string         `json:"noAkta" gorm:"type:varchar(50)"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal" gorm:"type:varchar(2)"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal" gorm:"type:varchar(4)"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir" gorm:"type:varchar(2)"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir" gorm:"type:varchar(4)"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan" gorm:"type:varchar(50)"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
	NominalSSP         *float64        `json:"nominalSSP"`
	gh.DateModel
}

type CreateDto struct {
	Sptpd_Id           *string         `json:"sptpd_id"`
	Ppat_Id            *string         `json:"ppat_id"`
	NoAkta             *string         `json:"noAkta"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
}

type UpdateDto struct {
	Id                 *uint64         `json:"id"`
	Sptpd_Id           *string         `json:"sptpd_id"`
	Ppat_Id            *string         `json:"ppat_id"`
	NoAkta             *string         `json:"noAkta"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
}

type FilterDto struct {
	Sptpd_Id           *string         `json:"sptpd_id"`
	Ppat_Id            *string         `json:"ppat_id"`
	NoAkta             *string         `json:"noAkta"`
	TglAkta            *datatypes.Date `json:"tglAkta"`
	PeriodeBulanAwal   *string         `json:"periodeBulanAwal"`
	PeriodeTahunAwal   *string         `json:"periodeTahunAwal"`
	PeriodeBulanAkhir  *string         `json:"periodeBulanAkhir"`
	PeriodeTahunAkhir  *string         `json:"periodeTahunAkhir"`
	TglLapor           *datatypes.Date `json:"tglLapor"`
	PihakYgMengalihkan *string         `json:"pihakYgMengalihkan"`
	TglSSP             *datatypes.Date `json:"tglSSP"`
	Page               int             `json:"page"`
	PageSize           int             `json:"page_size"`
}

type ResponsePelaporanPpat struct {
	Ppat_Id     *string `json:"ppat_id"`
	Ppat_Name   *string `json:"ppat_name"`
	TglLapor    *string `json:"tglLapor"`
	Sptpd_Id    *string `json:"sptpd_Id"`
	NilaiOp     *string `json:"nilaiOp"`
	JumlahSetor *string `json:"jumlahSetor"`
	Status      *string `json:"status"`
}

type ResponseDetilPelaporanPpat struct {
	Ppat_Id     *string `json:"ppat_id"`
	Ppat_Name   *string `json:"ppat_name"`
	TglLapor    *string `json:"tglLapor"`
	Sptpd_Id    *string `json:"sptpd_Id"`
	NilaiOp     *string `json:"nilaiOp"`
	JumlahSetor *string `json:"jumlahSetor"`
	Status      *string `json:"status"`
}

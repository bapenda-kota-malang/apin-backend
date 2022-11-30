package suratpemberitahuan

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan/detail"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

/*
Notes:
Enum Status:
0 Belum Terbit; 1 Terbit; 2 Jatuh Tempo; 3 Lunas; 4 Undangan; 5 Pemeriksaan; 6 Penetapan
"Belum Terbit : Belum dikirim ke email atau Belum Dicetak
Terbit: Sudah Dikirim ke Email atau Sudah Dicetak
Jatuh Tempo : SPT sudah Jatuh Tempo"
*/

type SuratPemberitahuan struct {
	Id           uint64          `json:"id" gorm:"primaryKey"`
	NoSurat      string          `json:"noSurat"`
	Tanggal      datatypes.Date  `json:"tanggal"`
	Npwpd_Id     uint64          `json:"npwpd_id"`
	JatuhTempo   datatypes.Date  `json:"jatuhTempo"`
	Nominal      float64         `json:"nominal"`
	TanggalLunas *datatypes.Date `json:"tanggalLunas"`
	Status       uint8           `json:"status"`
	gormhelper.DateModel
	Npwpd                    *npwpd.Npwpd                      `json:"npwpd" gorm:"foreignKey:Npwpd_Id"`
	SuratPemberitahuanDetail []detail.SuratPemberitahuanDetail `json:"suratPemberitahuanDetail" gorm:"foreignKey:SuratPemberitahuan_Id"`
}

type CreateDto struct {
	NoSurat      string          `json:"noSurat"`
	Tanggal      datatypes.Date  `json:"tanggal"`
	Npwpd_Id     uint64          `json:"npwpd_id"`
	JatuhTempo   datatypes.Date  `json:"jatuhTempo"`
	Nominal      float64         `json:"-"`
	TanggalLunas *datatypes.Date `json:"tanggalLunas"`
	Status       uint8           `json:"-"`
}

type UpdateBulk struct {
	Id      *uint64         `json:"id" validate:"required;min=1"`
	Tanggal *datatypes.Date `json:"tanggal"`
	Status  *uint8          `json:"status" validate:"min=1;max=2"`
}

type UpdateBulkDto struct {
	Datas []UpdateBulk `json:"datas" validate:"required"`
}

type UpdateDto struct {
	Tanggal *datatypes.Date `json:"tanggal"`
	Status  *uint8          `json:"status" validate:"min=1;max=2"`
}

type FilterDto struct {
	Npwpd_Id   uint64         `json:"npwpd_id"`
	JatuhTempo datatypes.Date `json:"jatuhTempo"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
}

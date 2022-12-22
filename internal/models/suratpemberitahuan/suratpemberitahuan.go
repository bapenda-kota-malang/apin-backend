package suratpemberitahuan

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/suratpemberitahuan/detail"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
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
	Id           uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	NoSurat      string          `json:"noSurat"`
	Tanggal      datatypes.Date  `json:"tanggal"`
	Npwpd_Id     uint64          `json:"npwpd_id"`
	JatuhTempo   datatypes.Date  `json:"jatuhTempo"`
	Nominal      float64         `json:"nominal"`
	TanggalLunas *datatypes.Date `json:"tanggalLunas"`
	Status       uint8           `json:"status"`
	Dokumen      string          `json:"dokumen"`
	gormhelper.DateModel
	Npwpd                    *npwpd.Npwpd                      `json:"npwpd" gorm:"foreignKey:Npwpd_Id"`
	SuratPemberitahuanDetail []detail.SuratPemberitahuanDetail `json:"suratPemberitahuanDetail" gorm:"foreignKey:SuratPemberitahuan_Id"`
}

type CreateDto struct {
	Id           uuid.UUID       `json:"id"`
	NoSurat      string          `json:"noSurat"`
	Tanggal      datatypes.Date  `json:"tanggal"`
	Npwpd_Id     uint64          `json:"npwpd_id"`
	JatuhTempo   datatypes.Date  `json:"jatuhTempo"`
	Nominal      float64         `json:"-"`
	TanggalLunas *datatypes.Date `json:"tanggalLunas"`
	Status       uint8           `json:"-"`
	Dokumen      string          `json:"-"`
}

type UpdateBulk struct {
	Id      *uuid.UUID      `json:"id" validate:"required"`
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

type DataTableSurat struct {
	No           uint16
	MasaPajak    string
	JatuhTempo   string
	Skpd         string
	Ketetapan    string
	Denda        string
	TelahDibayar string
	SisaPajak    string
}
type Pdf interface {
	GeneratePdf(outFile string) error
	AppendContent(DataTableSurat)
	SetTotal(map[string]float64)
}

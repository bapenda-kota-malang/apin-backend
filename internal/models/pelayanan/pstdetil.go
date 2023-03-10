package pstpermohonan

import (
	"fmt"

	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type PstDetail struct {
	Id                    uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string         `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  *string         `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID      *string         `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID *string         `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID *string         `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID      *string         `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon         *string         `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID      *string         `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	JenisPelayananID      *string         `json:"jenisPelayananID" gorm:"type:varchar(4)"`
	TahunPajakPemohon     *string         `json:"tahunPajakPemohon" gorm:"type:varchar(30)"`
	NIP                   *string         `json:"nip" gorm:"type:varchar(9)"`
	Notes                 *string         `json:"catatan" gorm:"type:varchar(75)"`
	StatusSelesai         *int            `json:"keterangan" gorm:"type:integer"`
	TanggalSelesai        *datatypes.Date `json:"tanggalSelesai"`
	SeksiBerkasID         *string         `json:"seksiBerkasID" gorm:"type:varchar(2)"`
	TanggalPenyerahan     *datatypes.Date `json:"tanggalPenyerahan"`
	gormhelper.DateModel
}

type GetByNoPelayananPstDetailDto struct {
	TahunPelayanan  string `json:"tahunPelayanan" validate:"required"`
	BundelPelayanan string `json:"jenisPelayanan" validate:"required"`
	NoUrutPelayanan string `json:"noUrutPelayanan" validate:"required"`
}

type PstDetailInput struct {
	Id                    uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId          *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string         `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  *string         `json:"permohonanProvinsiID" gorm:"type:varchar(2)"`
	PermohonanKotaID      *string         `json:"permohonanKotaID" gorm:"type:varchar(2)"`
	PermohonanKecamatanID *string         `json:"permohonanKecamatanID" gorm:"type:varchar(3)"`
	PermohonanKelurahanID *string         `json:"permohonanKelurahanID" gorm:"type:varchar(3)"`
	PermohonanBlokID      *string         `json:"permohonanBlokID" gorm:"type:varchar(3)"`
	NoUrutPemohon         *string         `json:"noUrutPemohon" gorm:"type:varchar(4)"`
	PemohonJenisOPID      *string         `json:"pemohonJenisOPID" gorm:"type:varchar(1)"`
	JenisPelayananID      *string         `json:"jenisPelayananID" gorm:"type:varchar(4)"`
	TahunPajakPemohon     *string         `json:"tahunPajakPemohon" gorm:"type:varchar(30)"`
	TanggalSelesai        *datatypes.Date `json:"tanggalSelesai"`
}

func (i PstDetail) GetDataPermohonanNOP() string {
	result := fmt.Sprintf("%s%s%s%s%s%s%s", *i.PermohonanProvinsiID, *i.PermohonanKotaID, *i.PermohonanKecamatanID, *i.PermohonanKelurahanID, *i.PermohonanBlokID, *i.NoUrutPemohon, *i.PemohonJenisOPID)
	return result
}

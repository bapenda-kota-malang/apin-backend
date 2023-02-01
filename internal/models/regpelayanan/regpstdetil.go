package regpstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type RegPstDetail struct {
	Id                uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId      *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId          *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId           *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan    *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan   *string         `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan   *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	JenisPelayananID  *string         `json:"jenisPelayananID" gorm:"type:varchar(4)"`
	TahunPajakPemohon *string         `json:"tahunPajakPemohon" gorm:"type:varchar(30)"`
	NIP               *string         `json:"nip" gorm:"type:varchar(9)"`
	Notes             *string         `json:"catatan" gorm:"type:varchar(75)"`
	StatusSelesai     *int            `json:"keterangan" gorm:"type:integer"`
	TanggalSelesai    *datatypes.Date `json:"tanggalSelesai"`
	SeksiBerkasID     *string         `json:"seksiBerkasID" gorm:"type:varchar(2)"`
	TanggalPenyerahan *datatypes.Date `json:"tanggalPenyerahan"`
	gormhelper.DateModel
}

type RegPstDetailInput struct {
	Id                uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	PermohonanId      *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId          *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId           *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan    *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan   *string         `json:"jenisPelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan   *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	JenisPelayananID  *string         `json:"jenisPelayananID" gorm:"type:varchar(4)"`
	TahunPajakPemohon *string         `json:"tahunPajakPemohon" gorm:"type:varchar(30)"`
	TanggalSelesai    *datatypes.Date `json:"tanggalSelesai"`
}

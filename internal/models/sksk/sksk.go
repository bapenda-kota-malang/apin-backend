package sksk

import (
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pelayanan"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type SkSk struct {
	Id            uint64          `json:"id" gorm:"primaryKey"`
	PermohonanId  *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId      *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId       *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	JnsSK         *string         `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK          *string         `json:"noSK" gorm:"type:varchar(30);unique"`
	TglSK         *datatypes.Date `json:"tglSK"`
	NoBaKantor    *string         `json:"noBaKantor" gorm:"type:varchar(30)"`
	TglBaKantor   *datatypes.Date `json:"tglBaKantor"`
	NoBaLapangan  *string         `json:"noBaLapangan" gorm:"type:varchar(30)"`
	TglBaLapangan *datatypes.Date `json:"tglBaLapangan"`
	TglCetak      *datatypes.Date `json:"tglCetak"`
	NipPenetak    *string         `json:"nipPenetak" gorm:"type:varchar(9)"`

	gormhelper.DateModel
}

type CreateDto struct {
	PermohonanId  *uint64         `json:"permohonanId"`
	KanwilId      *string         `json:"kanwilId"`
	KppbbId       *string         `json:"kppbbId"`
	JnsSK         *string         `json:"jnsSK"`
	NoSK          *string         `json:"noSK"`
	TglSK         *datatypes.Date `json:"tglSK"`
	NoBaKantor    *string         `json:"noBaKantor"`
	TglBaKantor   *datatypes.Date `json:"tglBaKantor"`
	NoBaLapangan  *string         `json:"noBaLapangan"`
	TglBaLapangan *datatypes.Date `json:"tglBaLapangan"`
	TglCetak      *datatypes.Date `json:"tglCetak"`
	NipPenetak    *string         `json:"nipPenetak"`
}

type UpdateDto struct {
	Id            uint64          `json:"id"`
	PermohonanId  *uint64         `json:"permohonanId"`
	KanwilId      *string         `json:"kanwilId"`
	KppbbId       *string         `json:"kppbbId"`
	JnsSK         *string         `json:"jnsSK"`
	NoSK          *string         `json:"noSK"`
	TglSK         *datatypes.Date `json:"tglSK"`
	NoBaKantor    *string         `json:"noBaKantor"`
	TglBaKantor   *datatypes.Date `json:"tglBaKantor"`
	NoBaLapangan  *string         `json:"noBaLapangan"`
	TglBaLapangan *datatypes.Date `json:"tglBaLapangan"`
	TglCetak      *datatypes.Date `json:"tglCetak"`
	NipPenetak    *string         `json:"nipPenetak"`
}

type FilterDto struct {
	Id            uint64          `json:"id"`
	PermohonanId  *uint64         `json:"permohonanId"`
	KanwilId      *string         `json:"kanwilId"`
	KppbbId       *string         `json:"kppbbId"`
	JnsSK         *string         `json:"jnsSK"`
	NoSK          *string         `json:"noSK"`
	TglSK         *datatypes.Date `json:"tglSK"`
	NoBaKantor    *string         `json:"noBaKantor"`
	TglBaKantor   *datatypes.Date `json:"tglBaKantor"`
	NoBaLapangan  *string         `json:"noBaLapangan"`
	TglBaLapangan *datatypes.Date `json:"tglBaLapangan"`
	TglCetak      *datatypes.Date `json:"tglCetak"`
	NipPenetak    *string         `json:"nipPenetak"`
	Page          int             `json:"page"`
	PageSize      int             `json:"page_size"`
}

type CetakDto struct {
	Provinsi_Kode  string          `json:"provinsi_kode" validate:"required;minLength=2;maxLength=2"`
	Daerah_Kode    string          `json:"daerah_kode" validate:"required;minLength=2;maxLength=2"`
	Kecamatan_Kode string          `json:"kecamatan_kode" validate:"required;minLength=3;maxLength=3"`
	Kelurahan_Kode string          `json:"kelurahan_kode" validate:"required;minLength=3;maxLength=3"`
	NamaPropinsi   *string         `json:"namaPropinsi"`
	NamaDati2      *string         `json:"namaDati2"`
	NamaKecamatan  *string         `json:"namaKecamatan"`
	NamaKelurahan  *string         `json:"namaKelurahan"`
	KanwilId       *string         `json:"kanwilId"`
	KppbbId        *string         `json:"kppbbId"`
	Tahun          *string         `json:"tahun"`
	JnsSK          *string         `json:"jenis"`
	NoSK           *string         `json:"nomor"`
	TglSK          *datatypes.Date `json:"tanggal"`
	TglCetak       *datatypes.Date `json:"tglCetak"`
	NipPenetak     *string         `json:"nipCetak"`
	TypeLaporan    *string         `json:"typeLaporan"`
}

func SetSkSk(i p.PstPermohonan) *SkSk {
	return &SkSk{
		PermohonanId: &i.Id,
		KanwilId:     i.KanwilId,
		KppbbId:      i.KppbbId,
	}
}

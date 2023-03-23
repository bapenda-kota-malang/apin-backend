package keberatan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type KeputusanKeberatanPbb struct {
	Id                    uint64          `json:"id" gorm:"primaryKey"`
	PermohonanId          *uint64         `json:"permohonanId" gorm:"type:integer"`
	KanwilId              *string         `json:"kanwilId" gorm:"type:varchar(2)"`
	KppbbId               *string         `json:"kppbbId" gorm:"type:varchar(2)"`
	TahunPelayanan        *string         `json:"tahunPelayanan" gorm:"type:varchar(4)"`
	BundelPelayanan       *string         `json:"bundlePelayanan" gorm:"type:varchar(4)"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan" gorm:"type:varchar(3)"`
	PermohonanProvinsiID  string          `json:"provinsiKode" gorm:"type:varchar(2)"`
	PermohonanKotaID      string          `json:"daerahKode" gorm:"type:varchar(2)"`
	PermohonanKecamatanID string          `json:"kecamatanKode" gorm:"type:varchar(3)"`
	PermohonanKelurahanID string          `json:"kelurahanKode" gorm:"type:varchar(3)"`
	PermohonanBlokID      string          `json:"blokKode" gorm:"type:varchar(3)"`
	NoUrutPemohon         string          `json:"noUrutKode" gorm:"type:varchar(4)"`
	PemohonJenisOPID      string          `json:"jenisOP" gorm:"type:varchar(1)"`
	JnsSK                 *string         `json:"jnsSK" gorm:"type:varchar(1)"`
	NoSK                  *string         `json:"noSK" gorm:"type:varchar(30)"`
	JnsKeputusan          JnsKeputusan    `json:"jnsKeputusan" gorm:"type:varchar(1);not null"`
	KlsTanah_Kode         *string         `json:"klsTanah_kode" gorm:"type:varchar(3)"`
	KlsTanah_TahunAwal    *string         `json:"klsTanah_tahunAwal" gorm:"type:varchar(4)"`
	KlsBangunan_Kode      *string         `json:"klsBangunan_kode" gorm:"type:varchar(3)"`
	KlsBangunan_TahunAwal *string         `json:"klsBangunan_tahunAwal" gorm:"type:varchar(4)"`
	LuasBumi              *int            `json:"luasBumi" gorm:"type:INTEGER"`
	LuasBangunan          *int            `json:"luasBangunan" gorm:"type:INTEGER"`
	NjopBumi              *int32          `json:"njopBumi" gorm:"type:BIGINT"`
	NjopBangunan          *int32          `json:"njopBangunan" gorm:"type:BIGINT"`
	PBB                   *int32          `json:"pbb" gorm:"type:BIGINT"`
	TanggalCetak          *datatypes.Date `json:"tanggalCetak"`
	NIPPencetak           *string         `json:"nipPencetak" gorm:"type:varchar(9)"`
	gormhelper.DateModel
}

type CreateDtoKepKebPbb struct {
	PermohonanId          *uint64         `json:"permohonanId" validate:"required"`
	KanwilId              *string         `json:"kanwilId"`
	KppbbId               *string         `json:"kppbbId"`
	TahunPelayanan        *string         `json:"tahunPelayanan"`
	BundelPelayanan       *string         `json:"bundlePelayanan"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan"`
	PermohonanProvinsiID  string          `json:"provinsiKode"`
	PermohonanKotaID      string          `json:"daerahKode"`
	PermohonanKecamatanID string          `json:"kecamatanKode"`
	PermohonanKelurahanID string          `json:"kelurahanKode"`
	PermohonanBlokID      string          `json:"blokKode"`
	NoUrutPemohon         string          `json:"noUrutKode"`
	PemohonJenisOPID      string          `json:"jenisOP"`
	JnsSK                 *string         `json:"jnsSK" validate:"required"`
	NoSK                  *string         `json:"noSK"`
	JnsKeputusan          *JnsKeputusan   `json:"jnsKeputusan"`
	KlsTanah_Kode         *string         `json:"klsTanah_Kode"`
	KlsTanah_TahunAwal    *string         `json:"klsTanah_TahunAwal"`
	KlsBangunan_Kode      *string         `json:"klsBangunan_Kode"`
	KlsBangunan_TahunAwal *string         `json:"klsBangunan_TahunAwal"`
	LuasBumi              *int            `json:"luasBumi"`
	LuasBangunan          *int            `json:"luasBangunan"`
	NjopBumi              *int32          `json:"njopBumi"`
	NjopBangunan          *int32          `json:"njopBangunan"`
	PBB                   *int32          `json:"pbb"`
	TanggalCetak          *datatypes.Date `json:"tanggalCetak"`
	NIPPencetak           *string         `json:"nipPencetak"`
}

type FilterDtoKepKebPbb struct {
	PermohonanId          *uint64         `json:"permohonanId"`
	KanwilId              *string         `json:"kanwilId"`
	KppbbId               *string         `json:"kppbbId"`
	TahunPelayanan        *string         `json:"tahunPelayanan"`
	BundelPelayanan       *string         `json:"bundlePelayanan"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan"`
	PermohonanProvinsiID  *string         `json:"provinsiKode"`
	PermohonanKotaID      *string         `json:"daerahKode"`
	PermohonanKecamatanID *string         `json:"kecamatanKode"`
	PermohonanKelurahanID *string         `json:"kelurahanKode"`
	PermohonanBlokID      *string         `json:"blokKode"`
	NoUrutPemohon         *string         `json:"noUrutKode"`
	PemohonJenisOPID      *string         `json:"jenisOP"`
	JnsSK                 *string         `json:"jnsSK"`
	NoSK                  *string         `json:"noSK"`
	JnsKeputusan          *string         `json:"jnsKeputusan"`
	KlsTanah_Kode         *string         `json:"klsTanah_Kode"`
	KlsTanah_TahunAwal    *string         `json:"klsTanah_TahunAwal"`
	KlsBangunan_Kode      *string         `json:"klsBangunan_Kode"`
	KlsBangunan_TahunAwal *string         `json:"klsBangunan_TahunAwal"`
	LuasBumi              *int            `json:"luasBumi"`
	LuasBangunan          *int            `json:"luasBangunan"`
	NjopBumi              *int32          `json:"njopBumi"`
	NjopBangunan          *int32          `json:"njopBangunan"`
	PBB                   *int32          `json:"pbb"`
	TanggalCetak          *datatypes.Date `json:"tanggalCetak"`
	NIPPencetak           *string         `json:"nipPencetak"`
	Page                  *int            `json:"page"`
	PageSize              *int            `json:"page_size"`
}

type UpdateDtoKepKebPbb struct {
	KanwilId              *string         `json:"kanwilId"`
	KppbbId               *string         `json:"kppbbId"`
	TahunPelayanan        *string         `json:"tahunPelayanan"`
	BundelPelayanan       *string         `json:"bundlePelayanan"`
	NoUrutPelayanan       *string         `json:"noUrutPelayanan"`
	PermohonanProvinsiID  *string         `json:"provinsiKode"`
	PermohonanKotaID      *string         `json:"daerahKode"`
	PermohonanKecamatanID *string         `json:"kecamatanKode"`
	PermohonanKelurahanID *string         `json:"kelurahanKode"`
	PermohonanBlokID      *string         `json:"blokKode"`
	NoUrutPemohon         *string         `json:"noUrutKode"`
	PemohonJenisOPID      *string         `json:"jenisOP"`
	JnsSK                 *string         `json:"jnsSK"`
	NoSK                  *string         `json:"noSK"`
	JnsKeputusan          *JnsKeputusan   `json:"jnsKeputusan"`
	KlsTanah_Kode         *string         `json:"klsTanah_Kode"`
	KlsTanah_TahunAwal    *string         `json:"klsTanah_TahunAwal"`
	KlsBangunan_Kode      *string         `json:"klsBangunan_Kode"`
	KlsBangunan_TahunAwal *string         `json:"klsBangunan_TahunAwal"`
	LuasBumi              *int            `json:"luasBumi"`
	LuasBangunan          *int            `json:"luasBangunan"`
	NjopBumi              *int32          `json:"njopBumi"`
	NjopBangunan          *int32          `json:"njopBangunan"`
	PBB                   *int32          `json:"pbb"`
	TanggalCetak          *datatypes.Date `json:"tanggalCetak"`
	NIPPencetak           *string         `json:"nipPencetak"`
}

type FilterNopSkDtoKepKebPbb struct {
	PermohonanProvinsiID  string `json:"provinsiKode" validate:"required"`
	PermohonanKotaID      string `json:"daerahKode" validate:"required"`
	PermohonanKecamatanID string `json:"kecamatanKode" validate:"required"`
	PermohonanKelurahanID string `json:"kelurahanKode" validate:"required"`
	PermohonanBlokID      string `json:"blokKode" validate:"required"`
	NoUrutPemohon         string `json:"noUrutKode" validate:"required"`
	PemohonJenisOPID      string `json:"jenisOP" validate:"required"`
	JnsSK                 string `json:"jnsSK" validate:"required"`
	NoSK                  string `json:"noSK" validate:"required"`
}

type VerifyDtoKepKebPbb struct {
	JnsKeputusan *JnsKeputusan `json:"jnsKeputusan" validate:"required"`
}

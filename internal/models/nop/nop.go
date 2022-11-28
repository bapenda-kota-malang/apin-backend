package nop

import (
	jp "github.com/bapenda-kota-malang/apin-backend/internal/models/jenisperolehan"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Nop struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	Provinsi_Kode     *uint64            `json:"provinsi_kode"`
	Daerah_Kode       *uint64            `json:"daerah_kode"`
	Kecamatan_Kode    *uint64            `json:"kecamatan_kode"`
	Kelurahan_Kode    *uint64            `json:"kelurahan_kode"`
	KodeBlok          *string            `json:"kodeBlok"`
	NoUrut            *string            `json:"noUrut"`
	KodeJenisOp       *string            `json:"kodeJenisOp"`
	OpRtRw            *string            `json:"opRtRw"`
	LuasTanahOp       *float64           `json:"luasTanahOp" gKodem:"type:decimal(10,0)"`
	LuasBangunanOp    *float64           `json:"luasBangunanOp" gorm:"type:decimal(10,0)"`
	NjopTanahOp       *float64           `json:"njopTanahOp" gorm:"type:decimal(20,0)"`
	NjopBangunanOp    *float64           `json:"njopBangunanOp" gorm:"type:decimal(20,0)"`
	NilaiOp           *float64           `json:"nilaiOp" gorm:"type:decimal(20,0)"`
	JenisPerolehan_Id *uint64            `json:"jenisPerolehan_id"`
	JenisPerolehan    *jp.JenisPerolehan `json:"jenisPerolehan,omitempty" gorm:"foreignKey:JenisPerolehan_Id"`
	NoSertifikat      *string            `json:"noSertifikat"`
	NjopPbb           *float64           `json:"njopPbb" gorm:"type:decimal(20,0)"`
	NamaPenjual       *string            `json:"namaPenjual"`
	AlamatPenjual     *string            `json:"alamatPenjual"`
	LokasiOp          *string            `json:"lokasiOp"`
	TahunPajakSppt    *string            `json:"tahunPajakSppt"`
	RefTanah          *float64           `json:"refTanah" gorm:"type:decimal(20,0)"`
	RefBangunan       *float64           `json:"refBangunan" gorm:"type:decimal(20,0)"`
}

type CreateDto struct {
	Provinsi_Kode     *uint64  `json:"provinsi_kode"`
	Daerah_Kode       *uint64  `json:"daerah_kode"`
	Kecamatan_Kode    *uint64  `json:"kecamatan_kode"`
	Kelurahan_Kode    *uint64  `json:"kelurahan_kode"`
	KodeBlok          *string  `json:"kodeBlok"`
	NoUrut            *string  `json:"noUrut"`
	KodeJenisOp       *string  `json:"kodeJenisOp"`
	OpRtRw            *string  `json:"opRtRw"`
	LuasTanahOp       *float64 `json:"luasTanahOp"`
	LuasBangunanOp    *float64 `json:"luasBangunanOp"`
	NjopTanahOp       *float64 `json:"njopTanahOp"`
	NjopBangunanOp    *float64 `json:"njopBangunanOp"`
	NilaiOp           *float64 `json:"nilaiOp"`
	JenisPerolehan_Id *uint64  `json:"jenisPerolehan_id"`
	NoSertifikat      *string  `json:"noSertifikat"`
	NjopPbb           *float64 `json:"njopPbb"`
	NamaPenjual       *string  `json:"namaPenjual"`
	AlamatPenjual     *string  `json:"alamatPenjual"`
	LokasiOp          *string  `json:"lokasiOp"`
	TahunPajakSppt    *string  `json:"tahunPajakSppt"`
	RefTanah          *float64 `json:"refTanah"`
	RefBangunan       *float64 `json:"refBangunan"`
}

type UpdateDto struct {
	Provinsi_Kode     *uint64  `json:"provinsi_kode"`
	Daerah_Kode       *uint64  `json:"daerah_kode"`
	Kecamatan_Kode    *uint64  `json:"kecamatan_kode"`
	Kelurahan_Kode    *uint64  `json:"kelurahan_kode"`
	KodeBlok          *string  `json:"kodeBlok"`
	NoUrut            *string  `json:"noUrut"`
	KodeJenisOp       *string  `json:"kodeJenisOp"`
	OpRtRw            *string  `json:"opRtRw"`
	LuasTanahOp       *float64 `json:"luasTanahOp"`
	LuasBangunanOp    *float64 `json:"luasBangunanOp"`
	NjopTanahOp       *float64 `json:"njopTanahOp"`
	NjopBangunanOp    *float64 `json:"njopBangunanOp"`
	NilaiOp           *float64 `json:"nilaiOp"`
	JenisPerolehan_Id *uint64  `json:"jenisPerolehan_id"`
	NoSertifikat      *string  `json:"noSertifikat"`
	NjopPbb           *float64 `json:"njopPbb"`
	NamaPenjual       *string  `json:"namaPenjual"`
	AlamatPenjual     *string  `json:"alamatPenjual"`
	LokasiOp          *string  `json:"lokasiOp"`
	TahunPajakSppt    *string  `json:"tahunPajakSppt"`
	RefTanah          *float64 `json:"refTanah"`
	RefBangunan       *float64 `json:"refBangunan"`
}

type FilterDto struct {
	Provinsi_Kode     *uint64  `json:"provinsi_kode"`
	Daerah_Kode       *uint64  `json:"daerah_kode"`
	Kecamatan_Kode    *uint64  `json:"kecamatan_kode"`
	Kelurahan_Kode    *uint64  `json:"kelurahan_kode"`
	KodeBlok          *string  `json:"kodeBlok"`
	NoUrut            *string  `json:"noUrut"`
	KodeJenisOp       *string  `json:"kodeJenisOp"`
	OpRtRw            *string  `json:"opRtRw"`
	LuasTanahOp       *float64 `json:"luasTanahOp"`
	LuasBangunanOp    *float64 `json:"luasBangunanOp"`
	NjopTanahOp       *float64 `json:"njopTanahOp"`
	NjopBangunanOp    *float64 `json:"njopBangunanOp"`
	NilaiOp           *float64 `json:"nilaiOp"`
	JenisPerolehan_Id *uint64  `json:"jenisPerolehan_id"`
	NoSertifikat      *string  `json:"noSertifikat"`
	NjopPbb           *float64 `json:"njopPbb"`
	NamaPenjual       *string  `json:"namaPenjual"`
	AlamatPenjual     *string  `json:"alamatPenjual"`
	LokasiOp          *string  `json:"lokasiOp"`
	TahunPajakSppt    *string  `json:"tahunPajakSppt"`
	RefTanah          *float64 `json:"refTanah"`
	RefBangunan       *float64 `json:"refBangunan"`
	Page              int      `json:"page"`
	PageSize          int      `json:"page_size"`
}

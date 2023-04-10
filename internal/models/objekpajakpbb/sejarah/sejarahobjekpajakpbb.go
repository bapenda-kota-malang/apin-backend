package sejarah

import (
	haop "github.com/bapenda-kota-malang/apin-backend/internal/models/anggotaobjekpajak/sejarah"
	hni "github.com/bapenda-kota-malang/apin-backend/internal/models/nilaiindividu/sejarah"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	hopbgn "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbangunan/sejarah"
	hopb "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakbumi/sejarah"
	ot "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajakpbb/types"
	p "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"time"
)

type SejarahObjekPajakPbb struct {
	nop.NopDetail
	WajibPajakPbb_Id      *uint64           `json:"wajibPajakPbb_id"`
	NoFormulirSpop        *string           `json:"noFormulirSpop" gorm:"type:char(11)"`
	NoPersil              *string           `json:"noPersil" gorm:"type:varchar(5)"`
	Jalan                 *string           `json:"jalan" gorm:"type:varchar(30)"`
	BlokKavNo             *string           `json:"blokKavNo" gorm:"type:varchar(15)"`
	Rw                    *string           `json:"rw" gorm:"type:char(2)"`
	Rt                    *string           `json:"rt" gorm:"type:char(3)"`
	StatusCabang          *int              `json:"statusCabang"`
	StatusWp              ot.StatusWp       `json:"statusWp" gorm:"type:char(1)"`
	TotalLuasBumi         *int              `json:"totalLuasBumi"`
	TotalLuasBangunan     *int              `json:"totalLuasBangunan"`
	NjopBumi              *int              `json:"njopBumi"`
	NjopBangunan          *int              `json:"njopBangunan"`
	StatusPeta            *int              `json:"statusPeta"`
	IndeksObjekPajak      *int              `json:"indeksObjekPajak"`
	JenisTransaksi        ot.JenisTransaksi `json:"jenisTransaksi" gorm:"type:varchar(2)"`
	TanggalPendataan      *time.Time        `json:"tanggalPendataan"`
	Pendata_Pegawai_Nip   *string           `json:"pendata_pegawai_nip" gorm:"type:char(9)"`
	PegawaiPendata        *p.Pegawai        `json:"pegawaiPendata,omitempty" gorm:"foreignKey:Pendata_Pegawai_Nip;references:Nip"`
	TanggalPemeriksaan    *time.Time        `json:"tanggalPemeriksaan"`
	Pemeriksa_Pegawai_Nip *string           `json:"pemeriksa_pegawai_nip" gorm:"type:char(9)"`
	PegawaiPemeriksa      *p.Pegawai        `json:"pegawaiPemeriksa,omitempty" gorm:"foreignKey:Pemeriksa_Pegawai_Nip;references:Nip"`
	TanggalPerekaman      *time.Time        `json:"tanggalPerekaman"`
	Perekam_Pegawai_Nip   *string           `json:"perekam_pegawai_nip" gorm:"type:char(9)"`
	PegawaiPerekam        *p.Pegawai        `json:"pegawaiPerekam,omitempty" gorm:"foreignKey:Perekam_Pegawai_Nip;references:Nip"`

	// kelurahan
	// sket
	gh.DateModel
}

type SejarahOpPbbResponse struct {
	StatusCabang      *int        `json:"statusCabang"`
	StatusWp          ot.StatusWp `json:"statusWp"`
	NamaWajibPajak    *string     `json:"namaWajibPajak"`
	TotalLuasBangunan *int        `json:"totalLuasBangunan"`
	TanggalPerekaman  *time.Time  `json:"tanggalPerekaman"`
	NamaPerekam       *string     `json:"namaPerekam"`
}

type SejarahOpResponse struct {
	NamaWajibPajak    *string                            `json:"namaWajibPajak"`
	JalanObjekPajak   *string                            `json:"jalanObjekPajak"`
	JalanWajibPajak   *string                            `json:"jalanWajibPajak"`
	ListOpPbb         []SejarahOpPbbResponse             `json:"listOpPbb"`
	ListOpBumi        []hopb.SejarahOpBumiResponse       `json:"listOpBumi"`
	ListOpBangunan    []hopbgn.SejarahOpBangunanResponse `json:"listOpBangunan"`
	ListNilaiIndividu []hni.SejarahNilaiIndividuResponse `json:"ListNilaiIndividu"`
	ListOpAnggota     []haop.SejarahOpAnggotaResponse    `json:"listOpAnggota"`
}

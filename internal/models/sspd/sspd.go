package sspd

import (
	"time"

	mj "github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mop "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	mp "github.com/bapenda-kota-malang/apin-backend/internal/models/pegawai"
	mr "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/spt"
	mtp "github.com/bapenda-kota-malang/apin-backend/internal/models/tempatpembayaran"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Sspd struct {
	Id                            uint64                `json:"id" gorm:"primaryKey"`
	NomorTahun                    int                   `json:"nomorTahun"`
	NomorUrut                     int                   `json:"nomorUrut"`
	NomorOutput                   string                `json:"nomorOutput"`
	TanggalBayar                  *time.Time            `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64               `json:"objekPajak_id"`
	ObjekPajak                    *mop.ObjekPajak       `json:"objekPajak" gorm:"foreignKey:ObjekPajak_Id"`
	CreatedBy_User_Id             *uint64               `json:"createdBy_user_id"`
	User                          *mu.User              `json:"user" gorm:"foreignKey:CreatedBy_User_Id"`
	BendaharaPenerima_Pegawai_Id  uint64                `json:"bendaharaPenerima_pegawai_id"`
	BendaharaPenerima             *mp.Pegawai           `json:"bendaharaPenerima,omitempty" gorm:"foreignKey:BendaharaPenerima_Pegawai_Id"`
	RekeningBendahara_Rekening_Id *uint64               `json:"rekeningBendahara_rekening_id"`
	Rekening                      *mr.Rekening          `json:"rekening" gorm:"foreignKey:RekeningBendahara_Rekening_Id"`
	IdBt                          *int                  `json:"idBt"`
	Jurnal_Id                     *uint64               `json:"jurnal_id"`
	Jurnal                        *mj.Jurnal            `json:"jurnal" gorm:"foreignKey:Jurnal_Id"`
	Note                          *string               `json:"note" gorm:"type:varchar(255)"`
	PenyetorName                  *string               `json:"penyetorName" gorm:"type:varchar(50)"`
	PenyetorAddress               *string               `json:"penyetorAddress" gorm:"type:varchar(255)"`
	Npwpd_Npwpd                   *string               `json:"npwpd_npwpd" gorm:"type:varchar(22)"`
	Npwpd                         *mn.Npwpd             `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Npwpd;references:Npwpd"`
	Nominal                       *float64              `json:"nominal" gorm:"type:decimal"`
	IsKetetapan                   *bool                 `json:"isKetetapan"`
	IdAktivitas                   *int                  `json:"idAktivitas"`
	IsAngsuran                    *bool                 `json:"isAngsuran"`
	JumlahAngsuran                *int                  `json:"jumlahAngsuran" gorm:"type:smallint"`
	KetetapanDate                 *time.Time            `json:"ketetapanDate"`
	NominalBunga                  *float64              `json:"nominalBunga" gorm:"type:decimal"`
	NominalDenda                  *float64              `json:"nominalDenda" gorm:"type:decimal"`
	Total                         *float64              `json:"total" gorm:"type:decimal"`
	TempatPembayaran_Id           *uint64               `json:"tempatPembayaran_id"`
	TempatPembayaran              *mtp.TempatPembayaran `json:"tempatPembayaran,omitempty" gorm:"foreignKey:TempatPembayaran_Id"`
	gh.DateModel
	IsCancelled   *bool      `json:"isCancelled"`
	CancelledDate *time.Time `json:"cancelledDate"`

	SspdDetails []*SspdDetail `json:"sspdDetail,omitempty" gorm:"foreignKey:Sspd_Id;references:Id"`
}

type CreateDto struct {
	TanggalBayar                  *string  `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64  `json:"objekPajak_id"`
	CreatedBy_User_Id             *uint64  `json:"createdBy_user_id"`
	RekeningBendahara_Rekening_Id *uint64  `json:"rekeningBendahara_rekening_id"`
	BendaharaPenerima_Pegawai_Id  *uint64  `json:"bendaharaPenerima_pegawai_id"`
	IdBt                          *int     `json:"idBt"`
	Jurnal_Id                     *uint64  `json:"jurnal_id"`
	Note                          *string  `json:"note"`
	PenyetorName                  *string  `json:"penyetorName"`
	PenyetorAddress               *string  `json:"penyetorAddress"`
	Npwpd_Npwpd                   *string  `json:"npwpd_npwpd"`
	Nominal                       *float64 `json:"nominal"`
	IsKetetapan                   *bool    `json:"isKetetapan"`
	IdAktivitas                   *int     `json:"idAktivitas"`
	IsAngsuran                    *bool    `json:"isAngsuran"`
	JumlahAngsuran                *int     `json:"jumlahAngsuran"`
	KetetapanDate                 *string  `json:"ketetapanDate"`
	NominalBunga                  *float64 `json:"nominalBunga"`
	NominalDenda                  *float64 `json:"nominalDenda"`
	Total                         *float64 `json:"total"`
	TempatPembayaran_Id           *uint64  `json:"tempatPembayaran_id"`
	gh.DateModel
	IsCancelled   *bool      `json:"isCancelled"`
	CancelledDate *time.Time `json:"cancelledDate"`

	SspdDetail *SspdDetailCreateDto `json:"sspdDetail"`
}

type UpdateDto struct {
	TanggalBayar                  *string              `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64              `json:"objekPajak_id"`
	CreatedBy_User_Id             *uint64              `json:"createdBy_user_id"`
	RekeningBendahara_Rekening_Id *uint64              `json:"rekeningBendahara_rekening_id"`
	BendaharaPenerima_Pegawai_Id  *uint64              `json:"bendaharaPenerima_pegawai_id"`
	IdBt                          *int                 `json:"idBt"`
	Jurnal_Id                     *uint64              `json:"jurnal_id"`
	Note                          *string              `json:"note"`
	PenyetorName                  *string              `json:"penyetorName"`
	PenyetorAddress               *string              `json:"penyetorAddress"`
	Npwpd_Npwpd                   *string              `json:"npwpd_npwpd"`
	Nominal                       *float64             `json:"nominal"`
	IsKetetapan                   *bool                `json:"isKetetapan"`
	IdAktivitas                   *int                 `json:"idAktivitas"`
	IsAngsuran                    *bool                `json:"isAngsuran"`
	JumlahAngsuran                *int                 `json:"jumlahAngsuran"`
	KetetapanDate                 *string              `json:"ketetapanDate"`
	NominalBunga                  *float64             `json:"nominalBunga"`
	NominalDenda                  *float64             `json:"nominalDenda"`
	Total                         *float64             `json:"total"`
	TempatPembayaran_Id           *uint64              `json:"tempatPembayaran_id"`
	SspdDetail                    *SspdDetailUpdateDto `json:"sspdDetail"`
}

type FilterDto struct {
	NomorTahun                    *int       `json:"nomorTahun"`
	NomorUrut                     *int       `json:"nomorUrut"`
	NomorOutput                   *string    `json:"nomorOutput"`
	TanggalBayar                  *string    `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64    `json:"objekPajak_id"`
	CreatedBy_User_Id             *uint64    `json:"createdBy_user_id"`
	BendaharaPenerima_Pegawai_Id  *uint64    `json:"bendaharaPenerima_pegawai_id"`
	RekeningBendahara_Rekening_Id *uint64    `json:"rekeningBendahara_rekening_id"`
	IdBt                          *int       `json:"idBt"`
	Jurnal_Id                     *uint64    `json:"jurnal_id"`
	Note                          *string    `json:"note"`
	PenyetorName                  *string    `json:"penyetorName"`
	PenyetorAddress               *string    `json:"penyetorAddress"`
	Npwpd_Npwpd                   *string    `json:"npwpd_npwpd"`
	Nominal                       *float64   `json:"nominal"`
	IsKetetapan                   *bool      `json:"isKetetapan"`
	IdAktivitas                   *int       `json:"idAktivitas"`
	IsAngsuran                    *bool      `json:"isAngsuran"`
	JumlahAngsuran                *int       `json:"jumlahAngsuran"`
	KetetapanDate                 *time.Time `json:"ketetapanDate"`
	NominalBunga                  *float64   `json:"nominalBunga"`
	NominalDenda                  *float64   `json:"nominalDenda"`
	Total                         *float64   `json:"total"`
	Status                        *string    `json:"status"`
	TempatPembayaran_Id           *uint64    `json:"tempatPembayaran_id"`
	CreatedAt                     *time.Time `json:"createdAt"`
	IsCancelled                   *bool      `json:"isCancelled"`
	CancelledDate                 *time.Time `json:"cancelledDate"`
	TanggalBayar_Opt              *string    `json:"tanggalBayar_opt"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type FilterWpDto struct {
	PeriodeAwal  *datatypes.Date `json:"periodeAwal"`
	PeriodeAkhir *datatypes.Date `json:"periodeAkhir"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type ListLogPayment struct {
	*Sspd
	Spt spt.Spt `json:"spt"`
}

type CancelDto struct {
	IsCancelled *bool `json:"isCancelled"`
}

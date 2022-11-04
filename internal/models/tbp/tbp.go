package tbp

import (
	"time"

	mj "github.com/bapenda-kota-malang/apin-backend/internal/models/jurnal"
	mn "github.com/bapenda-kota-malang/apin-backend/internal/models/npwpd"
	mop "github.com/bapenda-kota-malang/apin-backend/internal/models/objekpajak"
	mr "github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	mu "github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Tbp struct {
	Id                            uint64          `json:"id" gorm:"primaryKey"`
	TbpNumber                     *int            `json:"tbpNumber"`
	TanggalBayar                  *time.Time      `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64         `json:"objekPajak_id"`
	ObjekPajak                    *mop.ObjekPajak `json:"objekPajak" gorm:"foreignKey:ObjekPajak_Id"`
	CreatedBy_User_Id             *uint64         `json:"createdBy_user_id"`
	User                          *mu.User        `json:"user" gorm:"foreignKey:CreatedBy_User_Id"`
	RekeningBendahara_Rekening_Id *uint64         `json:"rekeningBendahara_rekening_id"`
	Rekening                      *mr.Rekening    `json:"rekening" gorm:"foreignKey:RekeningBendahara_Rekening_Id"`
	IdBt                          *int            `json:"idBt"`
	Jurnal_Id                     *uint64         `json:"jurnal_id"`
	Jurnal                        *mj.Jurnal      `json:"jurnal" gorm:"foreignKey:Jurnal_Id"`
	Note                          *string         `json:"note" gorm:"type:varchar(255)"`
	PenyetorName                  *string         `json:"penyetorName" gorm:"type:varchar(50)"`
	PenyetorAddress               *string         `json:"penyetorAddress" gorm:"type:varchar(255)"`
	Npwpd_Npwpd                   *string         `json:"npwpd_npwpd" gorm:"type:varchar(22)"`
	Npwpd                         *mn.Npwpd       `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Npwpd;references:Npwpd"`
	Nominal                       *float64        `json:"nominal" gorm:"type:decimal"`
	IsKetetapan                   *bool           `json:"isKetetapan"`
	IdAktivitas                   *int            `json:"idAktivitas"`
	IsAngsuran                    *bool           `json:"isAngsuran"`
	JumlahAngsuran                *int            `json:"jumlahAngsuran" gorm:"type:smallint"`
	KetetapanDate                 *time.Time      `json:"ketetapanDate"`
	NominalBunga                  *float64        `json:"nominalBunga" gorm:"type:decimal"`
	NominalDenda                  *float64        `json:"nominalDenda" gorm:"type:decimal"`
	Total                         *float64        `json:"total" gorm:"type:decimal"`
	TempatPembayaran              *string         `json:"tempatPembayaran" gorm:"type:varchar(50)"`
	gh.DateModel
	IsCancelled   *bool      `json:"isCancelled"`
	CancelledDate *time.Time `json:"cancelledDate"`

	RincianTbps []*RincianTbp `json:"rincianTbp,omitempty" gorm:"foreignKey:Tbp_Id;references:Id"`
}

type CreateDto struct {
	TbpNumber                     *int       `json:"tbpNumber"`
	TanggalBayar                  *time.Time `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64    `json:"objekPajak_id"`
	CreatedBy_User_Id             *uint64    `json:"createdBy_user_id"`
	RekeningBendahara_Rekening_Id *uint64    `json:"rekeningBendahara_rekening_id"`
	IdBt                          *int       `json:"idBt"`
	Jurnal_Id                     *uint64    `json:"jurnal_id"`
	Note                          *string    `json:"note"`
	PenyetorName                  *string    `json:"penyetorName"`
	PenyetorAddress               *string    `json:"penyetorAddress"`
	Npwpd                         *string    `json:"npwpd"`
	Nominal                       *float64   `json:"nominal"`
	IsKetetapan                   *bool      `json:"isKetetapan"`
	IdAktivitas                   *int       `json:"idAktivitas"`
	IsAngsuran                    *bool      `json:"isAngsuran"`
	JumlahAngsuran                *int       `json:"jumlahAngsuran"`
	KetetapanDate                 *time.Time `json:"ketetapanDate"`
	NominalBunga                  *float64   `json:"nominalBunga"`
	NominalDenda                  *float64   `json:"nominalDenda"`
	Total                         *float64   `json:"total"`
	TempatPembayaran              *string    `json:"tempatPembayaran"`
	gh.DateModel
	IsCancelled   *bool      `json:"isCancelled"`
	CancelledDate *time.Time `json:"cancelledDate"`

	RincianTbp *RincianTbpCreateDto `json:"rincianTbp"`
}

type UpdateDto struct {
	Id                *uint64              `json:"id"`
	Note              *string              `json:"note"`
	CreatedBy_User_Id *uint64              `json:"createdBy_user_id"`
	TempatPembayaran  *string              `json:"tempatPembayaran"`
	RincianTbp        *RincianTbpCreateDto `json:"rincianTbp"`
}

type FilterDto struct {
	TbpNumber                     *int       `json:"tbpNumber"`
	TanggalBayar                  *time.Time `json:"tanggalBayar"`
	ObjekPajak_Id                 *uint64    `json:"objekPajak_id"`
	CreatedBy_User_Id             *uint64    `json:"createdBy_user_id"`
	RekeningBendahara_Rekening_Id *uint64    `json:"rekeningBendahara_rekening_id"`
	IdBt                          *int       `json:"idBt"`
	Jurnal_Id                     *uint64    `json:"jurnal_id"`
	Note                          *string    `json:"note"`
	PenyetorName                  *string    `json:"penyetorName"`
	PenyetorAddress               *string    `json:"penyetorAddress"`
	Npwpd                         *string    `json:"npwpd"`
	Nominal                       *float64   `json:"nominal"`
	IsKetetapan                   *bool      `json:"isKetetapan"`
	IdAktivitas                   *int       `json:"idAktivitas"`
	IsAngsuran                    *bool      `json:"isAngsuran"`
	JumlahAngsuran                *int       `json:"jumlahAngsuran"`
	KetetapanDate                 *time.Time `json:"ketetapanDate"`
	NominalBunga                  *float64   `json:"nominalBunga"`
	NominalDenda                  *float64   `json:"nominalDenda"`
	Total                         *float64   `json:"total"`
	TempatPembayaran              *string    `json:"tempatPembayaran"`
	CreatedAt                     *time.Time `json:"createdAt"`
	IsCancelled                   *bool      `json:"isCancelled"`
	CancelledDate                 *time.Time `json:"cancelledDate"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

type CancelDto struct {
	IsCancelled *bool `json:"isCancelled"`
}

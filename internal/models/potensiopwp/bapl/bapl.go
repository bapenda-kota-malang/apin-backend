package bapl

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/lib/pq"
)

type Bapl struct {
	Id                        uint                   `json:"id" gorm:"primaryKey"`
	TanggalPeninjauan         time.Time              `json:"tanggalPeninjauan"`
	Potensiop_Id              uint                   `json:"potensiop_id"`
	PotensiPemilikWp_Id       uint                   `json:"potensiPemilikWp_id"`
	CreateBy_User_Id          uint                   `json:"createBy_user_id"`
	DetailPotensiPajak_Id     uint                   `json:"detailPotensiPajak_id"`
	Ref_DetailPotensiPajak_Id uint                   `json:"ref_detailPotensiPajak_id"`
	Koordinator_User_Id       uint                   `json:"koordinator_user_id"`
	Petugas_User_Id           *pq.Int64Array         `json:"petugas_user_id" gorm:"type:integer[]"`
	Status                    types.StatusVerifikasi `json:"status"`
	Kasubid_User_Id           *uint                  `json:"kasubid_user_id,omitempty"`
	Kabid_User_Id             *uint                  `json:"kabid_user_id,omitempty"`
	gormhelper.DateModel
	// PotensiPemilikWp *potensipemilikwp.PotensiPemilikWp `json:"potensiPemilikWp,omitempty"`
	CreateBy    *user.User `json:"createBy,omitempty" gorm:"foreignKey:CreateBy_User_Id"`
	Koordinator *user.User `json:"koordinator,omitempty" gorm:"foreignKey:Koordinator_User_Id"`
	Kasubid     *user.User `json:"kasubid,omitempty" gorm:"foreignKey:Kasubid_User_Id"`
	Kabid       *user.User `json:"kabid,omitempty" gorm:"foreignKey:Kabid_User_Id"`
}

type CreateDto struct {
	TanggalPeninjauan         *time.Time             `json:"tanggalPeninjauan"`
	Potensiop_Id              uint                   `json:"-"`
	PotensiPemilikWp_Id       uint                   `json:"potensiPemilikWp_id"`
	CreateBy_User_Id          uint                   `json:"-"`
	DetailPotensiPajak_Id     uint                   `json:"detailPotensiPajak_id"`
	Ref_DetailPotensiPajak_Id uint                   `json:"ref_detailPotensiPajak_id"`
	Koordinator_User_Id       uint                   `json:"koordinator_user_id"`
	Petugas_User_Id           *pq.Int64Array         `json:"petugas_user_id"`
	Status                    types.StatusVerifikasi `json:"-"`
}

type UpdateDto struct {
	TanggalPeninjauan   *time.Time     `json:"tanggalPeninjauan"`
	Koordinator_User_Id *uint          `json:"koordinator_user_id"`
	Petugas_User_Id     *pq.Int64Array `json:"petugas_user_id"`
}

type VerifyDto struct {
	Status          string `json:"status" validate:"required"`
	Kasubid_User_Id *uint  `json:"-"`
	Kabid_User_Id   *uint  `json:"-"`
}

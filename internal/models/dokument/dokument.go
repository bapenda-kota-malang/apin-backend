package dokument

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"gorm.io/datatypes"
)

type Dokument struct {
	Id             uint64          `json:"id" gorm:"primaryKey"`
	Kanwil_Kd      *string         `json:"kanwil_kd" gorm:"type:char(2)"`
	Kppbb_Kd       *string         `json:"kppbb_kd" gorm:"type:char(2)"`
	JenisDokumen   *string         `json:"jenisDokumen" gorm:"type:char(1)"`
	NoDokumen      *string         `json:"noDokumen" gorm:"type:varchar(11)"`
	TglPendataan   *datatypes.Date `json:"tglPendataan"`
	NipPendataan   *string         `json:"nipPendataan" gorm:"type:varchar(9)"`
	TglPemeriksaan *datatypes.Date `json:"tglPemeriksaan"`
	NipPemeriksaan *string         `json:"nipPemeriksaan" gorm:"type:varchar(9)"`
	NipPerekaman   *string         `json:"nipPerekaman" gorm:"type:varchar(9)"`
	gormhelper.DateModel
}

type CreateDto struct {
	Kanwil_Kd      *string         `json:"kanwil_kd"`
	Kppbb_Kd       *string         `json:"kppbb_kd"`
	JenisDokumen   *string         `json:"jenisDokumen"`
	NoDokumen      *string         `json:"noDokumen"`
	TglPendataan   *datatypes.Date `json:"tglPendataan"`
	NipPendataan   *string         `json:"nipPendataan"`
	TglPemeriksaan *datatypes.Date `json:"tglPemeriksaan"`
	NipPemeriksaan *string         `json:"nipPemeriksaan"`
	NipPerekaman   *string         `json:"nipPerekaman"`
}

type FilterDto struct {
	Kanwil_Kd      *string         `json:"kanwil_kd"`
	Kppbb_Kd       *string         `json:"kppbb_kd"`
	JenisDokumen   *string         `json:"jenisDokumen"`
	NoDokumen      *string         `json:"noDokumen"`
	TglPendataan   *datatypes.Date `json:"tglPendataan"`
	NipPendataan   *string         `json:"nipPendataan"`
	TglPemeriksaan *datatypes.Date `json:"tglPemeriksaan"`
	NipPemeriksaan *string         `json:"nipPemeriksaan"`
	NipPerekaman   *string         `json:"nipPerekaman"`
	Page           int             `json:"page"`
	PageSize       int             `json:"page_size"`
	// NoPagination   bool    `json:"no_pagination"`
}

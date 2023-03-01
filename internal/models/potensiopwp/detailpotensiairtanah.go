package potensiopwp

import (
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiAirTanah struct {
	Id           uint          `json:"id" gorm:"primaryKey"`
	Potensiop_Id uuid.UUID     `json:"potensiop_id" gorm:"type:uuid"`
	Peruntukan   mt.Peruntukan `json:"peruntukan" gorm:"type:varchar(100)"`
	JenisAbt     mt.JenisABT   `json:"jenisAbt" gorm:"type:varchar(20)"`
	JumlahSumber string        `json:"jumlahSumber"`
	DiameterPipa string        `json:"diameterPipa"`
	gormhelper.DateModel
}

type CreateDtoDPAirTanah struct {
	Potensiop_Id uuid.UUID     `json:"-"`
	Peruntukan   mt.Peruntukan `json:"peruntukan"`
	JenisAbt     mt.JenisABT   `json:"jenisAbt"`
	JumlahSumber string        `json:"jumlahSumber"`
	DiameterPipa string        `json:"diameterPipa"`
}

type CreateDtoAirTanah struct {
	CreateDto
	DetailPajakDtos []CreateDtoDPAirTanah `json:"detailPajaks"`
}

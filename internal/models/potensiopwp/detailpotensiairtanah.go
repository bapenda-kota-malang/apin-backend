package potensiopwp

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/hargadasarair"
	mt "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
)

type DetailPotensiAirTanah struct {
	Id               uint          `json:"id" gorm:"primaryKey"`
	Potensiop_Id     uuid.UUID     `json:"potensiop_id" gorm:"type:uuid"`
	HargaDasarAir_Id uint64        `json:"hargaDasarAir_id"`
	Peruntukan       mt.Peruntukan `json:"peruntukan" gorm:"type:varchar(100)"`
	JenisAbt         mt.JenisABT   `json:"jenisAbt" gorm:"type:varchar(20)"`
	JumlahSumber     string        `json:"jumlahSumber"`
	DiameterPipa     string        `json:"diameterPipa"`
	gormhelper.DateModel
	HargaDasarAir hargadasarair.HargaDasarAir `json:"hargaDasarAir,omitempty" gorm:"foreignKey:HargaDasarAir_Id"`
}

type CreateDtoDPAirTanah struct {
	Potensiop_Id     uuid.UUID     `json:"-"`
	HargaDasarAir_Id uint64        `json:"-"`
	Peruntukan       mt.Peruntukan `json:"peruntukan"`
	JenisAbt         mt.JenisABT   `json:"jenisAbt"`
	JumlahSumber     string        `json:"jumlahSumber"`
	DiameterPipa     string        `json:"diameterPipa"`
}

type CreateDtoAirTanah struct {
	CreateDto
	DetailPajakDtos CreateDtoDPAirTanah `json:"detailPajaks"`
}

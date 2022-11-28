package potensiopwp

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailobjek"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type PotensiOp struct {
	Id             uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Assessment     t.JenisPajak    `json:"assessment" gorm:"size:10"`
	Golongan       t.Golongan      `json:"golongan"`
	Npwp           *string         `json:"npwp" gorm:"size:50"`
	Rekening_Id    uint            `json:"rekening_id"`
	Status         t.Status        `json:"status"`
	ClosingDate    *time.Time      `json:"closingDate"`
	OpeningDate    *time.Time      `json:"openingDate"`
	StartDate      *datatypes.Date `json:"startDate"`
	User_Id        uint            `json:"user_id"`
	LuasBangunan   *string         `json:"luasBangunan" gorm:"size:50"`
	JamBuka        *string         `json:"jamBuka" gorm:"size:50"`
	JamTutup       *string         `json:"jamTutup" gorm:"size:50"`
	Visitors       *string         `json:"visitors" gorm:"size:50"`
	OmsetOp        *string         `json:"omsetOp" gorm:"size:50"`
	Genset         bool            `json:"genset"`
	AirTanah       bool            `json:"airTanah"`
	VendorEtax_Id  *uint           `json:"vendorEtax_id"`
	FotoKtp        *string         `json:"fotoKtp"`
	FotoObjek      *pq.StringArray `json:"fotoObjek" gorm:"type:varchar[]"`
	FormBapl       *string         `json:"formBapl"`
	DokumenLainnya *pq.StringArray `json:"dokumenLainnya" gorm:"type:varchar[]"`
	gormhelper.DateModel
	Rekening *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	User     *user.User         `json:"user,omitempty" gorm:"foreignKey:User_Id"`
	// VendorEtax             cm.VendorEtax           `gorm:"foreignKey:VendorEtax_Id"`
	DetailPotensiOp        *detailpotensiop.DetailPotensiOp       `json:"detailPotensiOp,omitempty" gorm:"foreignKey:Potensiop_Id"`
	PotensiPemilikWp       *[]potensipemilikwp.PotensiPemilikWp   `json:"potensiPemilikWp,omitempty" gorm:"foreignKey:Potensiop_Id"`
	PotensiNarahubung      *[]potensinarahubung.PotensiNarahubung `json:"potensiNarahubung,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiAirTanahs *[]detailobjek.DetailPotensiAirTanah   `json:"detailPotensiAirTanahs,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiHiburans  *[]detailobjek.DetailPotensiHiburan    `json:"detailPotensiHiburans,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiHotels    *[]detailobjek.DetailPotensiHotel      `json:"detailPotensiHotels,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiPPJs      *[]detailobjek.DetailPotensiPPJ        `json:"detailPotensiPpjs,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiParkirs   *[]detailobjek.DetailPotensiParkir     `json:"DetailPotensiParkirs,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiReklames  *[]detailobjek.DetailPotensiReklame    `json:"DetailPotensiReklames,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiRestos    *[]detailobjek.DetailPotensiResto      `json:"DetailPotensiRestos,omitempty" gorm:"foreignKey:Potensiop_Id"`
	Bapl                   *bapl.PotensiBapl                      `json:"bapl,omitempty" gorm:"foreignKey:Potensiop_Id"`
}

type CreatePotensiOpDto struct {
	Assessment     string          `json:"assessment"`
	Golongan       t.Golongan      `json:"golongan" validate:"required;min=1"`
	Npwp           *string         `json:"npwp"`
	Rekening_Id    uint            `json:"rekening_id" validate:"required;min=1"`
	ClosingDate    *time.Time      `json:"closingDate"`
	OpeningDate    *time.Time      `json:"openingDate"`
	StartDate      *datatypes.Date `json:"startDate"`
	User_Id        uint            `json:"-"`
	LuasBangunan   *string         `json:"luasBangunan"`
	JamBuka        *string         `json:"jamBuka"`
	JamTutup       *string         `json:"jamTutup"`
	Visitors       *string         `json:"visitors"`
	OmsetOp        *string         `json:"omsetOp"`
	Genset         bool            `json:"genset"`
	AirTanah       bool            `json:"airTanah"`
	VendorEtax_Id  *uint           `json:"vendorEtax_id"`
	FotoKtp        *string         `json:"fotoKtp" validate:"base64=image;b64size=1025"`
	FotoObjek      *pq.StringArray `json:"fotoObjek" validate:"base64=image;b64size=1025"`
	FormBapl       *string         `json:"formBapl" validate:"base64=pdf;b64size=1025"`
	DokumenLainnya *pq.StringArray `json:"dokumenLainnya" validate:"base64=pdf,image,excel;b64size=1025"`
}

type UpdatePotensiOpDto struct {
	Assessment            *string         `json:"assessment"`
	Golongan              *t.Golongan     `json:"golongan" validate:"min=1"`
	Npwp                  *string         `json:"npwp"`
	Rekening_Id           *uint           `json:"rekening_id" validate:"min=1"`
	ClosingDate           *time.Time      `json:"closingDate"`
	OpeningDate           *time.Time      `json:"openingDate"`
	StartDate             *datatypes.Date `json:"startDate"`
	User_Id               *uint           `json:"-"`
	LuasBangunan          *string         `json:"luasBangunan"`
	JamBuka               *string         `json:"jamBuka"`
	JamTutup              *string         `json:"jamTutup"`
	Visitors              *string         `json:"visitors"`
	OmsetOp               *string         `json:"omsetOp"`
	Genset                *bool           `json:"genset"`
	AirTanah              *bool           `json:"airTanah"`
	VendorEtax_Id         *uint           `json:"vendorEtax_id"`
	FotoKtp               *string         `json:"fotoKtp" validate:"base64=image;b64size=1025"`
	FotoObjek             *pq.StringArray `json:"fotoObjek" validate:"base64=image;b64size=1025"`
	FormBapl              *string         `json:"formBapl" validate:"base64=pdf;b64size=1025"`
	DokumenLainnya        *pq.StringArray `json:"dokumenLainnya" validate:"base64=pdf,image,excel;b64size=1025"`
	FotoObjekDeleted      *[]string       `json:"fotoObjekDeleted"`
	DokumenLainnyaDeleted *[]string       `json:"dokumenLainnyaDeleted"`
}

type FilterDto struct {
	// dynamic from column table field
	// potensiop
	Assessment   *string     `json:"assessment"`
	Golongan     *t.Golongan `json:"golongan"`
	Status       *t.Status   `json:"status"`
	ClosingDate  *time.Time  `json:"closingDate"`
	OpeningDate  *time.Time  `json:"openingDate"`
	LuasBangunan *string     `json:"luasBangunan"`
	JamBuka      *string     `json:"jamBuka"`
	JamTutup     *string     `json:"jamTutup"`
	Visitors     *string     `json:"visitors"`
	OmsetOp      *string     `json:"omsetOp"`
	Genset       *string     `json:"genset"`
	AirTanah     *string     `json:"airTanah"`
	// detail potensiop
	JenisOp  *string `json:"jenisOp"`
	JumlahOp *string `json:"jumlahOp"`
	TarifOp  *string `json:"tarifOp"`
	UnitOp   *string `json:"unitOp"`
	Notes    *string `json:"notes"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

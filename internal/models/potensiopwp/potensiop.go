package potensiopwp

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/bapl"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/detailpotensiop"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensinarahubung"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/potensiopwp/potensipemilikwp"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/tarifpajak"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type PotensiOp struct {
	Id               uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Assessment       t.JenisPajak    `json:"assessment" gorm:"size:10"`
	Golongan         t.Golongan      `json:"golongan"`
	Npwp             *string         `json:"npwp" gorm:"size:50"`
	Rekening_Id      uint64          `json:"rekening_id"`
	StartDate        *datatypes.Date `json:"startDate"`
	LuasBangunan     *string         `json:"luasBangunan" gorm:"size:50"`
	JamBuka          *string         `json:"jamBuka" gorm:"size:50"`
	JamTutup         *string         `json:"jamTutup" gorm:"size:50"`
	Visitors         *string         `json:"visitors" gorm:"size:50"`
	OmsetOp          float64         `json:"omsetOp"`
	JumlahPajak      float64         `json:"jumlahPajak"`
	TarifPajak_Id    *uint64         `json:"tarifPajak_id"`
	Genset           bool            `json:"genset"`
	AirTanah         bool            `json:"airTanah"`
	PajakKonsumenPct *float64        `json:"pajakKonsumenPcts"`
	User_Id          uint64          `json:"user_id"`
	Status           t.Status        `json:"status"`
	VendorEtax_Id    *uint64         `json:"vendorEtax_id"`
	FotoKtp          *string         `json:"fotoKtp"`
	FotoObjek        *pq.StringArray `json:"fotoObjek" gorm:"type:varchar[]"`
	FormBapl         *string         `json:"formBapl"`
	DokumenLainnya   *pq.StringArray `json:"dokumenLainnya" gorm:"type:varchar[]"`
	gormhelper.DateModel
	Rekening   *rekening.Rekening     `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
	User       *user.User             `json:"user,omitempty" gorm:"foreignKey:User_Id"`
	TarifPajak *tarifpajak.TarifPajak `json:"tarifPajak,omitempty" gorm:"foreignKey:TarifPajak_Id"`
	// VendorEtax             cm.VendorEtax           `gorm:"foreignKey:VendorEtax_Id"`
	DetailPotensiOp        *detailpotensiop.DetailPotensiOp      `json:"detailPotensiOp,omitempty" gorm:"foreignKey:Potensiop_Id"`
	PotensiPemilikWp       []potensipemilikwp.PotensiPemilikWp   `json:"potensiPemilikWp,omitempty" gorm:"foreignKey:Potensiop_Id"`
	PotensiNarahubung      []potensinarahubung.PotensiNarahubung `json:"potensiNarahubung,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiAirTanah  *DetailPotensiAirTanah                `json:"potensiAirTanah,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiHiburans  []DetailPotensiHiburan                `json:"potensiHiburans,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiHotels    []DetailPotensiHotel                  `json:"potensiHotels,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiParkirs   []DetailPotensiParkir                 `json:"potensiParkirs,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiPPJNonPLN *DetailPotensiPPJNonPLN               `json:"potensiPpjNonPLN,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiReklames  []DetailPotensiReklame                `json:"potensiReklames,omitempty" gorm:"foreignKey:Potensiop_Id"`
	DetailPotensiResto     *DetailPotensiResto                   `json:"potensiResto,omitempty" gorm:"foreignKey:Potensiop_Id"`
	Bapl                   []bapl.PotensiBapl                    `json:"bapl,omitempty" gorm:"foreignKey:Potensiop_Id"`
}

type CreatePotensiOpDto struct {
	Id             uuid.UUID       `json:"-"`
	Assessment     t.JenisPajak    `json:"assessment"`
	Golongan       t.Golongan      `json:"golongan" validate:"required;min=1"`
	Npwp           *string         `json:"npwp"`
	Rekening_Id    uint64          `json:"rekening_id" validate:"required;min=1"`
	StartDate      *datatypes.Date `json:"startDate"`
	LuasBangunan   *string         `json:"luasBangunan"`
	JamBuka        *string         `json:"jamBuka"`
	JamTutup       *string         `json:"jamTutup"`
	Visitors       *string         `json:"visitors"`
	OmsetOp        *float64        `json:"omsetOp" validate:"required"`
	JumlahPajak    float64         `json:"jumlahPajak"`
	TarifPajak_Id  *uint64         `json:"-"`
	Genset         bool            `json:"genset"`
	AirTanah       bool            `json:"airTanah"`
	PajakKonsumen  *float64        `json:"pajakKonsumen"`
	User_Id        uint64          `json:"user_id"`
	Status         t.Status        `json:"status"`
	VendorEtax_Id  *uint64         `json:"vendorEtax_id"`
	FotoKtp        *string         `json:"fotoKtp" validate:"base64=image;b64size=1025"`
	FotoObjek      *pq.StringArray `json:"fotoObjek" validate:"base64=image;b64size=1025"`
	FormBapl       *string         `json:"formBapl" validate:"base64=pdf;b64size=1025"`
	DokumenLainnya *pq.StringArray `json:"dokumenLainnya" validate:"base64=pdf,image,excel;b64size=1025"`
}

// THIS STRUCT DEPRECATED BECAUSE POTENSI OP WP CAN'T UPDATE DATA
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
	Assessment     *string         `json:"assessment"`
	Golongan       *t.Golongan     `json:"golongan"`
	Npwp           *string         `json:"npwp" gorm:"size:50"`
	Rekening_Id    *uint64         `json:"rekening_id"`
	Rekening_Objek *string         `json:"rekening_objek" refsource:"Rekening.Objek"`
	NamaOp         *string         `json:"namaOp" refsource:"DetailPotensiOp.Nama"`
	NamaOp_Opt     *string         `json:"namaOp_opt" refsource:"DetailPotensiOp.Nama"`
	StartDate      *datatypes.Date `json:"startDate"`
	LuasBangunan   *string         `json:"luasBangunan"`
	JamBuka        *string         `json:"jamBuka"`
	JamTutup       *string         `json:"jamTutup"`
	Visitors       *string         `json:"visitors"`
	OmsetOp        *string         `json:"omsetOp"`
	JumlahPajak    *float64        `json:"jumlahPajak"`
	TarifPajak_Id  *uint64         `json:"tarifPajak_id"`
	Genset         *string         `json:"genset"`
	AirTanah       *string         `json:"airTanah"`
	Status         *t.Status       `json:"status"`
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

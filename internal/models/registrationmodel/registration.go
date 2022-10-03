package registration

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Golongan int16
type StatusPendaftaran int16
type JenisPajak string
type StatusObjekPajak int16
type Mode int16
type VerifyPendaftaran int16
type StatusPemilik string
type StatusNarahubung string

const (
	GolonganOrangPribadi Golongan = 1 //orang pribadi
	GolonganBadan        Golongan = 2 //badan

	StatusAktif          StatusPendaftaran = 1 //aktif
	StatusTutupSementara StatusPendaftaran = 2 //tutup sementara
	StatusTutup          StatusPendaftaran = 3 //tutup

	JenisPajakSelfAssessment     JenisPajak = "SA" //self_assessment
	JenisPajakOfficialAssessment JenisPajak = "OS" //official_assessment

	StatusBaru StatusObjekPajak = 0 //baru
	StatusLama StatusObjekPajak = 1 //lama

	ModeSelf     Mode = 1 //auto
	ModeOperator Mode = 2 //manual

	VerifiyPendaftaranBaru              VerifyPendaftaran = 1 //baru
	VerifiyPendaftaranDisetujui         VerifyPendaftaran = 2 //disetujui
	VerifiyPendaftaranDisetujuiModeAuto VerifyPendaftaran = 3 //disetujui jika verifymode = auto

	StatusPemilikBaru StatusPemilik = "baru"
	StatusPemilikLama StatusPemilik = "lama"

	StatusNarahubungBaru StatusNarahubung = "baru"
	StatusNarahubungLama StatusNarahubung = "lama"
)

type (
	Registration struct {
		Id                uint64             `json:"id" gorm:"primarykey"`
		Golongan          Golongan           `json:"golongan"`
		Nomor             *string            `json:"nomor" gorm:"size:10"`
		NomorRegistrasi   *string            `json:"nomorRegistrasi" gorm:"size:50"`
		Npwp              *string            `json:"npwp" gorm:"size:50"`
		TanggalPengukuhan *time.Time         `json:"tanggalPengukuhan"`
		TanggalNpwpd      *time.Time         `json:"tanggalNpwpd"`
		Npwpd             *string            `json:"npwpd" gorm:"size:22"`
		Status            StatusPendaftaran  `json:"status"`
		TanggalTutup      *time.Time         `json:"tanggalTutup"`
		TanggalBuka       *time.Time         `json:"tanggalBuka"`
		JenisPajak        JenisPajak         `json:"jenisPajak" gorm:"size:2"`
		Skpd_Id           *uint64            `json:"skpd_id"`
		Skpd              *skpd.Skpd         `json:"skpd,omitempty" gorm:"foreignKey:Skpd_Id"`
		Rekening_Id       *uint64            `json:"rekening_id"`
		Rekening          *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_Id"`
		Nama              *string            `json:"nama" gorm:"size:20"`
		User_Id           *uint64            `json:"user_id"`
		User              *user.User         `gorm:"foreignKey:User_Id"`
		OmsetOp           *string            `json:"omsetOp" gorm:"size:50"`
		Genset            *bool              `json:"genset"`
		AirTanah          *bool              `json:"airTanah"`
		TanggalMulaiUsaha *time.Time         `json:"tanggalMulaiUsaha"`
		LuasBangunan      *string            `json:"luasBangunan" gorm:"size:50"`
		JamBukaUsaha      *string            `json:"jamBukaUsaha" gorm:"size:50"`
		JamTutupUsaha     *string            `json:"jamTutupUsaha" gorm:"size:50"`
		Pengunjung        *string            `json:"pengunjung" gorm:"size:50"`
		FotoKtp           *string            `json:"fotoKtp" gorm:"size:50"`
		SuratIzinUsaha    *string            `json:"suratIzinUsaha" gorm:"size:50"`
		LainLain          *string            `json:"lainLain" gorm:"size:50"`
		FotoObjek         *string            `json:"fotoObjek" gorm:"size:50"`
		Mode              Mode               `json:"mode"`
		VerifyStatus      *VerifyPendaftaran `json:"verifyStatus"`
		VerifyAt          *time.Time         `json:"verifyAt"`
		VendorEtaxID      *string            `json:"vendor_etax_id"`
		// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

		ObjekPajaks []*ObjekPajak `json:"objekPajak,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		PemilikWps  []*PemilikWp  `json:"pemilik,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		Narahubungs []*Narahubung `json:"narahubung,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`

		DetailOpAirTanah []*DetailOpAirTanah `json:"detail_op_air_tanah,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		DetailOpHiburan  []*DetailOpHiburan  `json:"detail_op_hiburan,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		DetailOpHotel    []*DetailOpHotel    `json:"detail_op_hotel,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		DetailOpParkir   []*DetailOpParkir   `json:"detail_op_parkir,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		DetailOpPpj      []*DetailOpPpj      `json:"detail_op_ppj,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		DetailOpReklame  []*DetailOpReklame  `json:"detail_op_reklame,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`
		DetailOpResto    []*DetailOpResto    `json:"detail_op_resto,omitempty" gorm:"foreignKey:Pendaftaran_Id;references:Id"`

		gormhelper.DateModel
	}

	ObjekPajak struct {
		Id             uint64                  `json:"id" gorm:"primarykey"`
		Pendaftaran_Id uint64                  `json:"pendaftaran_id"`
		Pendaftaran    *Registration           `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_Id"`
		Nama           *string                 `json:"nama" gorm:"size:200"`
		Nop            *string                 `json:"nop" gorm:"size:50"`
		Alamat         *string                 `json:"alamat" gorm:"size:200"`
		RtRw           *string                 `json:"rtRw" gorm:"size:10"`
		Kecamatan_Id   *uint64                 `json:"kecamatan_id"`
		Kecamatan      *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
		Kelurahan_Id   *uint64                 `json:"kelurahan_id"`
		Kelurahan      *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
		Telp           *string                 `json:"telp" gorm:"size:20"`
		Status         StatusObjekPajak        `json:"status"`
		IsNpwpd        uint64                  `json:"isNpwpd"`
		gormhelper.DateModel
	}

	DetailOp struct {
		Id                uint64        `json:"id" gorm:"primaryKey"`
		Pendaftaran_Id    uint64        `json:"pendaftaran_id"`
		Pendaftaran       *Registration `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_Id"`
		Pendaftaran_Npwpd *string       `json:"pendaftaran_npwpd" gorm:"size:22"`
		JenisOp           *string       `json:"jenisOp" gorm:"size:200"`
		JumlahOp          *string       `json:"jumlahOp" gorm:"size:200"`
		TarifOp           *string       `json:"tarifOp" gorm:"size:200"`
		UnitOp            *string       `json:"unitOp" gorm:"size:50"`
		Notes             *string       `json:"notes" gorm:"size:200"`
	}

	DetailOpAirTanah struct {
		DetailOp
	}

	DetailOpHiburan struct {
		DetailOp
	}

	DetailOpHotel struct {
		DetailOp
	}

	DetailOpParkir struct {
		DetailOp
	}

	DetailOpPpj struct {
		DetailOp
	}

	DetailOpReklame struct {
		DetailOp
	}

	DetailOpResto struct {
		DetailOp
	}

	PemilikWp struct {
		Id             uint64                  `json:"id" gorm:"primaryKey"`
		Pendaftaran_Id uint64                  `json:"pendaftaran_id"`
		Pendaftaran    *Registration           `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_Id"`
		Nama           *string                 `json:"nama" gorm:"size:50"`
		Alamat         *string                 `json:"alamat" gorm:"size:50"`
		RtRw           *string                 `json:"rtRw" gorm:"size:10"`
		Kecamatan_Id   *uint64                 `json:"kecamatan_id"`
		Kecamatan      *areadivision.Kecamatan `gorm:"foreignKey:Kecamatan_Id"`
		Kelurahan_Id   *uint64                 `json:"kelurahan_id"`
		Kelurahan      *areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
		Telp           *string                 `json:"telp" gorm:"size:20"`
		Status         StatusPemilik           `json:"status"`
		NoIdPemilik    *string                 `json:"noIdPemilik" gorm:"size:20"`
		Nik            *string                 `json:"nik" gorm:"size:20"`
	}

	Narahubung struct {
		Id             uint64                 `json:"id" gorm:"primaryKey"`
		Pendaftaran_Id uint64                 `json:"pendaftaran_id"`
		Pendaftaran    Registration           `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_Id"`
		Nama           string                 `json:"nama" gorm:"size:50"`
		Alamat         string                 `json:"alamat" gorm:"size:50"`
		RtRw           string                 `json:"rtRw" gorm:"size:10"`
		Kecamatan_Id   uint64                 `json:"kecamatan_id"`
		Kecamatan      areadivision.Kecamatan `gorm:"foreignKey:Kecamatan_Id"`
		Kelurahan_Id   uint64                 `json:"kelurahan_id"`
		Kelurahan      areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_Id"`
		Telp           string                 `json:"telp" gorm:"size:20"`
		Status         StatusNarahubung       `json:"status"`
		Nik            string                 `json:"nik" gorm:"size:20"`
	}

	RegisterByOperator struct {
		JenisPajak JenisPajak `json:"jenisPajak"`
		Golongan   Golongan   `json:"golongan"`
		Npwp       *string    `json:"npwp"`

		NomorRegistrasi       *string `json:"nomorRegistrasi"`
		IsNomorRegistrasiAuto bool    `json:"isNomorRegistrasiAuto"`

		Npwpd             *string `json:"npwpd"`
		TanggalPengukuhan *string `json:"tanggalPengukuhan"`
		TanggalNpwpd      *string `json:"tanggalNpwpd"`

		TanggalMulaiUsaha *string `json:"tanggalMulaiUsaha"`
		LuasBangunan      *string `json:"luasBangunan"`
		JamBukaUsaha      *string `json:"jamBukaUsaha"`
		JamTutupUsaha     *string `json:"jamTutupUsaha"`
		Pengunjung        *string `json:"pengunjung"`
		OmsetOp           *string `json:"omsetOp"`

		Rekening_Id uint64 `json:"rekening_id"`

		Genset   bool `json:"genset"`
		AirTanah bool `json:"airTanah"`

		DetailOp *[]DetailOp `json:"detail_op"`

		ObjekPajak *[]ObjekPajak `json:"objekPajak"`
		Pemilik    *[]PemilikWp  `json:"pemilik"`
		Narahubung *[]Narahubung `json:"narahubung"`
	}
)

type RegisterUpdate struct {
	JenisPajak JenisPajak `json:"jenisPajak"`
	Golongan   Golongan   `json:"golongan"`
	Npwp       *string    `json:"npwp"`

	NomorRegistrasi       *string `json:"nomorRegistrasi"`
	IsNomorRegistrasiAuto bool    `json:"isNomorRegistrasiAuto"`

	Npwpd             *string `json:"npwpd"`
	TanggalPengukuhan *string `json:"tanggalPengukuhan"`
	TanggalNpwpd      *string `json:"tanggalNpwpd"`

	TanggalMulaiUsaha *string `json:"tanggalMulaiUsaha"`
	LuasBangunan      *string `json:"luasBangunan"`
	JamBukaUsaha      *string `json:"jamBukaUsaha"`
	JamTutupUsaha     *string `json:"jamTutupUsaha"`
	Pengunjung        *string `json:"pengunjung"`
	OmsetOp           *string `json:"omsetOp"`

	Rekening_Id uint64 `json:"rekening_id"`

	Genset   bool `json:"genset"`
	AirTanah bool `json:"airTanah"`

	DetailOp *[]DetailOp `json:"detail_op"`

	ObjekPajak []ObjekPajak `json:"objekPajak"`
	Pemilik    []PemilikWp  `json:"pemilik"`
	Narahubung []Narahubung `json:"narahubung"`
}

type VerifikasiDto struct {
	VerifyStatus string `json:"verifyStatus"`
}

package registration

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/rekening"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/skpd"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
)

type Golongan string
type StatusPendaftaran string
type JenisPajak string
type StatusObjekPajak string
type Mode string
type VerifyPendaftaran string
type StatusPemilik string
type StatusNarahubung string

const (
	GolonganOrangPribadi Golongan = "orang_pribadi"
	GolonganBadan        Golongan = "badan"

	StatusAktif          StatusPendaftaran = "aktif"
	StatusTutupSementara StatusPendaftaran = "tutup_sementara"
	StatusTutup          StatusPendaftaran = "tutup"

	JenisPajakSelfAssessment     JenisPajak = "self_assessment"
	JenisPajakOfficialAssessment JenisPajak = "official_assessment"

	StatusBaru StatusObjekPajak = "baru"
	StatusLama StatusObjekPajak = "lama"

	ModeSelf     Mode = "self"
	ModeOperator Mode = "operator"

	VerifiyPendaftaranDisetujui      VerifyPendaftaran = "disetujui"
	VerifiyPendaftaranTidakDisetujui VerifyPendaftaran = "tidak_disetujui"

	StatusPemilikBaru StatusPemilik = "baru"
	StatusPemilikLama StatusPemilik = "lama"

	StatusNarahubungBaru StatusNarahubung = "baru"
	StatusNarahubungLama StatusNarahubung = "lama"
)

type (
	Registration struct {
		ID                uint64             `json:"id" gorm:"primarykey"`
		Golongan          Golongan           `json:"golongan"`
		Nomor             *string            `json:"nomor" gorm:"size:10"`
		NomorRegistrasi   *string            `json:"nomor_registrasi" gorm:"size:50"`
		Npwp              *string            `json:"npwp" gorm:"size:50"`
		TanggalPengukuhan *time.Time         `json:"tanggal_pengukuhan"`
		TanggalNpwpd      *time.Time         `json:"tanggal_npwpd"`
		Npwpd             *string            `json:"npwpd" gorm:"size:22"`
		Status            StatusPendaftaran  `json:"status"`
		TanggalTutup      *time.Time         `json:"tanggal_tutup"`
		TanggalBuka       *time.Time         `json:"tanggal_buka"`
		JenisPajak        JenisPajak         `json:"jenis_pajak"`
		Skpd_ID           *uint64            `json:"skpd_id"`
		Skpd              *skpd.Skpd         `json:"skpd,omitempty" gorm:"foreignKey:Skpd_ID"`
		Rekening_ID       *uint64            `json:"rekening_id"`
		Rekening          *rekening.Rekening `json:"rekening,omitempty" gorm:"foreignKey:Rekening_ID"`
		Nama              *string            `json:"nama" gorm:"size:20"`
		User_ID           *uint64            `json:"user_id"`
		User              *user.User         `gorm:"foreignKey:User_ID"`
		OmsetOp           *string            `json:"omset_op" gorm:"size:50"`
		Genset            *bool              `json:"genset"`
		AirTanah          *bool              `json:"air_tanah"`
		TanggalMulaiUsaha *time.Time         `json:"tanggal_mulai_usaha"`
		LuasBangunan      *string            `json:"luas_bangunan" gorm:"size:50"`
		JamBukaUsaha      *string            `json:"jam_buka_usaha" gorm:"size:50"`
		JamTutupUsaha     *string            `json:"jam_tutup_usaha" gorm:"size:50"`
		Pengunjung        *string            `json:"pengunjung" gorm:"size:50"`
		FotoKtp           *string            `json:"foto_ktp" gorm:"size:50"`
		SuratIzinUsaha    *string            `json:"surat_izin_usaha" gorm:"size:50"`
		LainLain          *string            `json:"lain_lain" gorm:"size:50"`
		FotoObjek         *string            `json:"foto_objek" gorm:"size:50"`
		Mode              Mode               `json:"mode"`
		VerifyStatus      *VerifyPendaftaran `json:"verify_status"`
		VerifyAt          *time.Time         `json:"verify_at"`
		VendorEtaxID      *string            `json:"vendor_etax_id"`
		// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`

		ObjekPajaks []*ObjekPajak `json:"objek_pajak,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		PemilikWps  []*PemilikWp  `json:"pemilik,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		Narahubungs []*Narahubung `json:"narahubung,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`

		DetailOpAirTanah []*DetailOpAirTanah `json:"detail_op_air_tanah,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		DetailOpHiburan  []*DetailOpHiburan  `json:"detail_op_hiburan,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		DetailOpHotel    []*DetailOpHotel    `json:"detail_op_hotel,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		DetailOpParkir   []*DetailOpParkir   `json:"detail_op_parkir,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		DetailOpPpj      []*DetailOpPpj      `json:"detail_op_ppj,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		DetailOpReklame  []*DetailOpReklame  `json:"detail_op_reklame,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`
		DetailOpResto    []*DetailOpResto    `json:"detail_op_resto,omitempty" gorm:"foreignKey:Pendaftaran_ID;references:ID"`

		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	ObjekPajak struct {
		ID             uint64                  `json:"id" gorm:"primarykey"`
		Pendaftaran_ID uint64                  `json:"pendaftaran_id"`
		Pendaftaran    *Registration           `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_ID"`
		Nama           *string                 `json:"nama" gorm:"size:200"`
		Nop            *string                 `json:"nop" gorm:"size:50"`
		Alamat         *string                 `json:"alamat" gorm:"size:200"`
		RtRw           *string                 `json:"rt_rw" gorm:"size:10"`
		Kecamatan_ID   *uint64                 `json:"kecamatan_id"`
		Kecamatan      *areadivision.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_ID"`
		Kelurahan_ID   *uint64                 `json:"kelurahan_id"`
		Kelurahan      *areadivision.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_ID"`
		Telp           *string                 `json:"telp" gorm:"size:20"`
		Status         StatusObjekPajak        `json:"status"`
		IsNpwpd        *bool                   `json:"is_npwpd"`
		CreatedAt      time.Time               `json:"created_at"`
		UpdatedAt      time.Time               `json:"updated_at"`
		DeletedAt      *time.Time              `json:"deleted_at"`
	}

	DetailOp struct {
		ID                 uint64        `json:"id" gorm:"primaryKey"`
		Pendaftaran_ID     uint64        `json:"pendaftaran_id"`
		Pendaftaran        *Registration `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_ID"`
		NpwpdPendaftaranID *string       `json:"npwpd_pendaftaran_id" gorm:"size:22"`
		JenisOp            *string       `json:"jenis_op" gorm:"size:200"`
		JumlahOp           *string       `json:"jumlah_op" gorm:"size:200"`
		TarifOp            *string       `json:"tarif_op" gorm:"size:200"`
		UnitOp             *string       `json:"unit_op" gorm:"size:50"`
		Notes              *string       `json:"notes" gorm:"size:200"`
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
		ID             uint64                  `json:"id" gorm:"primaryKey"`
		Pendaftaran_ID uint64                  `json:"pendaftaran_id"`
		Pendaftaran    *Registration           `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_ID"`
		Nama           *string                 `json:"nama" gorm:"size:50"`
		Alamat         *string                 `json:"alamat" gorm:"size:50"`
		RtRw           *string                 `json:"rt_rw" gorm:"size:10"`
		Kecamatan_ID   *uint64                 `json:"kecamatan_id"`
		Kecamatan      *areadivision.Kecamatan `gorm:"foreignKey:Kecamatan_ID"`
		Kelurahan_ID   *uint64                 `json:"kelurahan_id"`
		Kelurahan      *areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_ID"`
		Telp           *string                 `json:"telp" gorm:"size:20"`
		Status         StatusPemilik           `json:"status"`
		NoIdPemilik    *string                 `json:"no_id_pemilik" gorm:"size:20"`
		Nik            *string                 `json:"nik" gorm:"size:20"`
	}

	Narahubung struct {
		ID             uint64                 `json:"id" gorm:"primaryKey"`
		Pendaftaran_ID uint64                 `json:"pendaftaran_id"`
		Pendaftaran    Registration           `json:"pendaftaran,omitempty" gorm:"foreignKey:Pendaftaran_ID"`
		Nama           string                 `json:"nama" gorm:"size:50"`
		Alamat         string                 `json:"alamat" gorm:"size:50"`
		RtRw           string                 `json:"rt_rw" gorm:"size:10"`
		Kecamatan_ID   uint64                 `json:"kecamatan_id"`
		Kecamatan      areadivision.Kecamatan `gorm:"foreignKey:Kecamatan_ID"`
		Kelurahan_ID   uint64                 `json:"kelurahan_id"`
		Kelurahan      areadivision.Kelurahan `gorm:"foreignKey:Kelurahan_ID"`
		Telp           string                 `json:"telp" gorm:"size:20"`
		Status         StatusNarahubung       `json:"status"`
		Nik            string                 `json:"nik" gorm:"size:20"`
	}

	RegisterByOperator struct {
		JenisPajak JenisPajak `json:"jenis_pajak"`
		Golongan   Golongan   `json:"golongan"`
		Npwp       *string    `json:"npwp"`

		NomorRegistrasi       *string `json:"nomor_registrasi"`
		IsNomorRegistrasiAuto bool    `json:"is_nomor_registrasi_auto"`

		Npwpd             *string `json:"npwpd"`
		TanggalPengukuhan *string `json:"tanggal_pengukuhan"`
		TanggalNpwpd      *string `json:"tanggal_npwpd"`

		TanggalMulaiUsaha *string `json:"tanggal_mulai_usaha"`
		LuasBangunan      *string `json:"luas_bangunan"`
		JamBukaUsaha      *string `json:"jam_buka_usaha"`
		JamTutupUsaha     *string `json:"jam_tutup_usaha"`
		Pengunjung        *string `json:"pengunjung"`
		OmsetOp           *string `json:"omset_op"`

		RekeningID uint64 `json:"rekening_id"`

		Genset   bool `json:"genset"`
		AirTanah bool `json:"air_tanah"`

		DetailOp *[]DetailOp `json:"detail_op"`

		ObjekPajak *[]ObjekPajak `json:"objek_pajak"`
		Pemilik    *[]PemilikWp  `json:"pemilik"`
		Narahubung *[]Narahubung `json:"narahubung"`
	}
)

type RegisterUpdate struct {
	JenisPajak JenisPajak `json:"jenis_pajak"`
	Golongan   Golongan   `json:"golongan"`
	Npwp       string     `json:"npwp"`

	NomorRegistrasi       string `json:"nomor_registrasi"`
	IsNomorRegistrasiAuto bool   `json:"is_nomor_registrasi_auto"`

	Npwpd             string `json:"npwpd"`
	TanggalPengukuhan string `json:"tanggal_pengukuhan"`
	TanggalNpwpd      string `json:"tanggal_npwpd"`

	TanggalMulaiUsaha string `json:"tanggal_mulai_usaha"`
	LuasBangunan      string `json:"luas_bangunan"`
	JamBukaUsaha      string `json:"jam_buka_usaha"`
	JamTutupUsaha     string `json:"jam_tutup_usaha"`
	Pengunjung        string `json:"pengunjung"`
	OmsetOp           string `json:"omset_op"`

	RekeningID uint64 `json:"rekening_id"`

	Genset   bool `json:"genset"`
	AirTanah bool `json:"air_tanah"`

	DetailOp []DetailOp `json:"detail_op"`

	ObjekPajak []ObjekPajak `json:"objek_pajak"`
	Pemilik    []PemilikWp  `json:"pemilik"`
	Narahubung []Narahubung `json:"narahubung"`
}

type VerifikasiDto struct {
	VerifyStatus string `json:"verify_status"`
}

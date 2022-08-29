package registration

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/configurationmodel"
)

type Golongan string
type StatusPendaftaran string
type JenisPajak string
type StatusObjekPajak string
type Mode string
type VerifyPendaftaran string

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
)

type (
	Registration struct {
		ID                uint64                       `json:"id" gorm:"primarykey"`
		Golongan          Golongan                     `json:"golongan"`
		Nomor             *string                      `json:"nomor"`
		NomorRegistrasi   *string                      `json:"nomor_registrasi"`
		Npwp              *string                      `json:"npwp"`
		TanggalPengukuhan *time.Time                   `json:"tanggal_pengukuhan"`
		Npwpd             *string                      `json:"npwpd"`
		Status            StatusPendaftaran            `json:"status"`
		TanggalTutup      *time.Time                   `json:"tanggal_tutup"`
		TanggalBuka       *time.Time                   `json:"tanggal_buka"`
		JenisPajak        JenisPajak                   `json:"jenis_pajak"`
		SkpdID            uint64                       `json:"skpd_id"`
		Skpd              *configurationmodel.Skpd     `gorm:"foreignKey:SkpdID"`
		RekeningID        uint64                       `json:"rekening_id"`
		Rekening          *configurationmodel.Rekening `gorm:"foreignKey:RekeningID"`
		Nama              *string                      `json:"nama"`
		UserID            uint64                       `json:"user_id"`
		User              *user.User                   `gorm:"foreignKey:UserID"`
		OmsetOp           *string                      `json:"omset_op"`
		Genset            *string                      `json:"genset"`
		AirTanah          *string                      `json:"air_tanah"`
		TanggalMulaiUsaha *time.Time                   `json:"tanggal_mulai_usaha"`
		LuasBangunan      *string                      `json:"luas_bangunan"`
		JamBukaUsaha      *string                      `json:"jam_buka_usaha"`
		JamTutupUsaha     *string                      `json:"jam_tutup_usaha"`
		Pengunjung        *string                      `json:"pengunjung"`
		FotoKtp           *string                      `json:"foto_ktp"`
		SuratIzinUsaha    *string                      `json:"surat_izin_usaha"`
		LainLain          *string                      `json:"lain_lain"`
		FotoObjek         *string                      `json:"foto_objek"`
		Mode              Mode                         `json:"mode"`
		VerifyStatus      VerifyPendaftaran            `json:"verify_status"`
		VerifyAt          *time.Time                   `json:"verify_at"`
		VendorEtaxID      *string                      `json:"vendor_etax_id"`
		// VendorEtax         *configurationmodel.VendorEtax `gorm:"foreignKey:VendorEtaxID"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	ObjekPajak struct {
		ID            uint64                        `json:"id" gorm:"primarykey"`
		PendaftaranID uint64                        `json:"pendaftaran_id"`
		Pendaftaran   *Registration                 `gorm:"foreignKey:PendaftaranID"`
		Nama          *string                       `json:"nama"`
		Nop           *string                       `json:"nop"`
		Alamat        *string                       `json:"alamat"`
		RtRw          *string                       `json:"rt_rw"`
		KecamatanID   uint64                        `json:"kecamatanId"`
		Kecamatan     *configurationmodel.Kecamatan `gorm:"foreignKey:KecamatanID"`
		KelurahanID   uint64                        `json:"kelurahanId"`
		Kelurahan     *configurationmodel.Kelurahan `gorm:"foreignKey:KelurahanID"`
		Telp          *string                       `json:"telp"`
		Status        StatusObjekPajak              `json:"status"`
		IsNpwpd       *bool                         `json:"is_npwpd"`
		CreatedAt     time.Time                     `json:"created_at"`
		UpdatedAt     time.Time                     `json:"updated_at"`
		DeletedAt     *time.Time                    `json:"deleted_at"`
	}

	DetailOp struct {
		ID                 uint64        `json:"id" gorm:"primaryKey"`
		PendaftaranID      uint64        `json:"pendaftaran_id"`
		Pendaftaran        *Registration `gorm:"foreignKey:PendaftaranID"`
		NpwpdPendaftaranID *string       `json:"npwpd_pendaftaran_id"`
		JenisOp            *string       `json:"jenis_op"`
		JumlahOp           *string       `json:"jumlah_op"`
		TarifOp            *string       `json:"tarif_op"`
		UnitOp             *string       `json:"unit_op"`
		Notes              *string       `json:"notes"`
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
		ID            uint64                        `json:"id" gorm:"primaryKey"`
		PendaftaranID uint64                        `json:"pendaftaran_id"`
		Pendaftaran   *Registration                 `gorm:"foreignKey:PendaftaranID"`
		Nama          *string                       `json:"nama"`
		Alamat        *string                       `json:"alamat"`
		RtRw          *string                       `json:"rt_rw"`
		KecamatanID   uint64                        `json:"kecamatanId"`
		Kecamatan     *configurationmodel.Kecamatan `gorm:"foreignKey:KecamatanID"`
		KelurahanID   uint64                        `json:"kelurahanId"`
		Kelurahan     *configurationmodel.Kelurahan `gorm:"foreignKey:KelurahanID"`
		Telp          *string                       `json:"telp"`
		Status        *int                          `json:"status"`
		NoIdPemilik   *string                       `json:"no_id_pemilik"`
		Nik           *string                       `json:"nik"`
	}

	Narahubung struct {
		ID            uint64                        `json:"id" gorm:"primaryKey"`
		PendaftaranID uint64                        `json:"pendaftaran_id"`
		Pendaftaran   *Registration                 `gorm:"foreignKey:PendaftaranID"`
		Nama          *string                       `json:"nama"`
		Alamat        *string                       `json:"alamat"`
		RtRw          *string                       `json:"rt_rw"`
		KecamatanID   uint64                        `json:"kecamatanId"`
		Kecamatan     *configurationmodel.Kecamatan `gorm:"foreignKey:KecamatanID"`
		KelurahanID   uint64                        `json:"kelurahanId"`
		Kelurahan     *configurationmodel.Kelurahan `gorm:"foreignKey:KelurahanID"`
		Telp          *string                       `json:"telp"`
		Status        *int                          `json:"status"`
		Nik           *string                       `json:"nik"`
	}
)

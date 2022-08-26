package registration

import (
	"os/user"
	"time"
)

type Golongan string
type StatusPendaftaran string
type JenisPajak string
type StatusObjekPajak string

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
)

type (
	Skpd struct{}

	Rekening struct{}

	VendorEtax struct{}

	Kecamatan struct{}

	Kelurahan struct{}

	Registration struct {
		ID                 uint64            `json:"id" gorm:"primarykey"`
		Golongan           Golongan          `json:"golongan"`
		SerialNumber       *string           `json:"serial_number"`
		RegistrationNumber *string           `json:"registration_number"`
		Npwp               *string           `json:"npwp"`
		Nop                *string           `json:"nop"`
		TanggalPengukuhan  *time.Time        `json:"tanggal_pengukuhan"`
		Npwpd              *string           `json:"npwpd"`
		Status             StatusPendaftaran `json:"status"`
		ClosingDate        *time.Time        `json:"closing_date"`
		OpeningDate        *time.Time        `json:"opening_date"`
		JenisPajak         JenisPajak        `json:"jenis_pajak"`
		SkpdID             uint64            `json:"skpd_id"`
		Skpd               *Skpd             `gorm:"foreignKey:SkpdID"`
		RekeningID         uint64            `json:"rekening_id"`
		Rekening           *Rekening         `gorm:"foreignKey:RekeningID"`
		UserName           *string           `json:"user_name"`
		UserID             uint64            `json:"user_id"`
		User               *user.User        `gorm:"foreignKey:UserID"`
		OmsetOp            *string           `json:"omset_op"`
		Genset             *string           `json:"genset"`
		AirTanah           *string           `json:"air_tanah"`
		StartDate          *time.Time        `json:"start_date"`
		BuildingArea       *string           `json:"building_area"`
		OpeningHour        *string           `json:"openingHour"`
		ClosingHour        *string           `json:"closing_hour"`
		Visitor            *string           `json:"visitor"`
		VendorEtaxID       *string           `json:"vendor_etax_id"`
		VendorEtax         *VendorEtax       `gorm:"foreignKey:VendorEtaxID"`
		CreatedAt          time.Time         `json:"created_at"`
		UpdatedAt          time.Time         `json:"updated_at"`
		DeletedAt          *time.Time        `json:"deleted_at"`
	}

	ObjekPajak struct {
		ID            uint64           `json:"id" gorm:"primarykey"`
		PendaftaranID uint64           `json:"pendaftaran_id"`
		Pendaftaran   *Registration    `gorm:"foreignKey:PendaftaranID"`
		Name          *string          `json:"name"`
		Address       *string          `json:"address"`
		RtRw          *string          `json:"rt_rw"`
		KecamatanID   uint64           `json:"kecamatanId"`
		Kecamatan     *Kecamatan       `gorm:"foreignKey:KecamatanID"`
		KelurahanID   uint64           `json:"kelurahanId"`
		Kelurahan     *Kelurahan       `gorm:"foreignKey:KelurahanID"`
		NoTelp        *string          `json:"no_telp"`
		Status        StatusObjekPajak `json:"status"`
		IsNpwpd       *bool            `json:"is_npwpd"`
		CreatedAt     time.Time        `json:"created_at"`
		UpdatedAt     time.Time        `json:"updated_at"`
		DeletedAt     *time.Time       `json:"deleted_at"`
	}

	DetailOp struct {
		ID               uint64        `json:"id" gorm:"primaryKey"`
		PendaftaranID    uint64        `json:"pendaftaran_id"`
		Pendaftaran      *Registration `gorm:"foreignKey:PendaftaranID"`
		NpwpdPendaftaran *string       `json:"npwpd_pendaftaran"`
		JenisOp          *string       `json:"jenis_op"`
		JumlahOp         *string       `json:"jumlah_op"`
		TarifOp          *string       `json:"tarif_op"`
		UnitOp           *string       `json:"unit_op"`
		Notes            *string       `json:"notes"`
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
		ID            uint64        `json:"id" gorm:"primaryKey"`
		PendaftaranID uint64        `json:"pendaftaran_id"`
		Pendaftaran   *Registration `gorm:"foreignKey:PendaftaranID"`
		Name          *string       `json:"name"`
		Address       *string       `json:"address"`
		RtRw          *string       `json:"rt_rw"`
		KecamatanID   uint64        `json:"kecamatanId"`
		Kecamatan     *Kecamatan    `gorm:"foreignKey:KecamatanID"`
		KelurahanID   uint64        `json:"kelurahanId"`
		Kelurahan     *Kelurahan    `gorm:"foreignKey:KelurahanID"`
		NoTelp        *string       `json:"no_telp"`
		Status        *int          `json:"status"`
		NoIdPemilik   *string       `json:"no_id_pemilik"`
		Nik           *string       `json:"nik"`
	}

	Narahubung struct {
		ID            uint64        `json:"id" gorm:"primaryKey"`
		PendaftaranID uint64        `json:"pendaftaran_id"`
		Pendaftaran   *Registration `gorm:"foreignKey:PendaftaranID"`
		Name          *string       `json:"name"`
		Address       *string       `json:"address"`
		RtRw          *string       `json:"rt_rw"`
		KecamatanID   uint64        `json:"kecamatanId"`
		Kecamatan     *Kecamatan    `gorm:"foreignKey:KecamatanID"`
		KelurahanID   uint64        `json:"kelurahanId"`
		Kelurahan     *Kelurahan    `gorm:"foreignKey:KelurahanID"`
		NoTelp        *string       `json:"no_telp"`
		Status        *int          `json:"status"`
		NoIdPemilik   *string       `json:"no_id_pemilik"`
		Nik           *string       `json:"nik"`
	}
)

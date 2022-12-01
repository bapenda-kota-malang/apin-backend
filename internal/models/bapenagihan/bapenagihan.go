package bapenagihan

import (
	"time"

	"github.com/bapenda-kota-malang/apin-backend/internal/models/bapenagihan/detail"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/undanganpemeriksaan"
	"github.com/bapenda-kota-malang/apin-backend/internal/models/user"
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type BaPenagihan struct {
	Id               uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Undangan_Id      uint64         `json:"undangan_id"`
	Tanggal          datatypes.Date `json:"tanggal"`
	ObjekPajakPbb_Id uint64         `json:"objekPajakPbb_id"`
	NoBeritaAcara    string         `json:"noBeritaAcara"`
	Hasil            *string        `json:"hasil"`
	Catatan          *string        `json:"catatan"`
	Dokumentasi      datatypes.JSON `json:"dokumentasi" gorm:"type:jsonb"`
	DokumenLainLain  datatypes.JSON `json:"dokumenLainLain" gorm:"type:jsonb"`
	VerifyBy_User_Id *uint64        `json:"verifyBy_user_id"`
	VerifyAt         *time.Time     `json:"verifyAt"`
	gormhelper.DateModel
	VerifyBy          *user.User                               `json:"verifyBy,omitempty" gorm:"foreignKey:VerifyBy_User_Id"`
	Undangan          *undanganpemeriksaan.UndanganPemeriksaan `json:"undanganPemeriksaan,omitempty" gorm:"foreignKey:Undangan_Id"`
	BaPenagihanDetail *[]detail.BaPenagihanDetail              `json:"baPenagihanDetail,omitempty" gorm:"foreignKey:BaPenagihan_Id"`
}

type CreateDto struct {
	Undangan_Id       uint64         `json:"undangan_id" validate:"required;min=1"`
	ObjekPajakPbb_Id  uint64         `json:"objekPajakPbb_id" validate:"min=1"`
	Tanggal           datatypes.Date `json:"tanggal"`
	NoBeritaAcara     string         `json:"-"`
	Hasil             *string        `json:"hasil"`
	Dokumentasi       []string       `json:"dokumentasi" validate:"base64=image;b64size=1024"`
	DokumenLainLain   []string       `json:"dokumenLainLain" validate:"base64=pdf,image,excel;b64size=1024"`
	PetugasListDetail []uint64       `json:"petugasListDetail" validate:"required;min=1"`
}

type UpdateBulk struct {
	Id              *uuid.UUID     `json:"id"`
	Hasil           string         `json:"hasil"`
	Dokumentasi     datatypes.JSON `json:"dokumentasi" validate:"base64=image;b64size=1024"`
	DokumenLainLain datatypes.JSON `json:"dokumenLainLain" validate:"base64=pdf,image,excel;b64size=1024"`
}

type UpdateBulkDto struct {
	Datas []UpdateBulk `json:"datas" validate:"required"`
}

type FilesUpdateDokumentasi struct {
	NewFile *string `json:"newFile" validate:"base64=image;b64size=1024"`
	OldFile *string `json:"oldFile"`
	Deleted bool    `json:"deleted"`
}

type FilesUpdateDokumenLainLain struct {
	NewFile *string `json:"newFile" validate:"base64=pdf,image,excel;b64size=1024"`
	OldFile *string `json:"oldFile"`
	Deleted bool    `json:"deleted"`
}

type UpdateDto struct {
	Tanggal            datatypes.Date               `json:"tanggal"`
	Hasil              string                       `json:"hasil"`
	Dokumentasi        []FilesUpdateDokumentasi     `json:"dokumentasi"`
	DokumenLainLain    []FilesUpdateDokumenLainLain `json:"dokumenLainLain"`
	BaPenagihanDetails []detail.UpdateDto           `json:"baPenagihanDetails"`
}

type VerifikasiDto struct {
	Tanggal          *datatypes.Date `json:"tanggal"`
	Hasil            *string         `json:"hasil"`
	Catatan          *string         `json:"catatan"`
	VerifyBy_User_Id *uint64         `json:"-"`
	VerifyAt         *time.Time      `json:"-"`
}

type FilterDto struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

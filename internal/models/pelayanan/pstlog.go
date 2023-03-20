package pstpermohonan

import (
	"github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PstLogApproval struct {
	Id            uint64  `json:"id" gorm:"primaryKey"`
	Permohonan_Id uint64  `json:"permohonan_id"`
	User_id       string  `json:"user_id" gorm:"type:varchar(50)"`
	Catatan       *string `json:"catatan" gorm:"type:TEXT"`
	Keterangan    *string `json:"keterangan" gorm:"type:TEXT"`
	Status        string  `json:"status" gorm:"type:varchar(50)"`
	Jabatan       string  `json:"jabatan" gorm:"type:varchar(2)"`
	Divisi        *string `json:"divisi" gorm:"type:varchar(2)"`
	gormhelper.DateModel
}

type RequestPSTLogApproval struct {
	Permohonan_Id uint64  `json:"permohonan_id"`
	User_id       string  `json:"user_id"`
	Catatan       *string `json:"catatan"`
	Keterangan    *string `json:"keterangan"`
	Status        string  `json:"status"`
	Jabatan       string  `json:"jabatan"`
	Divisi        *string `json:"divisi"`
	gormhelper.DateModel
}

type ResponsePSTLogApproval struct {
	UserStaff   *string `json:"user_staff"`
	Catatan     *string `json:"catatan"`
	UserKasubid *string `json:"user_kasubid"`
	UserKabid   *string `json:"user_kabid"`
	UserSekban  *string `json:"user_sekban"`
	UserKaban   *string `json:"user_kaban"`
	TglStaff    *string `json:"tgl_staff"`
	TglKasubid  *string `json:"tgl_kasubid"`
	TglKabid    *string `json:"tgl_kabid"`
	TglSekban   *string `json:"tgl_sekban"`
	TglKaban    *string `json:"tgl_kaban"`
}

package kunjungan

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type KunjunganDetail struct {
	Id           uint64  `json:"id" gorm:"primaryKey"`
	Kunjungan_Id uint64  `json:"kunjungan_id" gorm:"type:BIGINT"`
	Pegawai_Id   *uint64 `json:"pegawai_id" gorm:"type:BIGINT"`
	Nip          *string `json:"nip" gorm:"type:varchar(50)"`
	Nama         *string `json:"nama" gorm:"type:varchar(50)"`
	JabatanId    *int    `json:"jabatan" gorm:"type:INTEGER"`
	gh.DateModel
}

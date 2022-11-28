package npwpd

import (
	"github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	t "github.com/bapenda-kota-malang/apin-backend/internal/models/types"
)

type PemilikWp struct {
	Id                    uint64                  `json:"id" gorm:"primaryKey"`
	Npwpd_Id              uint64                  `json:"npwpd_id"`
	Npwpd                 *Npwpd                  `json:"npwpd,omitempty" gorm:"foreignKey:Npwpd_Id"`
	Nama                  *string                 `json:"nama" gorm:"size:50"`
	Alamat                *string                 `json:"alamat" gorm:"size:50"`
	Npwp                  *string                 `json:"npwp" gorm:"size:50"`
	Daerah_Id             *uint64                 `json:"daerah_id"`
	Daerah                *areadivision.Daerah    `json:"daerah" gorm:"foreignKey:Daerah_Id"`
	Kelurahan_Id          *uint64                 `json:"kelurahan_id"`
	Kelurahan             *areadivision.Kelurahan `json:"kelurahan" gorm:"foreignKey:Kelurahan_Id"`
	Telp                  *string                 `json:"telp" gorm:"size:20"`
	Status                t.StatusBL              `json:"status"`
	NoIdPemilik           *string                 `json:"noIdPemilik" gorm:"size:20"`
	Nik                   *string                 `json:"nik" gorm:"size:20"`
	Direktur_Nama         *string                 `json:"direktur_nama" gorm:"type:varchar(50)"`
	Direktur_Nik          *string                 `json:"direktur_nik" gorm:"type:varchar(20)"`
	Direktur_Alamat       *string                 `json:"direktur_alamat" gorm:"type:varchar(50)"`
	Direktur_Daerah_Id    *uint64                 `json:"direktur_daerah_id"`
	Direktur_Daerah       *areadivision.Daerah    `json:"direktur_daerah" gorm:"foreignKey:Direktur_Daerah_Id"`
	Direktur_Kelurahan_Id *uint64                 `json:"direktur_kelurahan_id"`
	Direktur_Kelurahan    *areadivision.Kelurahan `json:"direktur_kelurahan" gorm:"foreignKey:Direktur_Kelurahan_Id"`
	Direktur_Telp         *string                 `json:"direktur_telp" gorm:"type:varchar(20)"`
}

type PemilikWpCreateDto struct {
	Nama                  *string `json:"nama" validate:"required"`
	Alamat                *string `json:"alamat" validate:"required"`
	Npwp                  *string `json:"npwp" gorm:"size:50"`
	Daerah_Id             *uint64 `json:"daerah_id" validate:"required"`
	Kelurahan_Id          *uint64 `json:"kelurahan_id" validate:"required"`
	Telp                  *string `json:"telp" validate:"required;nohp"`
	Nik                   *string `json:"nik" validate:"required;nik"`
	Direktur_Nama         *string `json:"direktur_nama"`
	Direktur_Nik          *string `json:"direktur_nik"`
	Direktur_Alamat       *string `json:"direktur_alamat"`
	Direktur_Daerah_Id    *uint64 `json:"direktur_daerah_id"`
	Direktur_Kelurahan_Id *uint64 `json:"direktur_kelurahan_id"`
	Direktur_Telp         *string `json:"direktur_telp"`
}

type PemilikWpUpdateDto struct {
	Id                    uint64  `json:"id" validate:"required"`
	Nama                  *string `json:"nama" validate:"required"`
	Npwp                  *string `json:"npwp" gorm:"size:50"`
	Alamat                *string `json:"alamat" validate:"required"`
	Daerah_Id             *uint64 `json:"daerah_id" validate:"required"`
	Kelurahan_Id          *uint64 `json:"kelurahan_id" validate:"required"`
	Telp                  *string `json:"telp" validate:"required;nohp"`
	Nik                   *string `json:"nik" validate:"required;nik"`
	Direktur_Nama         *string `json:"direktur_nama"`
	Direktur_Nik          *string `json:"direktur_nik"`
	Direktur_Alamat       *string `json:"direktur_alamat"`
	Direktur_Daerah_Id    *uint64 `json:"direktur_daerah_id"`
	Direktur_Kelurahan_Id *uint64 `json:"direktur_kelurahan_id"`
	Direktur_Telp         *string `json:"direktur_telp"`
	IsDeleted             bool    `json:"isDeleted"`
}

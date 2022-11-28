package reklas

import (
	ad "github.com/bapenda-kota-malang/apin-backend/internal/models/areadivision"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Reklas struct {
	Id                   uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Id          *uint64 `json:"provinsi_id"`
	Daerah_Id            *uint64 `json:"daerah_id"`
	Kecamatan_Id         *uint64 `json:"kecamatan_id"`
	Kelurahan_Id         *uint64 `json:"kelurahan_id"`
	IsSeluruhKelurahan   bool    `json:"isSeluruhKelurahan"`
	BlokNomor            *string `json:"blokNomor"`
	Mode                 *string `json:"mode"`
	TingkatKenaikanKelas *string `json:"tingkatKenaikanKelas"` //
	KelasTerendah        *string `json:"kelasTerendah"`        //
	KelasTertinggi       *string `json:"kelasTertinggi"`       //
	DariKelas            *string `json:"dariKelas"`            //
	KeKelas              *string `json:"keKelas"`              //
	JumlahRecord         *int    `json:"jumlahRecord"`         //
	RecordKe             *int    `json:"recordKe"`             //
	Record               *int    `json:"record"`               //
	gh.DateModel
	Provinsi  *ad.Provinsi  `json:"provinsi,omitempty" gorm:"foreignKey:Provinsi_Id"`
	Daerah    *ad.Daerah    `json:"daerah,omitempty" gorm:"foreignKey:Daerah_Id"`
	Kecamatan *ad.Kecamatan `json:"kecamatan,omitempty" gorm:"foreignKey:Kecamatan_Id"`
	Kelurahan *ad.Kelurahan `json:"kelurahan,omitempty" gorm:"foreignKey:Kelurahan_Id"`
}

type CreateDto struct {
	Provinsi_Id          *uint64 `json:"provinsi_id"`
	Daerah_Id            *uint64 `json:"daerah_id"`
	Kecamatan_Id         *uint64 `json:"kecamatan_id"`
	Kelurahan_Id         *uint64 `json:"kelurahan_id"`
	IsSeluruhKelurahan   bool    `json:"isSeluruhKelurahan"`
	BlokNomor            *string `json:"blokNomor"`
	Mode                 *string `json:"mode"`
	TingkatKenaikanKelas *string `json:"tingkatKenaikanKelas"` //
	KelasTerendah        *string `json:"kelasTerendah"`        //
	KelasTertinggi       *string `json:"kelasTertinggi"`       //
	DariKelas            *string `json:"dariKelas"`            //
	KeKelas              *string `json:"keKelas"`              //
	JumlahRecord         *int    `json:"jumlahRecord"`         //
	RecordKe             *int    `json:"recordKe"`             //
	Record               *int    `json:"record"`               //
}

type UpdateDto struct {
	Id                   uint64  `json:"id" gorm:"primaryKey"`
	Provinsi_Id          *uint64 `json:"provinsi_id"`
	Daerah_Id            *uint64 `json:"daerah_id"`
	Kecamatan_Id         *uint64 `json:"kecamatan_id"`
	Kelurahan_Id         *uint64 `json:"kelurahan_id"`
	IsSeluruhKelurahan   bool    `json:"isSeluruhKelurahan"`
	BlokNomor            *string `json:"blokNomor"`
	Mode                 *string `json:"mode"`
	TingkatKenaikanKelas *string `json:"tingkatKenaikanKelas"` //
	KelasTerendah        *string `json:"kelasTerendah"`        //
	KelasTertinggi       *string `json:"kelasTertinggi"`       //
	DariKelas            *string `json:"dariKelas"`            //
	KeKelas              *string `json:"keKelas"`              //
	JumlahRecord         *int    `json:"jumlahRecord"`         //
	RecordKe             *int    `json:"recordKe"`             //
	Record               *int    `json:"record"`               //
}

type FilterDto struct {
	Provinsi_Id          *uint64 `json:"provinsi_id"`
	Daerah_Id            *uint64 `json:"daerah_id"`
	Kecamatan_Id         *uint64 `json:"kecamatan_id"`
	Kelurahan_Id         *uint64 `json:"kelurahan_id"`
	IsSeluruhKelurahan   bool    `json:"isSeluruhKelurahan"`
	BlokNomor            *string `json:"blokNomor"`
	Mode                 *string `json:"mode"`
	TingkatKenaikanKelas *string `json:"tingkatKenaikanKelas"` //
	KelasTerendah        *string `json:"kelasTerendah"`        //
	KelasTertinggi       *string `json:"kelasTertinggi"`       //
	DariKelas            *string `json:"dariKelas"`            //
	KeKelas              *string `json:"keKelas"`              //
	JumlahRecord         *int    `json:"jumlahRecord"`         //
	RecordKe             *int    `json:"recordKe"`             //
	Record               *int    `json:"record"`               //
	Page                 int     `json:"page"`
	PageSize             int     `json:"page_size"`
}

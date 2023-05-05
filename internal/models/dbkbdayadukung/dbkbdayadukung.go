package dbkbdayadukung

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type DbkbDayaDukung struct {
	Id                  uint64   `json:"id" gorm:"primaryKey"`
	Provinsi_Kode       string   `json:"provinsi_kode" gorm:"type:varchar(2)"`
	Daerah_Kode         string   `json:"daerah_kode" gorm:"type:varchar(2)"`
	ThnDbkbDayaDukung   *string  `json:"thnDbkbDayaDukung" gorm:"type:varchar(4)"`
	TipeKonstruksi      *string  `json:"tipeKonstruksi" gorm:"type:varchar(1)"`
	NilaiDbkbDayaDukung *float64 `json:"nilaiDbkbDayaDukung"`
	gh.DateModel
}

type CreateDto struct {
	Provinsi_Kode       *string  `json:"provinsi_kode"`
	Daerah_Kode         *string  `json:"daerah_kode"`
	ThnDbkbDayaDukung   *string  `json:"thnDbkbDayaDukung"`
	TipeKonstruksi      *string  `json:"tipeKonstruksi"`
	NilaiDbkbDayaDukung *float64 `json:"nilaiDbkbDayaDukung"`
}

type UpdateDto struct {
	Provinsi_Kode       *string  `json:"provinsi_kode"`
	Daerah_Kode         *string  `json:"daerah_kode"`
	ThnDbkbDayaDukung   *string  `json:"thnDbkbDayaDukung"`
	TipeKonstruksi      *string  `json:"tipeKonstruksi"`
	NilaiDbkbDayaDukung *float64 `json:"nilaiDbkbDayaDukung"`
}

type FilterDto struct {
	Provinsi_Kode       *string  `json:"provinsi_kode"`
	Daerah_Kode         *string  `json:"daerah_kode"`
	ThnDbkbDayaDukung   *string  `json:"thnDbkbDayaDukung"`
	TipeKonstruksi      *string  `json:"tipeKonstruksi"`
	NilaiDbkbDayaDukung *float64 `json:"nilaiDbkbDayaDukung"`
	Page                int      `json:"page"`
	PageSize            int      `json:"page_size"`
}

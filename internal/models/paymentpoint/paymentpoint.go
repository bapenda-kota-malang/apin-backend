package paymentpoint

import (
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type PaymentPoint struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	gh.DateModel
	User_Id     *uint64 `json:"user_id"`
	Nama        *string `json:"nama"`
	Alamat      *string `json:"alamat"`
	Nama_Kepada *string `json:"nama_kepada"`
	Telepon     *string `json:"telepon"`
	IsDeleted   *bool   `json:"isDeleted"`
}

type CreateDto struct {
	User_Id     *uint64 `json:"user_id"`
	Nama        *string `json:"nama"`
	Alamat      *string `json:"alamat"`
	Nama_Kepada *string `json:"nama_kepada"`
	Telepon     *string `json:"telepon"`
	IsDeleted   *bool   `json:"isDeleted"`
}

type UpdateDto struct {
	Id          *uint64 `json:"id"`
	User_Id     *uint64 `json:"user_id"`
	Nama        *string `json:"nama"`
	Alamat      *string `json:"alamat"`
	Nama_Kepada *string `json:"nama_kepada"`
	Telepon     *string `json:"telepon"`
	IsDeleted   *bool   `json:"isDeleted"`
}

type FilterDto struct {
	User_Id     *uint64 `json:"user_id"`
	Nama        *string `json:"nama"`
	Alamat      *string `json:"alamat"`
	Nama_Kepada *string `json:"nama_kepada"`
	Telepon     *string `json:"telepon"`
	IsDeleted   *bool   `json:"isDeleted"`
	Page        int     `json:"page"`
	PageSize    int     `json:"page_size"`
}

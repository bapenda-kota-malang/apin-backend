package tempatpembayaran

import gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type TempatPembayaran struct {
	Id              uint64  `json:"id" gorm:"primaryKey"`
	AlamatKtp       *string `json:"alamatKtp" gorm:"type:varchar(50)"`
	BankPersepsi_Id *string `json:"bankPersepsi_id" gorm:"type:varchar(2)"`
	BankTunggal_Id  *string `json:"bankTunggal_id" gorm:"type:varchar(2)"`
	Institusi_Id    *string `json:"institusi_id" gorm:"type:varchar(2)"`
	Kanwil_Id       *string `json:"kanwil_id" gorm:"type:varchar(2)"`
	Kppbb_Id        *string `json:"kppbb_id" gorm:"type:varchar(2)"`
	Merchant_Id     *string `json:"merchant_id" gorm:"type:varchar(4)"`
	Tp_Id           *string `json:"tp_id" gorm:"unique;type:varchar(2)"`
	NamaTp          *string `json:"namaTp" gorm:"type:varchar(30)"`
	NoRekeningTp    *string `json:"noRekeningTp" gorm:"type:varchar(15)"`
	gh.DateModel
}

type CreateDto struct {
	AlamatKtp       *string `json:"alamatKtp"`
	BankPersepsi_Id *string `json:"bankPersepsi_id"`
	BankTunggal_Id  *string `json:"bankTunggal_id"`
	Institusi_Id    *string `json:"institusi_id"`
	Kanwil_Id       *string `json:"kanwil_id"`
	Kppbb_Id        *string `json:"kppbb_id"`
	Merchant_Id     *string `json:"merchant_id"`
	Tp_Id           *string `json:"tp_id"`
	NamaTp          *string `json:"namaTp"`
	NoRekeningTp    *string `json:"noRekeningTp"`
}

type UpdateDto struct {
	Id              *uint64 `json:"id"`
	AlamatKtp       *string `json:"alamatKtp"`
	BankPersepsi_Id *string `json:"bankPersepsi_id"`
	BankTunggal_Id  *string `json:"bankTunggal_id"`
	Institusi_Id    *string `json:"institusi_id"`
	Kanwil_Id       *string `json:"kanwil_id"`
	Kppbb_Id        *string `json:"kppbb_id"`
	Merchant_Id     *string `json:"merchant_id"`
	Tp_Id           *string `json:"tp_id"`
	NamaTp          *string `json:"namaTp"`
	NoRekeningTp    *string `json:"noRekeningTp"`
}

type FilterDto struct {
	AlamatKtp       *string `json:"alamatKtp"`
	BankPersepsi_Id *string `json:"bankPersepsi_id"`
	BankTunggal_Id  *string `json:"bankTunggal_id"`
	Institusi_Id    *string `json:"institusi_id"`
	Kanwil_Id       *string `json:"kanwil_id"`
	Kppbb_Id        *string `json:"kppbb_id"`
	Merchant_Id     *string `json:"merchant_id"`
	Tp_Id           *string `json:"tp_id"`
	NamaTp          *string `json:"namaTp"`
	NoRekeningTp    *string `json:"noRekeningTp"`
	Page            int     `json:"page"`
	PageSize        int64   `json:"page_size"`
}

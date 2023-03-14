package jenisusaha

type JenisUsaha struct {
	Id               uint64            `json:"id" gorm:"primaryKey"`
	Kode             *string           `json:"kode" gorm:"type:varchar(5)"`
	Uraian           *string           `json:"uraian" gorm:"type:varchar(100)"`
	JenisUsahaDetail *JenisUsahaDetail `json:"jenisUsahaDetail,omitempty" gorm:"foreignKey:JenisUsaha_Id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CreateDto struct {
	Kode             *string                    `json:"kode"`
	Uraian           *string                    `json:"uraian"`
	JenisUsahaDetail *JenisUsahaDetailCreateDto `json:"jenisUsahaDetail"`
}

type UpdateDto struct {
	Id               *uint64                    `json:"id"`
	Kode             *string                    `json:"kode"`
	Uraian           *string                    `json:"uraian"`
	JenisUsahaDetail *JenisUsahaDetailUpdateDto `json:"jenisUsahaDetail"`
}

type FilterDto struct {
	Kode     *string `json:"kode"`
	Uraian   *string `json:"uraian"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

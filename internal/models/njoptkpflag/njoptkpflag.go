package njoptkpflag

type NjoptkpFlag struct {
	Id                 uint64               `json:"id" gorm:"primaryKey"`
	Nop                *string              `json:"nop"`
	NoKtp              *string              `json:"noKtp"`
	AlamatObjekPajak   *string              `json:"alamatObjekPajak"`
	RtRw               *string              `json:"rtRw"`
	Tahun              *string              `json:"tahun"`
	NamaWajibPajak     *string              `json:"namaWajibPajak"`
	AlamatWajibPajak   *string              `json:"alamatWajibPajak"`
	NjoptkpFlagDetails *[]NjoptkpFlagDetail `json:"njoptkpFlagDetail,omitempty" gorm:"foreignKey:NjoptkpFlag_Id;references:Id"`
}

type CreateDto struct {
	Nop               *string                       `json:"nop"`
	NoKtp             *string                       `json:"noKtp"`
	AlamatObjekPajak  *string                       `json:"alamatObjekPajak"`
	RtRw              *string                       `json:"rtRw"`
	Tahun             *string                       `json:"tahun"`
	NamaWajibPajak    *string                       `json:"namaWajibPajak"`
	AlamatWajibPajak  *string                       `json:"alamatWajibPajak"`
	NjoptkpFlagDetail *[]NjoptkpFlagDetailCreateDto `json:"njoptkpFlagDetail"`
}

type UpdateDto struct {
	Id                uint64                        `json:"id"`
	Nop               *string                       `json:"nop"`
	NoKtp             *string                       `json:"noKtp"`
	AlamatObjekPajak  *string                       `json:"alamatObjekPajak"`
	RtRw              *string                       `json:"rtRw"`
	Tahun             *string                       `json:"tahun"`
	NamaWajibPajak    *string                       `json:"namaWajibPajak"`
	AlamatWajibPajak  *string                       `json:"alamatWajibPajak"`
	NjoptkpFlagDetail *[]NjoptkpFlagDetailUpdateDto `json:"njoptkpFlagDetail"`
}

type FilterDto struct {
	Nop              *string `json:"nop"`
	NoKtp            *string `json:"noKtp"`
	AlamatObjekPajak *string `json:"alamatObjekPajak"`
	RtRw             *string `json:"rtRw"`
	Tahun            *string `json:"tahun"`
	NamaWajibPajak   *string `json:"namaWajibPajak"`
	AlamatWajibPajak *string `json:"alamatWajibPajak"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

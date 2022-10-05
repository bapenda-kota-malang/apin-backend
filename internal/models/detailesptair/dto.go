package detailesptair

type CreateDto struct {
	Espt_Id uint `json:"espt_id" validate:"min=1"`
}

type CreateDetailAirDto struct {
	CreateDto
	Peruntukan string  `json:"peruntukan"  validate:"required"`
	JenisAbt   string  `json:"jenisAbt"  validate:"required"`
	Pengenaan  float32 `json:"pengenaan"  validate:"required"`
}

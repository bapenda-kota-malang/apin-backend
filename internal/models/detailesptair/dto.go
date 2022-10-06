package detailesptair

type EsptIdDto struct {
	Espt_Id uint `json:"espt_id"`
}

type CreateDto struct {
	EsptIdDto
	Peruntukan string  `json:"peruntukan"  validate:"required"`
	JenisAbt   string  `json:"jenisAbt"  validate:"required"`
	Pengenaan  float32 `json:"pengenaan"  validate:"required"`
}

package jaminanbongkar

type DetailJambong struct {
	Id                  uint64  `json:"id" gorm:"primaryKey;autoIncrement"`
	JaminanBongkar_Id   uint64  `json:"jaminanBongkar_id"`
	DetailSptReklame_id uint64  `json:"detailSptReklame_id"`
	TarifJambong_Id     *uint64 `json:"tarifJambong_id"`
}

type DetailJambongCreateDto struct {
	JaminanBongkar_Id   uint64  `json:"jaminanBongkar_id" validate:"required"`
	DetailSptReklame_id uint64  `json:"detailSptReklame_id" validate:"required"`
	TarifJambong_Id     *uint64 `json:"tarifJambong_id"`
}

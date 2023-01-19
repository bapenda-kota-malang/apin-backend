package jaminanbongkar

import "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"

type DetailJambong struct {
	Id                  uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	JaminanBongkar_Id   uint64 `json:"jaminanBongkar_id"`
	DetailSptReklame_id uint64 `json:"detailSptReklame_id"`
	TarifJambong_Id     *int64 `json:"tarifJambong_id"`
	gormhelper.DateModel
}

type DetailJambongCreateDto struct {
	DetailSptReklame_id uint64 `json:"detailSptReklame_id" validate:"required"`
	TarifJambong_Id     *int64 `json:"tarifJambong_id"`
}

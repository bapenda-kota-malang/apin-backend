package detailjambong

type DetailJambong struct {
	Id                  uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	JaminanBongkar_Id   uint64 `json:"jaminanBongkar_Id"`
	DetailSptReklame_Id uint64 `json:"detailSptReklame_Id"`
	TarifJambong_Id     uint64 `json:"tarifJambong_Id"`
}

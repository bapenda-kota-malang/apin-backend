package objekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb3 struct {
	nop.NopDetail
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	gh.DateModel
}

type Jpb3CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type Jpb3UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type Jpb3FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom3     *int    `json:"tinggiKolom3"`
	LebarBentang3    *int    `json:"lebarBentang3"`
	LuasMezzanine3   *int    `json:"luasMezzanine3"`
	KelilingDinding3 *int    `json:"kelilingDinding3"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

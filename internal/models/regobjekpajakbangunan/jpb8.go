package regobjekpajakbangunan

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type RegJpb8 struct {
	nop.NopDetail
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi" gorm:"type:char(1)"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	gh.DateModel
}

type RegJpb8CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type RegJpb8UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type RegJpb8FilterDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
	Page             int     `json:"page"`
	PageSize         int     `json:"page_size"`
}

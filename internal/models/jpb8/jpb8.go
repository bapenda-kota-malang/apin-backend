package jpb8

import (
	nop "github.com/bapenda-kota-malang/apin-backend/internal/models/nop"
	gh "github.com/bapenda-kota-malang/apin-backend/pkg/gormhelper"
)

type Jpb8 struct {
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

type CreateDto struct {
	nop.NopDetailCreateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type UpdateDto struct {
	nop.NopDetailUpdateDto
	NoBangunan       *int    `json:"noBangunan"`
	TipeKonstruksi   *string `json:"tipeKonstruksi"`
	TinggiKolom8     *int    `json:"tinggiKolom8"`
	LebarBentang8    *int    `json:"lebarBentang8"`
	LuasMezzanine8   *int    `json:"luasMezzanine8"`
	KelilingDinding8 *int    `json:"kelilingDinding8"`
	DayaDukungLantai *int    `json:"dayaDukungLantai"`
}

type FilterDto struct {
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

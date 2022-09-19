package group

type Group struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" gorm:"size:100"`
	TestDescription string `json:"description" gorm:"size:1000"`
	Position        byte   `json:"position"`
	Status          byte   `json:"status"`
}

type CreateDto struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Position    byte   `json:"position" validate:"required;min=1"`
}

type UpdateDto struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Position    byte   `json:"position" validate:"required;min=1"`
}

type FilterDto struct {
	// dynamic fields
	Name     *string `json:"name"`
	Position *int    `json:"position"`
	Status   *int    `json:"status"`
	// fixed fields
	Page      int   `json:"page"`
	Page_Size int64 `json:"page_size"`
}

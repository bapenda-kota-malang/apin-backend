package group

type Group struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" gorm:"size:100"`
	TestDescription string `json:"description" gorm:"size:1000"`
	Position        int16  `json:"position"`
	Status          int16  `json:"status"`
}

type Create struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Position    int16  `json:"position" validate:"required;min=1"`
}

type Update struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Position    int16  `json:"position" validate:"required;min=1"`
}

type Filter struct {
	Name     string `json:"name"`
	Position int16  `json:"position"`
	Status   int16  `json:"status"`
}

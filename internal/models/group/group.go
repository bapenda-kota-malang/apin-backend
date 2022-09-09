package group

type Group struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"name" gorm:"size:100"`
	TestDescription string `json:"description" gorm:"size:1000"`
	Position        string `json:"position"`
	Status          int16  `json:"status"`
}

type Create struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Position    string `json:"position" validate:"required"`
}

type Update struct {
	Create // copy from Create since it's the same
}

type Filter struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Status   int16  `json:"status"`
}

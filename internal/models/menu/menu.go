package menu

type Menu struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Parent_Id int    `json:"parent_id" gorm:"index"`
	Title     string `json:"title" gorm:"size:100"`
	Url       string `json:"url" gorm:"size:200"`
	Active    int16  `json:"status"`
}

type CreateDto struct {
	Parent_Id int    `json:"parent_id"`
	Title     string `json:"title" validate:"required"`
	Url       string `json:"url" validate:"required"`
	Active    int16  `json:"status" validate:"required;min=1"`
}

type UpdateDto struct {
	Parent_Id int    `json:"parent_id"`
	Title     string `json:"title" validate:"required"`
	Url       string `json:"url" validate:"required"`
	Active    int16  `json:"status" validate:"required;min=1"`
}

type FilterDto struct {
<<<<<<< HEAD
	Title *string `json:"title" gorm:"size:100"`

	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
=======
	Title    *string `json:"title" gorm:"size:100"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
>>>>>>> dev
}

package group

type Group struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"size:100"`
	Description string `json:"description" gorm:"size:1000"`
	Access      string `json:"access"`
	Status      byte   `json:"status"`
}

type CreateDto struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Access      string `json:"access"`
	Status      byte   `json:"status"`
}

type UpdateDto struct {
	Name        string `json:"name" validate:"required;maxLength=100"`
	Description string `json:"description"`
	Access      string `json:"access"`
	Status      byte   `json:"status"`
}

type FilterDto struct {
	// dynamic fields
	Name   *string `json:"name"`
	Status *int    `json:"status"`
	// fixed fields
	Page         int   `json:"page"`
	PageSize     int64 `json:"page_size"`
	NoPagination bool  `json:"no_pagination"`
}

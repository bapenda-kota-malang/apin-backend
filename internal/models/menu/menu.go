package user

type Menu struct {
	Title  string `json:"title"`
	Link   string `json:"link"`
	Active string `json:"active"`
}

type Create struct {
	Title  string `json:"title" validate:"required"`
	Link   string `json:"link" validate:"required"`
	Active string `json:"active" validate:"required"`
}

type Update struct {
	Create
}

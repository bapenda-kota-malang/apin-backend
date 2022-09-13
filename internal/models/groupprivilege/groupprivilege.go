package groupprivilege

type GroupPrivilege struct {
	Group_ID string `json:"group_id"`
	Menu_ID  string `json:"menu_id"`
	Action   string `json:"action"`
}

type Create struct {
	Group_ID string `json:"group_id" validate:"required"`
	Menu_ID  string `json:"menu_id" validate:"required"`
	Action   string `json:"action" validate:"required"`
}

type Update struct {
	Create
}

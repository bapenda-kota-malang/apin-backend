package esptd

type FilterDto struct {
	// fixed
	Page     int   `json:"page"`
	PageSize int64 `json:"page_size"`
}

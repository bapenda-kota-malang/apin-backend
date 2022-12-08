package bphtb

// DTO for verify from user PPAT
type VerifyPpatDto struct {
	Status uint8 `json:"status" validate:"required;min=1;max=2"`
}

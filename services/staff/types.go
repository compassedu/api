package staff

import "github.com/compassedu/api/services/types"

type StaffRequest struct {
	Limit float64 `json:"limit"`
	Page  float64 `json:"page"`
	Start float64 `json:"start"`
}
type StaffResult struct {
	D []types.User `json:"d"`
}

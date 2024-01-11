package calendar

import "api.compass.education/services/types"

type CalendarRequest struct {
	UserId    float64 `json:"userId"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Limit     float64 `json:"limit"`
	Start     float64 `json:"start"`
	Page      float64 `json:"page"`
}
type CalendarResult struct {
	D []types.CalendarTransport `json:"d"`
}

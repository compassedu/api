package attendance

import "github.com/compassedu/api/services/types"

type GetAttendanceSummaryRequest struct {
	StartDate            string   `json:"startDate"`
	EndDate              string   `json:"endDate"`
	StudentStatus        string   `json:"studentStatus"`
	InClass              []string `json:"inClass"`
	OkClass              []string `json:"okClass"`
	Vce                  []string `json:"vce"`
	Schl                 []string `json:"schl"`
	Perspective          string   `json:"perspective"`
	TotalWholeDayLimit   string   `json:"totalWholeDayLimit"`
	TotalPartialDayLimit string   `json:"totalPartialDayLimit"`
	UserId               string   `json:"userId"`
}
type GetAttendanceSummaryResponse struct {
	D []types.SL `json:"d"`
}

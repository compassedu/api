package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"

	"github.com/compassedu/api/services/attendance"
	"github.com/compassedu/api/services/types"
)

func getAttendanceSummary(cookies string, schoolId string, userId float64, startDate time.Time, endDate time.Time) ([]types.SL, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.SL{}, err
	}
	client := &http.Client{Jar: jar}
	requestBody := attendance.GetAttendanceSummaryRequest{
		StartDate:            startDate.Format("2006-01-02T15:04:05.000Z"),
		EndDate:              endDate.Format("2006-01-02T15:04:05.000Z"),
		StudentStatus:        "2",
		InClass:              []string{"0", "1"},
		OkClass:              []string{"0", "1"},
		Vce:                  []string{"0", "1"},
		Schl:                 []string{"0", "1"},
		Perspective:          "1",
		TotalWholeDayLimit:   "0",
		TotalPartialDayLimit: "0",
		UserId:               strconv.FormatFloat(userId, 'f', -1, 64),
	}
	reqbody, err := json.Marshal(requestBody)
	if err != nil {
		return []types.SL{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/Attendance.svc/GetAttendanceSummary"
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqbody))
	if err != nil {
		return []types.SL{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.SL{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u attendance.GetAttendanceSummaryResponse
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.SL{}, err
	}
	return u.D, nil
}
func GetAttendanceSummary(cookies string, schoolId string, userId float64, startDate time.Time, endDate time.Time) ([]types.SL, error) {
	getAttendanceSummary(cookies, schoolId, userId, startDate, endDate)
	locations, err := getAttendanceSummary(cookies, schoolId, userId, startDate, endDate)
	if err != nil {
		return []types.SL{}, err
	}
	return locations, nil
}

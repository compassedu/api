package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/compassedu/api/services/calendar"
	"github.com/compassedu/api/services/types"
)

func getCalendarEventsByUser(cookies string, schoolId string, userId float64, startDate time.Time, endDate time.Time) ([]types.CalendarTransport, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	client := &http.Client{Jar: jar}
	requestBody := calendar.CalendarRequest{
		UserId:    userId,
		StartDate: startDate.Format("2006-01-02"),
		EndDate:   endDate.Format("2006-01-02"),
		Limit:     50,
		Start:     0,
		Page:      1,
	}
	reqbody, err := json.Marshal(requestBody)
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/Calendar.svc/GetCalendarEventsByUser"
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqbody))
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u calendar.CalendarResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	return u.D, nil
}

// getCalendarEventsByUser retrieves calendar events for a specified user within a date range.
//
// The function takes the provided session 'cookies', 'schoolId', 'userId', 'startDate', and 'endDate'
// to make a request to the Compass Education API and fetch calendar events for the specified user within the given date range.
//
// Parameters:
//   - cookies:   Session cookies for authentication.
//   - schoolId:   The unique identifier for the school.
//   - userId:     The user ID for whom the calendar events are to be retrieved.
//   - startDate:  The start date of the range for which calendar events are requested.
//   - endDate:    The end date of the range for which calendar events are requested.
//
// Returns:
//   - []types.CalendarTransport: A slice of calendar events in the form of 'types.CalendarTransport'.
//   - error:                      An error, if any, that occurred during the API request or processing of the response.
//
// Example:
//
//	cookies := "your_session_cookies_here"
//	schoolId := "your_school_id_here"
//	userId := 12345
//	startDate := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
//	endDate := time.Date(2022, time.December, 31, 23, 59, 59, 999999999, time.UTC)
//
//	events, err := getCalendarEventsByUser(cookies, schoolId, userId, startDate, endDate)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use 'events' as needed.
func GetCalendarEventsByUser(cookies string, schoolId string, userId float64, startDate time.Time, endDate time.Time) ([]types.CalendarTransport, error) {
	getCalendarEventsByUser(cookies, schoolId, userId, startDate, endDate)
	events, err := getCalendarEventsByUser(cookies, schoolId, userId, startDate, endDate)
	if err != nil {
		return []types.CalendarTransport{}, err
	}
	return events, nil
}

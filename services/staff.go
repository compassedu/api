package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/compassedu/api/services/staff"
	"github.com/compassedu/api/services/types"
)

func getAllStaff(cookies string, schoolId string, userId float64) ([]types.User, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.User{}, err
	}
	client := &http.Client{Jar: jar}
	requestBody := staff.StaffRequest{
		Limit: 50,
		Page:  1,
		Start: 0,
	}
	reqbody, err := json.Marshal(requestBody)
	if err != nil {
		return []types.User{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/User.svc/GetAllStaff"
	req, err := http.NewRequest("POST", url, bytes.NewReader(reqbody))
	if err != nil {
		return []types.User{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.User{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u staff.StaffResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.User{}, err
	}
	return u.D, nil
}

// GetAllStaff retrieves a list of all staff members from the Compass Education API.
//
// The function takes the provided session 'cookies', 'schoolId', and 'userId' to make a request
// to the Compass Education API and fetch all staff members.
//
// Parameters:
//   - cookies:   Session cookies for authentication.
//   - schoolId:   The unique identifier for the school.
//   - userId:     The user ID associated with the request.
//
// Returns:
//   - []types.User: A slice containing information about all staff members in the form of 'types.User'.
//   - error:        An error, if any, that occurred during the API request or processing of the response.
//
// Example:
//
//	cookies := "your_session_cookies_here"
//	schoolId := "your_school_id_here"
//	userId := 12345
//
//	staffMembers, err := getAllStaff(cookies, schoolId, userId)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Use 'staffMembers' as needed.
func GetAllStaff(cookies string, schoolId string, userId float64) ([]types.User, error) {
	getAllStaff(cookies, schoolId, userId)
	staff, err := getAllStaff(cookies, schoolId, userId)
	if err != nil {
		return []types.User{}, err
	}
	return staff, nil
}

package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/compassedu/api/services/locations"
	"github.com/compassedu/api/services/types"
)

// getAllLocations retrieves all locations from the Compass Education API for a specific school.
//
// The function takes two parameters, cookies and schoolId, to authenticate and identify the school.
// It returns a slice of types.LC representing the locations and an error if any operation fails.
//
// The cookies parameter should contain the authentication cookies required for making requests to the API.
// The schoolId parameter is used to construct the API endpoint URL.
//
// Example usage:
//
//	cookies := "your_authentication_cookies"
//	schoolId := "your_school_id"
//	locations, err := getAllLocations(cookies, schoolId)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Use the 'locations' slice containing the retrieved locations.
//
// Note: This function uses the Compass Education API to fetch location data and requires a valid authentication session.
// For more information about the API, refer to the documentation of Compass Education.
func getAllLocations(cookies string, schoolId string) ([]types.LC, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.LC{}, err
	}
	client := &http.Client{Jar: jar}
	if err != nil {
		return []types.LC{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/ReferenceDataCache.svc/GetAllLocations?page=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []types.LC{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.LC{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u locations.LocationsResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.LC{}, err
	}
	return u.D, nil
}
func GetAllLocations(cookies string, schoolId string) ([]types.LC, error) {
	getAllLocations(cookies, schoolId)
	locations, err := getAllLocations(cookies, schoolId)
	if err != nil {
		return []types.LC{}, err
	}
	return locations, nil
}

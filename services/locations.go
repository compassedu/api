package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"

	"api.compass.education/services/locations"
	"api.compass.education/services/types"
)

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

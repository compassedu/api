package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/compassedu/api/services/events"
	"github.com/compassedu/api/services/types"
)

func getEvents(cookies string, schoolId string) ([]types.ActionCentreEvent, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	client := &http.Client{Jar: jar}
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/ActionCentre.svc/GetEvents"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString("{}"))
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u events.GetEventsResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	return u.D, nil
}
func getPastEvents(cookies string, schoolId string) ([]types.ActionCentreEvent, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	client := &http.Client{Jar: jar}
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/ActionCentre.svc/GetPastEvents"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString("{}"))
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u events.GetEventsResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	fmt.Println(u.D)
	return u.D, nil
}
func GetAllEvents(cookies string, schoolId string) ([]types.ActionCentreEvent, error) {
	getEvents(cookies, schoolId)
	// ucap, err := getEvents(cookies, schoolId)
	// if err != nil {
	// 	return []types.ActionCentreEvent{}, err
	// }
	getPastEvents(cookies, schoolId)
	past, err := getPastEvents(cookies, schoolId)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	var f []types.ActionCentreEvent

	f = append(f, past...)
	return f, nil
}
func GetEvents(cookies string, schoolId string) ([]types.ActionCentreEvent, error) {
	getEvents(cookies, schoolId)
	ucap, err := getEvents(cookies, schoolId)
	if err != nil {
		return []types.ActionCentreEvent{}, err
	}
	var f []types.ActionCentreEvent

	f = append(f, ucap...)
	return f, nil
}

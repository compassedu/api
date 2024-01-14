package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/compassedu/api/services/news"
	"github.com/compassedu/api/services/types"
)

func getNewsItems(cookies string, schoolId string) ([]types.NewsItem, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return []types.NewsItem{}, err
	}
	client := &http.Client{Jar: jar}
	if err != nil {
		return []types.NewsItem{}, err
	}
	url := "https://" + schoolId + ".compass.education/Services/NewsFeed.svc/GetMyNewsFeed"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return []types.NewsItem{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go API/v1")
	req.Header.Set("cookie", cookies)
	res, err := client.Do(req)
	if err != nil {
		return []types.NewsItem{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	var u news.NewsResult
	err = json.Unmarshal(body, &u)
	if err != nil {
		return []types.NewsItem{}, err
	}
	return u.D, nil
}
func GetNewsItems(cookies string, schoolId string) ([]types.NewsItem, error) {
	getNewsItems(cookies, schoolId)
	items, err := getNewsItems(cookies, schoolId)
	if err != nil {
		return []types.NewsItem{}, err
	}
	return items, nil
}

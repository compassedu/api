package news

import "github.com/compassedu/api/services/types"

type NewsResult struct {
	D []types.NewsItem `json:"d"`
}

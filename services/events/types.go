package events

import "github.com/compassedu/api/services/types"

type GetEventsResult struct {
	D []types.ActionCentreEvent `json:"d"`
}

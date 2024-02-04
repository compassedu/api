package services_test

import (
	"testing"

	"github.com/compassedu/api/services"
)

func TestGetAllEvents(t *testing.T) {
	c, _, err := services.Login("KOGO25", "Killian20*", "kilianssc-ie")
	if err != nil {
		t.Error(err)
	}
	t.Log(c)
	events, err := services.GetEvents(c, "kilianssc-ie")
	t.Log(events)
}

package services_test

import (
	"testing"
	"time"

	"github.com/compassedu/api/services"
)

func TestGetCalendarEventsByUser(t *testing.T) {
	c, u, err := services.Login("", "", "")
	if err != nil {
		t.Error(err)
	}
	got, err := services.GetCalendarEventsByUser(c, "", u, time.Now(), time.Now())
	if err != nil {
		t.Error(err)
	}
	if len(got) == 0 {
		t.Error("No items returned")
	}
	t.Log(got)
}

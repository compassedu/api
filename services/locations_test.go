package services_test

import (
	"testing"

	"github.com/compassedu/api/services"
)

func TestGetAllLocations(t *testing.T) {
	c, _, err := services.Login("", "", "")
	if err != nil {
		t.Error(err)
	}
	got, err := services.GetAllLocations(c, "")
	if err != nil {
		t.Error(err)
	}
	if len(got) == 0 {
		t.Error("No items returned")
	}
	t.Log(got)
}

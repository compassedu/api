package services_test

import (
	"os"
	"testing"

	"github.com/compassedu/api/services"
)

func TestGetAllLocations(t *testing.T) {
	p := os.Getenv("COMPASS_PASSWORD")
	un := os.Getenv("COMPASS_USERNAME")
	s := os.Getenv("COMPASS_SCHOOLID")
	c, _, err := services.Login(un, p, s)
	if err != nil {
		t.Error(err)
	}
	got, err := services.GetAllLocations(c, s)
	if err != nil {
		t.Error(err)
	}
	if len(got) == 0 {
		t.Error("No items returned")
	}
	t.Log(got)
}

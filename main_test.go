package compass

import (
	"os"
	"testing"
	"time"

	"github.com/compassedu/api/services"
)

func TestGetAllLocations(t *testing.T) {
	password := os.Getenv("COMPASS_PASSWORD")
	username := os.Getenv("COMPASS_USERNAME")
	schoolid := os.Getenv("COMPASS_SCHOOLID")
	c, _, err := services.Login(username, password, schoolid)
	if err != nil {
		t.Error(err)
	}
	locations, err := services.GetAllLocations(c, schoolid)
	if err != nil {
		t.Error(err)
	}
	t.Log(locations)
}
func TestGetAllStaff(t *testing.T) {
	password := os.Getenv("COMPASS_PASSWORD")
	username := os.Getenv("COMPASS_USERNAME")
	schoolid := os.Getenv("COMPASS_SCHOOLID")
	c, u, err := services.Login(username, password, schoolid)
	if err != nil {
		t.Error(err)
	}
	staff, err := services.GetAllStaff(c, schoolid, u)
	if err != nil {
		t.Error(err)
	}
	t.Log(staff)
}
func TestGetAllActivies(t *testing.T) {
	password := os.Getenv("COMPASS_PASSWORD")
	username := os.Getenv("COMPASS_USERNAME")
	schoolid := os.Getenv("COMPASS_SCHOOLID")
	c, u, err := services.Login(username, password, schoolid)
	if err != nil {
		t.Error(err)
	}
	staff, err := services.GetCalendarEventsByUser(c, schoolid, u, time.Now(), time.Now())
	if err != nil {
		t.Error(err)
	}
	t.Log(staff)
}

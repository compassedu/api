package services_test

import (
	"testing"

	"github.com/compassedu/api/services"
)

func TestGetAllStaff(t *testing.T) {
	c, u, err := services.Login("", "*", "")
	if err != nil {
		t.Error(err)
	}
	got, err := services.GetAllStaff(c, "", u)
	if err != nil {
		t.Error(err)
	}
	if len(got) == 0 {
		t.Error("No items returned")
	}
	t.Log(got)
}

package services_test

import (
	"testing"

	"github.com/compassedu/api/services"
)

func TestLogin(t *testing.T) {
	_, u, err := services.Login("KOGO25", "Killian20*", "kilianssc-ie")
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

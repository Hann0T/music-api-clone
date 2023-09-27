package models_test

import (
	"testing"

	"github.com/Hann0/music-api-clone/internal/models"
)

func TestPasswordValidation(t *testing.T) {
	user, err := models.NewUser("name", "email", "password")

	if err != nil {
		t.Fatalf(`models.NewUser("") = %q, %v, error`, user, err)
	}
}

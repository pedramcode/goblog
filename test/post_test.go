package test

import (
	logic "github.com/pedramcode/goblog/logics"
	"github.com/pedramcode/goblog/models"
	"testing"
)

var testUser models.User

func TestCreateUserForPost(t *testing.T) {
	Setup()
	user, err := logic.UserCreate("rose", "123", "Rose", "Doe", "rose@gmail.com", "+9811111111")
	if err != nil {
		t.Fatalf("Unable to create user: %s", err.Error())
	}
	testUser = user
}

func TestCreatePost(t *testing.T) {
	_, err := logic.PostCreate(testUser.ID, "This is a test", "Hello World!", nil)
	if err != nil {
		t.Fatalf("Unable to create post: %s", err.Error())
	}
}

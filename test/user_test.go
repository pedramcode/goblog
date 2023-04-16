package test

import (
	"github.com/pedramcode/goblog/core"
	logic "github.com/pedramcode/goblog/logics"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

func Setup() {
	log.Println("Preparing database connection")
	db, err := gorm.Open(sqlite.Open(core.TestSqlitePath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	core.SetDB(db)
	core.Migrate()
	log.Println("Database is ready")
}

func TestCreateUser(t *testing.T) {
	Setup()
	_, err := logic.UserCreate("john", "123", "John", "Doe", "john@gmail.com", "+9811111111")
	if err != nil {
		t.Fatalf("Unable to create user: %s", err.Error())
	}
}

func TestCreateUserBadEmail(t *testing.T) {
	_, err := logic.UserCreate("tom", "123", "Tomas", "Doe", "badEmail", "+9811111111")
	if err == nil {
		t.Fatal("Email format checking failed")
	}
}

func TestCreateNoEmail(t *testing.T) {
	_, err := logic.UserCreate("bill", "123", "Bill", "Doe", "", "+9811111111")
	if err != nil {
		t.Fatal("Email shouldn't be required")
	}
}

func TestTokenGeneration(t *testing.T) {
	_, err := logic.TokenCreate(1)
	if err != nil {
		t.Fatal("Token generation failed")
	}
}

func TestTokenGetForUser(t *testing.T) {
	token1, err := logic.TokenGetByUserIDCreate(2)
	if err != nil {
		t.Fatal("Token generation failed")
	}
	key1 := token1.Key

	token2, err := logic.TokenGetByUserIDCreate(2)
	if err != nil {
		t.Fatal("Token generation failed")
	}
	key2 := token2.Key

	if key1 != key2 {
		t.Fatal("Tokens aren't same, without logout")
	}

}

// This should be last test function
func TestCleanUp(t *testing.T) {
	err := os.Remove(core.TestSqlitePath)
	if err != nil {
		t.Fatal("Failed to cleanup test database")
	}
}

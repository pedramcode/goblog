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

// This should be last test function
func TestCleanUp(t *testing.T) {
	err := os.Remove(core.TestSqlitePath)
	if err != nil {
		t.Fatal("Failed to cleanup test database")
	}
}

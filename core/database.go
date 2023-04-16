package core

import (
	"log"

	"github.com/pedramcode/goblog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var databaseErr error

func InitDb() {
	log.Println("Preparing database connection")
	database, databaseErr = gorm.Open(sqlite.Open(SqlitePath), &gorm.Config{})
	if databaseErr != nil {
		panic(databaseErr)
	}
	log.Println("Database is ready")
}

func SetDB(dbN *gorm.DB) {
	database = dbN
}

func Migrate() {
	log.Println("Migrating models")
	if database == nil {
		panic("Database isn't ready for migration")
	}

	// Model Schemas
	err := database.AutoMigrate(
		&models.User{},
		&models.Comment{},
		&models.Post{},
		&models.Token{},
	)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Migration done")
}

func DB() *gorm.DB {
	return database
}

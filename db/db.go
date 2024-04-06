package db

import (
	"htmx-demo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect with database
func Connect() {
	gormDb, err:= gorm.Open(sqlite.Open("htmx-demo-dev.db"), &gorm.Config{})
	
	if err != nil {
	  panic("Failed to connect to the database.")
	}

	db = gormDb
}

func Get() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&models.Person{})
}
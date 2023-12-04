package database

import (
	"rentalmanagement/infrastructure/database/entities"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var InMemDB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// Connecting to SQLite in-memory database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	InMemDB = db
	return db, nil
}

func Migrate() error {
	err := InMemDB.AutoMigrate(&entities.RentalPersistenceEntity{}) // Add your model structs here
	if err != nil {
		return err
	}
	err = InMemDB.AutoMigrate(&entities.CustomerPersistenceEntity{}) // Add your model structs here
	if err != nil {
		return err
	}

	return nil
}

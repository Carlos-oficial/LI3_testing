package db

import (
	"LI3_testing/config"
	"LI3_testing/internal/models"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func setupSQLite(config *config.DatabaseEnvironmentVars) (*gorm.DB, error) {
	dbLocation := config.Path
	if dbLocation == "" {
		dbLocation = "../uber2.db"
	}

	// Create the sqlite file if it's not available
	if _, err := os.Stat(dbLocation); err != nil {
		if _, err = os.Create(dbLocation); err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	return db, err
}

func InitializeDatabase(config *config.DatabaseEnvironmentVars) error {
	db, err := setupSQLite(config)
	if err != nil {
		return err
	}
	err = models.AutoMigrate(db)
	if err != nil {
		return err
	}
	dbInstance = db
	return nil
}

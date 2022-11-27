package db

import (
	"LI3_testing/config"
	"LI3_testing/internal/models"
	"log"
)

func PopulateDB(db_vars config.DatabaseEnvironmentVars, seed_vars config.SeedEnvironmentVars) {
	log.Printf("Populating DB")
	db := GetDB()
	// models.LoadFile(db, &models.Driver{}, seed_vars.DriversFile)
	models.LoadFile(db, &models.User{}, seed_vars.UsersFile)
	// models.LoadFile(db, &models.Ride{}, seed_vars.RidesFile)
}

package main

import (
	"LI3_testing/config"
	"LI3_testing/internal/db"
	"LI3_testing/internal/models"
	"fmt"
)

func main() {
	config := config.InitEnvVars()
	db.InitializeDatabase(config.DbVars)
	db.PopulateDB(*config.DbVars, *config.SeedVars)
	database := db.GetDB()
	user := models.User{}
	fmt.Printf("%v", database.Find(&user))
}

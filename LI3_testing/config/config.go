package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type EnvLoader interface {
	Load() error
}

type DatabaseEnvironmentVars struct {
	Path string
}

type SeedEnvironmentVars struct {
	UsersFile   string
	DriversFile string
	RidesFile   string
}

type EnvironmentVars struct {
	DbVars   *DatabaseEnvironmentVars
	SeedVars *SeedEnvironmentVars
}

var environmentVars EnvironmentVars = EnvironmentVars{}

func InitEnvVars() *EnvironmentVars {
	log.Printf("Initializing environment variables")
	var db_vars = &DatabaseEnvironmentVars{}
	Init(db_vars)

	var seed_vars = &SeedEnvironmentVars{}
	Init(seed_vars)

	environmentVars.DbVars = db_vars
	environmentVars.SeedVars = seed_vars
	return &environmentVars
}

func Init(loadable EnvLoader) *EnvLoader {
	log.Printf("Loading: %v of type %T", loadable, loadable)

	viper.SetConfigFile("../.env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	errors := loadable.Load()
	if errors != nil {
		log.Fatalf("Some error occured. Err: %s", errors)
	}
	return &loadable
}

func (db_vars *DatabaseEnvironmentVars) Load() error {
	log.Printf("Loading db environment vars: %v", db_vars)
	db_vars.Path = viper.GetString("DB_PATH")
	log.Printf("Loaded db environment vars: %v", db_vars)
	return nil
}

func (seed_vars *SeedEnvironmentVars) Load() error {
	log.Printf("Loading seed environment vars: %v", seed_vars)
	seed_vars.UsersFile = viper.GetString("USERS")
	seed_vars.DriversFile = viper.GetString("DRIVERS")
	seed_vars.RidesFile = viper.GetString("RIDES")
	log.Printf("Loaded seed environment vars: %v", seed_vars)
	if seed_vars.DriversFile == "" || seed_vars.RidesFile == "" || seed_vars.UsersFile == "" {
		return fmt.Errorf("expected non empty strings got users:%s, drivers:%s, rides:%s", seed_vars.UsersFile, seed_vars.DriversFile, seed_vars.RidesFile)
	}
	return nil
}

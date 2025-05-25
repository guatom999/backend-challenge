package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App App
		Db  Db
		Jwt Jwt
	}
	App struct {
		Port string
	}

	Db struct {
		Uri string
	}

	Jwt struct {
		Secret string
	}
)

func GetConfig() *Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading env file %s", err.Error())
	}

	return &Config{
		App: App{
			Port: os.Getenv("APP_PORT"),
		},
		Db: Db{
			Uri: os.Getenv("DB_URI"),
		},
		Jwt: Jwt{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}

func GetMigrateConfig() *Config {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading env file %s", err.Error())
	}

	return &Config{
		Db: Db{
			Uri: os.Getenv("DB_URI"),
		},
	}
}

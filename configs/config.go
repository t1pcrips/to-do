package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Path PathConfig
}

type DbConfig struct {
	Dsn  string
	Port string
}

type PathConfig struct {
	Url  string
	Port string
}

func LoadConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Db: DbConfig{
			Dsn:  os.Getenv("DSN"),
			Port: os.Getenv("POSTGRES_PORT"),
		}, Path: PathConfig{
			Url:  os.Getenv("APP_PATH"),
			Port: os.Getenv("APP_PORT"),
		},
	}
}

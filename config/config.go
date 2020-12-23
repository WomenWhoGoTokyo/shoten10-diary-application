package config

import "os"

var Config *config

type config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string

	ServerPort string
}

func init() {
	Config = &config{
		DBUsername: os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DATABASE"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

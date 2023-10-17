package config

import "os"

type Config struct {
	//Postgres
	DBUser string
	DBPassword string
	DBHost string
	DBPort string
	DBName string
	//Server
	Port string
	IsDebug bool
}

func New() *Config {
	return &Config{
		DBUser: os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBHost: os.Getenv("POSTGRES_HOST"),
		DBPort: os.Getenv("POSTGRES_PORT"),
		DBName: os.Getenv("POSTGRES_DATABASE"),
		Port: os.Getenv("PORT"),
		IsDebug: os.Getenv("IS_DUBUG") == "true",
	}
}
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppVersion string
	AppPort    string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
	DbName     string
	DbSslMode  string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{
		AppName:    getEnvVar("APP_NAME"),
		AppVersion: getEnvVar("APP_VERSION"),
		AppPort:    getEnvVar("APP_PORT"),
		DbHost:     getEnvVar("DB_HOST"),
		DbUser:     getEnvVar("DB_USER"),
		DbPassword: getEnvVar("DB_PASSWORD"),
		DbPort:     getEnvVar("DB_PORT"),
		DbName:     getEnvVar("DB_NAME"),
		DbSslMode:  getEnvVar("DB_SSL_MODE"),
	}

	fmt.Println("Configuration loaded")

	return cfg
}

func getEnvVar(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		fmt.Printf("Environment variable %s is an empty string.\n", key)
		return ""
	}
}

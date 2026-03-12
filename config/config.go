package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	DBIP       string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func LoadConfig() *ConfigStruct {

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println(".env file not found")
	}

	return &ConfigStruct{
		DBIP:       os.Getenv("DB_IP"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

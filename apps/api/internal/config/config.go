package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	Name string
	Host string
	User string
	Pass string
}

type Email struct {
	Email string
	AppPassword string
}

type Config struct {
	DB DB
	JwtSecretKey string
	Port string
	Email Email
} 

func InitConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	if err := godotenv.Load(".env."+env); err != nil {
		log.Fatalf("Error loading env %s file", env)
	}


	cfg := &Config{
		DB: DB{
			Name: os.Getenv("DB_NAME"),
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
		},
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
		Port: os.Getenv("PORT"),
		Email: Email{
			Email: os.Getenv("EMAIL"),
			AppPassword: os.Getenv("APPPASSWORD"),
		},
	}

	return cfg
}
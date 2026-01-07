package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	JWTSecret   string
	JWTExpiry   string
	DatabaseURL string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:        os.Getenv("PORT"),
		Env:         os.Getenv("APP_ENV"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		JWTExpiry:   os.Getenv("JWT_EXPIRY"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	fmt.Println(cfg.Env)

	if cfg.Port == "" {
		log.Fatal("Port is not set")
	}

	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return cfg
}

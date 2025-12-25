package config

import (
	"os"
	"strings"
	"github.com/joho/godotenv"	
)

type AppConfig struct {
	Port       string
	Environment string
}

func Load() AppConfig {
	_ = godotenv.Load(".env")

	port := getEnv("Port","8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	
	environment := getEnv("Environment","development")

	return AppConfig{
		Port: port,
		Environment: environment,
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
		return v
	}
	return def
}

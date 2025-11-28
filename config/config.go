package config

import (
	"os"
	"strings"
	"fmt"
	"github.com/joho/godotenv"	
)

type AppConfig struct {
	Port 	string
	Environment string

	DBUser 	string
	DBPassword 	string
	DBHost 	string
	DBPort 	string
	DBName 	string
}

func Load() AppConfig {
	_ = godotenv.Load(".env")

	port := getEnv("Port","8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	
	environment := getEnv("Environment","development")

	dbUser := getEnv("DB_User","root")
	dbPassword := getEnv("DB_Password","")
	dbHost := getEnv("DB_Host","localhost")
	dbPort := getEnv("DB_Port","3306")
	dbName := getEnv("DB_Name","todo_app")

	return AppConfig{
		Port: port,
		Environment: environment,
		DBUser: dbUser,
		DBPassword: dbPassword,
		DBHost: dbHost,
		DBPort: dbPort,
		DBName: dbName,
	}
}

func (c AppConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && strings.TrimSpace(v) != "" {
		return v
	}
	return def
}
package config

import (
	"os"
	"strings"
	"fmt"
	"github.com/joho/godotenv"	
)

type AppConfig struct {
	Port       string
	Environment string

	DBUser string
	DBPassword string
	DBHost string
	DBPort string
	DBName string

	JWTSecret        string
    JWTExpiresMinute string
}

func Load() AppConfig {
	_ = godotenv.Load(".env")

	port := getEnv("Port","8080")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	
	environment := getEnv("Environment","development")

	dbUser := getEnv("DB_USER","root")
	dbPassword := getEnv("DB_PASSWORD","")
	dbHost := getEnv("DB_HOST","localhost")
	dbPort := getEnv("DB_PORT","3306")
	dbName := getEnv("DB_NAME","todo_app")
	jwtSecret := getEnv("JWT_SECRET","secret")
	jwtExpiresMinute := getEnv("JWT_EXPIRES_MINUTE","60")

	return AppConfig{
		Port: port,
		Environment: environment,

		DBUser: dbUser,
		DBPassword: dbPassword,
		DBHost: dbHost,
		DBPort: dbPort,
		DBName: dbName,
		JWTSecret: jwtSecret,
		JWTExpiresMinute: jwtExpiresMinute,
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

package database

import (
	"log"

	"github.com/Chien0903/Go-ToDo-App/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Nhận dsn string, KHÔNG nhận AppConfig nữa
func Connect(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate bảng User và Todo
	if err := DB.AutoMigrate(&models.User{}, &models.Todo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

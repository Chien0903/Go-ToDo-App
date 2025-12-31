package database

import (
	"log"

	"github.com/Chien0903/Go-ToDo-App/internal/config"
	"github.com/Chien0903/Go-ToDo-App/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB là connection GORM dùng chung trong app (nếu bạn muốn dùng kiểu global).
var DB *gorm.DB

// Connect khởi tạo kết nối MySQL dùng AppConfig.DSN() để tránh trùng logic build DSN.
func Connect(cfg config.AppConfig) error {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}

	DB = db
	log.Println("✅ MySQL database connected successfully")

	// Auto-migrate models
	if err := migrateModels(db); err != nil {
		return err
	}

	return nil
}

// migrateModels tự động tạo/update tables dựa trên models
func migrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	log.Println("✅ Database models migrated successfully")
	return nil
}

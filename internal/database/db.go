package database

import (
	"log"

	"github.com/Chien0903/Go-ToDo-App/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB là connection GORM dùng chung trong app (nếu bạn muốn dùng kiểu global).
var DB *gorm.DB

// Connect khởi tạo kết nối MySQL dùng AppConfig.DSN() để tránh trùng logic build DSN.
func Connect(cfg config.AppConfig) error {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	DB = db
	log.Println("✅ MySQL database connected successfully")
	return nil
}

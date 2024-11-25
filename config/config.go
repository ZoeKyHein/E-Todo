package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load() // 加载 .env 文件
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// 从环境变量读取 MySQL 配置信息
	dsn := os.Getenv("MYSQL_DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")
}

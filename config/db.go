package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	// 1. 拼接 DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	// 2. 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 3. 设置数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("获取底层数据库实例失败: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)           // 空闲连接
	sqlDB.SetMaxOpenConns(100)          // 最大连接
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最长存活时间

	// 赋值给全局 DB
	DB = db

	log.Println("✅ 数据库连接成功")
}

// โค้ดคำสั่ง connect database
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	username = "root"
	password = ""
	dbname   = "sau_db"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("DB connected successful")
	// ทำการ Migration Database กับตัว Model ที่สร้างขึ้นมา
	if err := db.AutoMigrate(&Student{}); err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("Migration successful....")

	return db, nil
}

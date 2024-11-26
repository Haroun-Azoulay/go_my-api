package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
    // if you want connect without docker uncomment this line bellow
	// dsn := "root:root@tcp(127.0.0.1:3306)/myapi?charset=utf8mb4&parseTime=True&loc=Local"
    dsn := "root:root@tcp(my-api-db:3306)/myapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	DB = db

	log.Println("Database connection successful")
	return nil
}

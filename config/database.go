package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// Default XAMPP/MySQL connection string
		dsn = "root:@tcp(127.0.0.1:3306)/akademik_db?charset=utf8mb4&parseTime=True&loc=Local"
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database MySQL:", err)
	}

	log.Println("Berhasil terhubung ke database MySQL!")
	DB = database
}

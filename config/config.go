package config

import (
	"demonstrasi/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	fmt.Printf("MYSQLUSER: %s\n", os.Getenv("MYSQLUSER"))
	fmt.Printf("MYSQLPASSWORD: %s\n", os.Getenv("MYSQLPASSWORD"))
	fmt.Printf("MYSQLHOST: %s\n", os.Getenv("MYSQLHOST"))
	fmt.Printf("MYSQLPORT: %s\n", os.Getenv("MYSQLPORT"))
	fmt.Printf("MYSQLDATABASE: %s\n", os.Getenv("MYSQLDATABASE"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQLUSER"),
		os.Getenv("MYSQLPASSWORD"),
		os.Getenv("MYSQLHOST"),
		os.Getenv("MYSQLPORT"), // Pastikan ini diperbaiki
		os.Getenv("MYSQLDATABASE"),
	)

	// Log connection string for debugging
	fmt.Printf("Connecting to database with DSN: %s\n", dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to database!")
}

func InitialMigration() {
	if DB == nil {
		log.Fatalf("Database connection is not initialized")
	}
	err := DB.AutoMigrate(
		&models.TbLogin{},
		&models.TbInformation{},
		&models.TbComment{},
		&models.TbGalery{},
		&models.TbMonument{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed!")
}

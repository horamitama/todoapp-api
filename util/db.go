package util

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load(".env.development")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	defer fmt.Println("DB connented")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("db connect failed:", err)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	defer fmt.Println("DB closed")
	if err := sqlDB.Close(); err != nil {
		log.Fatal("err")
	}
}

func SetupTestDB() *gorm.DB {
	port := "54320"
	name := "postgres"
	user := "postgres"
	password := "postgres"
	host := "localhost"

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user,
		password, host,
		port, name)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("db connect failed")
	}

	return db
}

func TeardownTestDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic("db close failed")
	}
}

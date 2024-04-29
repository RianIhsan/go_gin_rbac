package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dbhost := os.Getenv("DB_HOST")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbhost, dbuser, dbpass, dbname, dbport)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
		return nil
	}
	log.Println("`Successfully connected to the database")

	return db

}

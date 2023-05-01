package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func getConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
}

func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("postgres", getConnectionString())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	return db
}

func (User) TableName() string {
	return "User"
}
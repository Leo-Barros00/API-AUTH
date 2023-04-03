package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
		})
	})

	router.Run(":8080")
}
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	err := loadEnv()
	if err != nil {
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	db := getDatabaseConnection()
	defer db.Close()

	router := gin.Default()

	router.GET("/health", healthHandler)
	router.POST("/login", func(c *gin.Context) {
		loginHandler(c, db)
	})
	router.POST("/refresh", func(c *gin.Context) {
		refreshHandler(c, db)
	})

	router.Run(":3330")
}
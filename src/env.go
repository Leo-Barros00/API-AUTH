package main

import (
	"github.com/joho/godotenv"
)

func loadEnv() error {
	err := godotenv.Load()
	return err
}
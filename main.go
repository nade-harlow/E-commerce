package main

import (
	"github.com/joho/godotenv"
	"github.com/nade-harlow/E-commerce/adapter/api/server"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil && os.Getenv("ENV") != "dev" {
		log.Fatal("Error loading .env file!")
	}
	server.Start()
}

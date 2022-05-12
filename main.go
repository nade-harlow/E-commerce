package main

import (
	"github.com/joho/godotenv"
	"github.com/nade-harlow/E-commerce/adapter/api/server"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{PrettyPrint: true}
	log.SetOutput(logger.Writer())
	// If we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Llongfile)

	err := godotenv.Load()
	if err != nil && os.Getenv("ENV") != "dev" {
		log.Fatal("Error loading .env file!")
	}
	server.Start()
}

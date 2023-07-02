package main

import (
	"github.com/joho/godotenv"
	"log"
	"modules/internal/app"
)

// Иннициализация переменных окружения
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app.RunServer()
}

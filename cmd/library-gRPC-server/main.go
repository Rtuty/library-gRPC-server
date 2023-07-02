package main

import (
	"context"
	"database/sql"
	"log"
	"modules/internal/config"
	"modules/internal/db"
	"modules/internal/models"

	"github.com/joho/godotenv"
)

// Иннициализация переменных окружения
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := context.Background()
	cf, err := config.GetConnection()
	if err != nil {
		panic(err)
	}

	con, err := sql.Open("mysql", cf.User+":"+cf.Passwd+"@/"+cf.Dbname)
	if err != nil {
		panic(err)
	}

	strg := db.NewMySqlRepository(con)

	upd := models.Author{
		ID:      1,
		Name:    "А.Уе.121123541234",
		Country: "Россия",
	}

	if err = strg.AuthorMethodsHandler(ctx, "create", upd); err != nil {
		panic(err)
	}
}

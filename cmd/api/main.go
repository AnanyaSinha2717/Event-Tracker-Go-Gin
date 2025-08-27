package main

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"
)

type application struct {
	Port   int
	Secret string
	Models database.Models
}

func main() {
	// setting up connection with db
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // closes connection

	models := database.NewModels(db)

	app := &application{
		Port:   env.GetEnvInt("LOCALHOST:", 8080),
		Secret: env.GetenvString("SECRET", "yolo"),
		Models: models,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}

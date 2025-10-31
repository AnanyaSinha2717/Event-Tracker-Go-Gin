package main

import (
	"database/sql"
	_ "Event-Tracker-Go-Gin/docs"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"
)

type application struct {
	port   int
	jwtSecret string
	models database.Models
}

func main() {
	// setting up connection with db
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // closes connection

	m := database.NewModels(db)

	app := &application{
		port:   env.GetEnvInt("LOCALHOST:", 8080),
		jwtSecret: env.GetenvString("JWT_SECRET", "This is a JWT secret"),
		models: m,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}

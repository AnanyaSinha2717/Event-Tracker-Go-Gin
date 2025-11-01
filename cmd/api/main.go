package main

import (
	"database/sql"
	_ "event-tracker-go-gin/docs"
	"event-tracker-go-gin/internal/database"
	"event-tracker-go-gin/internal/env"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// @title Go Gin Rest API
// @version 1.0
// @description A rest API in Go using Gin framework to make an Event Tracker
// @securityDefinitions.apikey BearerAuth
// @in header
//@name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

type application struct {
	port      int
	jwtSecret string
	models    database.Models
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
		port:      env.GetEnvInt("LOCALHOST:", 8080),
		jwtSecret: env.GetenvString("JWT_SECRET", "This is a JWT secret"),
		models:    m,
	}

	if err := serve(app); err != nil {
		log.Fatal(err)
	}
}

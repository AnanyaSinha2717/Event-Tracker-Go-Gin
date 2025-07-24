package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4/database/sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3/", "./data.db")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance{db, &sqlite3.Config{}}
	if err != nil{
		log.Fatal(err)
	}
}

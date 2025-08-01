package main

import (
	"database/sql"
	"log"
)

func main() {
	// setting up connection with db
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close() // closes connection
}

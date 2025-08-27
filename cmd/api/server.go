package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func serve(app *application) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on port: %d", app.Port)

	return server.ListenAndServe()
}

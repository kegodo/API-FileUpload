package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// api version
const version = "1.0.0"

// config struct
type config struct {
	port int
	env  string //development, staging, production
}

// application struct
type application struct {
	config config
	logger *log.Logger
}

// main
func main() {
	var cfg config
	cfg.port = 4000
	cfg.env = "development"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

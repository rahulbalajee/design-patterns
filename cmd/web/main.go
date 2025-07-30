package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rahulbalajee/design-patterns/streamer"
)

const port = ":4000"

func main() {
	// Initialize video processing queue
	initStreamer()
	defer closeStreamer()

	var config appConfig
	flag.BoolVar(&config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&config.DSN, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN for database")
	flag.Parse()

	// Initialize database (dependency)
	db, err := initMySQLDB(config.DSN)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	// Use constructor injection instead of field assignment
	app := NewApplication(db, config) // ← Better dependency injection

	wp := streamer.New(videoQueue, numWorkers)
	wp.Run()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web server on port:", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

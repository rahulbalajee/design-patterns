package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateCache map[string]*template.Template
	config        appConfig
	DB            *sql.DB
}

type appConfig struct {
	useCache bool
	DSN      string
}

func main() {
	app := application{
		templateCache: map[string]*template.Template{},
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.DSN, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	// get database
	db, err := initMySQLDB(app.config.DSN)
	if err != nil {
		log.Panic(err)
	}

	app.DB = db

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web server on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

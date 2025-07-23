package main

import (
	"database/sql"
	"html/template"

	"github.com/rahulbalajee/design-patterns/models"
)

// application holds the application-wide dependencies
type application struct {
	templateCache map[string]*template.Template
	config        appConfig
	DB            *sql.DB
	Models        models.Models
}

// appConfig holds configuration settings
type appConfig struct {
	useCache bool
	DSN      string
}

// NewApplication is a constructor that uses dependency injection
func NewApplication(db *sql.DB, config appConfig) *application {
	return &application{
		DB:            db,
		config:        config,
		templateCache: make(map[string]*template.Template),
		Models:        *models.New(db),
	}
}

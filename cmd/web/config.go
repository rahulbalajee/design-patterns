package main

import (
	"database/sql"
	"html/template"

	"github.com/rahulbalajee/design-patterns/adapters"
	"github.com/rahulbalajee/design-patterns/configuration"
)

// application holds the application-wide dependencies
type application struct {
	templateCache map[string]*template.Template
	config        appConfig
	App           *configuration.Application
}

// appConfig holds configuration settings
type appConfig struct {
	useCache bool
	DSN      string
}

// NewApplication is a constructor that uses dependency injection
func NewApplication(db *sql.DB, config appConfig) *application {
	return &application{
		templateCache: make(map[string]*template.Template),
		config:        config,
		App:           configuration.New(db, &adapters.RemoteService{Remote: &adapters.XMLBackend{}}),
	}
}

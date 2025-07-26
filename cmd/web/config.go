package main

import (
	"database/sql"
	"html/template"

	"github.com/rahulbalajee/design-patterns/configuration"
)

// application holds the application-wide dependencies
type application struct {
	templateCache map[string]*template.Template
	config        appConfig
	App           *configuration.Application
	catService    *RemoteService
}

// appConfig holds configuration settings
type appConfig struct {
	useCache bool
	DSN      string
}

// NewApplication is a constructor that uses dependency injection
func NewApplication(db *sql.DB, config appConfig) *application {
	return &application{
		config:        config,
		templateCache: make(map[string]*template.Template),
		App:           configuration.New(db),
		catService:    &RemoteService{Remote: &JSONBackend{}},
	}
}

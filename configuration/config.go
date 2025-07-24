package configuration

import (
	"database/sql"
	"sync"

	"github.com/rahulbalajee/design-patterns/models"
)

type Application struct {
	Models *models.Models
}

var instance *Application
var once sync.Once
var db *sql.DB
var initialized bool
var initMutex sync.Mutex

func New(pool *sql.DB) *Application {
	initMutex.Lock()
	defer initMutex.Unlock()

	// Only set db on first call - prevents accidental overwrites!
	if !initialized && pool != nil {
		db = pool
		initialized = true
	}

	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
		}
	})
	return instance
}

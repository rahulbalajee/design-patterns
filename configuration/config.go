package configuration

import (
	"database/sql"
	"sync"

	"github.com/rahulbalajee/design-patterns/adapters"
	"github.com/rahulbalajee/design-patterns/models"
)

type Application struct {
	Models     *models.Models
	CatService *adapters.RemoteService
}

var instance *Application
var once sync.Once
var db *sql.DB
var initialized bool
var initMutex sync.Mutex
var catService *adapters.RemoteService

func New(pool *sql.DB, cs *adapters.RemoteService) *Application {
	initMutex.Lock()
	defer initMutex.Unlock()

	// Only set db on first call - prevents accidental overwrites!
	if pool != nil && cs != nil && !initialized {
		db = pool
		catService = cs
		initialized = true
	}

	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models:     models.New(db),
			CatService: catService,
		}
	})
	return instance
}

package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// Display our test page
	mux.Get("/test-patterns", app.TestPatterns)

	// Factory routes
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)
	mux.Get("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
	mux.Get("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	mux.Get("/api/dog-breeds", app.GetAllDogBreedsJSON)

	mux.Get("/api/dog-from-builder", app.CreateDogWithBuilder)
	mux.Get("/api/cat-from-builder", app.CreateCatWithBuilder)
	mux.Get("/api/cat-breeds", app.GetAllCatBreeds)

	mux.Get("/api/animal-from-abstract-factory/{species}/{breed}", app.AnimalFromAbstractFactory)

	return mux
}

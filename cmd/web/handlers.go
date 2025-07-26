package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rahulbalajee/design-patterns/pets"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	app.writeJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("some desc").
		SetColor("Black").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	app.writeJSON(w, http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("felis silverstris breed").
		SetWeight(15).
		SetDescription("some desc").
		SetColor("Black").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	app.writeJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	app.writeJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	// Get species from URL itself

	// Get breed from the URL

	// Create a pet from Abstract Factory

	// Write the result as JSON
}

package main

import (
	"os"
	"testing"

	"github.com/rahulbalajee/design-patterns/configuration"
	"github.com/rahulbalajee/design-patterns/models"
)

var testApp application

func TestMain(m *testing.M) {

	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}

	testApp = application{
		App:        configuration.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{
			ID:      1,
			Breed:   "Tomcat",
			Details: "something",
		},
	}

	return breeds, nil
}

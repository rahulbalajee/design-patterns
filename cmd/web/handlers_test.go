package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	// create a request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	// create a response recorder
	rr := httptest.NewRecorder()

	// create the handler
	handler := http.HandlerFunc(testApp.GetAllDogBreedsJSON)

	// serve the handler
	handler.ServeHTTP(rr, req)

	// check the response
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code; got %d wanted 200", rr.Code)
	}
}

func TestApplication_GetAllCatBreeds(t *testing.T) {
	// create a request
	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)

	// create a response recorder
	rr := httptest.NewRecorder()

	// create the handler
	handler := http.HandlerFunc(testApp.GetAllCatBreeds)

	// serve the handler
	handler.ServeHTTP(rr, req)

	// check the response
	if rr.Code != http.StatusOK {
		t.Errorf("wrong response code; got %d wanted 200", rr.Code)
	}
}

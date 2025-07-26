package main

import (
	"os"
	"testing"

	"github.com/rahulbalajee/design-patterns/adapters"
	"github.com/rahulbalajee/design-patterns/configuration"
)

var testApp application

func TestMain(m *testing.M) {

	testBackend := &adapters.TestBackend{}
	testAdapter := &adapters.RemoteService{Remote: testBackend}

	testApp = application{
		App: configuration.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}

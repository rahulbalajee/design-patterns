package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
	GetDogOfMonthByID(id int) (*DogOfMonth, error)
}

type mysqlRepository struct {
	DB *sql.DB
}

func newMySqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestMySqlRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}

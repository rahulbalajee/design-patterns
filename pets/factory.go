package pets

import "github.com/rahulbalajee/design-patterns/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no desc entered",
		Lifespan:    0,
	}

	return &pet
}

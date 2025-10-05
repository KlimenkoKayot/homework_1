package interfaces

import "github.com/klimenkokayot/homework_1/internal/domain/models/animal"

type Zoo interface {
	AddAnimal(animal animal.Animal) bool
	VetClinic
}

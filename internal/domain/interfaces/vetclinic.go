package interfaces

import "github.com/klimenkokayot/homework_1/internal/domain/models/animal"

type VetClinic interface {
	CheckAnimal(animal animal.Animal) bool
}

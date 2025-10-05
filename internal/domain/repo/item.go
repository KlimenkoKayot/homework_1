package repo

import "github.com/klimenkokayot/homework_1/internal/domain/models/animal"

type AliveRepo interface {
	AddAnimal(animal animal.Animal) error
	DeleteAnimal(animal.Animal)
}

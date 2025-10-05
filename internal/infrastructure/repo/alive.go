package repo

import "github.com/klimenkokayot/homework_1/internal/domain/interfaces"

func NewAnimalRepo[T interfaces.Animal]() *dataRepo[T] {
	return
}

type animalRepo struct {
}

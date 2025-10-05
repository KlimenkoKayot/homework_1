package service

import "github.com/klimenkokayot/homework_1/internal/domain/interfaces"

type AnimalManager struct {
	zoo       interfaces.Zoo
	vetClinic interfaces.VetClinic
}

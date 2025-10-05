package animal

func NewPredator(animal *Animal, agressiveLevel int) *Predator {
	return &Predator{
		Animal:         animal,
		agressiveLevel: agressiveLevel,
	}
}

type Predator struct {
	*Animal
	agressiveLevel int
}

func (h *Predator) AgressiveLevel() int {
	return h.agressiveLevel
}

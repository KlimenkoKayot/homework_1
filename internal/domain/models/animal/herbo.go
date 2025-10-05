package animal

func NewHerbo(animal *Animal, kindnessLevel int) *Herbo {
	return &Herbo{
		Animal:        animal,
		kindnessLevel: kindnessLevel,
	}
}

type Herbo struct {
	*Animal
	kindnessLevel int
}

func (h *Herbo) KindnessLevel() int {
	return h.kindnessLevel
}

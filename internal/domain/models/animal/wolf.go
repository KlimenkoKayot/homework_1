package animal

func NewWolf(predator *Predator) *Wolf {
	return &Wolf{
		Predator: predator,
	}
}

type Wolf struct {
	*Predator
}

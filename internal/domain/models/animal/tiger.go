package animal

func NewTiger(predator *Predator) *Tiger {
	return &Tiger{
		Predator: predator,
	}
}

type Tiger struct {
	*Predator
}

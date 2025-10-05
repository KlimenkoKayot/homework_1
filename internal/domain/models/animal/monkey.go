package animal

func NewMonkey(herbo *Herbo, intelligence int) *Monkey {
	return &Monkey{
		Herbo:        herbo,
		intelligence: intelligence,
	}
}

type Monkey struct {
	*Herbo
	intelligence int
}

func (h *Monkey) Intelligence() int {
	return h.intelligence
}

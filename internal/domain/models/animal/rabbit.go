package animal

func NewRabbit(herbo *Herbo) *Rabbit {
	return &Rabbit{
		Herbo: herbo,
	}
}

type Rabbit struct {
	*Herbo
}

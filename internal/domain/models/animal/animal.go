package animal

type HealthStatus int

const (
	Healthy HealthStatus = -1 + iota
	Unhealthy
)

func NewAnimal(id int, name string, status HealthStatus, foodIntake int) *Animal {
	return &Animal{
		id:           id,
		name:         name,
		healthStatus: status,
		foodIntake:   foodIntake,
	}
}

type Animal struct {
	id           int
	name         string
	healthStatus HealthStatus
	foodIntake   int
}

func (a *Animal) Food() int {
	return a.foodIntake
}

func (a *Animal) ID() int {
	return a.id
}

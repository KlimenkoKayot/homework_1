package animal

type HealthStatus int

const (
	Healthy HealthStatus = -1 + iota
	Unhealthy
)

func NewAnimal(number int, name string, status HealthStatus, foodIntake int) *Animal {
	return &Animal{
		number:       number,
		name:         name,
		healthStatus: status,
		foodIntake:   foodIntake,
	}
}

type Animal struct {
	number       int
	name         string
	healthStatus HealthStatus
	foodIntake   int
}

func (a *Animal) Food() int {
	return a.foodIntake
}

func (a *Animal) Number() int {
	return a.number
}

func (a *Animal) SetHealthStatus(status HealthStatus) bool {
	changed := a.healthStatus != status
	a.healthStatus = status
	return changed
}

package interfaces

type Alive interface {
	Food() int
}

type Inventory interface {
	ID() int
}

type Animal interface {
	Alive
	Inventory
}

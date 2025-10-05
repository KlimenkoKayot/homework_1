package item

func NewThing(number int) *Thing {
	return &Thing{
		number: number,
	}
}

type Thing struct {
	number int
}

func (t *Thing) Number() int {
	return t.number
}

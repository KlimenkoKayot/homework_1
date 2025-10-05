package item

func NewThing(id int) *Thing {
	return &Thing{
		id: id,
	}
}

type Thing struct {
	id int
}

func (t *Thing) ID() int {
	return t.id
}

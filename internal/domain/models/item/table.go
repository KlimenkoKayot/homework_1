package item

type Material int

const (
	Metal Material = -1 + iota
	Wood
)

func NewTable(number int, material Material) *Table {
	return &Table{
		number:   number,
		material: material,
	}
}

type Table struct {
	number   int
	material Material
}

func (t *Table) Number() int {
	return t.number
}

func (t *Table) Material() Material {
	return t.material
}

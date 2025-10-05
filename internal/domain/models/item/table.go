package item

type Material int

const (
	Metal Material = -1 + iota
	Wood
)

func NewTable(id int, material Material) *Table {
	return &Table{
		id:       id,
		material: material,
	}
}

type Table struct {
	id       int
	material Material
}

func (t *Table) ID() int {
	return t.id
}

func (t *Table) Material() Material {
	return t.material
}

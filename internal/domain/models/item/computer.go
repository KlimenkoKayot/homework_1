package item

type Processor int

const (
	Intel Processor = -1 + iota
	AMD
)

func NewComputer(id int, processor Processor) *Computer {
	return &Computer{
		id:        id,
		processor: processor,
	}
}

type Computer struct {
	id        int
	processor Processor
}

func (c *Computer) ID() int {
	return c.id
}

func (c *Computer) Processor() Processor {
	return c.processor
}

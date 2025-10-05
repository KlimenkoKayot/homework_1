package item

type Processor int

const (
	Intel Processor = -1 + iota
	AMD
)

func NewComputer(number int, processor Processor) *Computer {
	return &Computer{
		number:    number,
		processor: processor,
	}
}

type Computer struct {
	number    int
	processor Processor
}

func (c *Computer) Number() int {
	return c.number
}

func (c *Computer) Processor() Processor {
	return c.processor
}

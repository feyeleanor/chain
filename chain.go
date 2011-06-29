package chain

type Equatable interface {
	Equal(interface{}) bool
}

type Linear interface {
	Len() int
}

func Cons(items... interface{}) (c *Cell) {
	var n *Cell
	for i, v := range items {
		if i == 0 {
			c = &Cell{ Head: v }
			n = c
		} else {
			n.Tail = &Cell{ Head: v }
			n = n.Tail
		}
	}
	return
}
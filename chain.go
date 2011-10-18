package chain

type Equatable interface {
	Equal(interface{}) bool
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

func Append(c *Cell, x interface{}) *Cell {
	if c == nil {
		c = &Cell{ x, nil }
	} else {
		c.Append(x)
	}
	return c
}

func Prepend(c *Cell, x interface{}) *Cell {
	return &Cell{ x, c }
}
package chain

import "fmt"

type Cell struct {
	Head		interface{}
	Tail		*Cell
}

func (c Cell) Content() interface{} {
	return c.Head
}

func (c *Cell) End() (r *Cell) {
	if c != nil {
		for r = c; r.Tail != nil; r = r.Tail {}
	}
	return
}

func (c Cell) MoveTo(i int) (l Node) {
	switch {
	case i < 0:				break
	case i == 0:			l = &c
	default:				n := &c
							for ; i > 0 && n != nil; i-- {
								n = n.Tail
							}
							if n != nil {
								l = n
							}
	}
	return
}

func (c *Cell) Link(i int, l Node) (b bool) {
	if l == nil {
		if i == NEXT_NODE {
			c.Tail = nil
			b = true
		}
	} else {
		switch i {
		case CURRENT_NODE:		if n, ok := l.(*Cell); ok {
									c.Head = l.Content()
									c.Tail = n.Tail
									b = true
								} else {
									if t, ok := Next(l).(*Cell); ok {
										c.Head = l.Content()
										c.Tail = t
										b = true
									}
								}

		case NEXT_NODE:			if n, ok := l.(*Cell); ok {
									c.Tail = n
									b = true
								}
		}
	}
	return
}

func (c *Cell) Set(i int, v interface{}) bool {
	if i > PREVIOUS_NODE {
		if c == nil {
			*c = Cell{}
		}

		for ; i > 0; i-- {
			if c.Tail == nil {
				c.Tail = &Cell{}
			}
			c = c.Tail
		}
		c.Head = v
		return true
	}
	return false
}

func (c *Cell) Append(x interface{}) {
	c.Set(NEXT_NODE, x)
}

func (c *Cell) Prepend(x interface{}) {
	*c = Cell{ Head: x, Tail: c }
}

func (c Cell) equal(o Cell) (r bool) {
	defer func() {
		if x := recover(); x != nil {
			r = false
		}
	}()
	if v, ok := c.Head.(Equatable); ok {
		r = v.Equal(o.Head)
	} else {
		r = c.Head == o.Head
	}
	return
}

func (c *Cell) Equal(o interface{}) (r bool) {
	switch o := o.(type) {
	case *Cell:			r = o != nil && c.equal(*o)
	case Cell:			r = c.equal(o)
	default:				r = c.equal(Cell{ Head: o })
	}
	return
}

func (c Cell) String() (t string) {
	return fmt.Sprint(c.Head)
}

func (c *Cell) Car() (v interface{}) {
	if c != nil {
		v = c.Head
	}
	return
}

func (c *Cell) Cdr() (v *Cell) {
	if c != nil {
		v = c.Tail
	}
	return
}

func (c *Cell) Rplaca(i interface{}) {
	if c != nil {
		c.Head = i
	} else {
		*c = Cell{ Head: i }
	}
}

func (c *Cell) Rplacd(next *Cell) {
	if c != nil {
		c.Tail = next
	} else {
		*c = *next
	}
}
package chain

import (
	"fmt"
)

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
	default:			r = c.equal(Cell{ Head: o })
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

func (c *Cell) Caar() (v interface{}) {
	if c != nil && c.Tail != nil {
		v = c.Tail.Head
	}
	return
}

func (c *Cell) Car2() (x, y interface{}) {
	switch {
	case c == nil:			return nil, nil
	case c.Tail == nil:		return c.Head, nil
	}
	return c.Head, c.Tail.Head
}

func (c *Cell) Car3() (x, y, z interface{}) {
	switch {
	case c == nil:					return nil, nil, nil
	case c.Tail == nil:				return c.Head, nil, nil
	case c.Tail.Tail == nil:		return c.Head, c.Tail.Head, nil
	}
	return c.Head, c.Tail.Head, c.Tail.Tail.Head
}

func (c *Cell) Cdr() (v *Cell) {
	if c != nil {
		v = c.Tail
	}
	return
}

func (c *Cell) Cddr() (v *Cell) {
	if c != nil && c.Tail != nil {
		v = c.Tail.Tail
	}
	return
}

func (c *Cell) Cadr() (v *Cell) {
	if c != nil {
		if h, ok := c.Head.(*Cell); ok {
			v = h.Tail
		}
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

func (c *Cell) Each(f interface{}) {
	switch f := f.(type) {
	case func(interface{}):						for k := c; k != nil; k = k.Tail { f(k.Head) }
	case func(int, interface{}):				for i, k := 0, c; k != nil; k = k.Tail {
													f(i, k.Head)
													i++
												}
	case func(interface{}, interface{}):		for i, k := 0, c; k != nil; k = k.Tail {
													f(i, k.Head)
													i++
												}
	}
}

func (c *Cell) While(f interface{}) (i int, k *Cell) {
	switch f := f.(type) {
	case func(interface{}) bool:				for k = c; k != nil; k = k.Tail {
													if !f(k.Head) {
														break
													}
													i++
												}
	case func(int, interface{}) bool:			for k = c; k != nil; k = k.Tail {
													if !f(i, k.Head) {
														break
													}
													i++
												}
	case func(interface{}, interface{}) bool:	for k = c; k != nil; k = k.Tail {
													if !f(i, k.Head) {
														break
													}
													i++
												}
	case Equatable:								for k = c; k != nil; k = k.Tail {
													if !f.Equal(k.Head) {
														break
													}
													i++
												}
	case interface{}:							for k = c; k != nil; k = k.Tail {
													if f != k.Head {
														break
													}
													i++
												}
	}
	return
}

func (c *Cell) Until(f interface{}) (i int, k *Cell) {
	switch f := f.(type) {
	case func(interface{}) bool:				for k = c; k != nil; k = k.Tail {
													if f(k.Head) {
														break
													}
													i++
												}
	case func(int, interface{}) bool:			for k = c; k != nil; k = k.Tail {
													if f(i, k.Head) {
														break
													}
													i++
												}
	case func(interface{}, interface{}) bool:	for k = c; k != nil; k = k.Tail {
													if f(i, k.Head) {
														break
													}
													i++
												}
	case Equatable:								for k = c; k != nil; k = k.Tail {
													if f.Equal(k.Head) {
														break
													}
													i++
												}
	case interface{}:							for k = c; k != nil; k = k.Tail {
													if f == k.Head {
														break
													}
													i++
												}
	
	}
	return
}

func (c *Cell) Len() (l int) {
	for n := c; n != nil; n = n.Tail {
		l++
	}
	return
}

func (c *Cell) MinimumLength(l int) bool {
	for n := c; n != nil && l > 0; n = n.Tail {
		l--
	}
	return l == 0
}
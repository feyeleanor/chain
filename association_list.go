package chain

import (
	"fmt"
)

func (c *Cell) Define(key interface{}, value interface{}) (r *Cell) {
	if c == nil {
		r = &Cell{ Head: Cons(key, value) }
	} else {
		r = &Cell{ Head: Cons(key, value), Tail: c }
	}
	return
}

func (c *Cell) Assq(key interface{}) (r *Cell) {
	switch k := key.(type) {
	case Equatable:				for r = c; r != nil; r = r.Tail {
									if head, ok := r.Head.(*Cell); ok && k.Equal(head) {
										break
									}
								}
	case fmt.Stringer:			r = c.Assq(k.String())
	case interface{}:			for r = c; r != nil; r = r.Tail {
									if head, ok := r.Head.(*Cell); ok && k == head {
										break
									}
								}

	}
	return
}
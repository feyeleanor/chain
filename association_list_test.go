package chain

import (
	"testing"
)

func TestDefine(t *testing.T) {
	ConfirmDefine := func(c *Cell, k, v interface{}, r *Cell) {
		if n := c.Define(k, v); !n.Equal(r) {
			t.Fatalf("%v.Define(%v, %v) should be %v but is %v", c, k, v, r, n)
		}
	}

	ConfirmDefine(Cons(), "x", 17, Cons(Cons("x", 17)))
	ConfirmDefine(Cons(0), "x", 17, Cons(Cons("x", 17), 0))
	ConfirmDefine(Cons(Cons("y", 12)), "x", 17, Cons(Cons("x", 17), Cons("y", 12)))
}
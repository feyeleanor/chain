package chain

import (
	"testing"
)

func TestPrepend(t *testing.T) {
	ConfirmPrepend := func(c *Cell, v interface{}, r interface{}) {
		x := Prepend(c, v)
		if !x.Equal(r) {
			t.Fatalf("%v.Prepend(%v) should be %v but is %v", c.Tail, v, r, x)
		}
	}
	ConfirmPrepend(Cons(), 1, Cons(1))
	ConfirmPrepend(Cons(1), 2, Cons(2, 1))
}

func TestAppend(t *testing.T) {
	ConfirmAppend := func(c *Cell, v interface{}, r interface{}) {
		x := Append(c, v)
		if !x.Equal(r) {
			t.Fatalf("%v.Append(%v) should be %v but is %v", c.Tail, v, r, x)
		}
	}
	ConfirmAppend(Cons(), 1, Cons(1))
	ConfirmAppend(Cons(1), 2, Cons(1, 2))
}
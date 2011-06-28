package chain

import "testing"
import "reflect"

func TestCellEnd(t *testing.T) {
	ConfirmEnd := func(c *Cell, r interface{}) {
		x := c.End()
		switch {
		case x == nil:		t.Fatalf("%v.End() returned nil", c)
		case x.Head != r:	t.Fatalf("%v.End() should be '%v' but is '%v'", c, r, x.Head)
		}
	}
	RefuteEnd := func(c *Cell) {
		if x := c.End(); x != nil {
			t.Fatalf("%v.End() should be nil but is '%v'", c, x.Head)
		}
	}
	RefuteEnd(Cons())
	ConfirmEnd(Cons(0), 0)
	ConfirmEnd(Cons(0, 1), 1)
	ConfirmEnd(Cons(0, 1, 2), 2)
}

func TestCellMoveTo(t *testing.T) {
	ConfirmMoveTo := func(c *Cell, i int, r interface{}) {
		if x := c.MoveTo(i).(*Cell); !x.Equal(r) {
			t.Fatalf("%v.MoveTo(%v) should be '%v' but is '%v'", c, i, r, x.Content())
		}
	}
	RefuteMoveTo := func(c *Cell, i int) {
		if x := c.MoveTo(i); x != Node(nil) {
			t.Fatalf("%v.MoveTo(%v) should be nil but is %v of type %v", c, i, x, reflect.TypeOf(x))
		}
	}
	c := Cons(0, 1, 2, 3, 4)
	RefuteMoveTo(c, PREVIOUS_NODE)
	ConfirmMoveTo(c, CURRENT_NODE, 0)
	ConfirmMoveTo(c, NEXT_NODE, 1)
	ConfirmMoveTo(c, 2, 2)
	ConfirmMoveTo(c, 3, 3)
	ConfirmMoveTo(c, 4, 4)
	RefuteMoveTo(c, 5)
}

func TestCellSet(t *testing.T) {
	ConfirmSet := func(c *Cell, i int, v interface{}, r interface{}) {
		switch {
		case !c.Set(i, v):		t.Fatalf("Set(%v, %v) failed", i, v)
		case !c.Equal(r):			t.Fatalf("Set(%v, %v) should be %v but is %v", i, v, r, c)
		}
	}
	ConfirmSet(Cons(0), CURRENT_NODE, 1, Cons(1))
	ConfirmSet(Cons(0, 1), NEXT_NODE, 2, Cons(0, 2))
	ConfirmSet(Cons(0, 1, 2), NEXT_NODE + 1, 3, Cons(0, 1, 2, 3))
}
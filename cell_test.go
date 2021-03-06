package chain

import (
	"testing"
	"reflect"
)

func TestCellEnd(t *testing.T) {
	ConfirmEnd := func(c *Cell, r interface{}) {
		x := c.End()
		switch {
		case x == nil:
			t.Fatalf("%v.End() returned nil", c)
		case x.Head != r:
			t.Fatalf("%v.End() should be '%v' but is '%v'", c, r, x.Head)
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

func TestCellEqual(t *testing.T) {
	ConfirmEqual := func(c, o *Cell) {
		if !c.Equal(o) {
			t.Fatalf("%v.Equal(%v) should be true", c, o)
		}
	}
	RefuteEqual := func(c, o *Cell) {
		if c.Equal(o) {
			t.Fatalf("%v.Equal(%v) should be false", c, o)
		}
	}

	ConfirmEqual(Cons(), Cons())
	ConfirmEqual(Cons(1), Cons(1))
	ConfirmEqual(Cons(Cons(2, 3), 1), Cons(Cons(2, 3), 1))

	RefuteEqual(Cons(), Cons(1))
	RefuteEqual(Cons(), Cons(Cons(1)))
	RefuteEqual(Cons(1), Cons(2))
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
		case !c.Set(i, v):
			t.Fatalf("Set(%v, %v) failed", i, v)
		case !c.Equal(r):
			t.Fatalf("Set(%v, %v) should be %v but is %v", i, v, r, c)
		}
	}
	ConfirmSet(Cons(0), CURRENT_NODE, 1, Cons(1))
	ConfirmSet(Cons(0, 1), NEXT_NODE, 2, Cons(0, 2))
	ConfirmSet(Cons(0, 1, 2), NEXT_NODE + 1, 3, Cons(0, 1, 2, 3))
}

func TestCellAppend(t *testing.T) {
	ConfirmAppend := func(c *Cell, v interface{}, r interface{}) {
		c.Append(v)
		tail := c.Tail
		if !tail.Equal(r) {
			t.Fatalf("%v.Append(%v) should have tail %v but has %v", c, v, r, tail)
		}
	}
	ConfirmAppend(Cons(1), 2, Cons(2))
}

func TestCellPrepend(t *testing.T) {
	ConfirmPrepend := func(c *Cell, v interface{}, r interface{}) {
		c.Prepend(v)
		if !c.Equal(r) {
			t.Fatalf("%v.Prepend(%v) should be %v but is %v", c.Tail, v, r, c)
		}
	}
	ConfirmPrepend(Cons(), 1, Cons())
	ConfirmPrepend(Cons(1), 2, Cons(2))
}

func TestCellString(t *testing.T) {
	ConfirmString := func(c *Cell, r string) {
		if s := c.String(); s != r {
			t.Fatalf("%v.String() should be %v but is %v", c, r, s)
		}
	}

	ConfirmString(Cons(0, 1, 2), "(0 1 2)")
	ConfirmString(Cons(0, Cons(1, Cons(2))), "(0 (1 (2)))")
}

func TestCellEach(t *testing.T) {
	list := Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	count := 0
	list.Each(func(i interface{}) {
		if i != count {
			t.Fatalf("element %v erroneously reported as %v", count, i)
		}
		count++
	})

	list.Each(func(index int, i interface{}) {
		if i != index {
			t.Fatalf("element %v erroneously reported as %v", index, i)
		}
	})

	list.Each(func(key, i interface{}) {
		if i != key {
			t.Fatalf("element %v erroneously reported as %v", key, i)
		}
	})
}

func TestCellWhile(t *testing.T) {
	list := Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c *Cell, l int, f interface{}) {
		if count, _ := c.While(f); count != l {
			t.Fatalf("%v.While() should have iterated %v times not %v times", c, l, count)
		}
	}

	count := 0
	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		if count == limit {
			return false
		}
		count++
		return true
	})

	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index != limit
	})

	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key.(int) != limit
	})
}

func TestCellUntil(t *testing.T) {
	list := Cons(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ConfirmLimit := func(c *Cell, l int, f interface{}) {
		if count, _ := c.Until(f); count != l {
			t.Fatalf("%v.Until() should have iterated %v times not %v times", c, l, count)
		}
	}

	count := 0
	limit := 5
	ConfirmLimit(list, limit, func(i interface{}) bool {
		if count == limit {
			return true
		}
		count++
		return false
	})

	ConfirmLimit(list, limit, func(index int, i interface{}) bool {
		return index == limit
	})

	ConfirmLimit(list, limit, func(key, i interface{}) bool {
		return key.(int) == limit
	})
}

func TestCellLen(t *testing.T) {
	ConfirmLen := func(c *Cell, l int) {
		if r := c.Len(); r != l {
			t.Fatalf("%v.Len() should be %v but is %v", c, l, r)
		}
	}
	ConfirmLen(Cons(), 0)
	ConfirmLen(Cons(0), 1)
	ConfirmLen(Cons(0, 1), 2)
	ConfirmLen(Cons(Cons(0, 1), 1), 2)
}

func TestCellMinimumLength(t *testing.T) {
	ConfirmMinimumLength := func(c *Cell, l int) {
		if !c.MinimumLength(l) {
			t.Fatalf("%v.MinimumLength(%v) should be true", c, l)
		}
	}
	RefuteMinimumLength := func(c *Cell, l int) {
		if c.MinimumLength(l) {
			t.Fatalf("%v.MinimumLength(%v) should be false", c, l)
		}
	}
	ConfirmMinimumLength(Cons(), 0)
	RefuteMinimumLength(Cons(), 1)
	ConfirmMinimumLength(Cons(0), 0)
	ConfirmMinimumLength(Cons(0), 1)
	RefuteMinimumLength(Cons(0), 2)
	ConfirmMinimumLength(Cons(0, 1), 0)
	ConfirmMinimumLength(Cons(0, 1), 1)
	ConfirmMinimumLength(Cons(0, 1), 2)
	RefuteMinimumLength(Cons(0, 1), 3)
}
package chain

import "testing"

func TestLastElement(t *testing.T) {
	ConfirmLastElement := func(n Node, r interface{}) {
		x := LastElement(n)
		switch x := x.(type) {
		case Equatable: 	if !x.Equal(r) {
								t.Fatalf("Last(%v) Equatable: should be '%v' but is '%v'", n, r, x)
							}
		default:		 	if r != x {
								t.Fatalf("Last(%v) default: should be '%v' but is '%v'", n, r, x)
							}
		}
	}
	ConfirmLastElement(Cons(0), 0)
	ConfirmLastElement(Cons(0, 1), 1)
	ConfirmLastElement(Cons(0, 1, 2), 2)
}

func TestMoveToNode(t *testing.T) {
	ConfirmMoveTo := func(n Node, i int, r interface{}) {
		switch x := n.MoveTo(i); {
		case x.Content() != r:	t.Fatalf("%v.MoveTo(%v) should be '%v' but is '%v'", n, i, r, x.Content())
		}
	}
	RefuteMoveTo := func(n Node, i int) {
		if x := n.MoveTo(i); x != nil {
			t.Fatalf("%v.MoveTo(%v) should not succeed", n, i)
		}
	}
	c := Cons(1, 2, 3, 4, 5)
	ConfirmMoveTo(c, 0, 1)
	ConfirmMoveTo(c, 1, 2)
	ConfirmMoveTo(c, 2, 3)
	ConfirmMoveTo(c, 3, 4)
	ConfirmMoveTo(c, 4, 5)
	RefuteMoveTo(c, -1)
	RefuteMoveTo(c, 5)
}
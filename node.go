package chain

type Node interface {
	MoveTo(int) Node
	Content() interface{}
	Link(int, Node) bool
	Set(int, interface{}) bool
}

const(
	PREVIOUS_NODE = -1
	CURRENT_NODE = 0
	NEXT_NODE = 1
)

func Next(l Node) (r Node) {
	if l != nil {
		r = l.MoveTo(NEXT_NODE)
	}
	return
}

func Previous(l Node) Node {
	return l.MoveTo(PREVIOUS_NODE)
}

func LastElement(l Node) (r Node) {
	if r = l; r != nil {
		for {
			if n := Next(r); n != nil {
				r = n
			} else {
				break
			}
		}
	}
	return
}

func FirstElement(l Node) (r Node) {
	if r = l; r != nil {
		for {
			if n := Previous(r); n != nil {
				r = n
			} else {
				break
			}
		}
	}
	return
}
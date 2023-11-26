package node

type nodeInputTypes interface {
	string | int | float64 | bool
}

type Node[e nodeInputTypes] struct {
	element e
	next    *Node[e]
}

func NewNode[v nodeInputTypes](element v, next *Node[v]) *Node[v] {
	return &Node[v]{element: element, next: next}
}
func (n *Node[v]) GetElement() any {
	return n.element
}
func (n *Node[v]) GetNext() *Node[v] {
	return n.next
}

func (n *Node[v]) SetNext(no *Node[v]) {
	n.next = no
}

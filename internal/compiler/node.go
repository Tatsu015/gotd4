package compiler

type Node struct {
	val      int
	children []Node
}

func NewNode(val int) Node {
	return Node{val, []Node{}}
}

func (n *Node) add(c Node) {
	n.children = append(n.children, c)
}

func (n *Node) convert() []byte {
	ml := []byte{}
	return ml
}

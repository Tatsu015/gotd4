package compiler

type Node struct {
	val      int
	children []Node
}

func NewNode(val int) Node {
	return Node{val, []Node{}}
}

func (n *Node) Add(c Node) {
	n.children = append(n.children, c)
}

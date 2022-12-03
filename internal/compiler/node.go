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

func (n *Node) ntob() []byte {
	t := []byte{}
	for _, c := range n.children {
		b := c.ntob()
		t = append(t, b...)
	}
	return append(t, byte(n.val))
}

func (n *Node) toBinary() []byte {
	ml := n.ntob()
	return ml
}
